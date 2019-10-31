package main

import (
	"errors"
	"fmt"
	"point"
	"time"
)

func main() {
	var name, age, sex = getInfo()
	fmt.Println(name, age, sex, &address)
	getArray()
	getPoint()
	useStruct()
	slice := getSlice()
	rangeSlice(slice)
	getMap()
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(HtcPhone)
	phone.call()

	go printTime("Hello World!")
	printTime("Hello Go!")

	getChannel()
	//point.Say()

	var pc Pc
	pc = new(MacPc)
	pc.print("mac print")

	point.SendData()
	point.HttpGet()
	//http.ListenAndServe(":12345", http.FileServer(http.Dir("/")))

	//函数赋值给变量
	sq := Square
	println("Square:", sq(3))

	callback("Hi", Square)
	getSubTime()
}

//变量
func getInfo() (string, int, string) {
	var name string = "Hello World"
	var age = 8
	sex := "男"

	println("类型转换", float32(age)/float32(3))
	return name, age, sex
}

//全局变量
var (
	name string
	age  int
	//sex := "男" 全局声明中次方式会报错
	qq, email, address = 2940003, "ten@qq.com", "腾讯大厦"
)

//数组
func getArray() [3]string {
	var ageArray [3] int
	ageArray[1] = 18
	println(ageArray[1])

	var nameArray = [3]string{"Helen", "Bob", "Eric"}
	println(nameArray[1])

	return nameArray
}

//指针
func getPoint() {
	name := "萨莉"
	//获取变量的指针：&name

	//声明指针类型: *
	//可简写为：var namePoint = &name
	var namePoint *string
	namePoint = &name

	//使用指针访问变量：*变量
	println("name的指针：", namePoint, "指向的值为：", *namePoint)

	//若声明指针类型，不赋值，则为空指针
	var age *int
	println("空指针：", age)

	//空指针判断
	if (age == nil) {
		println("为空指针")
	}
}

//结构体定义
type Person struct {
	name  string
	age   int
	phone string
}

//结构体使用
func useStruct() {
	//赋值
	var person Person
	person.age = 18
	person.name = "Eric"
	person.phone = "13333223"

	//使用
	println("person---> name:", person.name, " age:", person.age, "phone", person.phone)

	//结构体指针
	var personPoint *Person
	personPoint = &person
	println("结构体Person的指针：", personPoint)

	//通过结构体指针赋值其变量
	personPoint.phone = "234"
	//通过结构体指针访问其变量
	println(personPoint.phone)
}

//切片Slice
func getSlice() []string {
	//声明切片
	//方法一：声明一个未指定大小的数组
	var name []string
	println(len(name))

	name = append(name, "萨莉", "海伦", "戴维斯")
	println(name[0], len(name), cap(name))

	name[0] = "嘉文"
	println(name[0])

	//方法二：使用make函数
	var age = make([]int, 7)
	age[2] = 18

	//方法三：指定容量，capacity为可选参数
	var sex = make([]string, 8)
	sex = []string{"男", "女"}
	println(sex[0])

	//切片拷贝
	sexDump := make([]string, len(sex), cap(sex)*2)
	copy(sexDump, sex)
	println("sexDump:", sexDump[1])
	return name
}

//range 打印slice
func rangeSlice(names []string) {
	for _, name := range names {
		print(name)
	}

	for i, name := range names {
		println("i:", i, " name:", name)
	}
}

func getMap() {
	//初始化
	countryMap := make(map[string]string)
	//赋值
	countryMap["China"] = "中国"
	countryMap["American"] = "美国"
	//取值
	println(countryMap["China"])

	for _, country := range countryMap {
		print(country)
	}

	for country := range countryMap {
		println(country, countryMap[country])
	}

	//查看元素是否存在map中
	country, ok := countryMap["France"]
	if ok {
		println("存在集合中", country)
	} else {
		errors.New("不存在集合中")
	}

	//删除
	delete(countryMap, "China")
}

//接口
type Phone interface {
	call() string
}

type NokiaPhone struct {
}

type HtcPhone struct {
}

func (nokiaPhone NokiaPhone) call() string {
	msg := "NokiaPhone call"
	println(msg)
	return msg
}

func (htcPhone HtcPhone) call() string {
	msg := "HtcPhone call"
	println(msg)
	return msg
}

//end
//接口.2
type Pc interface {
	print(msg string) (string, int)
}

type MacPc struct {
}

func (macPc MacPc) print(msg string) (string, int) {
	println(msg)
	return msg, 1
}

func printTime(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println(msg)
	}
}

//channel通道
func getChannel() {
	nums := []int{1, 2, 3}
	// make 的第二个参数指定缓冲区大小
	// 如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
	// 如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；
	// 如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
	channel := make(chan int, 100)

	getSum(nums, channel)
	go getSum(nums, channel)
	println(channel)

	x, y := <-channel, <-channel
	println(x, y)
	close(channel)
}

func getSum(nums []int, channel chan int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	channel <- sum
}

func Square(n int) int {
	//匿名函数
	square := func() int {
		return n * n
	}()

	return square
}

//函数作为参数
func callback(s string, f func(i int) int) {
	println(s, f(9))
}

//计算函数执行时间
func getSubTime() {
	start := time.Now()
	for i := 0; i < 100; i++ {

	}
	end := time.Now()
	duration :=end.Sub(start)
	println("duration:",duration)
}
