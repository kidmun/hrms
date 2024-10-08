package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	RoleAdmin           Role = "Admin"
	RoleHRManager       Role = "HRManager"
	// RoleLineManager     Role = "LineManager"
	RoleEmployee        Role = "Employee"
	// RoleSystemAdmin     Role = "SystemAdministrator"
	// RoleExternalAuditor Role = "ExternalAuditor"
)

type User struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Username     string              `bson:"username" json:"username"`
	Password     string              `bson:"password" json:"-"`
	Email        string              `bson:"email" json:"email"`
	Role         Role                `bson:"role" json:"role"`
	EmployeeID   *primitive.ObjectID `bson:"employee_id,omitempty" json:"employee_id,omitempty"`
	CreatedAt    time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time           `bson:"updated_at" json:"updated_at"`
}

type UserRepository interface {
    CreateUser(ctx context.Context, user *User) (primitive.ObjectID, error)
	GetUsers(ctx context.Context) ([]*User, error)
    GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
    GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
    UpdateUser(ctx context.Context, user *User) error
    DeleteUser(ctx context.Context, id primitive.ObjectID) error
}

type UserUsecase interface {
    CreateUser(ctx context.Context, user *User) (primitive.ObjectID, error)
	GetUsers(ctx context.Context) ([]*User, error)
    GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
    GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
    UpdateUser(ctx context.Context, user *User) error
    DeleteUser(ctx context.Context, id primitive.ObjectID) error
}


type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
}
