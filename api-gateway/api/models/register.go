package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterOwner struct {
	FullName    string `json:"full_name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	Tax         int64  `json:"tax"`
}

type RegisterOwnerResponse struct {
	Message string `json:"message"`
}

type OwnerResponse struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	Tax         int64  `json:"tax"`
	AccessToken string `json:"access_token"`
}

type ResponseAccessToken struct {
	AccessToken string `json:"access_token"`
	WorkerId    string `json:"worker_id"`
	Role        string `json:"role"`
	OwnerId     string `json:"owner_id"`
}

type ResponseOwnerLogin struct {
	AccessToken string `json:"access_token"`
	OwnerId     string `json:"owner_id"`
}

type LoginWorker struct {
	CompanyName string `json:"company_name"`
	LoginKey    string `json:"login_key"`
	Password    string `json:"password"`
}

func (rm *RegisterOwner) IsEmail() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Email, validation.Required, is.Email),
	)
}

func (rm *RegisterOwner) IsComplexPassword() error {
	return validation.Validate(
		&rm.Password,
		validation.Required,
		validation.Length(8, 30),
		validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]")),
	)
}
