package entities

import "time"

type BaseEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Entity interface {
	TableName() string
	ToDTO() *interface{}
	FromDTO(dto *interface{})
}
