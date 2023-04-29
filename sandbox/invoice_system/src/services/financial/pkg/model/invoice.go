package model

import (
	base "project/src/pkg/model"
	"time"
)

type Invoice struct {
	Base     base.Base `json:"base"`
	Value    float64   `json:"value"`
	ClientId uint32    `json:"clientId"`
	DueDate  time.Time `json:"dueDate"`
}
