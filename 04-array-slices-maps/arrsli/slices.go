package arrsli

import (
	"fmt"
)

func ImprimeSlices() {
	var sliceEnteros1 []int32 = []int32{10, 20, 30, 40, 50}
	fmt.Println("Slice de enteros:", sliceEnteros1)
	sliceEnteros1 = append(sliceEnteros1, 60, 70, 80, 90, 100)
	fmt.Println("Slice de enteros después de agregar elementos:", sliceEnteros1)
}
