/*manacher算法：给定一个字符串str,返回str中最长回文子串的长度。
思路：
1 对字符串进行处理，中间插入“#”，比如str="bcbaa"变成str1="#b#c#b#a#a#"可以保证字符串的字符个数总为奇数个
2 需要一个辅助数组pArr[]记录每一个i位置的最大回文半径，一个R记录能扩到的最右的位置，一个C记录最大回文半径的中心
3 计算pArr[]和R需要分成两种情况，一种是当前位置i大于R-1，则继续往右暴力扩。另一种是当前位置i小于R，则为可加速类型，具体又分为三种情况：
	3.1 i关于C的对称点i1的回文左范围超过了C的回文左范围，即i1=2*C-i,pArr[i1]>R-i，此时i的回文半径即pArr[i]=R-i
	3.2 i关于C的对称点i1的回文左范围在C的回文左范围之内，即pArr[i1]<R-i,此时i的回文半径即pArr[i]=pArr[i1]
	3.2 i关于C的对称点i1的回文左范围等于C的回文左范围，即pArr[i1]<R-i,则需要从R开始往外暴力扩
4 扩完之后更新R和C和最大的R值
5 最长回文子串的长度即为最大的R值
*/
package main

import "fmt"

func manacherString(str string) []rune {
	r1 := []rune(str)
	r2 := []rune{'#'}
	for i := 0; i < len(r1); i++ {
		r2 = append(r2, r1[i])
		r2 = append(r2, '#')
	}
	return r2
}
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func manacher(str string) int {
	r := manacherString(str)
	pArr := make([]int, len(r))
	R := -1
	C := -1
	Max := 0
	for i := 0; i < len(r); i++ {
		if i < R {
			pArr[i] = min(R-i, pArr[2*C-i])
		}
		for i+pArr[i] < len(r) && i-pArr[i] > -1 {
			if r[i+pArr[i]] == r[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		Max = max(Max, pArr[C])
	}
	return Max - 1
}
func main() {
	str := "bcbaa"
	r := manacherString(str)
	fmt.Println(string(r))
	fmt.Println(manacher(str))
}
