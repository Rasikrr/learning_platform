package files

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"os"
	"strings"
)

const (
	avatarDir = "./files/avatars"
	courseDir = "./files/courses"
	faqDir    = "./files/faq"
)

var (
	fileTypeMap = map[enum.FileType]string{
		enum.FileTypeAvatar: avatarDir,
		enum.FileTypeCourse: courseDir,
		enum.FileTypeFaq:    faqDir,
	}
)

type Service interface {
	DecodeAndSaveFile(ctx context.Context, base64Str string, fileType enum.FileType) (string, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) DecodeAndSaveFile(_ context.Context, base64Str string, fileType enum.FileType) (string, error) {
	base64Data := base64Str[strings.IndexByte(base64Str, ',')+1:]

	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}
	if !filetype.IsImage(data) {
		return "", fmt.Errorf("unknown file type")
	}

	kind, err := filetype.Match(data)
	if err != nil || kind == filetype.Unknown {
		return "", fmt.Errorf("unknown file type")
	}
	filePath := fmt.Sprintf("%s/%s.%s", fileTypeMap[fileType], uuid.New().String(), kind.Extension)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
