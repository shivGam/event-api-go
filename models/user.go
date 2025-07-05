package models

import (
	"time"
	"github.com/shivGam/event-api-go/db"
	"github.com/shivGam/event-api-go/utils"
)

type User struct{
	ID int64 `json:"user_id"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Save() error{
	createUserquery:=`INSERT INTO users (email,password) VALUES (?,?)`
	hashedPass,err:= utils.HashPassWord(u.Password)
	if err!=nil{
		return err
	}
	stmt, err:=db.DB.Prepare(createUserquery)
	if err!=nil {
		return err
	}
	defer stmt.Close()
	result , err:= stmt.Exec(u.Email,hashedPass)
	if err!=nil{
		return err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		return err
	}
	u.ID = id
	u.Password = hashedPass
	return nil
}

func (u *User) ValidateCredentials() error{
	selectQuery:=`SELECT id,password FROM users WHERE email=?`
	row:=db.DB.QueryRow(selectQuery,u.Email)
	var hashedPass string
	err:=row.Scan(&u.ID,&hashedPass)//ajeeb hai ye
	if err!=nil{
		return err
	}
	isValid:=utils.ComparePasswords(hashedPass,u.Password)
	if !isValid {
		return err
	}
	return nil
}
