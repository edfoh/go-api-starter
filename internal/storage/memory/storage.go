package memory

import (
	"context"

	"github.com/edfoh/go-api-starter/internal/adding"
	"github.com/edfoh/go-api-starter/pkg/log"
)

type Storage struct {
	logger  log.Logger
	tickets []*ticket
}

func NewStorage(logger log.Logger) *Storage {
	return &Storage{
		logger: logger,
	}
}

func (r *Storage) FindTicket(ctx context.Context, subject, description string) (string, *adding.Ticket, error) {
	r.logger.Debug("Entered memory.Storage: FindTicket", log.Field("subject", subject), log.Field("description", description))

	for _, ticket := range r.tickets {
		if ticket.Subject == subject && ticket.Description == description {
			return ticket.ID, &adding.Ticket{
				Subject:       ticket.Subject,
				Description:   ticket.Description,
				RequesterName: ticket.RequesterName,
				CreateDate:    ticket.CreateDate,
			}, nil
		}
	}
	return "", nil, nil
}

func (r *Storage) CreateTicket(ctx context.Context, tic *adding.Ticket) (string, error) {
	r.logger.Debug("Entered memory.Storage: CreateTicket", log.Field("ticket", tic))

	t := &ticket{
		Subject:       tic.Subject,
		Description:   tic.Description,
		RequesterName: tic.RequesterName,
		CreateDate:    tic.CreateDate,
	}
	r.tickets = append(r.tickets, t)
	return "", nil
}
