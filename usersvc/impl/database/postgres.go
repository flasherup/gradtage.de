package database

import (
	"database/sql"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/config"
	_ "github.com/lib/pq"
	"strconv"
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

//GetPlan(name string) (usersvc.Plan, error)
func (pg *Postgres) GetOrderById(id int) (usersvc.Order, error) {
	query := fmt.Sprintf("SELECT * FROM orders WHERE order_id = '%d';", id)
	//fmt.Println(query)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  usersvc.Order{},err
	}
	defer rows.Close()
	orders, err := parseOrdersRows(rows)
	if err != nil {
		return usersvc.Order{}, err
	}
	if len(orders) == 0 {
		return usersvc.Order{}, fmt.Errorf("oder %d not found", id)
	}
	return orders[0], nil
}

func (pg *Postgres) GetOrdersByUser(user string) ([]usersvc.Order, error) {
	query := fmt.Sprintf("SELECT * FROM orders WHERE email = '%s';", user)
	//fmt.Println(query)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  []usersvc.Order{},err
	}
	defer rows.Close()

	orders, err := parseOrdersRows(rows)
	if err != nil {
		return []usersvc.Order{}, err
	}

	return orders, nil
}

func (pg *Postgres) GetOrderByKey(key string) (usersvc.Order, error) {
	query := fmt.Sprintf("SELECT * FROM orders WHERE key = '%s';", key)
	//fmt.Println(query)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  usersvc.Order{},err
	}
	defer rows.Close()

	orders, err := parseOrdersRows(rows)
	if err != nil {
		return usersvc.Order{}, err
	}

	if len(orders) == 0 {
		return usersvc.Order{}, fmt.Errorf("oder for key %s, not found", key)
	}

	return orders[0], nil
}

/*
*Orders
*/

//Delete orders
func (pg *Postgres) DeleteOrders(orderIds []int) error {
	o := ordersToString(orderIds)
	query := fmt.Sprintf("DELETE FROM orders WHERE order_id = (%s);", o)
	//fmt.Println(query)
	_, err := pg.db.Query(query)
	return err
}

//Set/update order
func (pg *Postgres) SetOrder(order usersvc.Order) error {
	station := stationsToString(order.Stations)
	query := fmt.Sprintf("INSERT INTO orders " +
		"(order_id, key, email, plan, stations, request, req_count, admin) VALUES " +
		"( %d, '%s', '%s', '%s', '{%s}', '%s', %d, %t )",
		order.OrderId,
		order.Key,
		order.Email,
		order.Plan,
		station,
		order.RequestDate.Format(common.TimeLayout),
		order.Requests,
		order.Admin,
		)

	query += ` ON CONFLICT (order_id) DO UPDATE SET
			 (	
				order_id,
				key,
				email,
				plan,
				stations,
				request,
				req_count,
				admin
			) = (
				excluded.order_id,
				excluded.key,
				excluded.email,
				excluded.plan,
				excluded.stations,
				excluded.request,
				excluded.req_count,
				excluded.admin
			);`
	//fmt.Println(query)
	return writeToDB(pg.db, query)
}

