package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rrabit42/mvp-backend/internal/service"
)

func SetupRouter(policyService service.PolicyService, cardService service.CardService) *gin.Engine {
	router := gin.Default()

	policyHandler := NewPolicyHandler(policyService)
	cardHandler := NewCardHandler(cardService)

	// API v1 그룹
	apiV1 := router.Group("/api/v1")
	{
		policies := apiV1.Group("/policies")
		{
			policies.POST("/compile", policyHandler.CompilePolicy)
		}

		cards := apiV1.Group("/cards")
		{
			cards.POST("/issue", cardHandler.IssueCard)
		}
	}

	return router
}
