package hardware

import (
	"math/rand"
	"time"

	"github.com/ruraomsk/netudp/database"
	"github.com/ruraomsk/netudp/netware"
	"github.com/ruraomsk/netudp/setup"
)

func HardWare() {
	go runHW()
}
func runHW() {
	//Эмулируем сбор от датчиков
	rand.Seed(time.Now().Unix())
	oneCycleTicker := time.NewTicker(time.Duration(setup.Set.StepCycle) * time.Second)
	for {
		<-oneCycleTicker.C
		element := database.Element{UID: setup.Set.UID, Time: time.Now(),
			Data: database.Data{Temp: rand.Intn(40), Dipl: rand.Intn(100)}}
		database.InElements <- element //В базу данных слем для верности если вообще вся сеть упала
		netware.SendElement <- element //Рассылаем

	}

}
