package response

import (
	"go-learn-news-portal/internal/core/entity"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromEntityUser(e *entity.User) *User {
	return &User{
		ID:    e.ID,
		Name:  e.Name,
		Email: e.Email,
	}
}
