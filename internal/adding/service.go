package adding

import (
	"context"
	"fmt"

	"github.com/edfoh/go-api-starter/pkg/httperrors"
	"github.com/edfoh/go-api-starter/pkg/log"
)

type Service interface {
	CreateTicket(context.Context, *Ticket) (string, error)
}

type Repository interface {
	FindTicket(context.Context, string, string) (string, *Ticket, error)
	CreateTicket(context.Context, *Ticket) (string, error)
}

type service struct {
	r      Repository
	logger log.Logger
}

func NewService(r Repository, logger log.Logger) Service {
	return &service{r, logger}
}

func (s *service) CreateTicket(ctx context.Context, ticket *Ticket) (string, error) {
	s.logger.Debug("Entered adding.Service: CreateTicket", log.Field("ticket", ticket))

	ticketID, _, err := s.r.FindTicket(ctx, ticket.Subject, ticket.Description)
	if err != nil {
		return "", fmt.Errorf("CreateTicket: %w", err)
	}

	if ticketID != "" {
		return "", httperrors.BadRequest.Newf("ticket with subject '%s' and description '%s already exist", ticket.Subject, ticket.Description)
	}

	return s.r.CreateTicket(ctx, ticket)
}
