package mssql

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"telegram_bot/internal/account"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type sqlGetParam struct {
	id    int
	adres string
	cw1   int
	cw2   int
	hw1   int
	hw2   int
	ed    int
	en    int
	gas   int
	wdate string
	gdate string
	edate string
}

type sqlSetParam struct {
	id    int
	adres string
	cw1   int
	cw2   int
	hw1   int
	hw2   int
	ed1   int
	en1   int
	g1    int
	wdate time.Time
	gdate time.Time
	edate time.Time
}

// @chatId=?, @lastKey = ?, @kod_adres = ?, @adres = ?, @cw1 = ?, @cw2 = ?, @hw1 = ?, @hw2 = ?, @ed1 = ?, @en1 = ?, @g1 = ?,
// @wdate = ?,@gdate = ?,@edate = ?,@curcw1 = ?,@curcw2 = ?,@curhw1 = ?,@curhw2 = ?,@cured1 =? ,@curen1 = ?,@curg1 = ?,
// @create_date =?,@status=?

func GetAccountDeviceData(ctx context.Context, connString string, account *account.Account) error {
	kod_adres := account.Id
	data := sqlGetParam{}
	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "usp_web_api_test_get_indications", sql.Named("kod_adres", kod_adres))

	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(
			&data.id,
			&data.adres,
			&data.cw1,
			&data.cw2,
			&data.hw1,
			&data.hw2,
			&data.ed,
			&data.en,
			&data.gas,
			&data.wdate,
			&data.edate,
			&data.gdate,
		)
	}
	if err != nil {
		return err
	}

	account.Id = kod_adres
	account.Adres = data.adres
	account.AccountDevice.Water.Cw1.Prev = data.cw1
	account.AccountDevice.Water.Cw1.Cur = 0
	account.AccountDevice.Water.Cw2.Prev = data.cw2
	account.AccountDevice.Water.Cw2.Cur = 0
	account.AccountDevice.Water.Hw1.Prev = data.hw1
	account.AccountDevice.Water.Hw1.Cur = 0
	account.AccountDevice.Water.Hw2.Prev = data.hw2
	account.AccountDevice.Water.Hw2.Cur = 0
	account.AccountDevice.Water.Date = SqlFormatDateToGo(data.wdate)

	account.AccountDevice.Energie.Ed.Prev = data.ed
	account.AccountDevice.Energie.Ed.Cur = 0
	account.AccountDevice.Energie.En.Prev = data.en
	account.AccountDevice.Energie.En.Cur = 0
	account.AccountDevice.Energie.Date = SqlFormatDateToGo(data.edate)

	account.AccountDevice.Gas.Gas.Prev = data.gas
	account.AccountDevice.Gas.Date = SqlFormatDateToGo(data.gdate)

	account.AccountDevice.Check()

	return nil
}

func SqlFormatDateToGo(sqlDate string) time.Time {
	date := strings.Split(sqlDate, ".")
	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])

	return time.Date(year, time.Month(month), day, 1, 10, 30, 0, time.UTC)
}
