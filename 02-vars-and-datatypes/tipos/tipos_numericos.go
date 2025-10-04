package tipos

import (
	"fmt"
	"math"
)

func ImprimeTiposNumericos() {
	fmt.Println("int8  :", math.MinInt8, "a", math.MaxInt8)
	fmt.Println("int16 :", math.MinInt16, "a", math.MaxInt16)
	fmt.Println("int32 :", math.MinInt32, "a", math.MaxInt32)
	fmt.Println("int64 :", math.MinInt64, "a", math.MaxInt64)

	fmt.Println("uint8  :", 0, "a", math.MaxUint8)
	fmt.Println("uint16 :", 0, "a", math.MaxUint16)
	fmt.Println("uint32 :", 0, "a", math.MaxUint32)
	fmt.Println("uint64 :", 0, "a", uint64(math.MaxUint64))

	fmt.Println("float32 :", -math.MaxFloat32, "a", math.MaxFloat32)
	fmt.Println("float64 :", -math.MaxFloat64, "a", math.MaxFloat64)
}
