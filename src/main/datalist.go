package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

//struct User{
//	Name string,
//	age  string
//}

type User struct{
	name string `json:"User_name"`
	age int     `json:"User_age"`
}

type User1 struct{
	Name string `json:"User_name"`
	Age int     `json:"User_age"`
}

type Student struct{
	name string
	Age int
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

	print("\n")
	//print("=============================================================================")
	fmt.Println("=================================")

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
	fmt.Println("....",a1, b1, a2)

	fmt.Println("================================")
	//map的创建
	//字面量的创建


	//fmt.Println("ma[a] = ", ma["a"])		// ma[a] =  1
	//fmt.Println("ma[b] = ", ma["b"])		// ma[b] =  2


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

	fmt.Println("=======================================")
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
	fmt.Println("person:",person)

	// value 为User对象
	user := User{
		name : "zhangkejin",
		age : 12,
	}
	userMap := make(map[int]User)

	userMap[0] = user

	fmt.Println("=======================================")

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

	fmt.Println("数据类型=======================================")

	//格式化输出
	var n1 = 100
	fmt.Printf("n1 的类型是 %T \n",n1)

	//求变量的字节大小，和数据类型
	var n2 int64 = 12
	fmt.Printf("n2的类型 %T  n2的字节数的%d",n2,unsafe.Sizeof(n2))

	//字符
	var c1 = 'a'
	fmt.Println("c1=",c1)

	//类型转换
	var num1 int64 = 999999
	var b3 int8 = int8(num1)
	fmt.Println("b3:",b3)

	var str string
	//string 的类型转化  为 其他类型
	var num3 int = 54
	str = fmt.Sprintf("%d", num3)
	fmt.Printf("str的类型是：%T 值为：%q\n", str, str)

	var num5 int = 54
	str = strconv.FormatInt(int64(num5),10)
	fmt.Printf("str 的类型的为%T, 值为：%q\n",str,str)

	//string  转化为  其他类型
	var n22 int8
	var n11 int64
	var sr string = "111222333"
	n11, _ = strconv.ParseInt(sr,10,8)
	n22 = int8(n11)
	fmt.Printf("str类型为 %T, 值为：%v\n",n11,n11)
	fmt.Printf("str类型为 %T, 值为：%v\n",n22,n22)

	//指针类型
	fmt.Println("指针类型=======================================")
    var nu int = 9
    fmt.Printf("num的地址为:%v",&nu)

    var ptr *int
    ptr = &nu
    *ptr = 10
    fmt.Printf("nu的值为：%v\n",nu)


    //跳转
	fmt.Println("跳转方法=======================================")
    var rr int64
    fmt.Printf("输入值：")
    //fmt.Scanln(&rr)
    if rr>18{
	    fmt.Println("年龄太小")
    }else{
    	fmt.Println("is ok")
	}

    //for 循环
    var cen int = 9
    for i:=0; i<cen; i++{
    	for j:=1; j<=i+1; j++{
			fmt.Printf("%v * %v = %v ",i+1,j,j*(i+1))
		}
		fmt.Println()
	}

    //函数
    fmt.Println(sum(1,2,3,4,5,6,7))

    //匿名函数
	fmt.Println("匿名函数=======================================")
    res1 := func (a int ,b int)int {
    	return a+b
	}(12,2)

    fmt.Println("res:",res1)

    //函数的传地址
    //var er int64 = 53
    er := 64
    add1(&er)

    //时间类
	fmt.Println("时间类=======================================")

	now := time.Now()
    fmt.Println("now:",now)
    fmt.Printf("now=%v,now type =%T\n",now,now )
    fmt.Printf("年=%v\n", now.Year())

	//格式化时间
	fmt.Printf("当前时间的年月日为%d,%d,%d\n",now.Year(),now.Month(),now.Day())

	fmt.Printf(now.Format("2222-02-22 01:01:01"))
	fmt.Println()

	//时间戳
	fmt.Printf("unix的时间戳为 %v,unixnano的时间戳为：%v\n",now.Unix(),now.UnixNano())

	//异常
	fmt.Println("异常类=======================================")
	test()
	fmt.Println("异常测试")


	fmt.Println("map的使用=======================================")
    //map的使用
	//字面量创建
    ma := map[string]int{"a" : 1, "b" : 2}
	fmt.Println("ma:",ma)
	//make创建
    var map1 map[string]string
	map1 = make(map[string]string)
	map1["name"] = "zhangkejin"
	map1["age"] = "zhangkejin"
	map1["address"] = "zhangkejin"

	//遍历
	for k,v := range map1{
		fmt.Printf("key:%v,value:%v",k,v)
	}

	//切片创建 map
	var map2 []map[string]string
	map2 = make([]map[string]string, 2)
	if map2[0] == nil{
		map2[0] = make(map[string]string,2)
		map2[0]["name"] = "zhangkejin"
		map2[0]["age"] = "zhangkejin"
	}
	if map2[1] == nil{
		map2[1] = make(map[string]string,2)
		map2[0]["name"] = "zhangkejin"
		map2[0]["age"] = "zhangkejin"
	}

	//map的复杂类型 是map 但不是切片
	studentmap := make(map[string]map[string]string)
	studentmap["stu1"] = make(map[string]string, 3)
	studentmap["stu1"]["name"] = "zhangkejin"
	studentmap["stu1"]["age"] = "zhangkejin"
	studentmap["stu1"]["phone"] = "zhangkejin"

	studentmap["stu1"] = make(map[string]string, 3)
	studentmap["stu1"]["name"] = "zhangkejin"
	studentmap["stu1"]["age"] = "zhangkejin"
	studentmap["stu1"]["phone"] = "zhangkejin"

	for k1, v1 := range studentmap {
		fmt.Printf("k1=%v",k1)
		for k2, v2 := range v1 {
			fmt.Printf("k2=%v,v2=%v\n",k2,v2)
		}
	}


	//json的序列化和反序列化  map struct  slice
	fmt.Println("map的使用=======================================")

	testMap()
	testSlice()
	//struct 的序列化   会根据反射： 按照json的 反射出来
	user8 := User1{"zhangsan",34}
	jsonstr, err := json.Marshal(&user8)
	if err != nil{
		fmt.Println("json error")
	}
	fmt.Println("struct : ..jsonstr:", string(jsonstr))//struct : ..jsonstr: {"User_name":"zhangsan","User_age":34}

    // 反序列化   这里用到的是 json 后的定义
    usere := "{\"User_name\":\"zhangkejin\", \"User_age\":23}"
    var user9 User1
	err = json.Unmarshal([]byte(usere), &user9)
    if err != nil{
    	fmt.Println("json error ")
	}
	fmt.Println("jsonstr:", user9)



}


