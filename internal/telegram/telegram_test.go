package telegram

import (
	"os"
	"path/filepath"
	"telegram_bot/internal/config"
	"testing"
)

func TestTelegram(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("dir err %s", err)
	}
	conf, err := config.LoadConfiguration(filepath.Dir(filepath.Dir(dir)) + string(os.PathSeparator) + "configs" + string(os.PathSeparator) + "config.json")

	_, err = StartBot(conf.GetTgTokenString(), conf.GetTgWebHookString())

	if err != nil {
		t.Errorf("StartBot error %s", err)
	}

}
