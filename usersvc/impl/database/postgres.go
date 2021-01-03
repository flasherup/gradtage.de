package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/config"
	_ "github.com/lib/pq"
	"time"
)

//UserDB main structure
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

//SetUser(user usersvc.User) error
func (pg *Postgres) SetUser(user usersvc.User) error {
	stations := "{"
	for i,v := range user.Stations {
		stations += v
		if i < len(user.Stations)-1 {
			stations += ","
		}
	}
	stations += "}"
	query := fmt.Sprintf("INSERT INTO users " +
		"(key, name, renew, request, req_count, plan, stations, stripe) VALUES " +
		"( '%s', '%s', '%s', '%s', %d, '%s', '%s', '%s')",
		user.Key,
		user.Name,
		user.RenewDate.Format(common.TimeLayout),
		user.RequestDate.Format(common.TimeLayout),
		user.Requests,
		user.Plan,
		stations,
		user.Stripe)

	query += ` ON CONFLICT (name) DO UPDATE SET
			 (	name,
			 	renew,
			 	request,
			 	req_count,
			 	plan,
			 	stations,
				stripe
			) = (
				excluded.name,
			 	excluded.renew,
			 	excluded.request,
			 	excluded.req_count,
			 	excluded.plan,
			 	excluded.stations,
			 	excluded.stripe
			);`

	return writeToDB(pg.db, query)
}

//GetUserDataByName(userName string)  (usersvc.Parameters, error)
func (pg *Postgres) GetUserDataByName(userName string)  (res usersvc.Parameters, err error){
	res = usersvc.Parameters{}
	res.User, err = pg.getUserByName(userName)
	if err != nil {
		return res, err
	}

	res.Plan, err = pg.GetPlan(res.User.Plan)
	if err != nil {
		return res, err
	}

	return res, err
}

//GetUserDataByKey(key string)  (res usersvc.Parameters, err error)
func (pg *Postgres) GetUserDataByKey(key string)  (res usersvc.Parameters, err error){
	res = usersvc.Parameters{}
	res.User, err = pg.getUserByKey(key)
	if err != nil {
		return res, err
	}

	res.Plan, err = pg.GetPlan(res.User.Plan)
	if err != nil {
		return res, err
	}

	return res, err
}

//GetUserDataByStripe(stripe string)  (res usersvc.Parameters, err error)
func (pg *Postgres) GetUserDataByStripe(stripe string)  (res usersvc.Parameters, err error){
	res = usersvc.Parameters{}
	res.User, err = pg.getUserByStripe(stripe)
	if err != nil {
		return res, err
	}

	res.Plan, err = pg.GetPlan(res.User.Plan)
	if err != nil {
		return res, err
	}

	return res, err
}

//SetPlan(plan usersvc.Plan) error
func (pg *Postgres) SetPlan(plan usersvc.Plan) error {
	query := fmt.Sprintf("INSERT INTO plans " +
		"(name, stations, limitation, hdd, dd, cdd, stime, etime, period, admin) VALUES " +
		"( '%s', '%d', '%d', '%t', %t, '%t', '%s', '%s', %d, %t)",
			plan.Name,
			plan.Stations,
			plan.Limitation,
			plan.HDD,
			plan.DD,
			plan.CDD,
			plan.Start.Format(common.TimeLayout),
			plan.End.Format(common.TimeLayout),
			plan.Period,
			plan.Admin,
		)

	query += ` ON CONFLICT (name) DO UPDATE SET
			 (	
				name,
				stations,
				limitation,
				hdd,
				dd,
				cdd,
				stime,
				etime,
				period,
				admin
			) = (
				excluded.name,
				excluded.stations,
				excluded.limitation,
				excluded.hdd,
				excluded.dd,
				excluded.cdd,
				excluded.stime,
				excluded.etime,
				excluded.period,
				excluded.admin
			);`
	return writeToDB(pg.db, query)
}

//GetPlan(name string) (usersvc.Plan, error)
func (pg *Postgres) GetPlan(name string) (usersvc.Plan, error) {
	query := fmt.Sprintf("SELECT * FROM plans WHERE name = '%s';", name)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  usersvc.Plan{},err
	}
	defer rows.Close()
	var res usersvc.Plan
	for rows.Next() {
		p, err := parsePlanRow(rows)
		if err != nil {
			return res, err
		}

		res = p
	}
	return res, nil
}

func (pg *Postgres) getUserByName(userName string) (usersvc.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s';", userName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  usersvc.User{},err
	}
	defer rows.Close()
	var res usersvc.User
	for rows.Next() {
		u, err := parseUserRow(rows)
		if err != nil {
			return res, err
		}

		res = u
	}

	if res.Name == "" {
		err = errors.New("user not found")
	}
	return res, nil
}


