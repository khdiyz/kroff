package service

import (
	"context"
	"io"
	"kroff/config"
	"kroff/pkg/storage"
	"kroff/utils/response"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type fileService struct {
	storage *storage.Storage
	cfg     *config.Config
}

func NewFileService(storage *storage.Storage, cfg *config.Config) *fileService {
	return &fileService{
		storage: storage,
		cfg:     cfg,
	}
}

func (s *fileService) UploadFile(ctx context.Context, file io.Reader, fileSize int64, contentType string) (string, error) {
	// Generate a unique object name for the file
	objectName := uuid.NewString()

	// Upload the file directly using the storage service
	err := s.storage.UploadFile(ctx, objectName, file, fileSize, contentType)
	if err != nil {
		return "", response.ServiceError(err, codes.Internal)
	}

	// Return the object name (file ID)
	return objectName, nil
}

func (s *fileService) UploadWithName(ctx context.Context, file io.Reader, fileSize int64, contentType string, fileName string) error {
	// Upload the file directly using the storage service
	err := s.storage.UploadFile(ctx, fileName, file, fileSize, contentType)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}

	return nil
}
