package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/RaihanMalay21/api-service-riors/repository/admin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebsocketRiors struct {
	db       *admin.AdminUsersRepository
	upgrader websocket.Upgrader
}

func ConstructorWebsocket(db *admin.AdminUsersRepository, upgrader websocket.Upgrader) *WebsocketRiors {
	return &WebsocketRiors{
		db:       db,
		upgrader: upgrader,
	}
}

func GetUserIdFromCookie(e echo.Context) (uint, error) {
	cookie, err := e.Cookie("user_riors_token")
	if err != nil {
		return 0, err
	}

	tokenString := cookie.Value
	claim := &config.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*config.JWTClaim)
	if ok {
		return claims.Id, nil
	}

	return 0, fmt.Errorf("cannot find id user")
}

func (ws *WebsocketRiors) WebsocketDetectionStatusActiveUser(e echo.Context) error {
	conn, err := ws.upgrader.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return err
	}
	defer conn.Close()

	// var db repo.AdminUsersRepository

	userId, _ := GetUserIdFromCookie(e)

	fmt.Println(userId)
	if userId == 0 {
		log.Println("User ID not found in cookies")
		return e.JSON(http.StatusBadRequest, "User ID not found")
	}

	status := true
	if err := ws.db.UpdateUserActiveById(userId, status, time.Now()); err != nil {
		log.Println("Terjadi error saat update active user: ", err)
		return e.JSON(http.StatusInternalServerError, "error terjadi error when update user active")
	}

	active := make(chan bool)
	defer close(active)

	go func() {
		const timeoutDuration = 60 * time.Second
		timer := time.NewTimer(timeoutDuration)
		defer timer.Stop()

		for {
			select {
			case <-active:
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(timeoutDuration)
			case <-timer.C:
				if err := ws.db.UpdateUserActiveById(userId, false, time.Now()); err != nil {
					log.Println("Error updating inactive user: ", err)
				}
				log.Println("Connection timed out due to inactivity")
				conn.WriteMessage(websocket.CloseMessage, []byte("Timeout: Connection closed due to inactivity"))
				conn.Close()
				return
			}
		}
	}()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		active <- true

		fmt.Printf("Received message: %s\n", msg)

		if err := conn.WriteMessage(msgType, []byte("pong")); err != nil {
			log.Println("Error sending pong: ", err)
			break
		}
		fmt.Println("sent pong")
	}

	if err := ws.db.UpdateUserActiveById(userId, false, time.Now()); err != nil {
		log.Println("Error updating inactive user on disconnect:", err)
	}

	return nil
}
