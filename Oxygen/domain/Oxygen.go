package domain

//DTO for oxygen data

type SaveOxygenModel struct {
	User_id     int     `json:"user_id"`
	Measurement float64 `json:"measurement"`
	Device_id   int     `json:"device_id"`
}

type GetOxygenModel struct {
	Oxygen_id   int     `json:"oxygen_id"`
	User_id     int     `json:"user_id"`
	Measurement float64 `json:"measurement"`
	Recorded_at string  `json:"recorded_at"`
	Device_id   int     `json:"device_id"`
}

type GetOxygenByUserModel struct {
	Oxygen_id    int     `json:"oxygen_id"`
	User_id      int     `json:"user_id"`
	Measurement  float64 `json:"measurement"`
	Recorded_at  string  `json:"recorded_at"`
	Name_user    string  `json:"name_user"`
	Surname_user string  `json:"surname_user"`
	Email_user   string  `json:"email_user"`
	Premium_user bool    `json:"premium_user"`
	Device_id    int     `json:"device_id"`
}
