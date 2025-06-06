package config

import (
	"os"
	"testing"
)

func TestLoadConfigDefaultOnMissing(t *testing.T) {
	_, err := LoadConfig("nonexistent.json")
	if err == nil {
		t.Errorf("expected error on missing file")
	}
}

func TestLoadConfigValues(t *testing.T) {
	tmp, err := os.CreateTemp("", "cfg*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	data := `{"tower_damage":3,"base_health":5}`
	if _, err := tmp.Write([]byte(data)); err != nil {
		t.Fatal(err)
	}
	tmp.Close()
	cfg, err := LoadConfig(tmp.Name())
	if err != nil {
		t.Fatalf("load config: %v", err)
	}
	if cfg.TowerDamage != 3 || cfg.BaseHealth != 5 {
		t.Errorf("unexpected values %v", cfg)
	}
}
