package http

import (
	"net/http"

	"github.com/edfoh/go-api-starter/internal/adding"
	"github.com/edfoh/go-api-starter/pkg/log"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	logger       log.Logger
	router       *httprouter.Router
	serviceAdder adding.Service
}

func NewServer(
	logger log.Logger,
	r *httprouter.Router,
	serviceAdder adding.Service) http.Handler {
	s := &server{
		logger:       logger,
		router:       r,
		serviceAdder: serviceAdder,
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("server: serving http")
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.logger.Debug("server: registering routes")
	s.router.HandlerFunc(http.MethodPost, "/tickets", s.handleCreateTicket())
}

func (s *server) handleCreateTicket() http.HandlerFunc {
	type data struct {
		TicketID string `json:"ticket_id"`
	}
	type response struct {
		Data *data `json:"data"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		handler := newServerHandler("handleCreateTicket", w, r, s.logger)

		var newTicket *adding.Ticket
		handler.decode(w, r, newTicket)

		ticketID, err := s.serviceAdder.CreateTicket(r.Context(), newTicket)
		handler.handleServiceError(err)

		res := &response{
			Data: &data{
				TicketID: ticketID,
			},
		}
		handler.respondJSON(http.StatusCreated, res)
	}
}
