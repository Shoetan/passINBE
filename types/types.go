package types

import "time"


type Login struct {
	Email string `json:"email"`
	Password string `json:"master_password"`
}

type User struct {
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	MasterPassword string `json:"master_password"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	UserID int `json:"user_id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegisterPayload struct {
	Email string `json:"email"`
	Name string `json:"name"`
	MasterPassword string `json:"master_password"`
	
}

type AddRecordPayload struct {
	RecordEmail string `json:"record_email"`
	RecordName string `json:"record_name"`
	RecordPassword string `json:"record_password"`
}

