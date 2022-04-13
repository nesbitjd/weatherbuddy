package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"weatherbuddy/internal/wbclient"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	var c wbclient.Client

	err := c.ParseConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.DGo.Dg.AddHandler(c.Handler)

	err = c.DGo.Dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	c.DGo.Dg.Close()
}
