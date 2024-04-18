package repository

import (
	userDomain "myapp/module/user/domain"
	"time"

	"github.com/google/uuid"
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
	Avatar    *string   `gorm:"column:avatar;"`
}

func (dto *UserDTO) ToEntity() (*userDomain.User, error) {
	return userDomain.NewUser(dto.Id, dto.FirstName, dto.LastName, dto.Email, dto.Password, dto.Salt, userDomain.GetRole(dto.Role), dto.Status, StringFromPointer(dto.Avatar))
}

type SessionDTO struct {
	Id           uuid.UUID `gorm:"column:id;"`
	UserId       uuid.UUID `gorm:"column:user_id;"`
	RefreshToken string    `gorm:"column:refresh_token;"`
	AccessExpAt  time.Time `gorm:"column:access_exp_at;"`
	RefreshExpAt time.Time `gorm:"column:refresh_exp_at;"`
}

func (dto SessionDTO) ToEntity() (*userDomain.Session, error) {
	s := userDomain.NewSession(dto.Id, dto.UserId, dto.RefreshToken, dto.AccessExpAt, dto.RefreshExpAt)
	return s, nil
}

func StringFromPointer(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func GetStrPt(s string) *string {
	return &s
}
