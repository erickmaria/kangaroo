package main

import (
	"fmt"

	"github.com/erickmaria/kangaroo/core/pkg/handler"
	"github.com/erickmaria/kangaroo/core/pkg/logger"
	"github.com/erickmaria/kangaroo/core/pkg/profile"
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

	logger.Log(logger.INFO, "this is a log info exemple")
	logger.Log(logger.WARN, "this is a log warn exemple")
	logger.Log(logger.ERROR, "this is a log error exemple")

	var app Application

	err := profile.Init(&app, "configs/", "")

	if handler.Validate(err) {
		panic(err)
	}

	fmt.Println(app)

	server.NewServer(":9999").Listen()
}
