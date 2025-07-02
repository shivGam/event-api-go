package models


import (
	"time"
	"github.com/shivGam/event-api-go/db"
)

type Event struct {
	ID int64 `json:"id"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location string `json:"location" binding:"required"`
	UserID int `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

var events []Event

func (e *Event) Save() error {
	insertQuery := `INSERT INTO events(title,description,location,user_id,created_at) VALUES(?,?,?,?,?)`
	stmt,err:= db.DB.Prepare(insertQuery)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	result,err:=stmt.Exec(e.Title,e.Description,e.Location,e.UserID,e.CreatedAt)
	if err!=nil{
		return err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		return err
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event,error){
	selectQuery:= `SELECT id,title,description,location,user_id,created_at FROM events ORDER BY created_at DESC`
	stmt,err:=db.DB.Prepare(selectQuery)
	if err!=nil {
		return []Event{},err
	}
	defer stmt.Close()
	rows,err:=stmt.Query()
	if err!=nil{
		return []Event{},err
	}
	defer rows.Close()
	for rows.Next(){
		var e Event
		err:=rows.Scan(&e.ID,&e.Title,&e.Description,&e.Location,&e.UserID,&e.CreatedAt)
		if err!=nil{
			return []Event{},err
		}
		events = append(events,e)
	}
	return events,nil
}