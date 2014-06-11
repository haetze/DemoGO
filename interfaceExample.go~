package main

import "fmt"
//MyInt can be converted to Interface because it implements test()
//MyString it can't because it doesn't the test() method  
type Interface  interface{
	test() string
}

type MyInt struct {
	first int
	second int
}

type MyString struct {
	str string
	MyInt
}


func (d MyInt) test() string{
	return "test"
}


func  run(d Interface) string{
	return d.test()
}

func main(){
	a := MyInt{9, 10}
	b := MyString{"sss", MyInt{12,21}}
	//c := Interface(a)
	fmt.Println(b.first)
	//fmt.Println(a)
	//fmt.Println(run(c))
}
