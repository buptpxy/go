/*
比较器，在使用系统排序函数时可用比较器来排自定义类型
*/
package main

import (
	"fmt"
	"sort"
)

//以某一项为标准排序
type Student struct {
	name string
	age  int
}

func (s Student) String() string {
	return fmt.Sprintf("%s:%d", s.name, s.age)
}

type ByAge []Student

//ByAge 本身就是一个引用类型，故定义方法时不用传入它的指针
func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

///////////////////////////////////////////
//以每一项为排序标准，要定义多种Less()方法
type earthmass float64
type au float64
type Planet struct {
	name     string
	mass     earthmass
	distance au
}

//定义一个planetSorter类型，包含一个Planet类型的slice和一个排序标准By方法
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool
}

func (s *planetSorter) Len() int {
	return len(s.planets)
}
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

//对按不同性质排序的Less方法的定义
type By func(p1, p2 *Planet) bool

//定义在By上的Sort方法
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by,
	}
	sort.Sort(ps)
}

func main() {
	//以某一项为标准排序
	student := []Student{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(student)
	sort.Sort(ByAge(student)) //ByAge(student)把student转换成ByAge类型
	fmt.Println(student)

	// //以每一项为排序标准，要定义多种Less()方法
	var planets = []Planet{
		{"Mercury", 0.055, 0.4},
		{"Venus", 0.815, 0.7},
		{"Earth", 1.0, 1.0},
		{"Mars", 0.107, 1.5},
	}
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	decreasingDistance := func(p1, p2 *Planet) bool {
		return !distance(p1, p2)
	}
	//按照多种不同的标准为Planets组成的slice排序
	//By(name)把name转换成By类型
	By(name).Sort(planets)
	fmt.Println("by name: ", planets)
	By(mass).Sort(planets)
	fmt.Println("by mass: ", planets)
	By(distance).Sort(planets)
	fmt.Println("by distance: ", planets)
	By(decreasingDistance).Sort(planets)
	fmt.Println("by decreasingDistance: ", planets)
}
