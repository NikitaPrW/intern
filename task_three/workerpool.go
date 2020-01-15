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
	work := make(map[int]int)
	goroutineresult := make(chan task)
	for i := 0; i < numOfWorkers; i++ {
		a := time.Duration(rand.Intn(2) + 1)
		go func(timetoSleep time.Duration) {
			goroutineresult <- task{i, i + i}
			time.Sleep(a * time.Second)
		}(a)
	}
	for i := 0; i < numOfWorkers; i++ {
		task := <-goroutineresult
		work[task.id] = task.result
	}
	for i := 0; i < numOfWorkers; i++ { //прораб проверяет результаты, если все в порядке, то переходит к следуюущему, если нет-выводит информацию
		if work[i] != i+i {
			fmt.Println("Incorrect result", work[i], i+i)
		}
	}
	fmt.Println("////", work)
	return work
}
