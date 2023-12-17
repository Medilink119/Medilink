package data

import (
	"database/sql"
)

type Model struct {
	Um *UserModel
	Fm *FileModel
	Rm *ReminderModel
}

func NewModel(db *sql.DB) *Model {
	model := &Model{
		Um: &UserModel{DB: db},
		Fm: &FileModel{DB: db},
		Rm: &ReminderModel{DB: db},
	}

	return model
}
