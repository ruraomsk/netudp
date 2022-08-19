package netware

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/netudp/database"
	"github.com/ruraomsk/netudp/setup"
)

var SendElement chan database.Element
var ReciveElement chan database.Element

func NetWare() {
	SendElement = make(chan database.Element)
	ReciveElement = make(chan database.Element, 100)
	go sendUDP()
	go reciveUDP()
}
func sendUDP() {
	//Тут пересылаем свои данные
	ip := fmt.Sprintf("%s:%d", setup.Set.IP, setup.Set.Port)
	con, err := net.Dial("udp", ip)
	if err != nil {
		logger.Error.Printf("dial %s %s", ip, err.Error())
		return
	}
	defer con.Close()

	for {
		in := <-SendElement
		buf, err := json.Marshal(in)
		if err != nil {
			logger.Error.Printf("%v %s", in, err.Error())
			return
		}
		_, err = con.Write(buf)
		if err != nil {
			logger.Error.Printf("send udp %s", err.Error())
			return
		}
	}
}
func reciveUDP() {
	//Ждем UDP от всех кто ни попадя
	addr := net.UDPAddr{Port: setup.Set.Port, IP: net.ParseIP("0.0.0.0")}
	con, err := net.ListenUDP("udp", &addr)
	if err != nil {
		logger.Error.Printf("listen %s:%d %s", setup.Set.IP, setup.Set.Port, err.Error())
		return
	}
	defer con.Close()
	//Размер буфера можно поднять если мало
	buffer := make([]byte, 2048)
	for {
		for i := 0; i < len(buffer); i++ {
			buffer[i] = 0
		}
		_, remote_adr, err := con.ReadFromUDP(buffer)
		if err != nil {
			logger.Error.Printf("read from %s %s", remote_adr.String(), err.Error())
			return
		}
		var element database.Element
		buf := make([]byte, 0)
		for _, v := range buffer {
			if v == 0 {
				break
			}
			buf = append(buf, v)
		}
		err = json.Unmarshal(buf, &element)
		if err != nil {
			logger.Error.Printf("%v %s", string(buf), err.Error())
			return
		}
		//И в базу данных. Тут конечно можно прверить на свое ли это сообщение и не слать
		//его в базу данных
		database.InElements <- element
	}
}
