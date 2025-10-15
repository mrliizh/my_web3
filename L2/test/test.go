package main

import (
	"fmt"
	"time"
)

func convertTemperature(c float64) []float64 {
	return []float64{c + 273.15, c*1.80 + 32.00}
}

func arrayTest() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	l := len(a)
	fmt.Println(l)

	// var m = [5]int{0,1,2,4,5}
	c  := [...]int{0,1,2,3,4}
	fmt.Println(c)

	var twoD [2][3]int
	for i := range 2{
		for j := range 3{
			twoD [i][j] = i + j
		}
	}
	fmt.Println(twoD)

	twoD = [...][3]int{{1,2,3},{1,2,3}}
}
func slicesTest(){

	// var str string = "ello"
	str := "hello"
	fmt.Println(str[1])

	var s = []string{"hello","hah"}
	fmt.Println(s)

	s = make([]string, 3)
	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
	s = append(s, "d")
	fmt.Println(s, len(s), cap(s))
}


// func makeHandler(user string) func() {
//     return func() {
//         fmt.Println("处理来自用户：", user)
//     }
// }

// func onEvent(callback func()) {
//     fmt.Println("事件触发！")
//     callback()
// }

// func main() {
//     handler := makeHandler("Alice") // 创建闭包，记住 user = "Alice"
//     onEvent(handler)                // 在后面某个事件发生时调用
// }


// func main(){
// 	// slicesTest()

// 	fmt.Println("UtfTest():")
// 	s := "hi你好"
// 	fmt.Println("一共占用的字节数", len(s))
// 	for i := 0; i < len(s) ; i ++{
// 		fmt.Println(i, s[i])
// 	}
// 	fmt.Println(s[2],"是'你'所占用的三个字节中的第一个字节")
// 	for _, v := range s{
// 		fmt.Printf("%c : %d\n", v, v)
// 	}
// 	fmt.Println()

// 	byteSlice := []byte("Go语言")

//     // 将 byteSlice 转换为 string
//     // 这会将 byteSlice 中的 UTF-8 字节数据解码为一个 Unicode 字符串
//     str := string(byteSlice)

//     // 然后将 string 转换为 []rune
//     // 这会将每个 Unicode 字符（rune）提取出来
//     runeSlice := []rune(str)

//     fmt.Println("byteSlice:", byteSlice)  // 输出原始字节
//     fmt.Println("runeSlice:", runeSlice)  // 输出 Unicode 码点的切片

// 	fmt.Println(string(byteSlice)) //将utf8编码转换为字符串
// 	fmt.Println(string(runeSlice)) //将码点值切片转换为字符串

// }



func schedule(task func()) {
    go func() {
        time.Sleep(2 * time.Second)
        task() // 异步触发
    }()
}

func main() {
    // name := "张三"
    // schedule(func() {
    //     fmt.Println("异步任务由", name, "执行")
    // })
    // time.Sleep(3 * time.Second)
}
