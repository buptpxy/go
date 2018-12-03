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

【思路】轮廓线的变化一定是由于这些大楼在当前位置的最大高度发生了变化。故可以用一个搜索二叉树记录所有大楼的最大高度。
1. 把每栋楼的信息都拆解为上的信息和下的信息，比如[1,3,3]可拆解为{位置：1，高度：3，状态：up}，{位置：3，高度：3，状态：down}
2. 把所有得到的这些信息按位置排序，位置一样时状态为up的排在状态为down的前面
3. 遍历排序结果，把<高度，出现次数>存入heightMap中，以便计算最大高度。并把当前位置的<位置，最大高度>存入posMap中
4. 对heightMap的处理：当前位置状态为up，则往heightMap里面添加东西：当前高度未出现过，则令次数为1；当前高度出现过，则令次数加1；
					当前位置状态为down，则往heightMap里面删除东西：当前高度次数为1，则删除此记录；当前高度次数大于1，则次数减1；
5. 对posMap的处理：遍历posMap,如果当前pos的高度为0，则可记录轮廓的起始位置和高度
							如果当前pos的高度不为0且不等于当前最高高度，则当前位置即为一个轮廓的终止位置，当前高度为轮廓高度
*/
package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"sort"
)

type Node struct {
	pos    int
	height int
	up     bool
}
type Nodes []Node

func (n Nodes) Len() int {
	return len(n)
}
func (n Nodes) Less(i, j int) bool {
	if n[i].pos < n[j].pos {
		return true
	} else if n[i].pos == n[j].pos && n[i].up {
		return true
	} else {
		return false
	}
}
func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func buildingOutLine(building [][]int) [][]int {
	//生成buildings数组
	buildings := make(Nodes, len(building)*2)
	for i := 0; i < len(building); i++ {
		buildings[2*i] = Node{building[i][0], building[i][2], true}
		buildings[2*i+1] = Node{building[i][1], building[i][2], false}
	}
	//对buildings数组排序
	sort.Sort(buildings)
	//生成heightMap和posMap
	heightMap := treemap.NewWithIntComparator()
	posMap := treemap.NewWithIntComparator()
	//对heightMap的处理
	for i := 0; i < len(buildings); i++ {

		if buildings[i].up { //当前位置状态为up，则往heightMap里面添加东西
			if k, v := heightMap.Get(buildings[i].height); v == false { //当前高度未出现过，则令次数为1
				//if buildings[i].height>maxH.(int) {//当前高度大于最大高度则更新posMap；
				//	posMap.Put(buildings[i].pos,buildings[i].height)
				//}
				heightMap.Put(buildings[i].height, 1)
			} else { //当前高度出现过，则令次数加1；
				heightMap.Put(buildings[i].height, k.(int)+1)
			}
		} else { //当前位置状态为down，则往heightMap里面删除东西
			if k, _ := heightMap.Get(buildings[i].height); k == 1 { //当前高度次数为1，则删除此记录
				//if buildings[i].height==maxH.(int) {//删掉的为最大高度则更新posMap
				//	posMap.Put(buildings[i].pos,buildings[i].height)
				//}
				heightMap.Remove(buildings[i].height)
			} else { //当前高度次数大于1，则次数减1；
				heightMap.Put(buildings[i].height, k.(int)-1)
			}
		}
		//就算最大高度没有发生变化也添加到posMap中，然后在对posMap的遍历中去处理。
		maxH, _ := heightMap.Max()
		if heightMap.Empty() {
			posMap.Put(buildings[i].pos, 0)
		} else {
			posMap.Put(buildings[i].pos, maxH.(int))
		}
	}
	//对posMap的处理,遍历posMap
	res := make([][]int, posMap.Size())

	start := 0
	height := 0
	posH := posMap.Keys()
	for count, pos := range posH {
		res[count] = make([]int, 3)
		maxh, _ := posMap.Get(pos.(int))
		if maxh != height {
			//如果当前pos的高度不等于当前最高高度，则当前位置即为一个轮廓的终止位置，当前高度为轮廓高度
			res[count][0] = start
			res[count][1] = pos.(int)
			res[count][2] = height
			//如果当前pos的高度不等于当前最高高度，则可记录轮廓的起始位置和高度
			start = pos.(int)
			height = maxh.(int)
		}
	}
	//如果去掉高度为0的轮廓，则需要下面的操作
	res1 := [][]int{}
	for i := 0; i < len(res); i++ {
		if res[i][2] != 0 {
			res1 = append(res1, res[i])
		}
	}
	return res1
}
func main() {
	//buildings := Nodes{{4, 3, true}, {5, 3, false}, {1, 4, true}, {4, 4, false}}
	//sort.Sort(buildings)
	//fmt.Println(buildings)
	buildings := [][]int{
		{1, 3, 3},
		{2, 4, 4},
		{5, 6, 1},
	}
	res := buildingOutLine(buildings)
	fmt.Println(res)
}
