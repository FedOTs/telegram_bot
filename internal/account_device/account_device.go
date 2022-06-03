package account_device

import (
	"telegram_bot/internal/indication"
)

type AccountDevice struct {
	Water   indication.Water
	Energie indication.Energie
	Gas     indication.Gas
}

func (a *AccountDevice) Check() error {
	a.Water.IsActive = true
	a.Energie.IsActive = true
	a.Gas.IsActive = true
	if a.Water.Cw1.Prev < 0 && a.Water.Cw2.Prev < 0 && a.Water.Hw1.Prev < 0 && a.Water.Hw2.Prev < 0 {
		a.Water.IsActive = false
	}
	if a.Energie.Ed.Prev < 0 && a.Energie.En.Prev < 0 {
		a.Energie.IsActive = false
	}
	if a.Gas.Gas.Prev < 0 {
		a.Water.IsActive = false
	}

	return nil
}
