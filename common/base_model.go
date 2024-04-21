package common

import (
	"time"

	"github.com/google/uuid"
)

type EntityBase struct {
	id        uuid.UUID `gorm:"column:id;"`
	status    string    `gorm:"column:status"`
	createdAt time.Time `gorm:"column:created_at"`
	updatedAt time.Time `gorm:"column:updated_at"`
}

func GenNewEntityBase() EntityBase {
	now := time.Now().UTC()

	return EntityBase{
		id:        GenUUID(),
		status:    Activated,
		createdAt: now,
		updatedAt: now,
	}
}

func (u EntityBase) Id() uuid.UUID {
	return u.id
}

func (u EntityBase) Status() string {
	return u.status
}

func (u EntityBase) CreatedAt() time.Time {
	return u.createdAt
}

func (u EntityBase) UpdatedAt() time.Time {
	return u.updatedAt
}

func GenUUID() uuid.UUID {
	newId, _ := uuid.NewV7()
	return newId
}

func ParseUUID(s string) uuid.UUID {
	return uuid.MustParse(s)
}
