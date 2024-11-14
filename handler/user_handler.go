package handler

import (
    "loanapp/config"
    "loanapp/model"
    "loanapp/api"
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {
    var req api.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    user := model.User{
        Phone:    req.Phone,
        Password: req.Password,  // 实际开发中要对密码加密
    }
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, api.RegisterResponse{
        Status:  "success",
        Message: "User registered successfully",
        UserId:  fmt.Sprintf("%d", user.ID),   // 转换为字符串
        Token:   "valid_token",
    })
}

// 用户登录
func Login(c *gin.Context) {
    var req api.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    var user model.User
    if err := config.DB.Where("phone = ? AND password = ?", req.Phone, req.Password).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, api.LoginResponse{
        Status:  "success",
        Message: "Login successful",
        UserId:  fmt.Sprintf("%d", user.ID),   // 转换为字符串
        Token:   "valid_token",
    })
}
