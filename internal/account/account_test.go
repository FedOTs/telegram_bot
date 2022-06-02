package account

import (
	"telegram_bot/internal/account_device"
	"telegram_bot/internal/indication"
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	id := 12313
	cw1 := indication.Indication{Prev: 1, Cur: 2}
	date := time.Now().UTC()
	adres := "asdasasd"
	water := indication.Water{Cw1: cw1, Cw2: cw1, Hw1: cw1, Hw2: cw1, Date: date, IsActive: true}
	energie := indication.Energie{Ed: cw1, En: cw1, Date: date, IsActive: true}
	gas := indication.Gas{Gas: cw1, Date: date, IsActive: true}
	account := NewAccount()
	ad := account_device.AccountDevice{Water: water, Energie: energie, Gas: gas}

	t.Run("Set account data", func(t *testing.T) {
		err := account.Set(id, adres, ad)
		if err != nil {
			t.Errorf("Someone error %q", err)
		}
	})
	/*
		t.Run("Get account data", func(t *testing.T) {
			err := account.Set(id, adres, ad)
			if err != nil {
				t.Errorf("Someone error %q", err)
			}

			want := NewAccount()
			ad := account_device.AccountDevice{
				Water: indication.Water{
					Cw1:  indication.Indication{Prev: 1, Cur: 2},
					Cw2:  indication.Indication{Prev: 1, Cur: 2},
					Hw1:  indication.Indication{Prev: 1, Cur: 2},
					Hw2:  indication.Indication{Prev: 1, Cur: 2},
					Date: date, Error: nil},
				Energie: indication.Energie{
					Ed:   indication.Indication{Prev: 1, Cur: 2},
					En:   indication.Indication{Prev: 1, Cur: 2},
					Date: date, Error: nil},
				Gas: indication.Gas{
					Gas:  indication.Indication{Prev: 1, Cur: 2},
					Date: date, Error: nil},
			}
			want.id = 12313
			want.adres = "asdasasd"
			want.accountDevice = ad
			got := NewAccount()
			id_, adres_, ad_ := account.Get()
			got.id = id_
			got.adres = adres_
			got.accountDevice = ad_

			if *got != *want {
				t.Errorf("got %v, want %v", got, want)
			}
		})*/
}
