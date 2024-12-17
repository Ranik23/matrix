package postgres

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Storage interface {
    GetUserByID(id int) (*User, error)
    CreateUser(user *User) error
    UpdateUser(user *User) error 
    DeleteUser(id int) error 
}

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(connString string) (*PostgresStorage, error) {
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	// Автомиграция: создаёт таблицу, если её нет
	db.AutoMigrate(&User{})
	return &PostgresStorage{db: db}, nil
}

func (p *PostgresStorage) GetUserByID(id int) (*User, error) {
	var user User
	if err := p.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (p *PostgresStorage) CreateUser(user *User) error {
	if err := p.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresStorage) UpdateUser(user *User) error {
	if err := p.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresStorage) DeleteUser(id int) error {
	if err := p.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

type User struct {
	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}