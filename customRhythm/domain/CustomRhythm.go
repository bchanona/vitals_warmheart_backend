package domain

type CustomRhythmModel struct {
	UserId int `json:"userId"`
	LowBpm int `json:"LowBpm"`
	HighBpm int `json:"HighBpm"`
}