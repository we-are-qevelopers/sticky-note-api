package useCase

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/model"
	"github.com/we-are-qevelopers/sticky-note-api/domain/repository"
)

type StickyNoteUsecase interface {
	Get(userID string) ([]model.StickyNote, error)
}

type stickyNoteUsecase struct {
	stickyNoteRepo repository.StickyNoteRepository
}

func NewStickyNoteUsecase(stickyNoteRepo repository.StickyNoteRepository) StickyNoteUsecase {
	newStickyNoteUsecase := stickyNoteUsecase{stickyNoteRepo: stickyNoteRepo}

	return &newStickyNoteUsecase
}

func (usecase *stickyNoteUsecase) Get(userID string) ([]model.StickyNote, error) {
	STnotes, err := usecase.stickyNoteRepo.Get(userID)
	return STnotes, err
}
