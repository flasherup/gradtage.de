package database

import (
	"database/sql"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/config"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/parser"
	_ "github.com/lib/pq"
	"time"
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
func (pg *Postgres) PushData(stID string, wbd *parser.WeatherBitData) error {

	query := fmt.Sprintf("INSERT INTO %s " +
		"(date, rh, pod, pres, timezone, on_time, country_code, clouds, vis, wind_cdir, ob_time, solar_rad, wind_spd, " +
		"state_code, wind_cdir_full, city_name, app_temp, uv, lon, slp, h_angle, dewpt, snow, aqi, wind_dir, elev_angle, " +
		"ghi, datetime, lat, precip, sunset, temp, station, station, dni, sunrise) VALUES", stID)

	length := len(wbd.Data)
	for i, v := range wbd.Data {
		query += "("


		date := time.Unix(int64(v.TS), 0)
		time := date.Format(common.TimeLayout)
		query += fmt.Sprintf( "%s,", time)
		query += fmt.Sprintf( "%g,", v.Rh)
		query += fmt.Sprintf( "%s,", v.Pod)
		query += fmt.Sprintf( "%g,", v.Pres)
		query += fmt.Sprintf( "%g,", v.Timezone)



			/*%s, %s, %s, %g, %g, %s, %s, %g, %g, %s, %s, %s," +
				   " %g, %c, %s, %g, %g, %g, %g, %g, %g, %g, %g, %s, %g, %g, %s, %g, %s, %g, %s)",    //....


	,
	v.OnTime,
	v.CountryCode,
	v.Clouds,
	v.Vis ,
	v.WindCdir,
	v.ObTime,
	v.SolarRad,
	v.WindSPD,
	v.StateCode,
	v.WindCdirFull,
	v.CityName,
	v.AppTemp,
	v.UV,
	v.Lon,
	v.SLP,
	v.HAngle,
	v.Dewpt,
	v.Snow,
	v.AQI,
	v.WindDir,
	v.ElevAngle,
	v.GHI,
	v.DateTime,
	v.Lat,
	v.Precip,
	v.Sunset,
	v.Temp,
	v.Station,
	v.DNI,
	v.Sunrise*/

		query += ")"
		if i < length-1 {
			query += ","
		}
	}


	query += " ON CONFLICT (date) DO NOTHING;"
	fmt.Println("PushData query", query)
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
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ("+ //.......
		"	date timestamp UNIQUE,"+
		"	temperature real,"+
		"	rh real,"+
		"	pod VARCHAR(1),"+
		"	pres real,"+
		"	timezone VARCHAR,"+
		"	on_time timestamp,"+
		"	country_code,"+
		"	clouds real,"+
		"	vis real,"+
		"	wind_cdir,"+
		"	ob_time timestamp,"+
		"	solar_rad real,"+
		"	wind_spd real,"+
		"	state_code,"+
		"	wind_cdir_full real,"+
		"	city_name real,"+
		"	app_temp real,"+
		"	uv int,"+
		"	lon real,"+
		"	slp real,"+
		"	h_angle real,"+
		"	dewpt real,"+
		"	snow real,"+
		"	aqi real,"+
		"	wind_dir real,"+
		"	elev_angle real,"+
		"	ghi real,"+
		"	datetime timestamp,"+
		"	lat real,"+
		"	precip real,"+
		"	sunset real,"+
		"	temp real,"+
		"	station VARCHAR,"+
		"	dni real,"+
		"	sunrise,"+
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


























