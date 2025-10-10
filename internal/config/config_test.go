package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T){
	home, _ := os.MkdirTemp("", "gatorhome")
	defer os.RemoveAll(home)

	configPath := filepath.Join(home, ".gatorconfig.json")
	content := `{"db_url":"postgres://test","current_user_name":"mec"}`
	os.WriteFile(configPath, []byte(content), 0644)
	os.Setenv("HOME", home)

	cfg, err := Read()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.DBUrl != "postgres://test" {
		t.Errorf("expected db_url postgres://test, got %s", cfg.DBUrl)
	}
	if cfg.CurrentUserName != "mec" {
		t.Errorf("expected user mec, got %s", cfg.CurrentUserName)
	}
}