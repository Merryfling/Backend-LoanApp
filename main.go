package main

import (
    "loanapp/config"
    "loanapp/handler"
    "loanapp/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    config.InitDB()  // 初始化数据库

    router := gin.Default()

    // 用户注册和登录路由
    router.POST("/api/v1/register", handler.Register)
    router.POST("/api/v1/login", handler.Login)

    // 贷款申请路由（需要认证）
    loan := router.Group("/api/v1/loan", middleware.AuthMiddleware())
    {
        loan.POST("/apply", handler.ApplyLoan)
        loan.GET("/status/:application_id", handler.GetLoanStatus)
    }

    router.Run(":666")  // 启动服务器
}
