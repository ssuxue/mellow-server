package model

import "database/sql"

type UserModel struct {
	ID      int            `json:"id" form:"id"`
	Email    string         `json:"email" form:"email" binding:"email"`
	Password string         `json:"password" form:"password" `
	Avatar   sql.NullString `json:"avatar" form:"avatar"`
}
