package main

import (
	"log"

	"github.com/rrabit42/mvp-backend/internal/api"
	"github.com/rrabit42/mvp-backend/internal/config"
	"github.com/rrabit42/mvp-backend/internal/repository"
	"github.com/rrabit42/mvp-backend/internal/service"
)

func main() {
	// 1. 설정 로드
	cfg, err := config.LoadConfig(".") // config.toml이 있는 경로
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// 2. 리포지토리 초기화 (의존성 주입)
	// MVP에서는 인메모리 및 모의 리포지토리를 사용합니다.
	policyRepo := repository.NewInMemoryPolicyRepository()
	cardRepo := repository.NewMockCardRepository()

	// 3. 서비스 초기화 (의존성 주입)
	policyService := service.NewPolicyService()
	cardService := service.NewCardService(policyRepo, cardRepo)

	// 4. 라우터 설정
	router := api.SetupRouter(policyService, cardService)

	// 5. 서버 실행
	serverAddr := ":" + cfg.Server.Port
	log.Printf("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