func (pg *Postgres) getUserByKey(key string) (usersvc.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE key = '%s';", key)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  usersvc.User{},err
	}

	defer rows.Close()
	var res usersvc.User
	for rows.Next() {
		u, err := parseUserRow(rows)
		if err != nil {
			return res, err
		}

		res = u
	}

	if res.Name == "" {
		err = errors.New("user not found")
	}
	return res, err
}

func (pg *Postgres) getUserByStripe(stripe string) (usersvc.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE stripe = '%s';", stripe)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  usersvc.User{},err
	}
	defer rows.Close()
	var res usersvc.User
	for rows.Next() {
		u, err := parseUserRow(rows)
		if err != nil {
			return res, err
		}

		res = u
	}

	if res.Name == "" {
		err = errors.New("user not found")
	}
	return res, nil
}

//CreateUserTable() error
func (pg *Postgres) CreateUserTable() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS users (
			key 		varchar(%d),
			name 		varchar(50) UNIQUE,
			renew 		timestamp,
			request 	timestamp,
			req_count	integer,
			plan 		varchar(15),
			stations 	varchar(8)[],
			stripe 		varchar(20)
		);`, KeyLength)
	return writeToDB(pg.db, query)
}

//CreatePlanTable() error
func (pg *Postgres) CreatePlanTable() error {
	query := `CREATE TABLE IF NOT EXISTS plans (
				name 		varchar(15) UNIQUE,
				stations 	integer,
				limitation 	integer,
				hdd			bool,
				dd			bool,
				cdd			bool,
				stime 		timestamp,
				etime 		timestamp,
				period		integer,
				admin		bool
			);`
	return writeToDB(pg.db, query)
}

//RemoveUserTable remove users table from BD
func (pg *Postgres) RemoveUserTable() error {
	query := "DROP TABLE IF EXISTS users CASCADE;"
	return writeToDB(pg.db, query)
}

//RemovePlanTable remove plan table from BD
func (pg *Postgres) RemovePlanTable() error {
	query := "DROP TABLE IF EXISTS plans CASCADE;"
	return writeToDB(pg.db, query)
}

func parseUserRow(rows *sql.Rows) (user usersvc.User, err error) {
	u := struct {
		key 		string
		name 		string
		renew 		string
		request 	string
		req_count	int
		plan 		string
		stations 	[]uint8
		stripe		string
	}{}
	err = rows.Scan(
		&u.key,
		&u.name,
		&u.renew,
		&u.request,
		&u.req_count,
		&u.plan,
		&u.stations,
		&u.stripe,
	)

	renew, err := time.Parse(common.TimeLayout, u.renew)
	if err != nil {
		return user,err
	}

	request, err := time.Parse(common.TimeLayout, u.request)
	if err != nil {
		return user,err
	}

	str := parseToStringSlice(u.stations)

	user.Key = u.key
	user.Name = u.name
	user.RenewDate = renew
	user.RequestDate = request
	user.Requests = u.req_count
	user.Plan = u.plan
	user.Stations = str
	user.Stripe = u.stripe

	return user, err
}

func parseToStringSlice(slice []uint8) []string {
	if len(slice) < 3 {
		return []string{}
	}
	trim := slice[1:]
	res := make([]string, 0)
	word := make([]byte, 0)
	for _,v := range trim {
		if v == 44 || v == 125 {
			res = append(res, string(word))
			fmt.Println("parsed", string(v), v)
			word = make([]byte, 0)
			continue
		}
		word = append(word,  v)
	}
	return res
}

func parsePlanRow(rows *sql.Rows) (plan usersvc.Plan, err error) {
	p := struct {
		name 		string
		stations 	int
		limitation 	int
		hdd 		bool
		dd			bool
		cdd 		bool
		stime 		string
		etime		string
		period 		int
		admin		bool
	}{}
	err = rows.Scan(
		&p.name,
		&p.stations,
		&p.limitation,
		&p.hdd,
		&p.dd,
		&p.cdd,
		&p.stime,
		&p.etime,
		&p.period,
		&p.admin,
	)

	start, err := time.Parse(common.TimeLayout, p.stime)
	if err != nil {
		return plan,err
	}

	end, err := time.Parse(common.TimeLayout, p.etime)
	if err != nil {
		return plan,err
	}

	plan.Name = p.name
	plan.Stations = p.stations
	plan.Limitation = p.limitation
	plan.HDD = p.hdd
	plan.DD = p.dd
	plan.CDD = p.cdd
	plan.Start = start
	plan.End = end
	plan.Period = p.period
	plan.Admin = p.admin

	return plan, err
}

func writeToDB(db *sql.DB, query string) (err error){
	row, err := db.Query(query)
	if err != nil {
		return
	}
	row.Close()
	return
}



