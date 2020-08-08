package main

import (
	"encoding/json"
	"fmt"

)

type person struct{
	Name string
	Age int
	Scorces [5]float64
    ptr *int
	slice []int
	map1 map[string]string
}

type Dog struct{
	Name string
	Age int
}

func (d Dog) speak(a int) int {
	return a*a
}

func (d *Dog) eat(a int)int{
	return a*a
}


func main(){
	//stu := Student{"zangkeijn",34}
	//fmt.Println(stu)

	//stu :=Student {"zhangkeing",34}
    //fmt.Println(stu)

	var per person
	fmt.Println(per)

	if per.ptr == nil{
		fmt.Println("ok1")
	}
	//使用slice和  使用 map一定要  make
	per.map1 =  make(map[string]string)
	per.map1["name"] = "string"
	per.slice = make([]int,10)
	per.slice[1] = 111111

	fmt.Println("per:",per)

	//结构体的初始化
	Dog1 := Dog{"zhangkeing",34}
	fmt.Println("per:",Dog1)

	var dog1 *Dog = new(Dog)
	(*dog1).Age = 12

	var dog2 *Dog = &Dog1
	fmt.Printf("Dog1的地址是:%p, dog1的地址是%p\n", &dog1,&dog2)

	//struct 中每个字段可以加上一个tag
    dog3 := Dog{"zhang", 23}
    //进行dog3的序列化 json的格式字符串  根据反射的原理  序列化和反序列化
    jsonstr, err := json.Marshal(dog3)
    if err != nil{
    	fmt.Println("json error", err)
	}

	fmt.Println("jsonstr:",string(jsonstr)) //{"name":"zhang","age":23}

	fmt.Println("jsonstr:",jsonstr)// [123 34 110 97 109 101 34 58 34 122 104 97 110 103 34 44 34 97 103 101
	//34 58 50 51 125]

    dog4 := Dog{"zhangsan", 23}
	i := dog4.speak(dog4.Age)
    fmt.Println("i=",i)
    fmt.Println("dog4.age:",dog4.Age)

	fmt.Println(dog4.eat(dog4.Age))

	//
	////引用modle中的 person
	//per := modle.




}

