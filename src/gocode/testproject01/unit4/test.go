package main
import "fmt"

func main(){
	//类型转换，必须显式转换
	var n1 int = 100
	fmt.Println(n1)
	// var n2 float32 = n1 //报错
	var n2 float32 = float32(n1)
	fmt.Println(n2)

	fmt.Printf("%T ",n1)
	fmt.Printf("%T",n2)

	var n3 int64 = 88888
	var n4 int16 = int16(n3)
	fmt.Println(n4)

	var n5 int32 = 12
	var n6 int8 = int8(n5) + 30 //一定要匹配左右的数据类型
	fmt.Println(n6)

	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127 //编译通过，运行溢出
	// var n9 int8 = int8(n7) + 128 //128溢出，编译不通过
	fmt.Println(n8)
}