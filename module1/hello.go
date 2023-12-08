package module1

import "fmt"

func SayHello2() string {
	return "Hello, world 2!"
}

func sayHellopriv() string {
	fmt.Println("Hello, world! private")
	return ""
}
