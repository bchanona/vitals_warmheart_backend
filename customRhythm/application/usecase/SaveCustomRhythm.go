package usecase

import "github.com/bchanona/vitals_warmheart_backend/customRhythm/domain"

type SaveCustomRhytmUc struct {
	repo domain.ICustomRepository
}

func NewSaveCustomRhytmUc(repo domain.ICustomRepository)*SaveCustomRhytmUc{
	return &SaveCustomRhytmUc{repo: repo}
}

func (uc *SaveCustomRhytmUc) Execute(custom domain.CustomRhythmModel) error{
	return  uc.repo.Save(custom)
}