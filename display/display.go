package display

import (
	"fmt"
	"time"

	"github.com/ruraomsk/netudp/database"
	"github.com/ruraomsk/netudp/setup"
)

func Display() {
	go runDisplay()
}
func runDisplay() {
	//Создаем тикер для отображения
	oneTicker := time.NewTicker(time.Duration(setup.Set.StepDisplay) * time.Second)
	start := time.Now() //время начала диапазона
	for {
		<-oneTicker.C
		database.InSayMe <- start //Просим базу дать массив данных
		datas := <-database.Out
		//Считаем среднюю
		tm := 0.0
		bm := 0.0
		for _, v := range datas {
			tm += float64(v.Temp)
			bm += float64(v.Dipl)
		}
		if len(datas) != 0 {
			tm = tm / float64(len(datas))
			bm = bm / float64(len(datas))
		}
		fmt.Printf("%v\t%d\t%f\t%f\n", start, len(datas), tm, bm)
		//Меняем начало интервала
		start = time.Now()
	}
}
