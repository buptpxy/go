package main

/*
Q5. 平均值
	编写一个函数用于计算一个 float64 类型的 slice 的平均值。
*/
import (
	"fmt"
)

func average(fs []float64) float64 { //返回值带变量名时要写括号，只有类型时可不写括号func average(fs []float64) (avg float64){}
	var sum float64
	length := float64(len(fs))
	for _, v := range fs {
		sum += v
	}
	return sum / length
}
func main() {
	fmt.Printf("平均值为：%f \n", average([]float64{3.5, 6.7, 8.9, 3.4}))
}
