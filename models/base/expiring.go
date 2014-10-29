package base

import (
	"log"
	"time"
)

// Expiring is a base struct for resources that expires e.g. sessions.
type Expiring struct {
	// ExpDate is the RFC3339-encoded time when the resource will expire.
	ExpDate string `json:"exp_date" gorethink:"exp_date"`
}

// HasExpired returns true if the resource has expired (or if the ExpDate string is badly formatted)
func (e *Expiring) HasExpired() bool {
	t, err := time.Parse(time.RFC3339, e.ExpDate)
	if err != nil {
		log.Println("Bad format! The expiry date is not RFC3339-formatted.", err)
		return true
	}
	if time.Now().UTC().After(t) {
		return true
	}
	return false
}