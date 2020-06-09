package controllers

import (
	// "net/http"
	"strconv"
	//"log"
	 "fmt"
	 jwt "github.com/dgrijalva/jwt-go"
	  //"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo"
	//"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/labstack/echo/middleware"
	s "github.com/myrachanto/firsttemp/support"
	"github.com/myrachanto/firsttemp/models"
)
type Loginc struct{
	Email string
	Password string
	Confirmpassword string
}
func Register(c echo.Context) error {
	
	fusuer := models.User{}
	if err := c.Bind(&fusuer); err !=nil {
		return err
	}
	GormDB := s.Getconnected()
	fname := fusuer.FName
	lname := fusuer.LName
	uname := fusuer.UName
	phone := fusuer.Phone
	address := fusuer.Address
	dob := fusuer.Dob
	picture := fusuer.Picture
	email := fusuer.Email
	pword := fusuer.Password
	if fname == "" ||lname == "" || uname == "" || phone == "" || address == "" || picture == "" || email == "" || pword == "" {
		return c.JSON(500, "fill all the fields")
	}
	if s.ValidateEmail(email) == false {
		return c.JSON(501, "check your email")
	}
	
	// p, err := bcrypt.GenerateFromPassword([]byte(pword), 10)
	// if err != nil {
	// 	return err
	// }
	// password := string(p)
	encyKey := s.Enkey()
	p := s.Encrypt([]byte(pword), encyKey)
	password := string(p)
	
	auser := &models.User{FName:fname,LName:lname,UName:uname, Phone:phone, Address:address, Dob:dob,Picture:picture,Email:email, Password:password}
	GormDB.Create(&auser)
	s.DbClose(GormDB)
	return c.JSON(200, "Accout succesifully created")
}
func Login(c echo.Context) error {
	auser := Loginc{}
	if err := c.Bind(&auser); err !=nil {
		return err
	}
	GormDB := s.Getconnected()
	email := auser.Email
	pword := auser.Password
	cpword := auser.Confirmpassword
	if email == "" || pword == "" || cpword == "" {
		return c.JSON(500, "fill all the fields")
	}
	if s.ValidateEmail(email) == false {
		return c.JSON(501, "check your email")
	}
	if pword == cpword {
		return c.JSON(501, "your passwords do not match")
	}
	
	encyKey := s.Enkey()
	p := s.Encrypt([]byte(pword), encyKey)
	password := string(p) 

	fusuer := models.User{}
	GormDB.Model(&fusuer).Where("email = ?", email).First(&fusuer)
	if fusuer.Email != "" {
		if fusuer.Password == password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
					"email": fusuer.Email,
					"password": fusuer.Password,
				})
				d := fusuer.ID
				e := string(d) 
				t, err := token.SignedString(token)
					if err != nil {
						return err
					}	
				id, err := strconv.Atoi(e)
				if err != nil {
					fmt.Println(err)
				}
				auth := &models.Auth{UserID:id, Token:t}
				GormDB.Create(&auth)

				return c.JSON(200, "wellcome back")
		}
		
	return c.JSON(501, "wrong credetials")
	}
	s.DbClose(GormDB)
	//fmt.Println(customer, "got get this dude")
	return c.JSON(501, "wrong credetials")
}

func Logout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	if AuthExist(id) {
	GormDB := s.Getconnected()
	fauth := models.Auth{}
	GormDB.Model(&fauth).Delete(&fauth, "id = ?", id)
	s.DbClose(GormDB)
	}
	return c.JSON(501, "something went wrong")
 }
 func GetUsers(c echo.Context) error {
	
	users := []*models.User{}
	GormDB := s.Getconnected()
	c.Bind(GormDB.Model(&users).Find(&users).Association("Invoices"))
	
	s.DbClose(GormDB)
	return c.JSON(200, users)
}
 func AuthExist(id int) bool {
	auth := models.Auth{}
	GormDB := s.Getconnected()
	if GormDB.First(&auth, "id =?", id).RecordNotFound(){
	   return false
	}
	s.DbClose(GormDB)
	return true
	
}
func UserExist(email string) bool {
	auser := models.User{}
	GormDB := s.Getconnected()
	if GormDB.First(&auser, "email =?", email).RecordNotFound(){
	   return false
	}
	s.DbClose(GormDB)
	return true
	
}

