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
	Db     DBconfig
	// JWTConfig   JWTConfig
	EmailConfig EmailConfig
	ShowDocs    bool
	Token       TokenConfig
}

type DBconfig struct {
	Addr         string
	Port         string
	Name         string
	User         string
	Password     string
	MaxOpenConns int
}
type EmailConfig struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	SenderEmail  string
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

//	type JWTConfig struct {
//		SecretKey       string
//		TokenExpiration int64
//	}
type TokenConfig struct {
	Secret string
}
