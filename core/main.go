package main

import "github.com/erickmaria/kangaroo/core/pkg/logger"

func main() {

	logger.Log(logger.INFO, "this is a log info exemple")
	logger.Log(logger.WARN, "this is a log warn exemple")
	logger.Log(logger.ERROR, "this is a log error exemple")
}
