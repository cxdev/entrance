package platform

import (
	"testing"
	"time"
)

func TestBaseEntity_IsEntity(t *testing.T) {
	type fields struct {
		Id      int64
		Created time.Time
		Updated time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Empty case", fields{}, false},
		{"Normal case", fields{1, time.Now(), time.Now()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &BaseEntity{
				Id:      tt.fields.Id,
				Created: tt.fields.Created,
				Updated: tt.fields.Updated,
			}
			if got := entity.IsEntity(); got != tt.want {
				t.Errorf("BaseEntity.IsEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
