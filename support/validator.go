package support

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	//"time"

	_"encoding/json"
	// "github.com/jinzhu/gorm"
	// "os"
	// "github.com/joho/godotenv"
	// "log"
	// "fmt"
)
func Ifempty(t string) bool{
	if t == ""{
		return false
	}
	return true
}
func Notcornfirm(m, n string) bool{
	if m == n {
		return true
	}
	return false
}
//figure out how to get the length of a string.........................
// func maxlen(a string, b int) bool {
// 	if 
// }
func ValidateEmail(email string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(email)
	return
}
func ValidatePassword(password string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(password)
	return
}
func HashPassword(password string)(string, error){
/////////////////////////////////////////////
//hash, _ := HashPassword(password)
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(pwd), err
	
}