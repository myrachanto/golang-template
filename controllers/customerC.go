package controllers

import (
	// "net/http"
	"strconv"
	//"log"
	 "fmt"
	"github.com/labstack/echo"
	//"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/labstack/echo/middleware"
	s "github.com/myrachanto/firsttemp/support"
	"github.com/myrachanto/firsttemp/models"
)
//var GormDB = s.Getconnected()
//todo need to figure out how to pass an open connect with gorm..........................
func GetCustomers(c echo.Context) error {
	
	customers := []*models.Customer{}
	// err := c.Bind(s.GormDB.Find(&Customers))
	//DB.Model(&user).Association("Friends")
	//c.Bind(GormDB.Find(&Customers))
	GormDB := s.Getconnected()
	c.Bind(GormDB.Model(&customers).Find(&customers).Association("Invoices"))
	//err := c.Bind(s.GormDB.Model(&m.Customer{}).Find(Customers))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//err := c.Bind(s.GormDB.Model(&m.Customer{}).Find(Customers))
	
	s.DbClose(GormDB)
	return c.JSON(200, customers)
}
func GetCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	customer := models.Customer{}
	//db.Where("name = ?", "myrachanto").First(&user)
	GormDB := s.Getconnected()
	c.Bind(GormDB.Model(&customer).Where("id = ?", id).First(&customer))
	//c.Bind(GormDB.Model(&Customers).Find(&Customers).Association("Invoices"))
	//err := c.Bind(s.GormDB.Model(&m.Customer{}).Find(Customers))
	s.DbClose(GormDB)
	//fmt.Println(customer, "got get this dude")
	return c.JSON(200, customer)
}

func CreateCustomers(c echo.Context) error {
	
	findcustomer := models.Customer{}
	if err := c.Bind(&findcustomer); err !=nil {
		return err
	}
	GormDB := s.Getconnected()
	name := findcustomer.Name
	email := findcustomer.Email
	phone := findcustomer.Phone
	company := findcustomer.Company
	address := findcustomer.Address
	if name == "" || email == "" || phone == "" || company == "" || address == ""{
		return c.JSON(500, "fill all the fields")
	}
	if s.ValidateEmail(email) == false {
		return c.JSON(501, "check your email")
	}
	
	acustomer := &models.Customer{Name:name, Email:email, Phone:phone, Company:company, Address:address}
	GormDB.Create(&acustomer)
	s.DbClose(GormDB)
	return c.JSON(200, "account created succesifully")
}
func UpdateCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	findcustomer := models.Customer{}
	if err := c.Bind(&findcustomer); err !=nil {
		return err
	}
	GormDB := s.Getconnected()
	ucustomer := models.Customer{}
	GormDB.Model(&ucustomer).Where("id = ?", id).First(&ucustomer)
	name := findcustomer.Name
	email := findcustomer.Email
	phone := findcustomer.Phone
	company := findcustomer.Company
	address := findcustomer.Address
	if name == "" {
		name = ucustomer.Name
	}
	if name == "" {
		name = ucustomer.Name
	} else if email == "" {
		email = ucustomer.Email
	}else if phone == "" {
		phone = ucustomer.Phone
	}else if company == "" {
		company = ucustomer.Company
	} else	if address == "" {
		address = ucustomer.Address
	} else if s.ValidateEmail(email) == false {
		return c.JSON(501, "check your email")
	}

	acustomer := &models.Customer{Name:name, Email:email, Phone:phone, Company:company, Address:address}

	customer := models.Customer{}
	GormDB.Model(&customer).Where("id = ?", id).First(&customer).Update(&acustomer)
	s.DbClose(GormDB)
	return c.JSON(201, "updated succesifully")
}
func DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	customer := models.Customer{}
	//db.Where("name = ?", "myrachanto").First(&user)
	GormDB := s.Getconnected()
	//GormDB.Delete(&customer).Where("id = ?", id).First(&customer)
	GormDB.Model(&customer).Delete(&customer, "id = ?", id)
	//GormDB.Where("age = ?", id).First(&customer).Delete(&customer)
	//explicitelly update deleted at field
	// deleted_at := 
	// acustomer := &models.Customer{Name:name, Email:email, Phone:phone, Company:company, DeletedAt:address}

	// customer := models.Customer{}
	// GormDB.Model(&customer).Where("id = ?", id).First(&customer).Update(&acustomer)
	s.DbClose(GormDB)
	return c.JSON(201, "Deleted succesifully")
}

// type (
// 	user struct {
// 		ID   int    `json:"id"`
// 		Name string `json:"name"`
// 	}
// )

// var (
// 	users = map[int]*user{}
// 	seq   = 1
// )
// func createUser(c echo.Context) error {
// 	u := &user{
// 		ID: seq,
// 	}
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	users[u.ID] = u
// 	seq++
// 	return c.JSON(http.StatusCreated, u)
// }

// func getUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	return c.JSON(http.StatusOK, users[id])
// }

// func updateUser(c echo.Context) error {
// 	u := new(user)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	users[id].Name = u.Name
// 	return c.JSON(http.StatusOK, users[id])
// }

// func deleteUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	delete(users, id)
// 	return c.NoContent(http.StatusNoContent)
// }
