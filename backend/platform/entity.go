package platform

import (
	"time"
)

type BaseEntity struct {
	Id      int64
	Created time.Time
	Updated time.Time
}

func (entity *BaseEntity) IsEntity() bool {
	return BaseEntity{} != *entity
}
