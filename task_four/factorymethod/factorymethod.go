package factorymethod

import "log"

//Productcreator интерфейс для фабрик
type Productcreator interface {
	CreateProduct(action string) Producter
}

//Producter интерфейс для продуктов
type Producter interface {
	Rate() int
}

//фабрика
type concreteCreator struct {
}

//NewCreator инициализация фабрики
func NewCreator() Productcreator {
	return &concreteCreator{}
}

func (p *concreteCreator) CreateProduct(action string) Producter {
	var product Producter
	switch action {
	case "pc":
		product = &pc{id: 1, user: 1, processor: "i5", frequency: 2, cores: 4, department: "hr", ram: 8, storage: 200}

	case "monitor":
		product = &monitor{id: 1, user: 1, resolution: 2073600, diagonal: 24}
	default:
		log.Fatalln("Unknown Action")
	}
	return product
}
