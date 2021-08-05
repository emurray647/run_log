package db

type User struct {
	ID uint
}

type Activity struct {
	ID   uint
	User User

	StartTime    uint    `json:"start_time"` // seconds since epoch
	TotalTime    float32 `json:"total_time"`
	Distance     float32 `json:"distance"` // meters
	AvgHeartRate uint    `json:"avg_heart_rate"`
	AvgCadence   uint    `json:"avg_cadence"`

	DataGlob []byte
}
