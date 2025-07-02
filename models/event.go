package models


import (
	"errors"
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

func GetEventById(id int64) (Event,error){
	selectQuery:= `SELECT * FROM events WHERE id=?`
	row,err:=db.DB.Query(selectQuery,id)
	if err!=nil{
		return Event{},err
	}
	defer row.Close()
	var e Event
	if !row.Next(){
		return Event{},errors.New("event not found")
	}
	err=row.Scan(&e.ID,&e.Title,&e.Description,&e.Location,&e.UserID,&e.CreatedAt)
	if err!=nil{
		return Event{},err
	}
	return e,nil
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

func (e *Event) UpdateEvent() (Event,error){
	updateQuery:= `UPDATE events
	SET title=?,description=?,location=?,user_id=?
	WHERE id=?`
	stmt,err:=db.DB.Prepare(updateQuery)
	if err!=nil{
		return Event{},err
	}
	defer stmt.Close()
	row,err:=stmt.Exec(e.Title,e.Description,e.Location,e.UserID,e.ID)
	if err!=nil{
		return Event{},err
	}
	rowsAffected,err:=row.RowsAffected()
	if err!=nil{
		return Event{},err
	}
	if rowsAffected==0{
		return Event{},errors.New("event not found")
	}
	return *e,nil
}

func DeleteEvent(id int64) error {
	deleteQuery:=`DELETE FROM events WHERE id=?`
	stmt,err:=db.DB.Prepare(deleteQuery)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_ ,err = stmt.Exec(id)
	if err!=nil{
		return err
	}
	return nil
}