//将map进行序列化
func testMap(){
	var a map[string]interface{}
	//使用map 一定要make
	a = make(map[string]interface{})
	a["name"] = "zhangkejin"
	a["age"] = 333
	a["address"] = "luliang"

	jsonstr, err := json.Marshal(a)
	if err != nil{
		fmt.Println("json error ")
	}
	fmt.Println("Map : jsonstr:",jsonstr)
	fmt.Println("str(jsonstr):",string(jsonstr))

}

//将切片进行序列化
func testSlice() {
	//var slice []map[string]interface{}
	slice := make([]map[string]interface{},2)
	slice1 := make(map[string]interface{},3)
	slice1["name"] = "zhangkejin"
	slice1["age"] = 23
	slice1["add"] = "zhangkejin"

	slice = append(slice, slice1)
	//if slice[0] == nil {
	//	slice[0] = make(map[string]interface{},3)
	//	slice[0]["name"] = "zhangkejin"
	//	slice[0]["age"] = 12
	//	slice[0]["add"] = "zhangkejin"
	//}
	//if slice[1] == nil {
	//	slice[1] = make(map[string]interface{},3)
	//
	//	slice[1]["name"] = "222zhangkejin"
	//	slice[1]["age"] = 22212
	//	slice[1]["add"] = "222zhangkejin"
	//}

    jsonstr, err := json.Marshal(slice)
    if err != nil{
		fmt.Println("json error ")
	}else{
		fmt.Println("slice  jsonstr:", jsonstr)
		fmt.Println("string(jsonstr):", string(jsonstr))
	}






}







func test(){
	//处理异常
	//使用defer 和  recover  来处理错误
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println("err:",err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1/num2
	fmt.Println("res:",res)


}


func init() {
	fmt.Println("THIS IS init()")
}

func add1(a1 *int){
	*a1 = *a1+10
}


//args是切片类型
func sum(n1 int, args... int )int {
	sun := n1
	for i:=0; i<len(args); i++ {
		 sun += args[i]
	}
	return sun

}


