package handlers

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) index(c *gin.Context) {
	h.t.Render(c.Writer, "index.html", nil)
}