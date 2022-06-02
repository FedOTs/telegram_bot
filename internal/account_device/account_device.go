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

/*
func (a *AccountDevice) Get() (indication.Water, indication.Energie, indication.Gas) {
	return a.Water, a.Energie, a.Gas
}

func (a *AccountDevice) Set(water indication.Water, energie indication.Energie, gas indication.Gas) error {
	a.Water = water
	a.Energie = energie
	a.Gas = gas
	return nil
}

func (a *AccountDevice) SetDbWater(cw1 int, cw2 int, hw1 int, hw2 int, wdate time.Time) error {
	isNull := false
	if cw1 < 0 && cw2 < 0 && hw1 < 0 && hw2 < 0 {
		isNull = true
	}
	err := a.Water.Set(indication.WaterCw, cw1, 0, cw2, 0, wdate, isNull)

	if err != nil {
		return err
	}

	err = a.Water.Set(indication.WaterHw, hw1, 0, hw2, 0, wdate, isNull)

	if err != nil {
		return err
	}

	return nil
}

func (a *AccountDevice) SetDbEnergie(ed int, en int, edate time.Time) error {
	isNull := false
	if ed < 0 && en < 0 {
		isNull = true
	}
	err := a.Energie.Set(indication.EnergieEd, ed, 0, edate, isNull)
	if err != nil {
		return err
	}

	err = a.Energie.Set(indication.EnergieEn, en, 0, edate, isNull)
	if err != nil {
		return err
	}

	return nil
}

func (a *AccountDevice) SetDbGas(gas int, gdate time.Time) error {
	isNull := false
	if gas < 0 {
		isNull = true
	}
	err := a.Gas.Set(indication.GasGas, gas, 0, gdate, isNull)
	if err != nil {
		return err
	}
	return nil
}
*/
