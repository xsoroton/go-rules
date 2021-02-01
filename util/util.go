package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func FatalUnlessUnmarshal(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	FatalOnError(err, "fail to unmarshal")
}

func FatalOnError(err error, args ...interface{}) {
	if err == nil {
		return
	}
	if len(args) == 0 {
		log.Fatal(err.Error())
	}
	context := fmt.Sprint(args...)
	log.Fatalf("%s: %s", context, err.Error())
}

func EnvString(key string, dflt string) (setting string) {
	setting = os.Getenv(key)
	if setting == "" {
		setting = dflt
	}
	return
}
