package authentication

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	PushRedisRegister(data *dto.RegisterUser, expr time.Duration) error
	GetRediRegistrationByEmail(email string) (*dto.RegisterUser, error)
	DeleteRedisRegister(email string) error
	CreateRegistration(data *domain.User) error
}

type authenticationRepository struct {
	db     *gorm.DB
	client *redis.Client
}

func ConstructorAuthenticationRepository(db *gorm.DB, redisClient *redis.Client) AuthenticationRepository {
	return &authenticationRepository{
		db:     db,
		client: redisClient,
	}
}

func (d *authenticationRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := d.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *authenticationRepository) PushRedisRegister(data *dto.RegisterUser, expr time.Duration) error {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx := context.Background()

	if err := d.client.Set(ctx, data.Email, dataJson, expr).Err(); err != nil {
		return err
	}

	return nil
}

func (d *authenticationRepository) GetRediRegistrationByEmail(email string) (*dto.RegisterUser, error) {
	ctx := context.Background()
	val, err := d.client.Get(ctx, email).Result()
	if err != nil {
		return nil, err
	}

	var data dto.RegisterUser
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *authenticationRepository) DeleteRedisRegister(email string) error {
	ctx := context.Background()
	if err := d.client.Del(ctx, email).Err(); err != nil {
		return err
	}
	return nil
}

func (d *authenticationRepository) CreateRegistration(data *domain.User) error {
	if err := d.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
