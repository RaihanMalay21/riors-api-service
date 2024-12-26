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

type AuthenticationRepository interface {
	NewTransactionAuth() *gorm.DB
	GetUserByEmail(email string) (*domain.User, error)
	PushRedisRegister(data *dto.RegisterUser, expr time.Duration) error
	GetRediRegistrationByEmail(email string) (*dto.RegisterUser, error)
	DeleteRedisRegister(email string) error
	CreateUser(data *domain.User) error
	GetEmployeeByEmail(email string) (*domain.Employee, error)
	CreateEmployee(tx *gorm.DB, data *domain.Employee) error
	UpdateEmployeImage(tx *gorm.DB, data *domain.Employee) error
	// EmployeeUpdatePasswordById(data *domain.Employee) error
	// UserUpdatePasswordById(data *domain.User) error
	UpdatePasswordById(password string, id uint, structType string) error
	GetPasswordById(id uint, structType string) (interface{}, error)
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

func (d *authenticationRepository) NewTransactionAuth() *gorm.DB {
	return d.db.Begin()
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

func (d *authenticationRepository) CreateUser(data *domain.User) error {
	if err := d.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (d *authenticationRepository) GetEmployeeByEmail(email string) (*domain.Employee, error) {
	var employee domain.Employee
	if err := d.db.Where("email = ?", email).First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (d *authenticationRepository) CreateEmployee(tx *gorm.DB, data *domain.Employee) error {
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (dp *authenticationRepository) UpdateEmployeImage(tx *gorm.DB, data *domain.Employee) error {
	if err := tx.Model(data).Update("image", data.Image).Error; err != nil {
		return err
	}

	return nil
}

func (dp *authenticationRepository) GetPasswordById(id uint, structType string) (interface{}, error) {
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


func (dp *authenticationRepository) UpdatePasswordById(password string, id uint, structType string) error {
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

// func (dp *authenticationRepository) UserUpdatePasswordById(data *domain.User) error {
// 	if err := dp.db.Model(data).Update("password", data.Password).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
