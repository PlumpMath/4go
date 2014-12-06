package easy

// Say Hello
func SayHello(name string) string {
	helloMsg := "Hello, " + name
	return helloMsg
}

// Good Day Greet Message
// Used return named parameter
func GreetMsg(name string) (msg string) {
	msg = "Good Day, " + name
	return
}
