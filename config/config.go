package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sarahrajabazdeh/DreamPilot/dreamerr"
)

var Config Configuration

type Configuration struct {
	Server ServerConfig
	// Db contains db connection configuration.
	Db        DBconfig //nolint:stylecheck
	JWTConfig JWTConfig
}

type DBconfig struct {
	Addr         string // database address or hostname
	Port         string // database port
	Name         string // database name
	User         string // database user
	Password     string // database user password
	MaxOpenConns int    // max number of open connections

}

type ServerConfig struct {
	// The port to listen
	Port string
}

func Read() {
	var raw []byte
	var err error

	if raw, err = os.ReadFile("config.json"); err != nil {
		if raw, err = os.ReadFile("/app/config.json"); err != nil {
			err = fmt.Errorf("could not read config: %w", err)
			dreamerr.LogFatalError(err.Error())
		}
	}

	if err = json.Unmarshal(raw, &Config); err != nil {
		err = fmt.Errorf("could not read config: %w", err)
		dreamerr.LogFatalError(err.Error())
	}
}

type JWTConfig struct {
	SecretKey       string
	TokenExpiration int64
	// Add other JWT-specific configuration options as needed
}
