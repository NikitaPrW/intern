package workerpool

import (
	"fmt"
	"math/rand"
	"time"
)

type task struct {
	id     int
	result int
}

func workerpool(numOfWorkers int) map[int]int {
	//var wg sync.WaitGroup
	//wg.Add(numOfWorkers)
	work := make(map[int]int)
	goroutineresult := make(chan task)
	//goroutineresult := make(chan task, numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		rand.Seed(time.Now().UnixNano())
		a := time.Duration(rand.Intn(3) + 1)
		go func(timetoSleep time.Duration, index int) { //создается го рутина
			time.Sleep(timetoSleep * time.Second)
			goroutineresult <- task{index, index + index}
			//wg.Done()
		}(a, i) //в го рутину передаются a,i, чтобы в процессе работы го рутины не поменялось их значение для го рутины
	}
	//wg.Wait()
	for i := 0; i < numOfWorkers; i++ { //запись из канала в map
		task := <-goroutineresult
		work[task.id] = task.result
	}
	for i := 0; i < numOfWorkers; i++ { //прораб проверяет результаты, если все в порядке, то переходит к следуюущему, если нет-выводит информацию
		if work[i] != i+i {
			fmt.Println("Incorrect result", work[i], i+i)
		}
	}
	fmt.Println("////////////////////////////////////////////////")
	fmt.Println(work)
	return work
}
