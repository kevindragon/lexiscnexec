// 从history.history_{year}表里面查询用户的并发数

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kevindragon/lexiscnexec/config"
)

const (
	LongTimeFormat = "2006-01-02 15:04:05"

	//CUSTOMER_ID = "chinataxtrial"
	CUSTOMER_ID = "sndhrtrial"
)

func main() {
	databaseDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.DB_USER, config.DB_PWD, config.DB_HOST,
		config.DB_PORT, config.DB_HISTORY_NAME)
	db, err := sql.Open("mysql", databaseDSN)
	if err != nil {
		log.Fatal("open database error: ", err)
	}

	//sql := fmt.Sprintf("SELECT rid, ip, time FROM history_2014 WHERE user = ? "+
	//	"AND (time > '%s' AND time < '%s')", "2014-08-18", "2014-09-08")
	sql := fmt.Sprintf("SELECT rid, ip, time FROM history_2014 WHERE user = ?")
	rows, err := db.Query(sql, CUSTOMER_ID)

	cc := &Concurrency{}

	var lastStartTime time.Time
	var lastConcurrency int

	for rows.Next() {
		var rid int
		var ip, timeStr string
		rows.Scan(&rid, &ip, &timeStr)

		cc.AddActive(ip, timeStr)
		st, et := cc.GetStartEnd()

		if lastStartTime.Unix() != st.Unix() || lastConcurrency != cc.GetConcurrency() {
			fmt.Printf("%d,%s,%s\n", cc.GetConcurrency(),
				st.Format(LongTimeFormat), et.Format(LongTimeFormat))
			lastStartTime = st
			lastConcurrency = cc.GetConcurrency()
		}
	}
}

type Concurrency struct {
	onlineID map[string]time.Time // ip: last active time
}

func (c *Concurrency) AddActive(ip, timStr string) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(LongTimeFormat, timStr, loc)

	if c.onlineID == nil {
		c.onlineID = make(map[string]time.Time)
	}
	c.onlineID[ip] = t

	idsCopy := c.onlineID

	for k, v := range idsCopy {
		if t.Unix()-v.Unix() > 3600 {
			//fmt.Println(t, v, t.Unix()-v.Unix())
			delete(c.onlineID, k)
		}
	}
}
func (c *Concurrency) GetConcurrency() int {
	return len(c.onlineID)
}
func (c *Concurrency) GetStartEnd() (time.Time, time.Time) {
	var min int64 = -1
	var max int64 = -1
	for _, v := range c.onlineID {
		if min == -1 {
			min = v.Unix()
		}
		if v.Unix() < min {
			min = v.Unix()
		}

		if max == -1 {
			max = v.Unix()
		}
		if v.Unix() > max {
			max = v.Unix()
		}
	}

	var minTime, maxTime time.Time
	if min != -1 {
		minTime = time.Unix(min, 0)
	}
	if max != -1 {
		maxTime = time.Unix(max, 0)
	}

	return minTime, maxTime
}
