package repository

import (
	"github.com/google/uuid"
	userdomain "myapp/module/user/domain"
	"time"
)

type UserDTO struct {
	Id        uuid.UUID `gorm:"column:id;"`
	FirstName string    `gorm:"column:first_name;"`
	LastName  string    `gorm:"column:last_name;"`
	Email     string    `gorm:"column:email;"`
	Password  string    `gorm:"column:password;"`
	Salt      string    `gorm:"column:salt;"`
	Role      string    `gorm:"column:role;"`
	Status    string    `gorm:"column:status;"`
}

func (dto *UserDTO) ToEntity() (*userdomain.User, error) {
	return userdomain.NewUser(dto.Id, dto.FirstName, dto.LastName, dto.Email, dto.Password, dto.Salt, userdomain.GetRole(dto.Role), dto.Status)
}

type SessionDTO struct {
	Id           uuid.UUID `gorm:"column:id;"`
	UserId       uuid.UUID `gorm:"column:user_id;"`
	RefreshToken string    `gorm:"column:refresh_token;"`
	AccessExpAt  time.Time `gorm:"column:access_exp_at;"`
	RefreshExpAt time.Time `gorm:"column:refresh_exp_at;"`
}

func (dto SessionDTO) ToEntity() (*userdomain.Session, error) {
	s := userdomain.NewSession(dto.Id, dto.UserId, dto.RefreshToken, dto.AccessExpAt, dto.RefreshExpAt)
	return s, nil
}
