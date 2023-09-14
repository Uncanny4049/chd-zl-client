package types

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	ZLDBPath string `json:"old_zl_db_path"`
	DBPath   string `json:"new_zl_db_path"`
}

func ReadConfig() *ServerConfig {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		config := ServerConfig{
			ZLDBPath: "./.build/old.db",
			DBPath:   "./.build/new.db",
		}
		marshal, _ := json.Marshal(config)
		os.WriteFile("./config.json", marshal, os.FileMode(0777))
		return &config
	} else {
		t := ServerConfig{}
		err = json.Unmarshal(file, &t)
		return &t
	}
}
