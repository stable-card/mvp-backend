package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrabit42/mvp-backend/internal/domain"
	"github.com/rrabit42/mvp-backend/internal/service"
)

type CardHandler struct {
	cardService service.CardService
}

func NewCardHandler(cardService service.CardService) *CardHandler {
	return &CardHandler{cardService: cardService}
}

type IssueCardRequest struct {
	UserID string         `json:"userId"`
	Policy *domain.Policy `json:"policy"`
}

func (h *CardHandler) IssueCard(c *gin.Context) {
	var req IssueCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.Policy == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Policy is required"})
		return
	}

	card, err := h.cardService.IssueCard(c.Request.Context(), req.UserID, req.Policy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue card"})
		return
	}

	c.JSON(http.StatusCreated, card)
}
