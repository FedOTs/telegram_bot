package account

import (
	"telegram_bot/internal/account_device"
)

type Account struct {
	Id            int
	Adres         string
	AccountDevice account_device.AccountDevice
}

/*
func (a *Account) Set(id int, adres string, accountDevice account_device.AccountDevice) error {
	a.id = id
	a.adres = adres
	a.accountDevice = accountDevice
	return nil
}

func (a *Account) SetDb(id int, adres string, cw1 int, cw2 int, hw1 int, hw2 int, wdate time.Time, ed int, en int, edate time.Time, gas int, gdate time.Time) error {
	a.id = id
	a.adres = adres
	a.accountDevice.Water.Cw1.Prev = accountDevice{Water:}
	return nil
}

func (a *Account) Get() (id int, adres string, accountDevice account_device.AccountDevice) {
	return a.id, a.adres, a.accountDevice
}
*/
func NewAccount() *Account {
	return &Account{}
}
