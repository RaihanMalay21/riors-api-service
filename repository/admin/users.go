package admin

import (
	"time"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AdminUsersRepository struct {
	db     *gorm.DB
	client *redis.Client
}

func ConstructorAdminUsersRepository(db *gorm.DB, client *redis.Client) *AdminUsersRepository {
	return &AdminUsersRepository{
		db:     db,
		client: client,
	}
}

// func (ur *AdminUsersRepository) RedisGetDataUserActive(latestUserId *string, ctx context.Context) (*string, *[]string, error) {
// 	var message []redis.XStream
// 	var err error

// 	startID := "0"
// 	if *latestUserId != "" {
// 		startID = *latestUserId + ">"
// 	}

// 	message, err = ur.client.XRead(ctx, &redis.XReadArgs{
// 		Streams: []string{"active_stream_user", startID},
// 		Count:   10,
// 		Block:   1,
// 	}).Result()

// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to read from Redis stream: %v", err)
// 	}

// 	if len(message) == 0 || len(message[0].Messages) == 0 {
// 		return nil, nil, nil
// 	}

// 	var ids []string
// 	var latestId string
// 	for _, stream := range message {
// 		for _, msg := range stream.Messages {
// 			if userID, ok := msg.Values["user_id"].(string); ok {
// 				ids = append(ids, userID)
// 			}
// 			latestId = msg.ID
// 		}
// 	}

// 	return &latestId, &ids, nil
// }

func (ur *AdminUsersRepository) UpdateUserActiveById(userId uint, status bool, lastActive time.Time) error {
	if status {
		if err := ur.db.Model(&domain.User{}).Where("id = ?", userId).Update("active", status).Error; err != nil {
			return err
		}
	}

	if !status {
		if err := ur.db.Model(&domain.User{}).Where("id = ?", userId).Updates(map[string]interface{}{"active": status, "last_active": lastActive}).Error; err != nil {
			return err
		}
	}

	return nil
}

// func (ur *AdminUsersRepository) RedisGetDataUserActive(latestUserId *string, ctx context.Context) (*string, *[]string, error) {
// 	users, err :=
// }

// func (ur *AdminUsersRepository) GetUserById(ids *[]string) (*[]domain.User, error) {
// 	var users []domain.User
// 	if err := ur.db.Where("id IN ?", *ids).Omit("password").Find(&users).Error; err != nil {
// 		return nil, err
// 	}

// 	return &users, nil
// }

// func (ur *AdminUsersRepository) CleanupExpiredMessages() {
// 	ctx := context.Background()
// 	messages, _ := ur.client.XRange(ctx, "active_stream_user", "-", "+").Result()

// 	now := time.Now().Unix()
// 	for _, msg := range messages {
// 		if expireAtStr, ok := msg.Values["expire_at"].(string); ok {
// 			expireAt, _ := strconv.ParseInt(expireAtStr, 10, 64)
// 			if now > expireAt {
// 				ur.client.XDel(ctx, "active_stream_user", msg.ID)
// 			}
// 		}
// 	}
// }
