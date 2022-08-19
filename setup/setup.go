package setup

var (
	Set *Setup
)

type Setup struct {
	LogPath     string `toml:"logpath"`
	StepDisplay int    `toml:"display"` //Шаг с которым обновляем информацию на дисплее (секунды)
	StepCycle   int    `toml:"cycle"`   //Шаг цикла измерений
	IP          string `toml:"ip"`      //Адрес рассылки в видe "192.168.0.1"
	Port        int    `toml:"port"`
	UID         int    //Мой уникальный номер
}
