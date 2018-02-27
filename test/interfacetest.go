package test

/**
User:		wliangde
CreateTime:	18/2/27 下午3:50
Brief:		http://blog.csdn.net/uudou/article/details/52456133
http://blog.csdn.net/uudou/article/details/52556840
接口interface 是一系列方法的集合
**/

import (
	"fmt"
)

type PeopleGetter interface {
	GetName() string
	GetAge() int
}

//EmployeeGetter 接口嵌入PeopleGetter，表明EmployeeGetter 拥有PeopleGetter的所有方法
type EmployeeGetter interface {
	PeopleGetter
	GetSalay() int
	Help()
}

type Employee struct {
	name   string
	age    int
	salay  int
	gender string
}

func (this *Employee) GetName() string {
	return this.name
}
func (this *Employee) GetAge() int {
	return this.age
}

func (this *Employee) Help() {
	fmt.Println("This is help info.")
}

func (this *Employee) GetSalay() int {
	return this.salay
}

// 匿名接口可以被用作变量或者结构属性类型
//相当于结构体里有个gender接口指针
//type Man struct {
//	gender interface {
//		GetGender() string
//	}
//}

type Man struct {
	gender
}

//等价于下面
type gender interface {
	GetGender() string
}
type Woman struct {
	gender
}

func (this *Employee) GetGender() string {
	fmt.Println("Call Employee GetGender")
	return this.gender
}

func (this *Man) GetGender() string {
	fmt.Println("Call Man GetGender")
	return "Male"
}

type MagicError struct{}

func (this MagicError) Error() string {
	return "[Magic]"
}

func Generate() *MagicError {
	return nil
}

func Test() error {
	return Generate()
}

func InterfaceTest() {
	varEmployee := Employee{"wld", 27, 30000, "male"}
	var varEmpInter EmployeeGetter = &varEmployee
	fmt.Println("type %T", varEmpInter)
	switch varEmpInter.(type) {
	case nil:
		fmt.Println("nil")
	case PeopleGetter:
		fmt.Println("PeopleGetter")
	default:
		fmt.Println("Unknown")
	}
	varMan := Man{}
	varMan.gender = &varEmployee //可以赋值任何实现了gender接口的类型

	fmt.Println("man %T %s", varMan, varMan.gender.GetGender())
	varMan.GetGender() //调用的是Man里面的GetGender

	//++++++++++++++++++++++note
	errI := Test()
	if errI != nil { //接口的type为*test.MagicError  接口的value是nil，所以次数不为nil
		fmt.Printf("Hello, Mr. Pike! %T %V\n", errI, errI)
	}
}
