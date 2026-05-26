package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type ProcessAvatarPayload struct {
	UserID   uint   `json:"user_id"`
	FilePath string `json:"file_path"`
}

func NewProcessAvatarTask(payload ProcessAvatarPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeProcessAvatar, data, asynq.MaxRetry(2)), nil
}

type ProcessAvatarHandler struct{}

func NewProcessAvatarHandler() *ProcessAvatarHandler {
	return &ProcessAvatarHandler{}
}

func (h *ProcessAvatarHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload ProcessAvatarPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}

	// TODO: resize, compress, upload to S3...
	log.Printf("🖼️  Processing avatar for user %d: %s", payload.UserID, payload.FilePath)

	return nil
}
