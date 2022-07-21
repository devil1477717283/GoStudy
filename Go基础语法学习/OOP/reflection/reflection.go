package reflection

import (
	"fmt"
	"reflect"
)

type ForTest interface {
	Test()
}
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

func (c *Car) Test() {

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
func GetStructRelation() {
	car := &Car{}
	user := &User{}
	carType := reflect.TypeOf(car)
	userType := reflect.TypeOf(user)
	fortestType := reflect.TypeOf((*ForTest)(nil)).Elem() //接口不能创建实例,但是将nil转成一个指针是没有问题的，再用Elem获取原来的实例
	//输出都是false Go里面严格来说没有继承的概念，所以父类不能赋值给子类，也不能转换为子类
	println(carType.AssignableTo(userType))
	println(carType.ConvertibleTo(userType))
	println("-------------------------")

	//是否能赋值给该类型
	println(userType.AssignableTo(fortestType))
	println(carType.AssignableTo(fortestType))
	println("-------------------------")

	//是否能转换为该类型
	println(userType.ConvertibleTo(fortestType))
	println(carType.ConvertibleTo(fortestType))
	println("-------------------------")

	//是否实现了接口
	println(userType.Implements(fortestType))
	println(carType.Implements(fortestType))
	println("-------------------------")

}

func ValueOtherConversion() {
	iValue := reflect.ValueOf(1)
	sValue := reflect.ValueOf("hello")
	userValue := reflect.ValueOf(&User{
		Name: "zs",
		Sex:  2,
		age:  18,
	})
	fmt.Printf("%v\n", iValue)
	fmt.Printf("%v\n", sValue)
	fmt.Printf("%v\n", userValue)
	println("-------------------------")

	//reflect.Value.Type()可以将reflect.Value类型转换为reflect.Type   等价于reflect.Typeof()得到的reflect.Type
	//reflect.Type.New()可以将reflect.Type类型转换为reflect.Value    同上
	iType := iValue.Type()
	sType := sValue.Type()
	userType := userValue.Type()

	println(iValue.Kind() == iType.Kind(), iValue.Kind() == reflect.Int)               //true true
	println(sValue.Kind() == sType.Kind(), sValue.Kind() == reflect.String)            //true true
	println(userValue.Kind() == userType.Kind(), " ", userValue.Kind() == reflect.Ptr) //true true
	println("-------------------------")

	println(iValue.String(), " ", iType.String(), " ", iValue.String() == iType.String())             //<int Value> int false
	println(sValue.String(), " ", sType.String(), " ", sValue.String() == sType.String())             //hello  string  false
	println(userValue.String(), " ", userType.String(), " ", userValue.String() == userType.String()) //<*reflection.User Value> *reflection.User false
	println("-------------------------")

	//实现指针和结构体的互相转换
	userValue2 := userValue.Elem()
	println(userValue2.Kind() == reflect.Struct)
	userValue3 := userValue2.Addr()
	println(userValue3.Kind() == reflect.Ptr)
	println("-------------------------")

	//获取原始的值
	println("iValue的值为:", iValue.Interface().(int), " ", iValue.Int())
	println("sValue的值为:", sValue.Interface().(string), " ", sValue.String())
	println("-------------------------")

	user := userValue.Interface().(*User)
	println(user.Name, user.age, user.Sex)
	user2 := userValue2.Interface().(User)
	println(user2.Name, user2.age, user2.Sex)
}
func ValueIsEmpty() {
	var ifc interface{}
	v := reflect.ValueOf(ifc)
	println("v持有真实的值?", v.IsValid(), v.Kind() == reflect.Invalid)
	ifc = 8
	v1 := reflect.ValueOf(ifc)
	println("v持有真实的值?", v1.IsValid(), v1.Kind() == reflect.Int)

	var user *User = nil
	v = reflect.ValueOf(user)
	if v.IsValid() {
		println("v持有的值是nil", v.IsNil()) //调用IsNil之前必须保证IsValid必须是true，否则会panic
	} else {
		println("v没有持有值")
	}

	var u User
	v = reflect.ValueOf(u)
	if v.IsValid() {
		println("v持有的值是zero", v.IsZero()) //结构体不能调用IsNil
	} else {
		println("v没有持有值")
	}
}

//是否可寻址
func Addressable() {
	v1 := reflect.ValueOf(1)
	var x int
	v2 := reflect.ValueOf(x)
	v3 := reflect.ValueOf(&x)
	v4 := v3.Elem()
	println("v1 ", v1.CanAddr()) //false
	println("v2 ", v2.CanAddr()) //false
	println("v3 ", v3.CanAddr()) //false
	println("v4 ", v4.CanAddr()) //true

	slice := make([]int, 3, 5)
	v5 := reflect.ValueOf(slice)
	v6 := v5.Index(1)
	println("v5 ", v5.CanAddr()) //false  切片的value不可寻址
	println("v6 ", v6.CanAddr()) //true   切片中元素的value可以寻址

	mp := make(map[int]bool, 5)
	v7 := reflect.ValueOf(mp)
	println("v7 ", v7.CanAddr()) //false map的value不可寻址

}
func ChangeValue() {
	var i int = 10
	iValue := reflect.ValueOf(&i)
	if iValue.CanAddr() {
		iValue.SetInt(8)
		fmt.Printf("i=%d\n", i)
	} else {
		fmt.Println("iValue Can't Addressable!")
	}
	fmt.Printf("address of i %p\n", &i)
	iValue2 := iValue.Elem()
	if iValue2.CanAddr() {
		iValue2.SetInt(8)
		fmt.Printf("i=%d\n", i)
		fmt.Printf("address of i %p\n", &i)
	} else {
		fmt.Println("iValue2 Can't Addressable!")
	}

	var s string = "hello"
	sValue := reflect.ValueOf(&s) //必须传指针再调Elem()，否则会panic
	sValue.Elem().SetString("Go") //只有可寻址的才能调Set系列函数
	fmt.Println(s)

	user := &User{
		Name: "zs",
		age:  18,
		Sex:  2,
	}
	userValue := reflect.ValueOf(user)
	fmt.Printf("user is addressable %t\n", userValue.CanAddr())
	userValue2 := userValue.Elem()
	fmt.Printf("user2 is addressable %t\n", userValue2.CanAddr())
	fmt.Printf("change before %v\n", user)
	userValue2.FieldByName("Sex").SetUint(1) //只能修改可导出的变量，并且类型要一致,FieldByName()返回的Field是可寻址的
	fmt.Printf("change after %v\n", user)

	slice := make([]*User, 3, 5)
	slice[0] = &User{
		Name: "zs",
		Sex:  1,
	}
	fmt.Printf("%v\n", slice[0])
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() > 0 {
		Elem1 := sliceValue.Index(0)
		if Elem1.CanAddr() {
			Elem1.Elem().FieldByName("Sex").SetUint(20)
			Elem1.Elem().Field(0).SetString("张三")
			fmt.Printf("%v\n", slice[0])
		} else {
			fmt.Println("Elem1 Can't Addressable")
		}
	}
	sliceValue.Index(1).Set(reflect.ValueOf(&User{
		Name: "李四",
		age:  19,
		Sex:  2,
	}))
	fmt.Printf("%v\n", slice[1])

	//用反射更改切片的cap和len时，cap只能降不能升，len可降可升，但是不能超过cap
	//用反射更改map时，
}
