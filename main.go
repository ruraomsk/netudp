package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/netudp/database"
	"github.com/ruraomsk/netudp/display"
	"github.com/ruraomsk/netudp/hardware"
	"github.com/ruraomsk/netudp/netware"
	"github.com/ruraomsk/netudp/setup"
)

var (
	//go:embed config
	config embed.FS
)
var err error

func init() {
	// Вычитываем настройки
	setup.Set = new(setup.Setup)
	if _, err := toml.DecodeFS(config, "config/config.toml", &setup.Set); err != nil {
		fmt.Println("Dissmis config.toml")
		os.Exit(-1)
		return
	}
	//Создаем каталог для логов
	os.MkdirAll(setup.Set.LogPath, 0777)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := logger.Init(setup.Set.LogPath); err != nil {
		log.Panic("Error logger system", err.Error())
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("need uid for start!")
		return
	}
	setup.Set.UID, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("need uid for start!")
		return
	}
	database.DataBase() //Типа базы данных для хранения принятой информации
	netware.NetWare()   // Собственно обмен с сетью
	hardware.HardWare() // Типа сбора информации от датчиков
	display.Display()   //Типа индикатор

	fmt.Printf("Client %d start\n", setup.Set.UID)
	logger.Info.Printf("Client %d start\n", setup.Set.UID)
	//Ждем завершения
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP)
	<-c
	fmt.Printf("Client %d stopped\n", setup.Set.UID)
	logger.Info.Printf("Client %d stopped\n", setup.Set.UID)
}
