package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

// ----- Payload -----

type SendWelcomeEmailPayload struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

// ----- Enqueue (gọi từ service/handler) -----

func NewSendWelcomeEmailTask(payload SendWelcomeEmailPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	// asynq.MaxRetry(3): retry tối đa 3 lần nếu thất bại
	return asynq.NewTask(TypeSendWelcomeEmail, data, asynq.MaxRetry(3)), nil
}

// ----- Handler (worker xử lý) -----

type SendWelcomeEmailHandler struct {
	// inject dependencies ở đây: mailer, db, ...
}

func NewSendWelcomeEmailHandler() *SendWelcomeEmailHandler {
	return &SendWelcomeEmailHandler{}
}

func (h *SendWelcomeEmailHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload SendWelcomeEmailPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}

	// TODO: thay bằng logic gửi mail thật (SendGrid, SES, SMTP...)
	log.Printf("📧 Sending welcome email to %s <%s>", payload.Name, payload.Email)

	return nil
}
