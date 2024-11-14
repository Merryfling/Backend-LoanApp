package config

import (
    "loanapp/model"         // 导入 model 包
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("loanapp.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    // 自动迁移数据库结构
    DB.AutoMigrate(&model.User{}, &model.LoanApplication{})
}
