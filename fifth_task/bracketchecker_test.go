package bracketchecker

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"testing"
)

func TestBracketcheker(t *testing.T) {
	//want := make(map[string]bool)
	//got := make(map[string]bool)
	var errcounter int = 0
	var stringsfortest = []string{"{[(){]}", "{[()]}", "[](){()[]}", "{[])"}
	got := make([]bool, len(stringsfortest))
	var want = []bool{false, true, true, false}
	t.Run("bracket checker", func(t *testing.T) {
		for i := 0; i < len(stringsfortest); i++ {
			got[i] = isValid(&stringsfortest[i])
		}
		for i := 0; i < len(want); i++ {
			if got[i] != want[i] {
				log.Errorf("failed")
				errcounter++
			}
		}
	})
	if errcounter != 0 {
		fmt.Println("test failed")
	} else {
		fmt.Println("everything ok")
	}
}
