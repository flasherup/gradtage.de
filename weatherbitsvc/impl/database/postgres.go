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
	for i, v := range wbd.Data {
		query += "("
		length := len(wbd.Data)

		roundedTs := math.Floor(v.TS)
		date := time.Unix(int64(roundedTs), 0)
		time := date.Format(common.TimeLayout)
		query += fmt.Sprintf( "'%s',", time)
		query += fmt.Sprintf( "%g,", v.Rh)
		query += fmt.Sprintf( "'%s',", v.Pod)
		query += fmt.Sprintf( "%g,", v.Pres)
		query += fmt.Sprintf( "'%s',", v.Timezone)
		query += fmt.Sprintf( "'%s',", v.CountryCode)
		query += fmt.Sprintf( "%g,", v.Clouds)
		query += fmt.Sprintf( "%g,", v.Vis)
		query += fmt.Sprintf( "%g,", v.SolarRad)
		query += fmt.Sprintf( "%g,", v.WindSPD)
		query += fmt.Sprintf( "'%s',", v.StateCode)
		query += fmt.Sprintf( "'%s',", v.CityName)
		query += fmt.Sprintf( "%g,", v.AppTemp)
		query += fmt.Sprintf( "%g,", v.UV)
		query += fmt.Sprintf( "'%s',", v.Lon)
		query += fmt.Sprintf( "%g,", v.SLP)
		query += fmt.Sprintf( "%g,", v.HAngle)
		query += fmt.Sprintf( "%g,", v.Dewpt)
		query += fmt.Sprintf( "%g,", v.Snow)
		query += fmt.Sprintf( "%g,", v.AQI)
		query += fmt.Sprintf( "%g,", v.WindDir)
		query += fmt.Sprintf( "%g,", v.ElevAngle)
		query += fmt.Sprintf( "%g,", v.GHI)
		query += fmt.Sprintf( "'%s',", v.Lat)
		query += fmt.Sprintf( "%g,", v.Precip)
		query += fmt.Sprintf( "'%s',", v.Sunset)
		query += fmt.Sprintf( "%g,", v.Temp)
		query += fmt.Sprintf( "'%s',", v.Station)
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
		"	rh real,"+
		"	pod VARCHAR(1),"+
		"	pres real,"+
		"	timezone VARCHAR,"+
		"	country_code VARCHAR(4),"+
		"	clouds real,"+
		"	vis real,"+
		"	solar_rad real,"+
		"	wind_spd real,"+
		"	state_code VARCHAR(4),"+
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



func parseRow(rows *sql.Rows) (temp hourlysvc.Temperature, err error) {
	bdData := struct {
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
	}{}

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
	temp.Date = bdData.Date
	temp.Temperature = bdData.Temp
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


























