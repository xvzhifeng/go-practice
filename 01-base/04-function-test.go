package main

import "fmt"

func main() {
	fmt.Println(max(1,2))
	fmt.Println(min(1,2))
	var arr = []int{1,2,4,5,6,1,23,4,1,23}
	fmt.Println(arr)
	fmt.Println(sum(arr))
	fmt.Println(sort_d(arr, true))
	fmt.Println(sort_d(arr, false))
	fmt.Println(arr)
	fmt.Println(count(arr, 1))
	fmt.Println(sort_e(arr))

	var a1  = []int{1,2,3,5,5}

	sort_e(a1)
	cnt_0 := count(a1,0)
	n := len(a1)
	diff := 0
	for i:=cnt_0+1;i<n;i++ {
		if a1[i] == a1[i-1] {
			fmt.Println("F")
		}
		diff += a1[i] - a1[i-1] - 1
	}
	fmt.Println(diff, cnt_0)
	if diff <= cnt_0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func max(a, b int) (res int) {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) (res int) {

	if a > b {
		res = b 
	} else {
		res = a
	}
	return
}

func sum(a []int) (res int) {
	var n = len(a)
	for i:=0;i<n;i++ {
		res += a[i]
	}
	return

}

func sort_d(a []int, reverse bool) (res []int) {
	//  reverse is True 从大到小
	if reverse {
		for i := 0; i < len(a) ; i++ {
			for j := 0 ; j< len(a) - i-1; j++ {
				if a[j] > a[j+1] {
					a[j],a[j+1] = a[j+1],a[j]
				}
			}
		}
		res = a

	} else {
		for i := 0; i < len(a) ; i++ {
			for j := 0 ; j< len(a) - i-1; j++ {
				if a[j] < a[j+1] {
					a[j],a[j+1] = a[j+1],a[j]
				}
			}
		}
		res = a
	}
	return
}


func sort_e(a []int) (res []int) {
	n := len(a)
	for i:=0; i < n ;i++ {
		for j := 0 ; j < n-i-1;j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	res = a
	return 
}

func count(a []int, b int) (res int) {
	for _,v := range a{
		if v == b {
			res++
		}
	}
	return 
}