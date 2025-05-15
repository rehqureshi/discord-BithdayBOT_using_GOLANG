package main

import (
	"fmt"

	"github.com/rehqureshi/go-pingmod-Discord/bot"
	"github.com/rehqureshi/go-pingmod-Discord/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})
	return
}
