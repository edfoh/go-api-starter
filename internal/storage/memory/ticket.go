package memory

import "time"

type ticket struct {
	ID            string
	Subject       string
	RequesterName string
	Description   string
	CreateDate    time.Time
}
