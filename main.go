package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/gcfg.v1"
)

/*Config holds all values read from the config file*/
type Config struct {
	Radius struct {
		Listen string
		Secret string
	}
	Ldap struct {
		Host     string
		User     string
		Password string
		BaseDn   string
		Filter   string
		Secure   bool
	}
}

var config Config

func check(e error, msg string) {
	if e != nil {
		log.Println(msg)
		panic(e)
	}
}

func main() {
	err := gcfg.ReadFileInto(&config, "config.gcfg")
	check(err, "error reading config.gcfg")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	errChan := make(chan error)

	go func() {
		log.Println("[server] waiting for packets...")
		initRadius()
		err := radiusServer.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case <-signalChan:
		log.Println("[syscall] stopping server...")
		radiusServer.Stop()
	case err := <-errChan:
		log.Printf("[error] %v\n", err.Error())
	}
}
