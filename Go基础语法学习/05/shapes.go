package shapes

import "math"
type Shape interface {
	//Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
	//Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
	//Shape类型的形参就可以接收Rectangle和Circle以及Triangle类型的结构体来调用他们的Area()方法
    Area() float64
}
type Triangle struct {
	Base float64
	Height float64
}
func (T Triangle) Area() float64 {
	return (T.Base*T.Height)/2
}
type Rectangle struct{
	Width float64
	Height float64
}
func (r Rectangle)Area() float64 { 
	return r.Width * r.Height
}
type Circle struct{
	Radius float64
}
func (c Circle)Area() float64 {
	return c.Radius * c.Radius*math.Pi
}
func Perimeter(r Rectangle)float64 {
	return 2*(r.Width+r.Height)
}
