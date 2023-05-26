package dtos

import "time"

type AdminDetailResponse struct {
	AdministratorID uint      `json:"administrator_id" from:"administrator_id"`
	Nama            string    `json:"nama" from:"nama"`
	Email           string    `json:"email" from:"email"`
	Role            string    `json:"role" from:"role"`
	CreatedAt       time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt       time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

type RegisterAdminRequest struct {
	Nama     string `json:"nama" form:"nama" validate:"required" example:"Rahadina Budiman Sundara"`
	Email    string `json:"email" form:"email" validate:"required,email" example:"me@r4ha.com"`
	Password string `json:"password" form:"password" validate:"gte=6" example:"rahadinabudimansundara"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Admin', 'Super Admin');default:'Admin'; not-null" example:"Admin"`
}

type UpdateAdminRequest struct {
	Nama     string `json:"nama" form:"nama" validate:"required" example:"Rahadina Budiman Sundara"`
	Email    string `json:"email" form:"email" validate:"required,email" example:"me@r4ha.com"`
	Password string `json:"password" form:"password" validate:"gte=6" example:"rahadinabudimansundara"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Admin', 'Super Admin');default:'Admin'; not-null" example:"Admin"`
}

type DeleteAdminRequest struct {
	Password string `json:"password" form:"password" validate:"gte=6" example:"rahadinabudimansundara"`
}

type RegisterAdminResponse struct {
	Nama     string `json:"nama" form:"nama" example:"Rahadina Budiman Sundara"`
	Email    string `json:"email" form:"email" example:"me@r4ha.com"`
	Password string `json:"password" form:"password" example:"rahadinabudimansundara"`
	Role     string `json:"role" form:"role" example:"Admin"`
}

type UpdateAdminResponse struct {
	Nama     string `json:"nama" form:"nama" example:"Rahadina Budiman Sundara"`
	Email    string `json:"email" form:"email" example:"me@r4ha.com"`
	Password string `json:"password" form:"password" example:"rahadinabudimansundara"`
	Role     string `json:"role" form:"role" example:"Admin"`
}

type AdminProfileResponse struct {
	ID    uint   `json:"id" form:"id" example:"1"`
	Nama  string `json:"nama" form:"nama" example:"Rahadina Budiman Sundara"`
	Email string `json:"email" form:"email" example:"me@r4ha.com"`
	Role  string `json:"role" form:"role" example:"Admin"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email" example:"me@r4ha.com"`
	Password string `json:"password" form:"password" validate:"required" example:"rahadinabudimansundara"`
}

type LoginResponse struct {
	Email string `json:"email" form:"email" example:"me@r4ha.com"`
	Token string `json:"token" form:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}