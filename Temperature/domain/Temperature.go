package domain

//DTO fot temperature data

type SaveTemperatureModel struct {
	User_id     int     `json:"user_id"`
	Measurement float64 `json:"measurement"`
	Device_id   int     `json:"device_id"`
}

type GetTemperatureModel struct {
	Temperature_id int     `json:"temperature_id"`
	User_id        int     `json:"user_id"`
	Measurement    float64 `json:"measurement"`
	Device_id      int     `json:"device_id"`
	Recorded_at    string  `json:"recorded_at"`
}

type GetTemperatureByUserModel struct {
	Temperature_id int     `json:"temperature_id"`
	User_id        int     `json:"user_id"`
	Measurement    float64 `json:"measurement"`
	Recorded_at    string  `json:"recorded_at"`
	Name_user      string  `json:"name_user"`
	Surname_user   string  `json:"surname_user"`
	Email_user     string  `json:"email_user"`
	Premium_user   bool    `json:"premium_user"`
	Device_id      int     `json:"device_id"`
}
