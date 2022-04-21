package http

import (
	"encoding/json"
	"net/http"

	"github.com/edfoh/go-api-starter/pkg/log"
)

type clientError interface {
	ErrorDetails() (int, string)
}

type serverHandler struct {
	operation string
	w         http.ResponseWriter
	r         *http.Request
	logger    log.Logger
	err       error
}

func newServerHandler(operation string, w http.ResponseWriter, r *http.Request, logger log.Logger) *serverHandler {
	return &serverHandler{
		operation: operation,
		w:         w,
		r:         r,
		logger:    logger,
	}
}

func (s *serverHandler) handleServiceError(err error) {
	if s.err != nil || err == nil {
		return
	}

	if clientErr, ok := err.(clientError); ok {
		code, msg := clientErr.ErrorDetails()
		s.clientError(err, code, msg)
		return
	}

	s.unhandledError(err)
}

func (s *serverHandler) clientError(err error, status int, errMsgs ...string) {
	if s.err != nil {
		return
	}

	defer func() { s.err = err }()
	s.logger.Warnf("%s: %v", s.operation, err.Error())

	errs := []*errorPayload{}
	for _, errMsg := range errMsgs {
		errs = append(errs, &errorPayload{errMsg})
	}
	errRes := &errorResponse{
		Errors: errs,
	}

	s.respondJSON(status, errRes)
}

func (s *serverHandler) unhandledError(err error) {
	if s.err != nil {
		return
	}

	defer func() { s.err = err }()
	s.logger.Errorf("%s: %v", s.operation, err.Error())
	http.Error(s.w, "oops, something went wrong!", http.StatusInternalServerError)
}

func (s *serverHandler) respondJSON(status int, data interface{}) {
	if s.err != nil {
		return
	}

	s.w.WriteHeader(status)
	s.w.Header().Set("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(s.w).Encode(data)
		if err != nil {
			s.unhandledError(err)
			return
		}
	}
}

func (s *serverHandler) decode(w http.ResponseWriter, r *http.Request, v interface{}) {
	if s.err != nil {
		return
	}

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		s.clientError(err, http.StatusBadRequest, "failed to decode request body")
	}
}
