/*
猫狗队列 【 题目】 宠物、 狗和猫的类如下：
public class Pet {
	private String type;
	public Pet(String type) { this.type = type; }
}
public String getPetType() { return this.type; }
public class Dog extends Pet { public Dog() { super("dog"); } }
public class Cat extends Pet { public Cat() { super("cat"); } }

实现一种狗猫队列的结构， 要求如下：
1. 用户可以调用add方法将cat类或dog类的实例放入队列中；
2. 用户可以调用pollAll方法， 将队列中所有的实例按照进队列的先后顺序依次弹出；
3. 用户可以调用pollDog方法， 将队列中dog类的实例按照进队列的先后顺序依次弹出；
4. 用户可以调用pollCat方法， 将队列中cat类的实例按照进队列的先后顺序依次弹出；
5. 用户可以调用isEmpty方法， 检查队列中是否还有dog或cat的实例；
6. 用户可以调用isDogEmpty方法， 检查队列中是否有dog类的实例；
7. 用户可以调用isCatEmpty方法， 检查队列中是否有cat类的实例。
*/

package main

import (
	"fmt"
	"queue"
)

type Pet struct {
	Type string
}

type petEnter struct {
	Pet
	count int
}

type Dogcatq struct {
	dogq queue.Queue
	catq queue.Queue
	size int
}

func (d *Dogcatq) New() {
	d.dogq = *(queue.NewQueue())
	d.catq = *(queue.NewQueue())
	d.size = 0
}

func (d *Dogcatq) add(p petEnter) {
	if p.Type == "cat" {
		d.size++
		p.count = d.size
		d.catq.Push(p)
	} else if p.Type == "dog" {
		d.size++
		p.count = d.size
		d.dogq.Push(p)
	} else {
		fmt.Println("no cat or dog!")
	}
}
func (d *Dogcatq) printq() {
	fmt.Println("dog: ")
	d.dogq.Print()
	fmt.Println("cat: ")
	d.catq.Print()
}
func (d *Dogcatq) pollAll() {
	if d.isDogEmpty() && d.isCatEmpty() {
		fmt.Println("no dog or cat!")
		return
	}

	for !d.isDogEmpty() && !d.isCatEmpty() {
		if d.catq.Top().(petEnter).count < d.dogq.Top().(petEnter).count {
			fmt.Println(d.catq.Pop().(petEnter).Pet)
		} else {
			fmt.Println(d.dogq.Pop().(petEnter).Pet)
		}
	}

	for !d.isDogEmpty() {

		fmt.Println(d.dogq.Pop().(petEnter).Pet)

	}
	for !d.isCatEmpty() {

		fmt.Println(d.catq.Pop().(petEnter).Pet)

	}

}
func (d *Dogcatq) pollDog() {
	if d.isDogEmpty() {
		fmt.Println("no dog!")
		return
	}
	for i := 0; i < d.dogq.Size(); i++ {
		fmt.Println(d.dogq.Pop().(petEnter).Pet)
	}
}
func (d *Dogcatq) pollCat() {
	if d.isCatEmpty() {
		fmt.Println("no cat!")
		return
	}
	for i := 0; i < d.catq.Size(); i++ {
		fmt.Println(d.catq.Pop().(petEnter).Pet)
	}
}
func (d *Dogcatq) isEmpty() bool {
	if d.dogq.IsEmpty() && d.catq.IsEmpty() {
		return true
	}
	return false
}
func (d *Dogcatq) isDogEmpty() bool {
	if d.dogq.IsEmpty() {
		return true
	}
	return false
}
func (d *Dogcatq) isCatEmpty() bool {
	if d.catq.IsEmpty() {
		return true
	}
	return false
}
func main() {
	var cat Pet
	cat.Type = "cat"
	var dog Pet
	dog.Type = "dog"
	var d Dogcatq
	d.New()
	var p petEnter
	p = petEnter{cat, 0}
	d.add(p)
	// d.printq()
	p = petEnter{dog, 0}
	d.add(p)
	// d.printq()
	d.pollAll()
}
