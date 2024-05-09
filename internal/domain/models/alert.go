package models

import "time"

type Alert struct {
	SensorID  string    `json:"sensor_id"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	Message   string    `json:"message"`
}

type AlertType string

const (
	AlertTypeHighVibration   AlertType = "high_vibration"
	AlertTypeMediumVibration AlertType = "medium_vibration"
	AlertTypeLowTemperature  AlertType = "low_temperature"
)
