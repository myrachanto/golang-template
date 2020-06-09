package support

import (
	 "github.com/jinzhu/gorm"
	 "os"
	  "github.com/joho/godotenv"
	 "log"
	// "fmt"
	m "github.com/myrachanto/firsttemp/models"

	_"github.com/jinzhu/gorm/dialects/mysql"
	// _"github.com/jinzhu/gorm/dialects/postgres"
	// _"github.com/jinzhu/gorm/dialects/sqlite"
	//_"github.com/jinzhu/gorm/dialects/mssql"
)
// type Config struct {
// 	EncryptionKey string
// 	DbHost string
// 	DbPort string
// 	DbName string
// 	DbUsername string
// 	DbPassword string
// }
//var config Config
//var GormDB *gorm.DB
func Configs(){
	GormDB := Getconnected()
// 	//config := m.GetConfig()
// 	//gormParams := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.DbHost, config.DbPort, config.DbName, config.DbUsername, config.DbPassword)
// 	//gormParams := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.DbHost, config.DbPort, config.DbName, config.DbUsername, config.DbPassword)
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	fmt.Println(err)
// 	//, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
// dbURI := "root@/search?charset=utf8&parseTime=True&loc=Local"
// //GormDB, err := gorm.Open("mysql", gormParams)
// GormDB, err := gorm.Open("mysql", dbURI)
// if err != nil {
// 	panic("failed to connect database")
// }

// defer GormDB.Close()

	
	// //port := os.Getenv("PORT")
	// EncryptionKey := os.Getenv("EncryptionKey")
	// DbHost := os.Getenv("DbHost")
	// DbName := os.Getenv("DbName")
	// DbPort := os.Getenv("DbPort")
	// DbUsername := os.Getenv("DbUsername")
	// DbPassword := os.Getenv("DbPassword")
	// config = Config{
	// 	EncryptionKey,
	// 	DbHost,
	// 	DbPort,
	// 	DbName,
	// 	DbUsername,
	// 	DbPassword,
	// }
	
	//migrate tables
	GormDB.AutoMigrate(&m.User{})
	GormDB.AutoMigrate(&m.Auth{})
	GormDB.AutoMigrate(&m.Customer{})
	GormDB.AutoMigrate(&m.Invoice{})
	GormDB.AutoMigrate(&m.InvoiceItem{})
	// GormDB.AutoMigrate(&m.User{}, &m.Auth{}, &m.Customer{},&m.Invoice{},&m.InvoiceItem{})
	///////////////////////////////////////////////////////////////
	//rememdber to include soft delete
	//////////////////////////////////////////////////////////////////////////\
	GormDB.Model(&m.Auth{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	GormDB.Model(&m.Invoice{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
	GormDB.Model(&m.InvoiceItem{}).AddForeignKey("invoice_id", "invoices(id)", "CASCADE", "CASCADE")
	GormDB.Model(&m.Customer{}).Related(&m.Invoice{})
	GormDB.Model(&m.Invoice{}).Related(&m.InvoiceItem{})
	
	DbClose(GormDB)
}
func Getconnected()(GormDB *gorm.DB){
	//, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	dbURI := "root@/search?charset=utf8&parseTime=True&loc=Local"
	//GormDB, err := gorm.Open("mysql", gormParams)
	GormDB, err := gorm.Open("mysql", dbURI)
	if err != nil {
		panic("failed to connect database")
	}

	//defer GormDB.Close()
	return GormDB
}
func DbClose(GormDB *gorm.DB){
	defer GormDB.Close()
}
func Enkey()(EncryptionKey string){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}	
	EncryptionKey = os.Getenv("EncryptionKey")
	return EncryptionKey
}