package main

import "fmt"

const PI float32 =3.141592657

type Shape interface{
	Area() float32
	Perimeter() float32
}

type Rectangle struct{
	Width float32
	Height float32
	
}

type Circle struct{
	Radius int
}

func (rec Rectangle) Area() float32{
	return rec.Width*rec.Height
}

func (rec Rectangle) Perimeter() float32{
	return (rec.Height+rec.Width)*2
}

func (cir Circle) Area() float32{
	return float32(cir.Radius*cir.Radius)*PI
}


func (rec Circle) Perimeter() float32{
	return float32(rec.Radius)*2*PI
}


func main(){
	/*

题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
	*/

	var rectangle =Rectangle{Width: 3,Height: 5}
	
	recArea:=rectangle.Area()
	fmt.Println("recArea:",recArea)

	recPerimeter:=rectangle.Perimeter()
	fmt.Println("recPerimeter:",recPerimeter)

	var circle = Circle{Radius: 4}
	cirArea :=circle.Area()
	fmt.Printf("cirArea :%.2f\n",cirArea)

	cirPerimeter:=circle.Perimeter()
	fmt.Printf("cirPerimeter: %.2f\n",cirPerimeter)

}