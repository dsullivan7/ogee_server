package store

import (
	"go_server/internal/models"

	"github.com/google/uuid"
)

type Store interface {
	GetUser(userID uuid.UUID) (*models.User, error)
	ListUsers(query map[string]interface{}) ([]models.User, error)
	CreateUser(userPayload models.User) (*models.User, error)
	ModifyUser(userID uuid.UUID, userPayload models.User) (*models.User, error)
	DeleteUser(userID uuid.UUID) error
}
