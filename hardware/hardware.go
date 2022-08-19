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
	rand.Seed(time.Now().Unix())
	oneCycleTicker := time.NewTicker(time.Duration(setup.Set.StepCycle) * time.Second)
	for {
		<-oneCycleTicker.C
		element := database.Element{UID: setup.Set.UID, Time: time.Now(),
			Data: database.Data{Temp: rand.Intn(40), Dipl: rand.Intn(100)}}
		database.InElements <- element
		netware.SendElement <- element

	}

}
