package logwriter

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DatabaseWriter struct {
	Name 			string
	db            	*sql.DB

}

func NewDatabaseWriter(name string, host string, port int, user string, password string) (*DatabaseWriter, error){
	dbw := DatabaseWriter{ Name:name }
	err := dbw.connect(host, port, user, password)
	if err != nil {
		return nil,err
	}
	return &dbw,nil
}

func (dbw DatabaseWriter)Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	n = len(p)
	return
}

func (dbw *DatabaseWriter)connect( host string, port int, user string, password string) error {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbw.Name)
	res, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dbw.db = res
	return nil
}
