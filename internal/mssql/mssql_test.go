package mssql

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"telegram_bot/internal/account"
	"telegram_bot/internal/config"
	"testing"
)

func TestMssql(t *testing.T) {
	dir, err := os.Getwd()
	got, err := config.LoadConfiguration(filepath.Dir(filepath.Dir(dir)) + string(os.PathSeparator) + "configs" + string(os.PathSeparator) + "config.json")

	if err != nil {
		t.Errorf("Someone error %q", err)
	}

	account := account.Account{Id: 6998559}

	connString := got.GetConnString()
	fmt.Println(connString)

	ctx := context.Background()

	err = GetAccountDeviceData(ctx, connString, &account)

	if err != nil {
		t.Errorf("Someone error %q", err)
	}

	fmt.Println(account)
}
