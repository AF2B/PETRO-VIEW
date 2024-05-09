package models

import "time"

type SensorReading struct {
	SensorID  string           `json:"sensor_id"` // ID of the sensor that generated the reading
	Timestamp time.Time        `json:"timestamp"` // Time the reading was taken
	Value     float64          `json:"value"`     // Numerical value of the reading
	Status    SensorStatusType `json:"status"`    // Status of the reading (e.g., "normal", "alert", "error")
}

type SensorStatusType string

const (
	StatusNormal SensorStatusType = "normal"
	StatusAlert  SensorStatusType = "alert"
	StatusError  SensorStatusType = "error"
)
