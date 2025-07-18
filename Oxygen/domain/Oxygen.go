package domain

//DTO for oxygen data

type SaveOxygenModel struct {
	User_id     int     `json:"user_id"`
	Measurement float64 `json:"measurement"`
	Device_id   int     `json:"device_id"`
}

type GetOxygenModel struct {
	Measurement    float64 `json:"measurement"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
}

type GetOxygenByUserModel struct {
	Measurement    float64 `json:"measurement"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
	Name_user      string  `json:"name_user"`
	Surname_user   string  `json:"surname_user"`
	Email_user     string  `json:"email_user"`
}
