package easy

import "fmt"

// Say Hello
func SayHello(name string) string {
	helloMsg := "Hello, " + name
	fmt.Printf(helloMsg)
	return helloMsg
}

// Good Day Greet Message
// Used return named parameter
func GreetMsg(name string) (msg string) {
	msg = "Good Day, " + name
	return
}
