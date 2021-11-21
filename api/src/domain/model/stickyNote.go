package model

import "time"

type StickyNote struct {
	ID        int
	Content   string
	Position  Matrix
	Size      Matrix
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Matrix struct {
	X int
	Y int
}
