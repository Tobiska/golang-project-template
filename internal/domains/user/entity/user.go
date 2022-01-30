package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int    `json:"id,omitempty"`
	Role              string `json:"role,omitempty"`
	Email             string `json:"email,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	PasswordEncrypted string `json:"-"`
	GroupID           string `json:"group_id"`
}

const (
	Admin  string = "admin"
	Client string = "client"
)

func (u *User) EncryptPassword() error {
	enc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.PasswordEncrypted = string(enc)
	if err := u.SanitizeUser(); err != nil {
		return err
	}
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordEncrypted), []byte(password)) == nil
}

func (u *User) SanitizeUser() error {
	u.Password = ""
	return nil
}

func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

func (u *User) IsClient() bool {
	return u.Role == Client
}
