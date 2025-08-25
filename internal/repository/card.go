package repository

import "context"

// CardRepository는 카드 발급과 관련된 온체인 작업을 정의합니다.
type CardRepository interface {
	// SavePolicyOnChain은 사용자의 정책 해시를 블록체인에 기록합니다.
	// 성공 시 트랜잭션 해시를 반환합니다.
	SavePolicyOnChain(ctx context.Context, userAddress string, policyHash string) (txHash string, err error)
}

// mockCardRepository는 실제 블록체인 연동 없이 더미 데이터를 반환하는 구현체입니다.
type mockCardRepository struct{}

// NewMockCardRepository는 새로운 mockCardRepository를 생성합니다.
func NewMockCardRepository() CardRepository {
	return &mockCardRepository{}
}

// SavePolicyOnChain은 더미 트랜잭션 해시를 반환합니다.
func (r *mockCardRepository) SavePolicyOnChain(ctx context.Context, userAddress string, policyHash string) (string, error) {
	// 실제로는 여기에 블록체인 트랜잭션 전송 로직이 들어갑니다.
	// (e.g., go-ethereum 라이브러리 사용)
	return "0x54321fedcba" + userAddress + policyHash, nil
}
