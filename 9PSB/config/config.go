package config

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
)

type AppTools struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	Validate    *validator.Validate
}

func NewAppTools() *AppTools {
	return &AppTools{
		log.New(os.Stdout, "[ Error ]", log.LstdFlags|log.Lshortfile),
		log.New(os.Stdout, "[ info ]", log.LstdFlags|log.Lshortfile),
		validator.New(),
	}
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
