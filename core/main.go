package main

import (
	"github.com/erickmaria/kangaroo/core/pkg/server"
)

type Application struct {
	App struct {
		Addr      string `yaml:"addr"`
		Port      int    `yaml:"port"`
		Internval string `yaml:"internval"`
		Info      struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"info"`
		Ips []string `yaml:"ips"`
	} `yaml:"app"`

	Log string
}

func main() {

	var app Application
	server.NewKangarooServer(":9999").
		SetProperties(&app, "configs/", "").
		SetLoggerModule("Core")

	// logger.Log(logger.INFO, "this is a log info exemple")
	// logger.Log(logger.WARN, "this is a log warn exemple")
	// logger.Log(logger.ERROR, "this is a log error exemple")

	// fmt.Println(profile.GetProperties("app.i"))

	// svr.Listen()
}
