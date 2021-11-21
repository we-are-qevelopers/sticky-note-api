package get

import (
	"sticky-note-api/interfaceAdaptor/gin/controllers"
	"sticky-note-api/useCase"
)

type StickyNoteController struct {
	stickyNoteUseCase useCase.StickyNoteUsecase
}

func NewStickyNoteController(stickyNoteUseCase useCase.StickyNoteUsecase) *StickyNoteController {
	return &StickyNoteController{stickyNoteUseCase: stickyNoteUseCase}
}

func (controller *StickyNoteController) Get(c controllers.Context) {

	c.JSON(200, "get STnotes")
}
