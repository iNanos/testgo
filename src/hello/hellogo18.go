package main

import "fmt"

//值传递与指针传递的区别

//1.如果接受者是一个 map，func 或者 chan，使用值类型(因为它们本身就是引用类型)。
//2.如果接受者是一个 slice，并且方法不执行 reslice 操作，也不重新分配内存给 slice，使用值类型。
//3.如果接受者是一个小的数组或者原生的值类型结构体类型(比如 time.Time 类型)，而且没有可修改的字段和指针，又或者接受者是一个简单地基本类型像是 int 和 string，使用值类型就好了。
//
//一个值类型的接受者可以减少一定数量的垃圾生成，如果一个值被传入一个值类型接受者的方法，一个栈上的拷贝会替代在堆上分配内存(但不是保证一定成功)，所以在没搞明白代码想干什么之前，别因为这个原因而选择值类型接受者。
//
//何时使用指针类型
//
//1.如果方法需要修改接受者，接受者必须是指针类型。
//2.如果接受者是一个包含了 sync.Mutex 或者类似同步字段的结构体，接受者必须是指针，这样可以避免拷贝。
//3.如果接受者是一个大的结构体或者数组，那么指针类型接受者更有效率。(多大算大呢？假设把接受者的所有元素作为参数传给方法，如果你觉得参数有点多，那么它就是大)。
//4.从此方法中并发的调用函数和方法时，接受者可以被修改吗？一个值类型的接受者当方法调用时会创建一份拷贝，所以外部的修改不能作用到这个接受者上。如果修改必须被原始的接受者可见，那么接受者必须是指针类型。
//5.如果接受者是一个结构体，数组或者 slice，它们中任意一个元素是指针类型而且可能被修改，建议使用指针类型接受者，这样会增加程序的可读性
//
//当你看完这个还是有疑虑，还是不知道该使用哪种接受者，那么记住使用指针接受者。

type Rect struct {
	X, Y, Area, Length float64
}

func (r *Rect) getArea() {
	r.Area = r.X * r.Y
}

func (r Rect) getLength() {
	r.Length = r.X + r.Y
}

func main() {
	r1 := Rect{
		X:      2,
		Y:      3,
		Area:   0,
		Length: 0,
	}
	fmt.Println("before")
	fmt.Printf("r1 area : %f", r1.Area)
	fmt.Printf("   r1 length : %f", r1.Length)

	r1.getLength()
	r1.getArea()

	fmt.Println("\nafter")
	fmt.Printf("r1 area : %f", r1.Area)
	fmt.Printf("  r1 length : %f", r1.Length)
}
