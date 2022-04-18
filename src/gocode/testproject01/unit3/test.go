package main
import "fmt"

func main(){
	var s1 string = "hello golang !!!"
	fmt.Println(s1)

	//字符串是不可变的，一旦定义好，其中的字符的值不能改变
	var s2 string = "abc"
	s2 = "def"
	fmt.Println(s2)
	// s2[0] = 't' //报错

	//如果字符串中没有特殊字符，用"",否则用``
	var s3 string = "sgargagar"
	fmt.Println(s3)

	var s4 string = `
	func main(){
		var s1 string = "hello golang !!!"
		fmt.Println(s1)
	
		//字符串是不可变的，一旦定义好，其中的字符的值不能改变
		var s2 string = "abc"
		s2 = "def"
		fmt.Println(s2)
		// s2[0] = 't' //报错
	`
	fmt.Println(s4)

	var s5 string = "abc" + "def"
	s5 += "sgsggrag"
	fmt.Println(s5)

	var s6 string = "abc" + "def" + "abc" + "def" +
	"abc" + "def" + "abc" + "def" + "abc" + "def" +
	"abc" + "def"
	fmt.Println(s6)
}