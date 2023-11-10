package dto

import "time"

type UserResponseDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Number    string    `json:"number"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreateDTO struct {
	Name     string `json:"name"`
	Number   string `json:"number"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResBodyDTO struct {
	Token   string          `json:"token"`
	User    UserResponseDTO `json:"user"`
	Message string          `json:"message"`
	Error   string          `json:"error"`
}
