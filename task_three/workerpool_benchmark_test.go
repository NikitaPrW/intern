package workerpool

import (
	"log"
	"math/rand"
	"os"
	"testing"
)

func BenchmarkWorkerpool(b *testing.B) {
	var iscorrect bool = true
	numofworkers := rand.Intn(10)
	result := make(map[int]int)
	b.Run("Run the task", func(b *testing.B) {
		result = workerpool(numofworkers)
	})
	printresult := func(b *testing.B, iscorrect bool) {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
			log.Println("Counldn't open the file")
		}
		defer f.Close()
		log.SetOutput(f)
		switch iscorrect {
		case true:
			log.Println("Correct")

		default:
			log.Println("False")

		}
		return
	}
	b.Run("Check for completed tasks", func(b *testing.B) {
		for i := 0; i < numofworkers; i++ {
			if i+i != result[i] {
				iscorrect = false
			}
		}
		printresult(b, iscorrect)
	})

}
