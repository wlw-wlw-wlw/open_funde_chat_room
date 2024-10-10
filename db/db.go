package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var (
	dsn string

	stmt_usr_insert *sql.Stmt
	stmt_usr_query  *sql.Stmt
	pre_usr_insert  string
	pre_usr_query   string

	stmt_msg_insert *sql.Stmt
	stmt_msg_query  *sql.Stmt
	pre_msg_query   string
	pre_msg_insert  string

	stmt_ai_chat_history_insert *sql.Stmt
	stmt_ai_chat_history_query  *sql.Stmt
	pre_ai_chat_history_query   string
	pre_ai_chat_history_insert  string
)

type User struct {
	Name   string
	Passwd string
}
type Msg struct {
	Name string `json:"username"`
	Msg  string `json:"message"`
	Time string `json:"time"`
}

func init() {
	dsn = "root:W526452wlw@tcp(localhost:3306)/chat_room"
	pre_usr_insert = "insert into user_info(uname,passwd) values(?,?)"
	pre_usr_query = "select * from user_info where uname=?"

	pre_msg_insert = "insert into msg_table(uname,create_time,msg) values(?,?,?)"
	pre_msg_query = "select * from msg_table order by create_time"

	pre_ai_chat_history_insert = "insert into ai_chat_history(uname,speaker,msg,chat_time) values(?,?,?,?)"
	pre_ai_chat_history_query = "select speaker,msg,chat_time from ai_chat_history where uname=? order by chat_time"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	stmt_usr_insert, err = db.Prepare(pre_usr_insert)
	if err != nil {
		panic(err)
	}
	stmt_usr_query, err = db.Prepare(pre_usr_query)
	if err != nil {
		panic(err)
	}
	stmt_msg_insert, err = db.Prepare(pre_msg_insert)
	if err != nil {
		panic(err)
	}
	stmt_msg_query, err = db.Prepare(pre_msg_query)
	if err != nil {
		panic(err)
	}
	stmt_ai_chat_history_insert, err = db.Prepare(pre_ai_chat_history_insert)
	if err != nil {
		panic(err)
	}
	stmt_ai_chat_history_query, err = db.Prepare(pre_ai_chat_history_query)
	if err != nil {
		panic(err)
	}
}
func InsertUser(username string, passwd string) bool {
	_, err := stmt_usr_insert.Exec(username, passwd)
	return err == nil
}
func Usersearch(username string) bool {
	rows, err := stmt_usr_query.Query(username)
	if err != nil {
		return false
	}
	defer rows.Close()
	var check int = 0
	for rows.Next() {
		check++
	}
	return check != 0
}
func UserPasswdMatch(username string, passwd string) bool {
	rows, err := stmt_usr_query.Query(username)
	if err != nil {
		return false
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Name, &u.Passwd)
		if err != nil {
			return false
		}
		if u.Name == username && u.Passwd == passwd {
			return true
		}
	}
	return false
}

func Insert_msg(m Msg) {
	_, err := stmt_msg_insert.Exec(m.Name, m.Time, m.Msg)
	if err != nil {
		panic(err)
	}
}
func Query_msg() []Msg {
	rows, err := stmt_msg_query.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var msgs []Msg
	for rows.Next() {
		var m Msg
		err := rows.Scan(&m.Name, &m.Time, &m.Msg)
		if err != nil {
			panic(err)
		}
		msgs = append(msgs, m)
	}
	return msgs
}

func Insert_ai_chat(uname string, speaker string, msg string, chat_time string) {
	_, err := stmt_ai_chat_history_insert.Exec(uname, speaker, msg, chat_time)
	if err != nil {
		panic(err)
	}
}
func Query_ai_chat_history(uname string) []Msg {
	rows, err := stmt_ai_chat_history_query.Query(uname)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var msgs []Msg
	for rows.Next() {
		var m Msg
		err := rows.Scan(&m.Name, &m.Msg, &m.Time)
		if err != nil {
			panic(err)
		}
		msgs = append(msgs, m)
	}
	return msgs	
}
