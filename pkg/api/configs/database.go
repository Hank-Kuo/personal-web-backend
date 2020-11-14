package configs

import (
	models "WebBackend/pkg/api/core/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func ConnectDB() *gorm.DB {
	db, err = gorm.Open("sqlite3", "../../gorm.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	db.AutoMigrate(&models.Person{}, &models.User{}, &models.Company{}, &models.Peo{})
	return db
}

func CloseDB() error {
	return db.Close()
}

func GetServerConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)
	serverConfig["HOST"] = "127.0.0.1"
	serverConfig["PORT"] = "8080"
	// serverConfig["VIEWS_PATTERN"] = "easy-gin/views"
	serverConfig["ENV"] = "debug"
	return serverConfig
}

func GetSecretKey() []byte {
	var jwtSecret = []byte("secret")
	return jwtSecret
}

/*
func GetDbConfig() map[string]string {
	dbConfig := make(map[string]string)
	dbConfig["DB_HOST"] = "127.0.0.1" //主机
	dbConfig["DB_PORT"] = "3306"      //端口
	dbConfig["DB_NAME"] = "golang"    //数据库
	dbConfig["DB_USER"] = "root"      //用户名
	dbConfig["DB_PWD"] = ""           //密码
	dbConfig["DB_CHARSET"] = "utf8"
	dbConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return dbConfig
}
*/
