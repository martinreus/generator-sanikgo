package logging

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"os"
	"test/pkg/env"
)

func Configure() {
	log.SetHandler(json.New(os.Stderr))
	level, err := log.ParseLevel(env.GetStringOrDefault("LEVEL", "info"))
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)
}