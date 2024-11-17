package user

import "time"

//struct yang terhubung dengan db user
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdateAt       time.Time
}
