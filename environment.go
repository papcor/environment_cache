package environment_cache

import (
	"log"
	"os"
)

func CacheEnvironmentVariable(name string) (f func() string) {
	value := os.Getenv(name)
	f = func() string {
		return value
	}
	return
}

func CacheRequiredEnvironmentVariable(name string) (f func() string) {
	check := CacheEnvironmentVariable(name)
	if check() == "" {
		log.Fatal("$" + name + " must be set.")
	} else {
		log.Print("$" + name + " set.")
	}
	return check
}
