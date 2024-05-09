package models

type Equipament struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Location string   `json:"location"`
	Sensors  []Sensor `json:"sensors"`
}

type Sensor struct {
	ID   string     `json:"id"`
	Type SensorType `json:"type"`
	Unit string     `json:"unit"`
}

type UnitType string

const (
	UnitMillimeterPerSecond UnitType = "mm/s" // Vibração
	UnitCelsius             UnitType = "°C"   // Temperatura
	UnitBar                 UnitType = "bar"  // Pressão
	UnitKilogram            UnitType = "kg"   // Massa
)

type SensorType string

const (
	SensorTypeVibration   SensorType = "vibration"
	SensorTypeTemperature SensorType = "temperature"
	SensorTypePressure    SensorType = "pressure"
)
