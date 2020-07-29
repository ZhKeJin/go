package modle

import "fmt"

type person struct{
	Name string
	age int
	sal float64
}


func NewPerson(name string) *person{
	return &person{
		Name : name,
	}
}

func (p *person) SetAge(age int){
	if age >0 && age<159 {
		p.age = age
	}else{
		fmt.Println("年龄错误" )
	}

}