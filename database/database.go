package database

import (
	"time"

	"github.com/ruraomsk/ag-server/logger"
)

/*
Хранилище данных
*/
type Data struct {
	Temp int `json:"t"`
	Dipl int `json:"d"`
}
type Element struct {
	UID  int       `json:"uid"`
	Time time.Time `json:"time"`
	Data Data      `json:"data"`
}

var els map[int]Element
var InElements chan Element
var InSayMe chan time.Time
var Out chan []Data

func DataBase() {
	els = make(map[int]Element)
	InElements = make(chan Element, 300)
	InSayMe = make(chan time.Time)
	Out = make(chan []Data)
	go rundb()
}
func rundb() {

	for {
		select {
		case in := <-InElements:
			//Пришли данные
			els[in.UID] = in
			logger.Debug.Printf("%v", in)
		case t := <-InSayMe:
			//Запрос данных
			datas := make([]Data, 0)
			for _, v := range els {
				if v.Time.After(t) {
					datas = append(datas, v.Data)
				}
			}
			Out <- datas
		}
	}
}
