package models

import "time"

type Todo struct {
	ID          uint
	Name        string
	Note        string
	IsCaomplete bool
	CreatedAt   time.Time
}
