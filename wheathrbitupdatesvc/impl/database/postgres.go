package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/wheathrbitupdatesvc/config"
	"github.com/flasherup/gradtage.de/wheathrbitupdatesvc/impl/parser"
	_ "github.com/lib/pq"
	"math"
	"time"
)

//HourlyDB main structure
type Postgres struct {
	db  *sql.DB
}

func (pg *Postgres) GetUpdateDateList(names []string) (temps map[string]string, err error) {
	panic("implement me")
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
	if len(wbd.Data) == 0 {
		return errors.New("weather push error, data is empty")
	}
	query := fmt.Sprintf("INSERT INTO %s " +
		"(date, " +
		"rh, " +
		"pod, " +
		"pres, " +
		"timezone, " +
		"country_code, " +
		"clouds, " +
		"vis, " +
		"solar_rad, " +
		"wind_spd, " +
		"state_code, " +
		"city_name," +
		" app_temp, " +
		"uv, " +
		"lon, " +
		"slp, " +
		"h_angle, " +
		"dewpt, " +
		"snow, " +
		"aqi, " +
		"wind_dir, " +
		"elev_angle, " +
		"ghi, " +
		"lat, " +
		"precip, " +
		"sunset, " +
		"temp, " +
		"station, " +
		"dni, " +
		"sunrise) VALUES", stID)
	length := len(wbd.Data)
	for i, v := range wbd.Data {
		query += "("
		roundedTs := math.Floor(v.TS)
		date := time.Unix(int64(roundedTs), 0)
		dateStr := date.Format(common.TimeLayout)
		query += fmt.Sprintf( "'%s',", dateStr)
		query += fmt.Sprintf( "%g,", v.Rh)
		query += fmt.Sprintf( "'%s',", v.Pod)
		query += fmt.Sprintf( "%g,", v.Pres)
		query += fmt.Sprintf( "'%s',", wbd.Timezone)
		query += fmt.Sprintf( "'%s',", wbd.CountryCode)
		query += fmt.Sprintf( "%g,", v.Clouds)
		query += fmt.Sprintf( "%g,", v.Vis)
		query += fmt.Sprintf( "%g,", v.SolarRad)
		query += fmt.Sprintf( "%g,", v.WindSPD)
		query += fmt.Sprintf( "'%s',", wbd.StateCode)
		query += fmt.Sprintf( "'%s',", wbd.CityName)
		query += fmt.Sprintf( "%g,", v.AppTemp)
		query += fmt.Sprintf( "%g,", v.UV)
		query += fmt.Sprintf( "'%g',", wbd.Lon)
		query += fmt.Sprintf( "%g,", v.SLP)
		query += fmt.Sprintf( "%g,", v.HAngle)
		query += fmt.Sprintf( "%g,", v.Dewpt)
		query += fmt.Sprintf( "%g,", v.Snow)
		query += fmt.Sprintf( "%g,", wbd.AQI)
		query += fmt.Sprintf( "%g,", v.WindDir)
		query += fmt.Sprintf( "%g,", v.ElevAngle)
		query += fmt.Sprintf( "%g,", v.GHI)
		query += fmt.Sprintf( "'%g',", wbd.Lat)
		query += fmt.Sprintf( "%g,", v.Precip)
		query += fmt.Sprintf( "'%s',", v.Sunset)
		query += fmt.Sprintf( "%g,", v.Temp)
		query += fmt.Sprintf( "'%s',", wbd.Station)
		query += fmt.Sprintf( "%g,", v.DNI)
		query += fmt.Sprintf( "'%s'", v.Sunrise)

		query += ")"
		if i < length-1 {
			query += ","

		}
	}

	query += " ON CONFLICT (date) DO NOTHING;"
	return writeToDB(pg.db, query)
}

