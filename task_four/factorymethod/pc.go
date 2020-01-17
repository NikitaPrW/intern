package factorymethod

type pc struct {
	id         int //id компютера
	user       int //id пользователя
	processor  string
	frequency  int //Ghz
	cores      int //number of cores
	department string
	ram        int //GB
	storage    int //GB
}

//система рейтингования, считает рейтинг на основе тактовой частоты, количества ядер, объема оперативной памяти, объема жестокого диска, рейтинг от 0 до 5
func (p *pc) Rate() int {
	rate := 0
	if p.frequency > 2 && p.frequency < 3 {
		rate = rate + 1
	} else if p.frequency > 3 {
		rate = rate + 2
	}
	if p.cores >= 4 {
		rate = rate + 1
	}
	if p.ram >= 16 {
		rate = rate + 1
	}
	if p.storage > 256 {
		rate = rate + 1
	}
	return rate
}
