package account

import (
	"telegram_bot/internal/account_device"
)

type Account struct {
	Id            int
	Adres         string
	AccountDevice account_device.AccountDevice
}

func NewAccount() *Account {
	return &Account{}
}
