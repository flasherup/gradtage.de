package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"github.com/flasherup/gradtage.de/autocompletesvc/config"
	_ "github.com/lib/pq"
)

//Postgres database
type Postgres struct {
	db *sql.DB
}

const tableName = "autocomplete"

//NewPostgres create and initialize database and return it or error
func NewPostgres(config config.DatabaseConfig) (pg *Postgres, err error) {
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
		db: db,
	}
	return
}

//AddSources
func (pg *Postgres) GetAutocomplete(text string) (map[string][]autocompletesvc.Autocomplete, error) {
	result := make(map[string][]autocompletesvc.Autocomplete)
	query := "(SELECT *, 'icao' as column " +
		"FROM " + tableName + " " +
		"WHERE icao ILIKE '%" + text + "%') " +
		"UNION ALL " +
		"(SELECT *, 'id' as column " +
		"FROM " + tableName + " " +
		"WHERE city_name_english ILIKE '%" + text + "%') " +
		"UNION ALL " +
		"(SELECT *, 'wmo' as column " +
		"FROM " + tableName + " " +
		"WHERE wmo ILIKE '%" + text + "%')" +
		"UNION ALL " +
		"(SELECT *, 'cwop' as column " +
		"FROM " + tableName + " " +
		"WHERE cwop ILIKE '%" + text + "%');"
	rows, err := pg.db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	row := struct {
		autocompletesvc.Autocomplete
		Column string
	}{}

	for rows.Next() {
		err = rows.Scan(
			&row.ID,
			&row.SourceID,
			&row.Latitude,
			&row.Longitude,
			&row.Source,
			&row.Reports,
			&row.ISO2Country,
			&row.ISO3Country,
			&row.Prio,
			&row.CityNameEnglish,
			&row.CityNameNative,
			&row.CountryNameEnglish,
			&row.CountryNameNative,
			&row.ICAO,
			&row.WMO,
			&row.CWOP,
			&row.Maslib,
			&row.National_ID,
			&row.IATA,
			&row.USAF_WBAN,
			&row.GHCN,
			&row.NWSLI,
			&row.Elevation,
			&row.Column,
		)
		if err == nil {
			_, ok := result[row.Column]
			if !ok {
				result[row.Column] = make([]autocompletesvc.Autocomplete, 0)
			}
			result[row.Column] = append(result[row.Column], row.Autocomplete)
		}
	}
	return result, err
}

func (pg *Postgres) GetStationId(text string) (map[string][]autocompletesvc.Autocomplete, error) {
	result := make(map[string][]autocompletesvc.Autocomplete)
	query := "(SELECT *, 'id' as column " +
		"FROM " + tableName + " " +
		"WHERE id ILIKE '" + text + "') " +
		"UNION ALL " +
		"(SELECT *, 'icao' as column " +
		"FROM " + tableName + " " +
		"WHERE icao ILIKE '" + text + "') " +
		"UNION ALL " +
		"(SELECT *, 'station' as column " +
		"FROM " + tableName + " " +
		"WHERE city_name_english ILIKE '" + text + "') " +
		"UNION ALL " +
		"(SELECT *, 'wmo' as column " +
		"FROM " + tableName + " " +
		"WHERE wmo ILIKE '" + text + "')" +
		"UNION ALL " +
		"(SELECT *, 'cwop' as column " +
		"FROM " + tableName + " " +
		"WHERE cwop ILIKE '" + text + "');"

	rows, err := pg.db.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	row := struct {
		autocompletesvc.Autocomplete
		Column string
	}{}

	for rows.Next() {
			err = rows.Scan(
			&row.ID,
			&row.SourceID,
			&row.Latitude,
			&row.Longitude,
			&row.Source,
			&row.Reports,
			&row.ISO2Country,
			&row.ISO3Country,
			&row.Prio,
			&row.CityNameEnglish,
			&row.CityNameNative,
			&row.CountryNameEnglish,
			&row.CountryNameNative,
			&row.ICAO,
			&row.WMO,
			&row.CWOP,
			&row.Maslib,
			&row.National_ID,
			&row.IATA,
			&row.USAF_WBAN,
			&row.GHCN,
			&row.NWSLI,
			&row.Elevation,
			&row.Column,
		)
		if err == nil {
			_, ok := result[row.Column]
			if !ok {
				result[row.Column] = make([]autocompletesvc.Autocomplete, 0)
			}
			result[row.Column] = append(result[row.Column], row.Autocomplete)
		}
	}
	return result, err
}

//GetAllStations get a list of station
func (pg Postgres) GetAllStations() (map[string]*acrpc.Source,error) {
	sts := make(map[string]*acrpc.Source)
	query := fmt.Sprintf("SELECT * FROM %s;", tableName)

	rows, err := pg.db.Query(query)
	defer rows.Close()
	if err != nil {
		return sts, err
	}

	for rows.Next() {
		st,err := parseSourceRow(rows)
		if err != nil {
			return sts, err
		}
		sts[st.ID] = &st
	}
	return sts, nil
}


func parseSourceRow(rows *sql.Rows) (source acrpc.Source, err error) {
	err = rows.Scan(
			&source.ID,
			&source.SourceID,
			&source.Latitude,
			&source.Longitude,
			&source.Source,
			&source.Reports,
			&source.ISO2Country,
			&source.ISO3Country,
			&source.Prio,
			&source.CityNameEnglish,
			&source.CityNameNative,
			&source.CountryNameEnglish,
			&source.CountryNameNative,
			&source.ICAO,
			&source.WMO,
			&source.CWOP,
			&source.Maslib,
			&source.National_ID,
			&source.IATA,
			&source.USAF_WBAN,
			&source.GHCN,
			&source.NWSLI,
			&source.Elevation,
		)
	return
}

