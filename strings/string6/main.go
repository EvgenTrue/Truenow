package main
import (
    "fmt"
    "regexp"
)

func main() {
    fmt.Println(isEmailValid("test44@gmail.com")) // true
    fmt.Println(isEmailValid("test$@gmail.com")) // true -- expected "false" 
}


// isEmailValid checks if the email provided is valid by regex.
func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegex.MatchString(e)
}
