package user

import (
	// "database/sql"
	"time"
)

const (
	TABLE_USERS = "users"
)

type User struct {
	Id        int        `db:"id" json:"id" insert:"-" update:"-"`
	Email     string     `db:"email" json:"email" insert:"email" update:"email"`
	Username  string     `db:"username" json:"username" insert:"username" update:"username"`
	Password  string     `db:"password" json:"password" insert:"password" update:"password"`
	CreatedAt time.Time  `db:"created_at" json:"created_at" insert:"-" update:"-"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at" insert:"-" update:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at" insert:"-" update:"deleted_at"`
}
