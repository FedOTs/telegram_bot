package config

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	want := Config{
		Database: Database{
			Host:     "host",
			User:     "user",
			Password: "password",
			DbName:   "dbname",
		},
	}

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	got, err := LoadConfiguration(filepath.Dir(filepath.Dir(dir)) + string(os.PathSeparator) + "configs" + string(os.PathSeparator) + "config_example.json")

	if err != nil {
		t.Errorf("Someone error %q", err)
	}

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}
