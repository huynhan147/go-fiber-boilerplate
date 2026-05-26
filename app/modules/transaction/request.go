package transaction

// CreateTransactionRequest gửi sang Project A để tạo giao dịch mới
type CreateTransactionRequest struct {
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`     // VND, USD...
	Description string  `json:"description"`
	ReferenceID string  `json:"reference_id"` // ID nội bộ để đối soát
}

// RefundTransactionRequest gửi sang Project A để hoàn tiền
type RefundTransactionRequest struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Reason        string  `json:"reason"`
}
