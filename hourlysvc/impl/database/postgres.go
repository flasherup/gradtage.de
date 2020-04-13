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

//GetUpdateDate return latest date of update for station with @name
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

//Request example
//(SELECT *, 'de00044' as name
//FROM de00044
//ORDER BY date
//DESC LIMIT 1)
//UNION ALL
//(SELECT *, 'de00071' as name
//FROM de00071
//ORDER BY date
//DESC LIMIT 1);
//GetUpdateDateList return latest dates of update for stations with specified in @names
func (pg *Postgres)GetUpdateDateList(names []string) (temps map[string]string, err error) {
	query := ""
	for i,v := range names {
		query += fmt.Sprintf("(SELECT *, '%s' as name FROM %s ORDER BY date DESC LIMIT 1)",
			v, v)

		if i < len(names)-1 {
			query += " UNION ALL "
		} else {
			query += ";"
		}
	}

	rows, err := pg.db.Query(query)
	if err != nil {
		return temps,err
	}
	defer rows.Close()

	temps = map[string]string{}

	row := struct {
		Date 		string
		Temperature float64
		Name 		string
	}{}

	for rows.Next() {
		err = rows.Scan(
			&row.Date,
			&row.Temperature,
			&row.Name,
		)

		if err == nil {
			temps[row.Name] = row.Date
		}
	}
	return temps, nil
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


//Request example
//(SELECT *, 'de00044' as name
//FROM de00044
//ORDER BY date
//DESC LIMIT 1)
//UNION ALL
//(SELECT *, 'de00071' as name
//FROM de00071
//ORDER BY date
//DESC LIMIT 1);
//GetLatestList return latest temperatures data
//for station specified in @[]strings
func (pg *Postgres)GetLatestList(names []string) (temps map[string]hourlysvc.Temperature, err error) {
	query := ""

	for i,v := range names {
		query += fmt.Sprintf("(SELECT *, '%s' as name FROM %s ORDER BY date DESC LIMIT 1)",
			v, v)

		if i < len(names)-1 {
			query += " UNION ALL "
		} else {
			query += ";"
		}
	}

	rows, err := pg.db.Query(query)
	if err != nil {
		return temps,err
	}
	defer rows.Close()

	temps = map[string]hourlysvc.Temperature{}

	row := struct {
		Date 		string
		Temperature float64
		Name 		string
	}{}

	for rows.Next() {
		err = rows.Scan(
			&row.Date,
			&row.Temperature,
			&row.Name,
		)

		if err == nil {

			temps[row.Name] = hourlysvc.Temperature{
				Date:row.Date,
				Temperature:row.Temperature,
			}
		}
	}
	return temps, nil
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



