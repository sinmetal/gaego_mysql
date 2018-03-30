package main

import (
	"database/sql"
	"net"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/appengine"
	"google.golang.org/appengine/socket"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	dial := func(addr string) (net.Conn, error) {
		return socket.Dial(appengine.NewContext(r), "tcp", addr)
	}

	mysql.RegisterDial("external", dial)

	db, err := sql.Open("mysql", "root:d2QY.usK@external(35.188.168.168:3306)/db1")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtSec, err := db.Prepare("SELECT * FROM sample LIMIT 10")
	if err != nil {
		panic(err.Error())
	}
	defer stmtSec.Close()

	rows, err := stmtSec.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	w.Write([]byte("DONE."))
}
