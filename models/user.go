package models

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
