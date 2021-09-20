package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/metricssvc/config"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
	_ "github.com/lib/pq"
	"math"
)

const metricsTable = "station_metrics"
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

func (pg *Postgres)GetMetrics(ids []string) (map[string]*mtrgrpc.Metrics, error) {
	length := len(ids)
	if length == 0 {
		return map[string]*mtrgrpc.Metrics{}, nil
	}

	idString := ""
	for _,v := range ids {
		idString += fmt.Sprintf("'%s',",v)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s);", metricsTable, idString[:len(idString)-1])

	rows, err := pg.db.Query(query)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	res := make(map[string]*mtrgrpc.Metrics)

	for rows.Next() {
		metrics, err := parseRow(rows)
		if err != nil {
			return res, err
		}
		res[metrics.Id] = &metrics.Metrics
	}
	return res,nil
}

func (pg *Postgres)PushMetrics(metrics map[string]*mtrgrpc.Metrics) error {
	length := len(metrics)
	if length == 0 {
		return errors.New("metrics push error, data is empty")
	}

	ids := make([]string, length)
	i := 0
	for k,_ := range metrics {
		ids[i] = k
		i++
	}

	iterationStep := 100
	steps := int(math.Floor(float64(length/iterationStep))) + 1
	left := length%iterationStep
	for i:=0; i<steps; i++ {
		query := fmt.Sprintf(`INSERT INTO %s (
			id,
			date,
			last_update,
			first_update,
			records_all,
			records_clean
			) VALUES`, metricsTable)
		s := iterationStep * i
		e := s + iterationStep
		if e >= length {
			e = s + left
		}
		cId := ids[s:e]
		for j,id := range  cId {
			v := metrics[id]
			query += "("
			query += fmt.Sprintf( "'%s',", id)
			query += fmt.Sprintf( "'%s',", v.Date)
			query += fmt.Sprintf( "'%s',", v.LastUpdate)
			query += fmt.Sprintf( "'%s',", v.FirstUpdate)
			query += fmt.Sprintf( "%d,", v.RecordsAll)
			query += fmt.Sprintf( "%d", v.RecordsClean)
			query += ")"

			if s+j < length-1 {
				query += ","
			}

		}
		query += ` ON CONFLICT (id) DO UPDATE SET (
					date,
					last_update,
					first_update,
					records_all,
					records_clean
				) = (
					excluded.date,
					excluded.last_update,
					excluded.first_update,
					excluded.records_all,
					excluded.records_clean
				);`
		err := writeToDB(pg.db, query)
		if err != nil{
			return err
		}
	}
	return nil
}

//CreateTable create a table with name @icao + tPrefix if not exist
func (pg *Postgres) CreateTable() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id varchar(15) UNIQUE,
		date timestamp,
		last_update timestamp,
		first_update timestamp,
		records_all integer,
		records_clean integer
		);`,
		metricsTable)
	return writeToDB(pg.db, query)
}

//RemoveTable remove stations table from BD
func (pg *Postgres) RemoveTable() error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;",
		metricsTable)
	return writeToDB(pg.db, query)
}


//Dispose and disconnect
func (pg *Postgres) Dispose() {
	pg.db.Close()
	pg.db = nil
}

func writeToDB(db *sql.DB, query string) (err error){
	row, err := db.Query(query)
	if err != nil {
		return
	}
	row.Close()
	return
}

type parsedRow struct {
	Id string
	mtrgrpc.Metrics
}

func parseRow(rows *sql.Rows) (row parsedRow, err error) {
	err = rows.Scan(
		&row.Id,
		&row.Date,
		&row.LastUpdate,
		&row.FirstUpdate,
		&row.RecordsAll,
		&row.RecordsClean,
	)
	return
}

























