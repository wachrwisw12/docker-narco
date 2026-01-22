package models

import "time"

type NacorticsReport struct {
	ID           int8   `json:"id"`
	TrackingCode string `json:"tracking_code"`
	// Category     string `json:"category"`
	Details   string     `json:"details"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
}
