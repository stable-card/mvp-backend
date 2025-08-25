package domain

// Card는 발급된 사용자의 카드를 나타냅니다.
type Card struct {
	CardID        string `json:"cardId"`
	UserID        string `json:"userId"`
	PolicyHash    string `json:"policyHash"`
	OnchainTxHash string `json:"onchainTxHash"`
}
