package main

import (
	"encoding/json"
	"fmt"
)

//struct User{
//	Name string,
//	age  string
//}

type User struct{
	name string `json:"num"`
	age int     `json:"age"`
}


type A struct{
	name string
}
type B struct{
	name string
}

type AA A



func main() {

	//数组。。。。。
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("a:%T",a)


	print("=============================================================================")

	//silce。。。。。
	d := [...]int{1,3,3,4,55,6,6}
	//数组创建
	s1 := d[1:4]
	fmt.Println(a, s1)

	//make 创建
	a3 := [...]int{1,2,3,4,5,5,6,7}
	a1 := make([]int, 3)
	b1 := make([]int, 2, 3)
	a2 :=a3[1:6]

	b1 = append(b1,a2...)
	fmt.Println(len(b1),",,,",cap(b1)) // 自动扩容，和数组的区别
	fmt.Println(a1, b1, a2)

	print("=============================================================================")
	//map的创建
	//字面量的创建
	ma := map[string]int{"a" : 1, "b" : 2}

	fmt.Println("ma[a] = ", ma["a"])		// ma[a] =  1
	fmt.Println("ma[b] = ", ma["b"])		// ma[b] =  2


	// make(map[K]T) 	  map 的容量使用默认值
	// make(map[K]T, len) map 的容量使用给定的 len 值
	// make 表示分配内存空间，len 表示可以存放见指对的个数

	//make创建
	mp1 := make(map[int]string)
	mp2 := make(map[int]string, 2)
	mp1[1] = "zhangkejin"
	mp2[3] = "zhang"
	mp2[4] = "zhang"
	mp2[5] = "zhang"
	mp2[6] = "zhang"

	fmt.Println(mp1)
	fmt.Println(mp2)

	//查找
	val, is := mp2[4]
	if is {
		print(val)
	}else{
		print("is null")
	}

	//遍历  //输出的顺序是不确定的
	for key, val := range mp2 {
		fmt.Print(key,val)
	}

	print("=============================================================================\n")
	// map的 slice   把map变成了一个数组使用  然后通过 make 的方式创建了 切片   每一个map是一个数组中的值
	var person []map[string]string
	person = make([]map[string]string, 2)

	if person[0] == nil{
		person[0] = make(map[string]string,2)    //person 为一个数组
		person[0]["name"] = "zhangkejin"
		person[0]["name1"] = "zhangkejin2"
		person[0]["name2"] = "zhangkejin3"
		person[0]["name3"] = "zhangkejin3"
	}
	if person[1] == nil{
		person[1] = make(map[string]string,2)
		person[1]["name"] = "zhangkejin"
		person[1]["name1"] = "zhangkejin2"
		person[1]["name2"] = "zhangkejin3"
		person[1]["name3"] = "zhangkejin3"
	}

	// 越界
	//if person[2] == nil{
	//	person[2] = make(map[string]string,2)
	//	person[2]["name"] = "zhangkejin"
	//	person[2]["name1"] = "zhangkejin2"
	//	person[2]["name2"] = "zhangkejin3"
	//	person[2]["name3"] = "zhangkejin3"
	//}

	p := map[string]string{
		"name" : "zhangkejin",
		"age" : "12",
	}

	person = append(person, p)
	fmt.Println(person)

	// value 为User对象
	user := User{
		name : "zhangkejin",
		age : 12,
	}
	userMap := make(map[int]User)

	userMap[0] = user

	print("=============================================================================\n")

	//struct 的使用
	var user1 User
	var user2 User = User{"zhangkejin", 23}
	fmt.Println(user1,"...", user2)

	//var a4 A
	//	//var b4 B
	//	//var aa AA
	//	//aa = a4
	//	//
	//	//fmt.Print(a,b4, a)

	//struct 的序列化
	user2 = User{"zhangkejin", 23}

	marshal, err := json.Marshal(user2)
	if err != nil{
		fmt.Println("cuowu ")
	}else{
		fmt.Println(marshal)
		fmt.Println(string(marshal))
	}



}