package reflection

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Sex  byte   `xml:"sex" json:"sex"`
	age  int    `place:"abc"`
	Car
}

func (u *User) Test() {

}
func (u User) NonPtr() {

}
func (u *User) anonymous() {

}

type Car struct {
	speed int
}

func GetType() {
	typeI := reflect.TypeOf(1)
	fmt.Println(typeI)
	fmt.Println(typeI.String())
	fmt.Println(typeI.Kind()) //对于内置类型来说都是一样的,int
	println("-----------------------")

	typeS := reflect.TypeOf("hello")
	fmt.Println(typeS)
	fmt.Println(typeS.String())
	fmt.Println(typeS.Kind()) //对于内置类型来说都是一样的,string
	println("-----------------------")

	u1 := User{}
	typeUser := reflect.TypeOf(u1)
	fmt.Println(typeUser)          //包名.类型  reflection.User
	fmt.Println(typeUser.String()) //包名.类型
	fmt.Println(typeUser.Kind())   //struct
	println("-----------------------")

	u2 := &User{}
	typeUser2 := reflect.TypeOf(u2)
	fmt.Println(typeUser2)          //对于指针类型来说，*包名.类型  *reflection.User
	fmt.Println(typeUser2.String()) //对于指针类型来说，*包名.类型  *reflection.User
	fmt.Println(typeUser2.Kind())
	fmt.Println(typeUser2.Kind() == reflect.Pointer) //ptr
	fmt.Println(typeUser2.Elem())                    //相当于对指针解引用 为struct
	println("-----------------------")

	typeUser3 := typeUser2.Elem()
	fmt.Println(typeUser3)          //reflection.User
	fmt.Println(typeUser3.String()) //reflection.User
	fmt.Println(typeUser3.Kind())
	fmt.Println(typeUser3.Kind() == reflect.Struct) //struct
	//fmt.Println(typeUser3.Elem())   //报错，对不是指针的元素解引用
	println("-----------------------")
}
func GetField() {
	u1 := User{}
	typeUser := reflect.TypeOf(u1)
	FieldNum := typeUser.NumField()
	for i := 0; i < FieldNum; i++ {
		field := typeUser.Field(i)
		fmt.Println("Name:", field.Name, " Anonymous:", field.Anonymous, " Offset:", field.Offset,
			" Type:", field.Type, " Json", field.Tag.Get("json"), " Xml:", field.Tag.Get("xml"), " Place:", field.Tag.Get("place"),
			" IsExported:", field.IsExported(), " Index", field.Index, " PkgPath", field.PkgPath)
	}
	println("-----------------------")
	//u2 := &User{}
	//typeUser2 := reflect.TypeOf(u2)
	//FieldNum2 := typeUser2.NumField()  //不能对指针使用，会报异常，只能对结构体使用
}
func GetMethod() {
	u1 := User{}
	typeUser := reflect.TypeOf(&u1) //传入结构体只能获得结构体receiver的方法不能获得*receiver的方法，传入结构体指针那么receiver和*receiver的方法都能获得
	MethodNum := typeUser.NumMethod()
	for i := 0; i < MethodNum; i++ {
		method := typeUser.Method(i)
		fmt.Printf("Name:%s Type:%s PkgPath:%s IsExported:%t \n", method.Name, method.Type, method.PkgPath, method.IsExported())
	}
	println("-----------------------")
}
func Test(a int, b string, c User) int {
	return a
}
func GetFunc() {
	typeFunc := reflect.TypeOf(Test)
	argInNum := typeFunc.NumIn()
	argOutNum := typeFunc.NumOut()
	fmt.Println(typeFunc.Kind())
	fmt.Println(typeFunc.String())
	for i := 0; i < argInNum; i++ {
		argIn := typeFunc.In(i)
		fmt.Println("InArg", i, ":", argIn.Name())
	}
	for i := 0; i < argOutNum; i++ {
		argOut := typeFunc.In(i)
		fmt.Println("InArg", i, ":", argOut.Name())
	}
}
