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

var els map[int]Element     //Собствеено хранилище данных
var InElements chan Element //Прием данных для хранения
var InSayMe chan time.Time  //Тут приходят запросы на выборку
var Out chan []Data         //Тут отправляем данные

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
			//Пришли данные со стороны
			els[in.UID] = in
			logger.Debug.Printf("%v", in)
		case t := <-InSayMe:
			//Запрос данных в запросе время начала сбора
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
