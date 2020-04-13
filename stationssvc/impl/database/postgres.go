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
	db *sql.DB
}

const tableName = "stations"

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
	query := fmt.Sprintf(`INSERT INTO %s 
		(id, name, timezone, source_type, source_id) 
		VALUES ( '%s', '%s', '%s', '%s', '%s') `,
		tableName, station.ID, station.Name, station.Timezone, station.SourceType, station.SourceID)

	return writeToDB(pg.db, query)
}

//AddStations write stations data into DB
func (pg Postgres) AddStations(stations []stationssvc.Station) error {
	query := fmt.Sprintf(`INSERT INTO %s 
		(id, name, timezone, source_type, source_id) VALUES`, tableName)

	length := len(stations)
	for i, v := range stations {
		query += fmt.Sprintf(
			"( '%s', '%s', '%s', '%s', '%s') ",
			v.ID, v.Name, v.Timezone, v.SourceType, v.SourceID)
		if i < length-1 {
			query += ","
		}
	}
	query += ` ON CONFLICT (id) DO UPDATE SET 
		(name, timezone, source_type, source_id) = (excluded.name, excluded.timezone, excluded.source_type, excluded.source_id);`
	return writeToDB(pg.db, query)
}

//DeleteStation remove station by icao ID
func (pg Postgres) DeleteStation(id string) error {
	query := fmt.Sprintf(	`DELETE FROM %s 
									WHERE id = '%s'`,
									tableName, id)
	return writeToDB(pg.db, query)
}

//GetStations get a list of station
func (pg Postgres) GetStations(ids []string) ([]stationssvc.Station,error) {
	sts := make([]stationssvc.Station, 0)
	query := fmt.Sprintf(`SELECT * FROM %s
								 WHERE id IN ( `, tableName)
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
	query := fmt.Sprintf("SELECT * FROM %s;", tableName)

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

//GetAllStations get a list of station
func (pg Postgres) GetStationsBySrcType(types []string) ([]stationssvc.Station,error) {
	sts := make([]stationssvc.Station, 0)
	query := fmt.Sprintf("SELECT * FROM %s WHERE ", tableName)

	length := len(types)
	for i, v := range types {
		query += fmt.Sprintf("source_type='%s' ",v)
		if i < length-1 {
			query += "OR "
		}
	}
	query += ";"

	fmt.Println(query)

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
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s;", tableName)
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
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id varchar(8) UNIQUE,
			name varchar(50),
			timezone varchar(8),
			source_type varchar(4),
			source_id varchar(8)
		);`, tableName)
	return writeToDB(pg.db, query)
}

//RemoveTable remove stations table from BD
func (pg *Postgres) RemoveTable() error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", tableName)
	return writeToDB(pg.db, query)
}


func parseRow(rows *sql.Rows) (row stationssvc.Station, err error) {
	err = rows.Scan(
		&row.ID,
		&row.Name,
		&row.Timezone,
		&row.SourceType,
		&row.SourceID,
	)
	return
}

func writeToDB(db *sql.DB, query string) (err error){
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	rows.Close()
	return
}