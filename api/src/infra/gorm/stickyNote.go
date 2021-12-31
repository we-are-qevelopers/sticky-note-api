package gorm

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/model"
	"github.com/we-are-qevelopers/sticky-note-api/domain/repository"
	"time"

	"gorm.io/gorm"
)

type StickyNoteRepository struct {
	sqlHandler *gorm.DB
}

func NewStickyNoteRepository(sqlHandler *gorm.DB) repository.StickyNoteRepository {
	stickyNoteRepository := StickyNoteRepository{sqlHandler}
	return &stickyNoteRepository
}

func (repo *StickyNoteRepository) Get(userID string) ([]model.StickyNote, error) {
	stickyNotes := []model.StickyNote{
		{
			ID:        1,
			Content:   "sss",
			Position:  model.Matrix{X: 100, Y: 100},
			Size:      model.Matrix{X: 100, Y: 100},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return stickyNotes, nil
}
