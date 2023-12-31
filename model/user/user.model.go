package user

import "time"

type User struct {
	id         int64
	username   string
	email      string
	password   string
	bio        string
	is_deleted bool
	is_active  bool
	created_at time.Time
	updated_at *time.Time
}
