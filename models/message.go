package models

import (
	"database/sql"
	"time"
)

type Message struct {
	Id       *string    `sql:id`
	DeviceId *string    `sql:device_id`
	Message  *string    `sql:message`
	MsgTime  *time.Time `sql:msg_time`
}

const (
	SaveMessageQuery = `INSERT INTO messages(id, device_id, message, msg_time) VALUES(gen_random_uuid(), $1, $2, $3)`
)

func (m *Message) SaveToDb(db *sql.DB) error {
	_, err := db.Exec(SaveMessageQuery, m.DeviceId, m.Message, m.MsgTime)
	return err
}
