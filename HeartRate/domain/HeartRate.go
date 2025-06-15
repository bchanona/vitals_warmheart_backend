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
	Heart_rate_id int     `json:"heart_rate_id"`
	User_id       int     `json:"user_id"`
	Measurement   float64 `json:"measurement"`
	Recorded_at   string  `json:"recorded_at"`
	Device_id     int     `json:"device_id"`
}

//model to get heart rate data by user
type GetHeartRateByUserModel struct {
	Heart_rate_id int     `json:"heart_rate_id"`
	User_id       int     `json:"user_id"`
	Measurement   float64 `json:"measurement"`
	Recorded_at   string  `json:"recorded_at"`
	Name_user     string  `json:"name_user"`
	Surname_user  string  `json:"surname_user"`
	Email_user    string  `json:"email_user"`
	Premium_user  bool    `json:"premium_user"`
	Device_id     int     `json:"device_id"`
}
