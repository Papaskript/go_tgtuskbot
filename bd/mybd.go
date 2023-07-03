package bd

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "12345"
	Dbname   = "postgres"
)

var SqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

type DbWrapper struct {
	Md *sql.DB
}
type Task struct {
	Id   int
	Task string
}

var Testdb DbWrapper

func ConnectDB() (DbWrapper, error) {
	var err error
	Db, err := sql.Open("postgres", SqlInfo)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	Testdb.Md = Db

	return Testdb, err

}

func InsertDB(t DbWrapper, us_id int64, st string) {
	var newT = t.Md
	ins := `INSERT INTO tusks (user_id,note) VALUES ($1,$2)`
	d, err := newT.Exec(ins, us_id, st)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data added successfully", d)
}

func ShowNote(s DbWrapper, id int64) []Task {
	var res string
	var idTask int
	var allTusks []Task
	var newS = s.Md
	ins, err := newS.Query(`SELECT id, note FROM tusks WHERE user_id = ($1)`, id)
	if err != nil {
		panic(err)
	}
	for ins.Next() {
		err = ins.Scan(&idTask, &res)
		currTask := Task{Id: idTask, Task: res}
		allTusks = append(allTusks, currTask)
	}
	return allTusks
}
func RemoveNote(s DbWrapper, userId int, num int) {
	newQ := s.Md
	ins, err := newQ.Exec(`DELETE FROM tusks WHERE user_id = ($1) and id = ($2)`, userId, num)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted", ins)
}
func UpdateNote(s DbWrapper, us_id int64, id int64, st string) {
	var upNote = s.Md
	ins, err := upNote.Exec(`UPDATE tusks SET note = ($1) WHERE id = ($2) and user_id = ($3)`, st, id, us_id)
	if err != nil {
		log.Fatal(err, "Couldn't update")

	}
	fmt.Println("Updated", ins)
}
