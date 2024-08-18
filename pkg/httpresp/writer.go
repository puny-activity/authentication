package httpresp

import (
	"encoding/json"
	"errors"
	"github.com/puny-activity/authentication/pkg/base/headerbase"
	"net/http"
)

type Writer struct {
}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Write(writer http.ResponseWriter, statusCode int, payload any) error {
	writer.Header().Set(headerbase.ContentType, "application/json")

	writer.WriteHeader(statusCode)

	if payload != nil {
		err := json.NewEncoder(writer).Encode(payload)
		if err != nil {
			return errors.New("failed to encode payload")
		}
	}

	return nil
}