//Create Orders Table
func (pg *Postgres) CreateOrdersTable() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS orders (
				order_id 	integer UNIQUE,
				key 		varchar(%d),
				email 		varchar(50),
				plan 		varchar(15),
				stations 	varchar(15)[],
				request 	timestamp,
				req_count	integer,
				admin 		boolean
			);`, KeyLength)
	//fmt.Println(query)
	return writeToDB(pg.db, query)
}

//Remove Orders table
func (pg *Postgres) RemoveOrdersTable() error {
	query := "DROP TABLE IF EXISTS orders CASCADE;"
	return writeToDB(pg.db, query)
}

func parseOrdersRows(rows *sql.Rows) ([]usersvc.Order, error) {
	var err error
	u := struct {
		order_id 	int
		key			string
		email		string
		plan 		string
		stations 	[]uint8
		request 	string
		req_count	int
		admin 		bool
	}{}

	orders := make([]usersvc.Order, 0)

	for rows.Next() {
		err = rows.Scan(
			&u.order_id,
			&u.key,
			&u.email,
			&u.plan,
			&u.stations,
			&u.request,
			&u.req_count,
			&u.admin,
		)

		stations := parseToStringSlice(u.stations)

		request, err := time.Parse(common.TimeLayout, u.request)
		if err != nil {
			return orders,err
		}

		order := usersvc.Order{
			OrderId: u.order_id,
			Key: u.key,
			Email: u.email,
			Plan: u.plan,
			Stations: stations,
			RequestDate: request,
			Requests: u.req_count,
			Admin: u.admin,
		}
		orders = append(orders, order)
	}
	return orders, err
}

/*
* Plans
*/
//SetPlan(plan usersvc.Plan) error
func (pg *Postgres) SetPlan(plan usersvc.Plan) error {
	query := fmt.Sprintf("INSERT INTO plans " +
		"(name, stations, limitation, hdd, dd, cdd, stime, etime, period) VALUES " +
		"( '%s', '%d', '%d', '%t', %t, '%t', '%s', '%s', %d)",
		plan.Name,
		plan.Stations,
		plan.Limitation,
		plan.HDD,
		plan.DD,
		plan.CDD,
		plan.Start.Format(common.TimeLayout),
		plan.End.Format(common.TimeLayout),
		plan.Period,
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
				period
			) = (
				excluded.name,
				excluded.stations,
				excluded.limitation,
				excluded.hdd,
				excluded.dd,
				excluded.cdd,
				excluded.stime,
				excluded.etime,
				excluded.period
			);`
	//fmt.Println(query)

	return writeToDB(pg.db, query)
}

//GetPlan(name string) (usersvc.Plan, error)
func (pg *Postgres) GetPlans(plans []string) ([]usersvc.Plan, error) {
	p := plansToString(plans)
	query := fmt.Sprintf("SELECT * FROM plans WHERE name in (%s);", p)
	//fmt.Println(query)
	rows, err := pg.db.Query(query)
	if err != nil {
		return  []usersvc.Plan{},err
	}
	defer rows.Close()
	return parsePlanRows(rows)
}

//CreatePlanTable() error
func (pg *Postgres) CreatePlansTable() error {
	query := `CREATE TABLE IF NOT EXISTS plans (
				name 		varchar(15) UNIQUE,
				stations 	integer,
				limitation 	integer,
				hdd			bool,
				dd			bool,
				cdd			bool,
				stime 		timestamp,
				etime 		timestamp,
				period		integer
			);`
	//fmt.Println(query)
	return writeToDB(pg.db, query)
}

//RemovePlanTable remove plan table from BD
func (pg *Postgres) RemovePlansTable() error {
	query := "DROP TABLE IF EXISTS plans CASCADE;"
	return writeToDB(pg.db, query)
}

func parsePlanRows(rows *sql.Rows) ([]usersvc.Plan, error) {
	var err error
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
	}{}

	plans := make([]usersvc.Plan, 0)

	for rows.Next() {
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
		)

		start, err := time.Parse(common.TimeLayout, p.stime)
		if err != nil {
			return []usersvc.Plan{},err
		}

		end, err := time.Parse(common.TimeLayout, p.etime)
		if err != nil {
			return []usersvc.Plan{},err
		}

		plan := usersvc.Plan{
			Name:p.name,
			Stations:p.stations,
			Limitation:p.limitation,
			HDD:p.hdd,
			DD:p.dd,
			CDD:p.cdd,
			Start:start,
			End:end,
			Period:p.period,
		}
		plans = append(plans, plan)

	}

	return plans, err
}

//Dispose and disconnect
func (pg *Postgres) Dispose() {
	pg.db.Close()
	pg.db = nil
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
			word = make([]byte, 0)
			continue
		}
		word = append(word,  v)
	}
	return res
}


func writeToDB(db *sql.DB, query string) (err error){
	row, err := db.Query(query)
	if err != nil {
		return
	}
	row.Close()
	return
}

func ordersToString(orders []int) string {
	ordersCount := len(orders)-1
	res := ""
	for i,v := range orders {
		res += strconv.Itoa(v)
		if i < ordersCount {
			res += ","
		}
	}
	return res
}

func stationsToString(stations []string) string {
	stationsCount := len(stations)-1
	res := ""
	for i,v := range stations {
		res += v
		if i < stationsCount {
			res += ","
		}
	}
	return res
}

func plansToString(plans []string) string {
	count := len(plans)-1
	res := ""
	for i,v := range plans {
		res += "'" + v + "'"
		if i < count {
			res += ","
		}
	}
	return res
}



