package database

import (
	"database/sql"
	"fmt"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/config"
	_ "github.com/lib/pq"
)

//HourlyDB main structure
type Postgres struct {
	db  *sql.DB
}

//NewPostgres create and initialize database and return it or error
func NewPostgres(config config.DatabaseConfig) (pg *Postgres, err error){
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	pg = &Postgres{
		db:db,
	}
	return
}

//Dispose and disconnect
func (pg *Postgres) Dispose() {
	pg.db.Close()
	pg.db = nil
}

//PushPeriod write a list of temperatures in to DB
func (pg *Postgres) PushPeriod(name string, temperatures []hourlysvc.Temperature) error {

	query := fmt.Sprintf("INSERT INTO %s " +
		"(date, temperature) VALUES", name)

	length := len(temperatures)
	for i, v := range temperatures {
		query += fmt.Sprintf(
			" ( '%s', %g)",
			v.Date, v.Temperature)
		if i < length-1 {
			query += ","
		}
	}

	query += ` ON CONFLICT (date) DO UPDATE SET
			 temperature = excluded.temperature;`

	return writeToDB(pg.db, query)
}


//GetPeriod get a list of temperatures form table @name (station Id)
func (pg *Postgres) GetPeriod(name string, start string, end string) (temps []hourlysvc.Temperature, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE date >= '%s' AND date < '%s' ORDER BY date::timestamp ASC;",
		name, start, end)

	rows, err := pg.db.Query(query)
	if err != nil {
		return temps,err
	}
	defer rows.Close()


	for rows.Next() {
		st,err := parseRow(rows)
		if err != nil {
			return temps, err
		}
		temps = append(temps, st)
	}
	return temps,err
}


//CreateTable create a table with name @icao + tPrefix if not exist
func (pg *Postgres) CreateTable(name string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ("+
		"	date timestamp UNIQUE,"+
		"	temperature real"+
		");",
		name)
	return writeToDB(pg.db, query)
}

//RemoveTable remove stations table from BD
func (pg *Postgres) RemoveTable(name string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;",
		name)
	return writeToDB(pg.db, query)
}

//GetUpdateDate ...
func (pg *Postgres) GetUpdateDate(name string) (date string, err error) {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY date::timestamp DESC LIMIT 1;",
		name)

	rows, err := pg.db.Query(query)
	if err != nil {
		return date,err
	}
	defer rows.Close()


	for rows.Next() {
		temp,err := parseRow(rows)
		if err != nil {
			return date, err
		}

		date = temp.Date

	}
	return date,err
}

//GetLatest return latest temperature data
//for station with name @name
func (pg *Postgres)GetLatest(name string) (temp hourlysvc.Temperature, err error) {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY date::timestamp DESC LIMIT 1;",
		name)

	rows, err := pg.db.Query(query)
	if err != nil {
		return temp,err
	}
	defer rows.Close()

	rows.Next()
	return parseRow(rows)
}



func parseRow(rows *sql.Rows) (row hourlysvc.Temperature, err error) {
	err = rows.Scan(
		&row.Date,
		&row.Temperature,
	)
	return
}

func writeToDB(db *sql.DB, query string) (err error){
	row, err := db.Query(query)
	if err != nil {
		return
	}
	row.Close()
	return
}