func (pg *Postgres) GetWBData(name string, start string, end string) (wbd []weatherbitsvc.WBData, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE date >= '%s' AND date < '%s' ORDER BY date::timestamp ASC;",
		name, start, end)

	rows, err := pg.db.Query(query)
	if err != nil {
		return wbd,err
	}
	defer rows.Close()
	wbd = make([]weatherbitsvc.WBData, 0)

	for rows.Next() {
		dbWBD, err := parseRow(rows)
		if err != nil {
			return wbd, err
		}
		wbd = append(wbd, dbWBD)
	}

	return wbd,err
}

func (pg *Postgres) PushWBData(stID string, wbd []weatherbitsvc.WBData) (err error) {
	length := len(wbd)
	if length == 0 {
		return errors.New("weather push error, data is empty")
	}
	iterationStep := 100;
	for i:=iterationStep; i<length; i+=iterationStep{
		query := fmt.Sprintf("INSERT INTO %s " +
			"(date, " +
			"rh, " +
			"pod, " +
			"pres, " +
			"timezone, " +
			"country_code, " +
			"clouds, " +
			"vis, " +
			"solar_rad, " +
			"wind_spd, " +
			"state_code, " +
			"city_name," +
			" app_temp, " +
			"uv, " +
			"lon, " +
			"slp, " +
			"h_angle, " +
			"dewpt, " +
			"snow, " +
			"aqi, " +
			"wind_dir, " +
			"elev_angle, " +
			"ghi, " +
			"lat, " +
			"precip, " +
			"sunset, " +
			"temp, " +
			"station, " +
			"dni, " +
			"sunrise) VALUES", stID)
		for j := i-iterationStep; j<i; j++  {
			v := wbd[j]
			query += "("
			query += fmt.Sprintf( "'%s',", v.Date)
			query += fmt.Sprintf( "%g,", v.Rh)
			query += fmt.Sprintf( "'%s',", v.Pod)
			query += fmt.Sprintf( "%g,", v.Pres)
			query += fmt.Sprintf( "'%s',", v.Timezone)
			query += fmt.Sprintf( "'%s',", v.CountryCode)
			query += fmt.Sprintf( "%g,", v.Clouds)
			query += fmt.Sprintf( "%g,", v.Vis)
			query += fmt.Sprintf( "%g,", v.SolarRad)
			query += fmt.Sprintf( "%g,", v.WindSpd)
			query += fmt.Sprintf( "'%s',", v.StateCode)
			query += fmt.Sprintf( "'%s',", v.CityName)
			query += fmt.Sprintf( "%g,", v.AppTemp)
			query += fmt.Sprintf( "%g,", v.Uv)
			query += fmt.Sprintf( "'%g',", v.Lon)
			query += fmt.Sprintf( "%g,", v.Slp)
			query += fmt.Sprintf( "%g,", v.HAngle)
			query += fmt.Sprintf( "%g,", v.Dewpt)
			query += fmt.Sprintf( "%g,", v.Snow)
			query += fmt.Sprintf( "%g,", v.Aqi)
			query += fmt.Sprintf( "%g,", v.WindDir)
			query += fmt.Sprintf( "%g,", v.ElevAngle)
			query += fmt.Sprintf( "%g,", v.Ghi)
			query += fmt.Sprintf( "'%g',", v.Lat)
			query += fmt.Sprintf( "%g,", v.Precip)
			query += fmt.Sprintf( "'%s',", v.Sunset)
			query += fmt.Sprintf( "%g,", v.Temp)
			query += fmt.Sprintf( "'%s',", v.Station)
			query += fmt.Sprintf( "%g,", v.Dni)
			query += fmt.Sprintf( "'%s'", v.Sunrise)
			query += ")"

			if j < i-1 {
				query += ","
			}

		}
		query += " ON CONFLICT (date) DO NOTHING;"
		err = writeToDB(pg.db, query)
		if err != nil{
			return err
		}
	}

	return nil
}

