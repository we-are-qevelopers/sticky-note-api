package repository

import (
	"sticky-note-api/domain/model"
)

type StickyNoteRepository interface {
	Get(userID string) ([]model.StickyNote, error)
}
