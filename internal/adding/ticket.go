package adding

import "time"

type Ticket struct {
	Subject       string    `json:"subject"`
	RequesterName string    `json:"requester_name"`
	Description   string    `json:"description"`
	CreateDate    time.Time `json:"create_date"`
}
