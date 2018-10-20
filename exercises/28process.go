package main

/*
Q28. 进程
	编写一个程序，列出所有正在运行的进程，并打印每个进程执行的子进程个数。
	输出应当类似：
		Pid 0 has 2 children: [1 2]
		Pid 490 has 2 children: [1199 26524]
		Pid 1824 has 1 child: [7293]
	为了获取进程列表，需要得到 ps -e -opid,ppid,comm 的输出。输出类似：
		PID PPID COMMAND
		9024 9023 zsh
		19560 9024 ps
	如果父进程有一个子进程， 就打印 child， 如果多于一个， 就打印children；
	进程列表要按照数字排序，这样就以 pid 0 开始，依次展示。
************************************************************
解题思路
	1. 获得所有运行的进程
	2. 从输出结果解析其子进程
	2. 对进程号PID排序
	3. 判断是child还是children，输出
*/
import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	output, _ := ps.Output()
	//fmt.Println(string(output))
	child := make(map[int][]int)
	for i, s := range strings.Split(string(output), "\n") {
		if i == 0 {
			continue
		} //skip the first line
		if len(s) == 0 {
			continue
		} //skip the last line
		fields := strings.Fields(s)
		pidf, _ := strconv.Atoi(fields[0])  //string to int 得到子pid
		ppidf, _ := strconv.Atoi(fields[1]) //得到父pid
		child[ppidf] = append(child[ppidf], pidf)
	}
	//fmt.Println(child)
	ppidslice := make([]int, len(child))
	ppidCount := 0
	for ppidInt, _ := range child {
		ppidslice[ppidCount] = ppidInt
		ppidCount++
	}
	sort.Ints(ppidslice)
	fmt.Println(ppidslice)
	for _, ppidInt := range ppidslice {
		fmt.Printf("Pid %d has %d child", ppidInt, len(child[ppidInt]))
		if len(child[ppidInt]) == 1 {
			fmt.Printf(": %v\n", child[ppidInt])
			continue
		}
		fmt.Printf("ren: %v\n", child[ppidInt])
	}
}
