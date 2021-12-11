package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bariasabda/test-flash-coffee/config"
)

type AWS struct {
	cfg     *config.Config
	session *session.Session
}

func New(cfg *config.Config, session *session.Session) Interface {
	return AWS{cfg, session}
}

type Interface interface {
	AddFileToS3(file []byte, fileName, contentType string, size int64) error
	GetUrl(fileName string) string
}