//GetPeriod get a list of temperatures form table @name (station Id)
func (pg *Postgres) GetPeriod(name string, start string, end string) (temps []common.Temperature, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE date >= '%s' AND date < '%s' ORDER BY date::timestamp ASC;",
		name, start, end)

	rows, err := pg.db.Query(query)
	if err != nil {
		return temps,err
	}
	defer rows.Close()


	for rows.Next() {
		st,err := parseTempRow(rows)
		if err != nil {
			return temps, err
		}
		temps = append(temps, st)
	}

	return temps,err
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
		temp,err := parseTempRow(rows)
		if err != nil {
			return date, err
		}

		date = temp.Date

	}
	return date,err
}


//CreateTable create a table with name @icao + tPrefix if not exist
func (pg *Postgres) CreateTable(name string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ("+ //.......
		"	date timestamp UNIQUE,"+
		"	rh real,"+
		"	pod VARCHAR(1),"+
		"	pres real,"+
		"	timezone VARCHAR,"+
		"	country_code VARCHAR(4),"+
		"	clouds real,"+
		"	vis real,"+
		"	solar_rad real,"+
		"	wind_spd real,"+
		"	state_code VARCHAR(8),"+
		"	city_name VARCHAR,"+
		"	app_temp real,"+
		"	uv real,"+
		"	lon real,"+
		"	slp real,"+
		"	h_angle real,"+
		"	dewpt real,"+
		"	snow real,"+
		"	aqi real,"+
		"	wind_dir real,"+
		"	elev_angle real,"+
		"	ghi real,"+
		"	lat real,"+
		"	precip real,"+
		"	sunset VARCHAR(5),"+
		"	temp real,"+
		"	station VARCHAR,"+
		"	dni real,"+
		"	sunrise VARCHAR(5)"+
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

//RemoveTable remove stations table from BD
func (pg *Postgres) CountTableRows(name string) (int, error) {
	query := fmt.Sprintf("SELECT count(*) AS exact_count FROM %s;",
		name)

	rows, err := pg.db.Query(query)
	if err != nil {
		return 0,err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err = rows.Scan(
			&count,
		)
		if err != nil {
			return 0, err
		}
	}

	return count,nil
}

//GetListOfTables return a list of all stations
func (pg *Postgres) GetListOfTables() ([]string, error) {
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema='public'"

	list := []string{}

	rows, err := pg.db.Query(query)
	if err != nil {
		return list,err
	}
	defer rows.Close()

	var station string
	for rows.Next() {
		err = rows.Scan(
			&station,
		)
		if err != nil {
			return list, err
		}

		list = append(list, station)
	}

	return list,nil
}

func parseTempRow(rows *sql.Rows) (common.Temperature, error) {
	bdData, err := parseRow(rows)
	temp := common.Temperature{}
	temp.Date = bdData.Date
	temp.Temp= bdData.Temp
	return temp, err
}

func parseRow(rows *sql.Rows) (bdData weatherbitsvc.WBData, err error) {
	err = rows.Scan(
		&bdData.Date,
		&bdData.Rh,
		&bdData.Pod,
		&bdData.Pres,
		&bdData.Timezone,
		&bdData.CountryCode,
		&bdData.Clouds,
		&bdData.Vis,
		&bdData.SolarRad,
		&bdData.WindSpd,
		&bdData.StateCode,
		&bdData.CityName,
		&bdData.AppTemp,
		&bdData.Uv,
		&bdData.Lon,
		&bdData.Slp,
		&bdData.HAngle,
		&bdData.Dewpt,
		&bdData.Snow,
		&bdData.Aqi,
		&bdData.WindDir,
		&bdData.ElevAngle,
		&bdData.Ghi,
		&bdData.Lat,
		&bdData.Precip,
		&bdData.Sunset,
		&bdData.Temp,
		&bdData.Station,
		&bdData.Dni,
		&bdData.Sunrise,
	)
	return bdData, err
}

func writeToDB(db *sql.DB, query string) (err error){
	row, err := db.Query(query)
	if err != nil {
		return
	}
	row.Close()
	return
}


























