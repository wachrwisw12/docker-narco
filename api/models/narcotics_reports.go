package models

type NacorticsReport struct {
	ID           string `json:"id"`
	Details      string `json:"details"`
	TrackingCode string `json:"tracking_code"`
}
