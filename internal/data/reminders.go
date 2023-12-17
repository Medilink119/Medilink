package data

import (
	"database/sql"
	"time"
)

type Reminder struct {
	Id          int
	UserId      int
	Appointment time.Time
	Email       string
	Note        string
	Sent        bool
}

type ReminderModel struct {
	DB *sql.DB
}

func (rm *ReminderModel) Insert(reminder *Reminder) error {
	query := `INSERT into reminders (user_id, appointment, email, note, send) VALUES ($1, $2, $3, $4, $5) RETURNING reminders_id`
	args := []interface{}{reminder.UserId, reminder.Appointment, reminder.Email, reminder.Note, false}
	err := rm.DB.QueryRow(query, args...).Scan(&reminder.Id)
	return err
}

func (rm *ReminderModel) GetRemindersByUser(userID int) ([]Reminder, error) {
	var reminders []Reminder
	query := `SELECT * from reminders WHERE user_id = ($1)`
	rows, err := rm.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Reminder
		rows.Scan(
			&r.Id,
			&r.UserId,
			&r.Appointment,
			&r.Email,
			&r.Note,
			&r.Sent,
		)
		reminders = append(reminders, r)
	}

	return reminders, nil
}

func (rm *ReminderModel) GetAllUnsentReminders() ([]Reminder, error) {
	var reminders []Reminder

	query := `SELECT * from reminders where send = ($1)`
	rows, err := rm.DB.Query(query, false)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r Reminder
		rows.Scan(
			&r.Id,
			&r.UserId,
			&r.Appointment,
			&r.Email,
			&r.Note,
			&r.Sent,
		)
		reminders = append(reminders, r)
	}

	return reminders, nil
}

func (rm *ReminderModel) UpdateReminderSent(id int) error {
	query := `UPDATE reminders set send = TRUE where reminders_id = ($1)`
	_, err := rm.DB.Exec(query, id)
	return err
}

func NewReminder(appointment time.Time, email, note string, userId int) *Reminder {
	return &Reminder{
		UserId:      userId,
		Appointment: appointment,
		Email:       email,
		Note:        note,
	}
}
