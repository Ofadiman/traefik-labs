package request_id

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type RequestIdPlugin struct {
	logger *log.Logger
	name   string
	next   http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &RequestIdPlugin{
		logger: log.New(os.Stdout, "[RequestIdPlugin] ", 0),
		name:   name,
		next:   next,
	}, nil
}

func (p *RequestIdPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Request-Id") == "" {
		id := uuid.New().String()
		p.logger.Printf("X-Request-Id header is missing, attaching %s id to the request.", id)
		req.Header.Set("X-Request-Id", id)
	}
	p.next.ServeHTTP(rw, req)
}
