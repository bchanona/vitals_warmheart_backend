package domain

//DTO for heart rate data

//model to save heart rate data

type SaveHeartRateModel struct {
	User_id     int     `json:"user_id"`
	Measurement float64 `json:"measurement"`
	Device_id   int     `json:"device_id"`
}

//model to get heart rate data
type GetHeartRateModel struct {
	Measurement    float64 `json:"measurement"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
}

//model to get heart rate data by user
type GetHeartRateByUserModel struct {
	Measurement    float64 `json:"measurement"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
	Name_user      string  `json:"name_user"`
	Surname_user   string  `json:"surname_user"`
	Email_user     string  `json:"email_user"`
}
