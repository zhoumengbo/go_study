package main
import (
	"fmt"
	"strconv"
)


func main(){
	//基本数据类型转换string
	var n1 int = 19
	var n2 float64 = 4.78
	var n3 bool = false
	var n4 byte = 's'

	//方法1： fmt.Sprintf("%参数",表达式) 推荐！！！
	var s1 string = fmt.Sprintf("%d",n1)
	fmt.Printf("s1对应的类型是：%T，s1 = %q \n",s1,s1)

	var s2 string = fmt.Sprintf("%f",n2)
	fmt.Printf("s2对应的类型是：%T，s2 = %q \n",s2,s2)

	var s3 string = fmt.Sprintf("%t",n3)
	fmt.Printf("s3对应的类型是：%T，s3 = %q \n",s3,s3)

	var s4 string = fmt.Sprintf("%c",n4)
	fmt.Printf("s4对应的类型是：%T，s4 = %q \n",s4,s4)

	//方法2： strconv
	var s5 string = strconv.FormatInt(int64(n1),10)
	fmt.Printf("s5对应的类型是：%T，s5 = %q \n",s5,s5)

	var s6 string = strconv.FormatFloat(n2,'f',9,64) //9:保留9位小数
	fmt.Printf("s6对应的类型是：%T，s6 = %q \n",s6,s6)


	//string转换基本数据类型
	var s7 string = "true"
	var b bool
	b , _ = strconv.ParseBool(s7)
	fmt.Printf("b的类型是：%T，b = %v \n",b,b)

	var s8 string = "19"
	var num1 int64
	num1 , _ = strconv.ParseInt(s8,10,64)
	fmt.Printf("num1的类型是：%T，num1 = %v",num1,num1)

}