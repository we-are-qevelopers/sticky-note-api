package repository

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/model"
)

type StickyNoteRepository interface {
	Get(userID string) ([]model.StickyNote, error)
}
