package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rrabit42/mvp-backend/internal/domain"
	"github.com/rrabit42/mvp-backend/internal/repository"
)

// CardService는 카드 발급과 관련된 비즈니스 로직을 담당합니다.
type CardService interface {
	IssueCard(ctx context.Context, userID string, policy *domain.Policy) (*domain.Card, error)
}

type cardService struct {
	policyRepo repository.PolicyRepository
	cardRepo   repository.CardRepository
}

// NewCardService는 새로운 CardService를 생성합니다.
func NewCardService(policyRepo repository.PolicyRepository, cardRepo repository.CardRepository) CardService {
	return &cardService{
		policyRepo: policyRepo,
		cardRepo:   cardRepo,
	}
}

// IssueCard는 새로운 카드를 발급하고 정책을 온체인에 등록합니다.
func (s *cardService) IssueCard(ctx context.Context, userID string, policy *domain.Policy) (*domain.Card, error) {
	// 1. 정책을 JSON으로 직렬화합니다.
	policyBytes, err := json.Marshal(policy)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal policy: %w", err)
	}

	// 2. 정책 데이터의 keccak256 해시를 계산합니다.
	policyHash := crypto.Keccak256Hash(policyBytes)
	policyHashStr := policyHash.Hex()

	// 3. 정책 원본을 오프체인 DB에 저장합니다. (MVP에서는 인메모리)
	if err := s.policyRepo.Save(ctx, userID, policy); err != nil {
		return nil, fmt.Errorf("failed to save policy off-chain: %w", err)
	}

	// 4. 정책 해시를 온체인에 기록합니다. (MVP에서는 모의 리포지토리 사용)
	// 실제로는 userId로부터 userAddress를 조회해야 합니다. MVP에서는 userId를 그대로 사용합니다.
	txHash, err := s.cardRepo.SavePolicyOnChain(ctx, userID, policyHashStr)
	if err != nil {
		return nil, fmt.Errorf("failed to save policy on-chain: %w", err)
	}

	// 5. 발급된 카드 정보를 생성합니다.
	newCard := &domain.Card{
		CardID:        fmt.Sprintf("FLEX-%d-xxxx-xxxx", time.Now().UnixMilli()%10000),
		UserID:        userID,
		PolicyHash:    policyHashStr,
		OnchainTxHash: txHash,
	}

	return newCard, nil
}
