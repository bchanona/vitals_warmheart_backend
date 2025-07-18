package domain

type ICustomRepository interface {
	Save(customBpm CustomRhythmModel) error
	Update(userId int, updateBpm CustomRhythmModel) error
}