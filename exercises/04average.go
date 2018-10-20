package main

/*
Q4. 平均值
	编写计算一个类型是 float64 的 slice 的平均值的代码。
*/
import (
	"fmt"
)

func main() {
	floatslice := []float64{4.4, 24.2, 55.5, 67.3}
	var sum float64
	length := float64(len(floatslice))
	for _, v := range floatslice {
		sum += v
	}
	fmt.Printf("平均值为： %f \n", sum/length)
}
