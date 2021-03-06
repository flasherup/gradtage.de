package database

import (
	"database/sql"
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/autocompletesvc/config"
	_ "github.com/lib/pq"
)

//Postgres database
type Postgres struct {
	db *sql.DB
}

const tableName = "autocomplete"

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

//Query example
//(SELECT *, 'icao' as column
//FROM autocomplete
//WHERE icao LIKE '%frank%')
//UNION ALL
//(SELECT *, 'station' as column
//FROM autocomplete
//WHERE station LIKE '%frank%')
//UNION ALL
//(SELECT *, 'dwd' as column
//FROM autocomplete
//WHERE dwd LIKE '%frank%')
//UNION ALL
//(SELECT *, 'wmo' as column
//FROM autocomplete
//WHERE wmo LIKE '%frank%');

//AddSources
func (pg *Postgres) GetAutocomplete(text string) (map[string][]autocompletesvc.Source, error) {
	result := make(map[string][]autocompletesvc.Source)
	query := "(SELECT *, 'icao' as column " +
	"FROM autocomplete " +
	"WHERE icao ILIKE '%" + text + "%') " +
	"UNION ALL " +
	"(SELECT *, 'station' as column " +
	"FROM autocomplete " +
	"WHERE name ILIKE '%" + text + "%') " +
	"UNION ALL " +
	"(SELECT *, 'dwd' as column " +
	"FROM autocomplete " +
	"WHERE dwd ILIKE '%" + text + "%') " +
	"UNION ALL " +
	"(SELECT *, 'wmo' as column " +
	"FROM autocomplete " +
	"WHERE wmo ILIKE '%" + text + "%');"
	rows, err := pg.db.Query(query)
	if err != nil {
		return result,err
	}
	defer rows.Close()

	row := struct {
		ID 		string
		Name	string
		Icao 	string
		Dwd 	string
		Wmo 	string
		Column  string
	}{}

	for rows.Next() {
		err = rows.Scan(
			&row.ID,
			&row.Name,
			&row.Icao,
			&row.Dwd,
			&row.Wmo,
			&row.Column,
		)
		if err == nil {
			_,ok := result[row.Column]
			if !ok {
				result[row.Column] = make([]autocompletesvc.Source, 0)
			}
			result[row.Column] = append(result[row.Column],autocompletesvc.Source{
				ID:row.ID,
				Name:row.Name,
				Icao:row.Icao,
				Dwd:row.Dwd,
				Wmo:row.Wmo,
			})
		}
	}
	return result, err
}

func (pg *Postgres) GetStationId(text string) (map[string][]autocompletesvc.Source, error) {
	result := make(map[string][]autocompletesvc.Source)
	query := "(SELECT *, 'id' as column " +
		"FROM autocomplete " +
		"WHERE id ILIKE '" + text + "') " +
		"UNION ALL " +
		"(SELECT *, 'icao' as column " +
		"FROM autocomplete " +
		"WHERE icao ILIKE '" + text + "') " +
		"UNION ALL " +
		"(SELECT *, 'station' as column " +
		"FROM autocomplete " +
		"WHERE name ILIKE '" + text + "') " +
		"UNION ALL " +
		"(SELECT *, 'dwd' as column " +
		"FROM autocomplete " +
		"WHERE dwd ILIKE '" + text + "') " +
		"UNION ALL " +
		"(SELECT *, 'wmo' as column " +
		"FROM autocomplete " +
		"WHERE wmo ILIKE '" + text + "');"
	rows, err := pg.db.Query(query)
	if err != nil {
		return result,err
	}
	defer rows.Close()

	row := struct {
		ID 		string
		Name	string
		Icao 	string
		Dwd 	string
		Wmo 	string
		Column  string
	}{}

	for rows.Next() {
		err = rows.Scan(
			&row.ID,
			&row.Name,
			&row.Icao,
			&row.Dwd,
			&row.Wmo,
			&row.Column,
		)
		if err == nil {
			_,ok := result[row.Column]
			if !ok {
				result[row.Column] = make([]autocompletesvc.Source, 0)
			}
			result[row.Column] = append(result[row.Column],autocompletesvc.Source{
				ID:row.ID,
				Name:row.Name,
				Icao:row.Icao,
				Dwd:row.Dwd,
				Wmo:row.Wmo,
			})
		}
	}
	return result, err
}

//AddSources
func (pg *Postgres) AddSources(sources []autocompletesvc.Source) (err error) {
	query := fmt.Sprintf("INSERT INTO %s " +
		"(id, name, icao, dwd, wmo) VALUES", tableName)


	length := len(sources)
	for i, v := range sources {
		query += fmt.Sprintf(
			" ( '%s', '%s', '%s', '%s', '%s')",
			v.ID, v.Name, v.Icao, v.Dwd, v.Wmo)
		if i < length-1 {
			query += ","
		}
	}
	query += ` ON CONFLICT (id) DO UPDATE SET (name, icao, dwd, wmo) = (excluded.name, excluded.icao, excluded.dwd, excluded.wmo);`
	return writeToDB(pg.db, query)
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
			icao varchar(4),
			dwd varchar(8),
			wmo varchar(8)
		);`, tableName)
	return writeToDB(pg.db, query)
}

//RemoveTable remove stations table from BD
func (pg *Postgres) RemoveTable() error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", tableName)
	return writeToDB(pg.db, query)
}

func writeToDB(db *sql.DB, query string) (err error){
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	rows.Close()
	return
}