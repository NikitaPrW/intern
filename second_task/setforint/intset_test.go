package intset

import (
	"fmt"
	"testing"
)

/*type Method interface {
	Addtoset() map[int]bool
	Deletefromset() map [int]bool
	Checkforelem() bool
}*/

func TestIntset(t *testing.T) {
	//got := Intsetstruct{}
	t.Run("adding element to set", func(t *testing.T) {
		num := 1
		testmap := intsetstruct{}
		got := testmap.Addtoset(num)
		if _, ok := got.elem[1]; ok {
			fmt.Println("ok adding elem")
		} else {
			fmt.Println("couldn't add elem", num)
		}
	})
	t.Run("deleting elem from set", func(t *testing.T) {
		var testmap = intsetstruct{}
		testmap = testmap.Addtoset(1)
		testmap = testmap.Addtoset(123)
		got := testmap.Deletefromset(123)
		if _, ok := got.elem[123]; ok {
			fmt.Println("smth went wrong, couldn't delete element")
		}
		if _, ok := got.elem[1]; ok == false {
			fmt.Println("smth went wrong, deleted what it shouldn't")
		} else {
			fmt.Println("ok deleting elem")
		}
		t.Run("checking for element", func(t *testing.T) {
			testmap := intsetstruct{}
			testmap = testmap.Addtoset(2)
			got := testmap.Checkforelem(2)
			if got != true {
				fmt.Println("bad chacking for elem")
			}
			got = testmap.Checkforelem(3)
			if got == true {
				fmt.Println("this elem didn't exist")
			}
		})
	})

}

/*func Testinset(t *testing.T) {
	testarray := []struct {
		name string
		method Method
		num int
		want map[int]bool
	} {name: "adding element",method: Addtoset, num: 1, want: {1:true}},
		{name: "deleting element", method: Deletefromset, num: 1, want : {}},
		{name: "check for existance", num: 1, want :false}},
		{name: "adding element", num: 15, want: {15:true}},
}
	var a intsetstruct
	t.Run("test adding elements to set", func addelem(t *testing.T) {
		a.addtoset(a)
	})
	t.Run("test for delete element from set", func deleteelem(t *testing.T){
		a.deletefromset(num)
	})
	t.Run("check for existance in set", func (t *testing.T){
		a.checkforelem(num)
	})
}*/
