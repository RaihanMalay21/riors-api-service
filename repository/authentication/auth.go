package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthenticationRepository struct {
	db     *gorm.DB
	client *redis.Client
}

func ConstructorAuthenticationRepository(db *gorm.DB, redisClient *redis.Client) *AuthenticationRepository {
	return &AuthenticationRepository{
		db:     db,
		client: redisClient,
	}
}

func (d *AuthenticationRepository) NewTransactionAuth() *gorm.DB {
	return d.db.Begin()
}

func (d *AuthenticationRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := d.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *AuthenticationRepository) PushRedisRegister(data *dto.RegisterUser, expr time.Duration) error {
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

func (d *AuthenticationRepository) GetRediRegistrationByEmail(email string) (*dto.RegisterUser, error) {
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

func (d *AuthenticationRepository) DeleteRedisRegister(email string) error {
	ctx := context.Background()
	if err := d.client.Del(ctx, email).Err(); err != nil {
		return err
	}
	return nil
}

func (d *AuthenticationRepository) CreateUser(data *domain.User) error {
	if err := d.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (d *AuthenticationRepository) GetEmployeeByEmail(email string) (*domain.Employee, error) {
	var employee domain.Employee
	if err := d.db.Where("email = ?", email).First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (d *AuthenticationRepository) CreateEmployee(tx *gorm.DB, data *domain.Employee) error {
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (dp *AuthenticationRepository) UpdateEmployeImage(tx *gorm.DB, data *domain.Employee) error {
	if err := tx.Model(data).Update("image", data.Image).Error; err != nil {
		return err
	}

	return nil
}

func (dp *AuthenticationRepository) GetPasswordById(id uint, structType string) (interface{}, error) {
	var result interface{}

	switch structType {
	case "user":
		result = &domain.User{}
	case "employee":
		result = &domain.Employee{}
	default:
		return nil, errors.New("invalid user type")
	}

	if err := dp.db.Select("password").Where("id = ?", id).First(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (dp *AuthenticationRepository) UpdatePasswordById(password string, id uint, structType string) error {
	var result interface{}

	switch structType {
	case "user":
		result = &domain.User{}
	case "employee":
		result = &domain.Employee{}
	default:
		return errors.New("invalid user type")
	}

	if err := dp.db.Model(result).Where("id = ?", id).Update("password", password).Error; err != nil {
		return err
	}

	return nil
}
