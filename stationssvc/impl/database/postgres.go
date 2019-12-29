package database

import (
	"database/sql"
	"fmt"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/stationssvc/config"
	_ "github.com/lib/pq"
)

//Postgres database
type Postgres struct {
	db       	*sql.DB
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


//AddStation write single line of temperature in to DB
func (pg Postgres) AddStation(station stationssvc.Station) error {
	query := fmt.Sprintf("INSERT INTO stations "+
		"(id, name, timezone) "+
		"VALUES ( '%s', '%s', '%s') ",
		station.ID, station.Name, station.Timezone)

	return writeToDB(pg.db, query)
}

//AddStations write stations data into DB
func (pg Postgres) AddStations(stations []stationssvc.Station) error {
	query := fmt.Sprintf("INSERT INTO stations "+
		"(id, name, timezone) VALUES")

	length := len(stations)
	for i, v := range stations {
		query += fmt.Sprintf(
			"( '%s', '%s', '%s') ",
			v.ID, v.Name, v.Timezone)
		if i < length-1 {
			query += ","
		}
	}
	query += " ON CONFLICT (id) DO UPDATE SET " +
		"(name, timezone) = (excluded.name, excluded.timezone);"
	return writeToDB(pg.db, query)
}

//DeleteStation remove station by icao ID
func (pg Postgres) DeleteStation(id string) error {
	query := fmt.Sprintf("DELETE FROM stations "+
		" WHERE id = '%s'",
		id)
	return writeToDB(pg.db, query)
}

//GetStations get a list of station
func (pg Postgres) GetStations(ids []string) ([]stationssvc.Station,error) {
	sts := make([]stationssvc.Station, 0)
	query := "SELECT * FROM stations"
	query += " WHERE id IN ( "
	length := len(ids)
	for i, v := range ids {
		query += fmt.Sprintf("'%s'",v)
		if i < length-1 {
			query += ","
		}
	}
	query += " );"

	rows, err := pg.db.Query(query)
	if err != nil {
		return sts,err
	}
	defer rows.Close()


	for rows.Next() {
		st,err := parseRow(rows)
		if err != nil {
			return sts, err
		}
		sts = append(sts, st)
	}
	return sts,err
}

//GetAllStations get a list of station
func (pg Postgres) GetAllStations() ([]stationssvc.Station,error) {
	sts := make([]stationssvc.Station, 0)
	query := "SELECT * FROM stations;"

	rows, err := pg.db.Query(query)
	if err != nil {
		return sts, err
	}
	defer rows.Close()

	for rows.Next() {
		st,err := parseRow(rows)
		if err != nil {
			return sts, err
		}
		sts = append(sts, st)
	}
	return sts,err
}

//GetCount return number of stored stations
func (pg Postgres) GetCount() (int, error) {
	query := "SELECT COUNT(*) FROM stations;"
	rows, err := pg.db.Query(query)
	if err != nil {
		return 0, err
	}

	count := 0
	for rows.Next() {
		err = rows.Scan(
			&count,
		)
	}
	return count, err
}

//Dispose and disconnect
func (pg *Postgres) Dispose() {
	pg.db.Close()
	pg.db = nil
}

//CreateTable create a "Stations" table if not exist
func (pg Postgres) CreateTable() error {
	query := "CREATE TABLE IF NOT EXISTS stations ("+
		"	id char(4) UNIQUE,"+
		"	name varchar(30),"+
		"	timezone varchar(8)"+
		");"
	return writeToDB(pg.db, query)
}

//RemoveTable remove stations table from BD
func (pg *Postgres) RemoveTable() error {
	query := "DROP TABLE IF EXISTS stations CASCADE;"
	return writeToDB(pg.db, query)
}


func parseRow(rows *sql.Rows) (row stationssvc.Station, err error) {
	err = rows.Scan(
		&row.ID,
		&row.Name,
		&row.Timezone,
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