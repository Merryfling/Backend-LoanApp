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

    // 生成 token
    token, err := config.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, api.RegisterResponse{
        Status:  "success",
        Message: "User registered successfully",
        UserId:  fmt.Sprintf("%d", user.ID),
        Token:   token,  // 返回生成的 token
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

    // 生成 token
    token, err := config.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, api.LoginResponse{
        Status:  "success",
        Message: "Login successful",
        UserId:  fmt.Sprintf("%d", user.ID),
        Token:   token,  // 返回生成的 token
    })
}

// 用户资料查询
func GetProfile(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
        return
    }

    var user model.User
    // 根据 userID 查询用户信息
    if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    c.JSON(http.StatusOK, api.UserProfileResponse{
        UserId: user.ID,
        Name: user.Name,
        Phone: user.Phone,
        IdNumber: user.IdNumber
    })
}