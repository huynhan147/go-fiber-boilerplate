package transaction

// TransactionResponse nhận về sau khi tạo giao dịch thành công
type TransactionResponse struct {
	TransactionID string  `json:"transaction_id"`
	Status        string  `json:"status"`   // pending, success, failed
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	CreatedAt     string  `json:"created_at"`
}

// TransactionDetailResponse nhận về khi query chi tiết giao dịch
type TransactionDetailResponse struct {
	TransactionID string  `json:"transaction_id"`
	Status        string  `json:"status"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	ReferenceID   string  `json:"reference_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// RefundResponse nhận về sau khi hoàn tiền
type RefundResponse struct {
	RefundID      string  `json:"refund_id"`
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"` // pending, completed
	CreatedAt     string  `json:"created_at"`
}

// apiResponse wrapper chung nếu Project A trả về dạng { success, data, message }
type apiResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}
