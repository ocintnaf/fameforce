package entities

import "time"

type BaseEntity[TID any] struct {
	ID        TID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Entity interface {
	TableName() string
	ToDTO() any
	FromDTO(dto any)
}
