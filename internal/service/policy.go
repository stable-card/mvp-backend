package service

import (
	"context"

	"github.com/rrabit42/mvp-backend/internal/domain"
)

// PolicyService는 정책 생성과 관련된 비즈니스 로직을 담당합니다.
type PolicyService interface {
	CompilePolicy(ctx context.Context, userID, prompt string) (*domain.Policy, error)
}

type policyService struct {
	// llmClient llm.Client // 실제로는 LLM 클라이언트가 여기에 주입됩니다.
}

// NewPolicyService는 새로운 PolicyService를 생성합니다.
func NewPolicyService() PolicyService {
	return &policyService{}
}

// CompilePolicy는 사용자의 자연어 프롬프트를 받아 정책으로 컴파일합니다.
// MVP 단계에서는 LLM 호출 대신 더미 정책을 생성합니다.
func (s *policyService) CompilePolicy(ctx context.Context, userID, prompt string) (*domain.Policy, error) {
	// 실제 LLM 호출 로직을 여기에 구현합니다.
	// 예시: policy, err := s.llmClient.Compile(ctx, prompt)

	// MVP를 위한 더미 구현
	mockPolicy := &domain.Policy{
		Version: 1,
		Rules: []domain.Rule{
			{
				Category:    "BAKERY",
				BenefitType: "CASHBACK",
				Value:       1000, // 10%
			},
			{
				Category:    "CAFE",
				BenefitType: "CASHBACK",
				Value:       1000, // 10%
			},
		},
	}

	return mockPolicy, nil
}
