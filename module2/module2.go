package module2

import "goProjectTemplate/module1"

func SayGreetings() string {
	println("Greetings, world!")
	println(module1.SayHello())
	return "Greetings, world!"
}
