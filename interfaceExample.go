package main

import "fmt"
//MyInt can be converted to Interface because it implements test()
//MyString it can't because it doesn't the test() method  
type Interface  interface{
	test() string


}

type MyInt int
type MyString string

func (d MyInt) test() string{
	return "test"
}

func main(){
	a := MyInt(9)
	b := MyString("test")
	c := Interface(a)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)
}
