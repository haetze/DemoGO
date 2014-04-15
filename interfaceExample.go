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

func (d MyString) test() string{
	return "test2"
}

func  run(d Interface) string{
	return d.test()
}

func main(){
	a := MyInt(9)
	b := MyString("rrrr")
	c := Interface(a)
	d := Interface(b)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(run(c))
	fmt.Println(run(d))
}
