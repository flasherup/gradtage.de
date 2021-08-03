package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/config"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/parser"
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
		time := date.Format(common.TimeLayout)
		query += fmt.Sprintf( "'%s',", time)
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

func (pg *Postgres) GetWBData(name string, start string, end string) (wbd *parser.WeatherBitData, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE date >= '%s' AND date < '%s' ORDER BY date::timestamp ASC;",
		name, start, end)

	rows, err := pg.db.Query(query)
	if err != nil {
		return wbd,err
	}
	defer rows.Close()
	wbd = &parser.WeatherBitData{}
	wbd.Data = []parser.Data{}

	for rows.Next() {
		dbWBD, err := parseDataRow(rows)
		if err != nil {
			return wbd, err
		}
		wbd.Timezone = dbWBD.Timezone
		wbd.CountryCode = dbWBD.CountryCode
		wbd.StateCode = dbWBD.StateCode
		wbd.CityName = dbWBD.CityName
		wbd.Lon = dbWBD.Lon
		wbd.AQI = dbWBD.AQI
		wbd.Station = dbWBD.Station
		wbd.Data = append(wbd.Data, dbWBD.Data...)
	}

	return wbd,err
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

type DBRow struct {
	Date string
	Temp float64
	pod string
	pres float64
	timezone string
	country_code string
	clouds float64
	vis float64
	solar_rad float64
	wind_spd float64
	state_code string
	city_name string
	app_temp float64
	uv float64
	lon float64
	slp float64
	h_angle float64
	dewpt float64
	snow float64
	aqi float64
	wind_dir float64
	elev_angle float64
	ghi float64
	lat float64
	precip float64
	sunset string
	temp float64
	station string
	dni float64
	sunrise string
}

func parseTempRow(rows *sql.Rows) (hourlysvc.Temperature, error) {
	bdData, err := parseRow(rows)
	temp := hourlysvc.Temperature{}
	temp.Date = bdData.Date
	temp.Temperature = bdData.Temp
	return temp, err
}


func parseDataRow(rows *sql.Rows) (parser.WeatherBitData, error) {
	bdData, err := parseRow(rows)

	data := parser.Data{}
	wbd := parser.WeatherBitData{}
	wbd.Data = []parser.Data{data}
	date, dateErr := time.Parse(common.TimeLayout, bdData.Date)
	if dateErr != nil {
		date = time.Now()
	}
	data.TS = float64(date.Unix())
	data.Temp = bdData.Temp
	data.Pod = bdData.pod
	data.Pres = bdData.pres
	wbd.Timezone = bdData.timezone
	wbd.CountryCode = bdData.country_code
	data.Clouds = bdData.clouds
	data.Vis = bdData.vis
	data.SolarRad = bdData.solar_rad
	data.WindSPD = bdData.wind_spd
	wbd.StateCode = bdData.state_code
	wbd.CityName = bdData.city_name
	data.AppTemp = bdData.app_temp
	data.UV = bdData.uv
	wbd.Lon = bdData.lon
	data.SLP = bdData.slp
	data.HAngle = bdData.h_angle
	data.Dewpt = bdData.dewpt
	data.Snow = bdData.snow
	wbd.AQI = bdData.aqi
	data.WindDir = bdData.wind_dir
	data.ElevAngle = bdData.elev_angle
	data.GHI = bdData.ghi
	data.Precip = bdData.precip
	data.Sunset = bdData.sunset
	data.Temp = bdData.temp
	wbd.Station = bdData.station
	data.DNI = bdData.dni
	data.Sunrise = bdData.sunrise
	return wbd,err
}

func parseRow(rows *sql.Rows) (bdData DBRow, err error) {
	err = rows.Scan(
		&bdData.Date,
		&bdData.Temp,
		&bdData.pod,
		&bdData.pres,
		&bdData.timezone,
		&bdData.country_code,
		&bdData.clouds,
		&bdData.vis,
		&bdData.solar_rad,
		&bdData.wind_spd,
		&bdData.state_code,
		&bdData.city_name,
		&bdData.app_temp,
		&bdData.uv,
		&bdData.lon,
		&bdData.slp,
		&bdData.h_angle,
		&bdData.dewpt,
		&bdData.snow,
		&bdData.aqi,
		&bdData.wind_dir,
		&bdData.elev_angle,
		&bdData.ghi,
		&bdData.lat,
		&bdData.precip,
		&bdData.sunset,
		&bdData.temp,
		&bdData.station,
		&bdData.dni,
		&bdData.sunrise,
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


