//AddSources
func (pg *Postgres) AddSources(sources []autocompletesvc.Autocomplete) (err error) {
	length := len(sources)
	if length == 0 {
		return errors.New("add sources error, sources list is empty")
	}

	var query string
	iterationStep := 100
	for i:=0; i<length; i++ {
		if i%iterationStep == 0  {
			query = fmt.Sprintf("INSERT INTO %s ("+
				"id,"+
				"source_id,"+
				"latitude,"+
				"longitude,"+
				"source,"+
				"reports,"+
				"iso_2_country,"+
				"iso_3_country,"+
				"prio,"+
				"city_name_english,"+
				"city_name_native,"+
				"country_name_english,"+
				"country_name_native,"+
				"icao,"+
				"wmo,"+
				"cwop,"+
				"maslib,"+
				"national_id,"+
				"iata,"+
				"usaf_wban,"+
				"ghcn,"+
				"nwsli,"+
				"elevation"+
				") VALUES", tableName)
		}
		v := sources[i]
		query += "("
		query += fmt.Sprintf("'%s',", v.ID)
		query += fmt.Sprintf("'%s',", v.SourceID)
		query += fmt.Sprintf("'%g',", v.Latitude)
		query += fmt.Sprintf("'%g',", v.Longitude)
		query += fmt.Sprintf("'%s',", v.Source)
		query += fmt.Sprintf("'%s',", v.Reports)
		query += fmt.Sprintf("'%s',", v.ISO2Country)
		query += fmt.Sprintf("'%s',", v.ISO3Country)
		query += fmt.Sprintf("'%s',", v.Prio)
		query += fmt.Sprintf("'%s',", v.CityNameEnglish)
		query += fmt.Sprintf("'%s',", v.CityNameNative)
		query += fmt.Sprintf("'%s',", v.CountryNameEnglish)
		query += fmt.Sprintf("'%s',", v.CountryNameNative)
		query += fmt.Sprintf("'%s',", v.ICAO)
		query += fmt.Sprintf("'%s',", v.WMO)
		query += fmt.Sprintf("'%s',", v.CWOP)
		query += fmt.Sprintf("'%s',", v.Maslib)
		query += fmt.Sprintf("'%s',", v.National_ID)
		query += fmt.Sprintf("'%s',", v.IATA)
		query += fmt.Sprintf("'%s',", v.USAF_WBAN)
		query += fmt.Sprintf("'%s',", v.GHCN)
		query += fmt.Sprintf("'%s',", v.NWSLI)
		query += fmt.Sprintf("'%g'", v.Elevation)
		query += ")"

		if (i+1)%iterationStep != 0 && i < length -1 {
			query += ","
		} else {
			query += " ON CONFLICT (id) DO NOTHING;"
			/*query += ` ON CONFLICT (id) DO UPDATE SET (
					source_id,
					latitude,
					longitude,
					source,
					reports,
					iso_2_country,
					iso_3_country,
					prio,
					city_name_english,
					city_name_native,
					country_name_english,
					country_name_native,
					icao,
					wmo,
					cwop,
					maslib,
					national_id,
					iata,
					usaf_wban,
					ghcn,
					nwsli,
					elevation
				) = (
					excluded.source_id,
					excluded.latitude,
					excluded.longitude,
					excluded.source,
					excluded.reports,
					excluded.iso_2_country,
					excluded.iso_3_country,
					excluded.prio,
					excluded.city_name_english,
					excluded.city_name_native,
					excluded.country_name_english,
					excluded.country_name_native,
					excluded.icao,
					excluded.wmo,
					excluded.cwop,
					excluded.maslib,
					excluded.national_id,
					excluded.iata,
					excluded.usaf_wban,
					excluded.ghcn,
					excluded.nwsli,
					excluded.elevation
				);`*/
			err := writeToDB(pg.db, query)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//Dispose and disconnect
func (pg *Postgres) Dispose() {
	pg.db.Close()
	pg.db = nil
}

//CreateTable create a "Stations" table if not exist
func (pg Postgres) CreateTable() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id varchar(15) UNIQUE,  
			source_id varchar(15),  
			latitude real, 
			longitude real, 
			source varchar(8),  
			reports varchar(15),  
			iso_2_country varchar(3),  
			iso_3_country varchar(3),  
			prio varchar(1),  
			city_name_english varchar(70),  
			city_name_native varchar(70),  
			country_name_english varchar(70),  
			country_name_native varchar(70),  
			icao varchar(4),  
			wmo varchar(8),  
			cwop varchar(8),  
			maslib varchar(8),  
			national_id varchar(15),  
			iata varchar(4),  
			usaf_wban varchar(15),  
			ghcn varchar(15),  
			nwsli varchar(8),  
			elevation real 
		);`, tableName)
	return writeToDB(pg.db, query)
}

//RemoveTable remove stations table from BD
func (pg *Postgres) RemoveTable() error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", tableName)
	return writeToDB(pg.db, query)
}

func writeToDB(db *sql.DB, query string) (err error) {
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	rows.Close()
	return
}
