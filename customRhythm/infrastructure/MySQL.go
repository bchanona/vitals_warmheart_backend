package infrastructure

import (
	"database/sql"

	"github.com/bchanona/vitals_warmheart_backend/customRhythm/domain"
	"github.com/bchanona/vitals_warmheart_backend/customRhythm/queries"
)

type MySQL struct{
	db *sql.DB
}
func NewMySQL(db *sql.DB) *MySQL{
	return &MySQL{db: db}
}

func (sql *MySQL) Save(custom domain.CustomRhythmModel) error {
	_, err := sql.db.Exec(queries.SaveCustomQuery,custom.UserId, custom.LowBpm, custom.HighBpm)
	return  err
}
func (sql *MySQL) Update(userId int, custom domain.CustomRhythmModel) error{
	_, err := sql.db.Exec(queries.UpdateCustom,custom.LowBpm,custom.HighBpm,userId)
	if err != nil {
		return err
	}
	return nil
}

