package controllers

import (
	"github.com/we-are-qevelopers/sticky-note-api/useCase"
)

type StickyNoteController struct {
	stickyNoteUseCase useCase.StickyNoteUsecase
}

func NewStickyNoteController(stickyNoteUseCase useCase.StickyNoteUsecase) *StickyNoteController {
	return &StickyNoteController{stickyNoteUseCase: stickyNoteUseCase}
}

func (controller *StickyNoteController) Get(c Context) {

	// stickyNote, err := controller.stickyNoteUseCase.Get("2")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(stickyNote)

	c.JSON(200, "stickyNote")
}
