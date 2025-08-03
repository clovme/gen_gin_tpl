package public

import (
	"gen_gin_tpl/pkg/enums/status"
	"time"
)

type ApiEnumsVO struct {
	ID          int64         `json:"id"`
	Category    string        `json:"category"`
	Key         string        `json:"key"`
	Name        string        `json:"name"`
	Value       int           `json:"value"`
	ValueT      int           `json:"valuet"`
	Sort        int           `json:"sort"`
	Status      status.Status `json:"status"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}
