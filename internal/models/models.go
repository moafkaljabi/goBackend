package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// change account to devices.

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	EncryptedPassword string    `json:"-"` // never send to client
	Number            int64     `json:"number"`
	Balance           float64   `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)) == nil
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:                99,
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encryptedPassword),
		Number:            100,
		Balance:           1000,
		CreatedAt:         time.Now().UTC(),
	}, nil
}

type Device struct {
	DeviceID int
	Name     string
	Status   string
}

func NewDevice(name, status string) *Device {
	return &Device{
		DeviceID: 100,
		Name:     name,
	}
}
