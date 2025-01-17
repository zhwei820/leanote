package controllers

import (
	"github.com/revel/revel"
)

type NoteContentHistory struct {
	BaseController
}

// 得到list
func (c NoteContentHistory) ListHistories(noteId string) {
	histories := noteContentHistoryService.ListHistories(noteId, c.GetUserId())

	return c.RenderJSON(histories)
}
