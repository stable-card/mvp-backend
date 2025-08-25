package repository

import (
	"context"

	"github.com/rrabit42/mvp-backend/internal/domain"
)

// PolicyRepository는 정책 데이터에 대한 영속성 작업을 정의합니다.
type PolicyRepository interface {
	Save(ctx context.Context, userID string, policy *domain.Policy) error
}

// inMemoryPolicyRepository는 메모리를 사용하여 정책을 저장하는 구현체입니다.
type inMemoryPolicyRepository struct {
	policies map[string]*domain.Policy
}

// NewInMemoryPolicyRepository는 새로운 inMemoryPolicyRepository를 생성합니다.
func NewInMemoryPolicyRepository() PolicyRepository {
	return &inMemoryPolicyRepository{
		policies: make(map[string]*domain.Policy),
	}
}

// Save는 정책을 메모리에 저장합니다.
func (r *inMemoryPolicyRepository) Save(ctx context.Context, userID string, policy *domain.Policy) error {
	// 실제로는 여기에 DB 저장 로직이 들어갑니다.
	r.policies[userID] = policy
	return nil
}
