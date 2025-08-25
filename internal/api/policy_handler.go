package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrabit42/mvp-backend/internal/service"
)

type PolicyHandler struct {
	policyService service.PolicyService
}

func NewPolicyHandler(policyService service.PolicyService) *PolicyHandler {
	return &PolicyHandler{policyService: policyService}
}

type CompileRequest struct {
	UserID string `json:"userId"`
	Prompt string `json:"prompt"`
}

func (h *PolicyHandler) CompilePolicy(c *gin.Context) {
	var req CompileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	policy, err := h.policyService.CompilePolicy(c.Request.Context(), req.UserID, req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compile policy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"policy": policy})
}
