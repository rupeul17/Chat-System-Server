package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func proc_chg_admin_passwd() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/member_db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	fmt.Printf("Entered New admin Passwd : ")
	admin_pwd := input_string()

	result, err := db.Exec("UPDATE member SET PASSWD=? where ID=?", admin_pwd, "admin")
	if err != nil {
		log.Println(err.Error())
		return
	}

	nRow, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("update count : ", nRow)
}

func proc_crte_user_account() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/member_db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	fmt.Printf("Entered New User ID : ")
	user_id := input_string()

	fmt.Printf("Enter New User Passwd : ")
	user_pwd := input_string()

	result, err := db.Exec("INSERT INTO member value (?,?)", user_id, user_pwd)
	if err != nil {
		log.Println(err.Error())
		return
	}

	nRow, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("insert count : ", nRow)
}

func proc_chg_user_account() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/member_db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	fmt.Printf("Entered User ID : ")
	user_id := input_string()

	fmt.Printf("Enter New User Passwd : ")
	user_pwd := input_string()

	result, err := db.Exec("UPDATE member SET PASSWD=? where ID=?", user_pwd, user_id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	nRow, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("update count : ", nRow)
}

func proc_del_user_account() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/member_db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	fmt.Printf("Entered Deleted User ID : ")
	user_id := input_string()

	result, err := db.Exec("delete from member where ID=?", user_id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	nRow, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("delete count : ", nRow)
}

func proc_chg_max_sess_cnt() {

	ClearTerminal()
	fmt.Print("Entered Session Cnt: ")

	session_cnt := input_number()

	fmt.Printf("Max Session Count Changed : (%d) -> (%d)\n", Max_Session_Cnt, session_cnt)
	Max_Session_Cnt = session_cnt
}

func proc_chg_conf() {
	ClearTerminal()
	fmt.Println("==========Chat System Command==========")
	fmt.Println(" 1. Change Max Session Count.")
	fmt.Println(" 2. Change Admin Passwd.")
	fmt.Println(" 3. Create User Account.")
	fmt.Println(" 4. Change User Account.")
	fmt.Println(" 5. Delete User Account.")
	fmt.Println("=======================================")

	cmd_num := input_number()

	switch cmd_num {
	case 1:
		proc_chg_max_sess_cnt()
	case 2:
		proc_chg_admin_passwd()
	case 3:
		proc_crte_user_account()
	case 4:
		proc_chg_user_account()
	case 5:
		proc_del_user_account()
	default:

	}
}
