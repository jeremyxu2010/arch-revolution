package dal

import (
	"fmt"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/config"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	dbConfig := config.GetAppConfig().Application.Datasource
	dbURL := fmt.Sprintf("%s:%s@%s", dbConfig.Username, dbConfig.Password, dbConfig.URL)
	db, err = gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&model.User{})
}
