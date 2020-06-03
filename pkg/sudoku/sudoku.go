package sudoku

import (
	"github.com/wailsapp/wails"
)

type MyStruct struct {
	log *wails.CustomLogger
}

// WailsInit .
func (d *MyStruct) WailsInit(runtime *wails.Runtime) error {
	d.log = runtime.Log.New("MyStruct")
	return nil
}
func (d *MyStruct) DisplayArray() [9][9]int {
	ListNumbers := [81]int{2,5,4,9,1,7,3,8,6,
		6,8,9,2,0,5,4,0,7,
		7,3,0,6,4,8,2,5,9,
		9,4,0,8,7,0,5,2,3,
		0,7,2,3,9,4,0,6,8,
		3,0,8,5,2,0,9,7,4,
		8,6,3,0,5,9,7,4,2,
		0,9,7,4,6,2,8,3,5,
		4,2,0,7,8,3,6,9,0,
	}

	var result [9][9]int
	x:= 0
	for x < len(ListNumbers) {

		for i := 0; i < 9; i++ {

			for j := 0; j < 9; j++ {
				result[i][j] = ListNumbers[x]
				//fmt.Print(result[i][j])
				x++
			}
			//fmt.Println();
		}


	}
	//fmt.Print(x);
	/*mensaje := "Hola otra vez"
	return mensaje*/
	return result
}
