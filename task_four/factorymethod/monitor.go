package factorymethod

type monitor struct {
	id         int     //id монитора
	user       int     //id пользователя
	resolution int     //разрешение экрана
	diagonal   float64 //диагональ экрана
}

// рейтинг для монитора, учитываются разрешение и диагональ, рейтинг может быть от 0 до 4
func (m *monitor) Rate() int {
	rate := 0
	if m.resolution > 2073600 {
		rate = rate + 2
	} else if m.resolution == 2073600 {
		rate = rate + 1
	}
	if m.diagonal > 23 && m.diagonal < 27 {
		rate = rate + 1
	} else if m.diagonal >= 27 {
		rate = rate + 2
	}
	return rate
}
