/* 【此代码中用到了重定义比较器的comparator】
给定一个N行3列二维数组， 每一行表示有一座大楼， 一共有N座大楼。
所有大楼的底部都坐落在X轴上， 每一行的三个值(a,b,c)代表每座大楼的从(a,0)点开始， 到(b,0)点结束， 高度为c。
输入的数据可以保证a<b,且a， b， c均为正数。 大楼之间可以有重合。请输出整体的轮廓线。
例子： 给定一个二维数组
[[1, 3, 3],
[2, 4, 4],
[5, 6, 1]]
输出为轮廓线
[[1, 2, 3],
[2, 4, 4],
[5, 6, 1]]

【思路】轮廓线的变化一定是由于这些大楼在当前位置的最大高度发生了变化。故可以用一个搜索二叉树treeMap（不能用heap，因为没法知道哪个key是否已存在）生成所有大楼的最大高度，
再使用一个treeMap（不能用普通的map，因为没法实时知道map中的最大值）把每个位置和它对应的当前最大高度记录下来，最后遍历这个t，记录下最大高度有变化的位置，就是轮廓边界。
*/
package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"sort"
)

type Build struct {
	pos    int
	height int
	up     bool
}
type builds []Build

func (arr builds) Len() int {
	return len(arr)
}

func (arr builds) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr builds) Less(i, j int) bool {
	if arr[i].pos < arr[j].pos {
		return true
	} else if arr[i].pos == arr[j].pos && arr[i].up {

		return true

	} else {
		return false
	}
}

func buildingOutLine(arr [][]int) [][]int {
	buildings := make(builds, len(arr)*2)
	for i := 0; i < len(arr); i++ {
		buildings[2*i].pos = arr[i][0]
		buildings[2*i].height = arr[i][2]
		buildings[2*i].up = true
		buildings[2*i+1].pos = arr[i][1]
		buildings[2*i+1].height = arr[i][2]
		buildings[2*i+1].up = false
	}
	sort.Sort(buildings)
	hMap := treemap.NewWithIntComparator()
	phMap := treemap.NewWithIntComparator()
	for i := 0; i < len(buildings); i++ {
		if buildings[i].up {
			if v, t := hMap.Get(buildings[i].height); t == false {
				hMap.Put(buildings[i].height, 1)
			} else {
				hMap.Put(buildings[i].height, v.(int)+1)
			}
		} else {
			if v, _ := hMap.Get(buildings[i].height); v == 1 {
				hMap.Remove(buildings[i].height)
			} else {
				hMap.Put(buildings[i].height, v.(int)-1)
			}
		}
		if hMap.Empty() {
			phMap.Put(buildings[i].pos, 0)
		} else {
			maxH, _ := hMap.Max()
			phMap.Put(buildings[i].pos, maxH)
		}
	}
	res := make([][]int, phMap.Size())
	posH := phMap.Keys()
	start := 0
	height := 0
	for i, pos := range posH {
		res[i] = make([]int, 3)
		if maxH, _ := phMap.Get(pos.(int)); height != maxH {

			res[i][0] = start
			res[i][1] = pos.(int)
			res[i][2] = height
			start = pos.(int)
			height = maxH.(int)
		}
	}
	res1 := [][]int{}
	for i := 0; i < len(res); i++ {
		if res[i][2] != 0 {
			res1 = append(res1, res[i])
		}
	}
	return res1
}

func main() {
	buildings := [][]int{
		{1, 3, 3},
		{2, 4, 4},
		{5, 6, 1},
	}
	res := buildingOutLine(buildings)
	fmt.Println(res)
}
