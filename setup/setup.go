package setup

var (
	Set *Setup
)

type Setup struct {
	LogPath     string `toml:"logpath"`
	StepDisplay int    `toml:"display"` //Шаг с которым обновляем информацию на дисплее (секунды)
	StepSend    int    `toml:"send"`    //Шаг с которым шлем сообщения (секунды)
	StepCycle   int    `toml:"cycle"`   //Шаг цикла измерений
	IP          string `toml:"ip"`      //Адрес рассылки в видe "192.168.0.1"
	Port        int    `toml:"port"`
	UID         int    `toml:"uid"` //Мой уникальный номер
}
