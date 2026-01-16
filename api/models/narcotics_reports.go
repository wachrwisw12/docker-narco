package models

type NacorticsReport struct {
	ID            string `json:"id"`
	ContentDetail string `json:"details"`
	TrackingCode  string `json:"tracking_code"`
}
