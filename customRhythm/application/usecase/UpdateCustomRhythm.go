package usecase

import "github.com/bchanona/vitals_warmheart_backend/customRhythm/domain"

type UpdateCustomRhytmUc struct {
	repo domain.ICustomRepository
}

func NewUpdateCustomRhytmUc(repo domain.ICustomRepository) *UpdateCustomRhytmUc {
	return &UpdateCustomRhytmUc{repo: repo}
}

func (uc *UpdateCustomRhytmUc) Execute(userId int, custom domain.CustomRhythmModel) error {
	return uc.repo.Update(userId, custom)
}
