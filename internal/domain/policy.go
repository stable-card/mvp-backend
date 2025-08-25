package domain

// Rule은 단일 혜택 규칙을 정의합니다.
type Rule struct {
	Category    string `json:"category"`
	BenefitType string `json:"benefitType"`
	Value       int    `json:"value"` // 베이시스 포인트(bp), 1bp = 0.01%
}

// Policy는 사용자가 설정한 혜택 정책의 집합입니다.
type Policy struct {
	Version int    `json:"version"`
	Rules   []Rule `json:"rules"`
}
