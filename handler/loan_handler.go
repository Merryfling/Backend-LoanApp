package handler

import (
    "loanapp/config"
    "loanapp/model"
    "loanapp/api"
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
)

// 创建贷款申请
func ApplyLoan(c *gin.Context) {
    var req api.LoanApplicationRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    application := model.LoanApplication{
        UserID:      1,                          // 示例值
        Income:      float64(req.Income),        // 类型转换
        LoanAmount:  float64(req.LoanAmount),    // 类型转换
        LoanTerm:    int(req.LoanTerm),          // 类型转换
        LoanPurpose: req.LoanPurpose,
        Status:      "Pending",
    }
    if err := config.DB.Create(&application).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan application"})
        return
    }

    c.JSON(http.StatusOK, api.LoanApplicationResponse{
        Status:         "success",
        Message:        "Loan application submitted",
        ApplicationId:  fmt.Sprintf("%d", application.ID),  // 转换为字符串
        ApplicationStatus: application.Status,
    })
}

// 查询贷款状态
func GetLoanStatus(c *gin.Context) {
    applicationID := c.Param("application_id")
    var application model.LoanApplication
    if err := config.DB.Where("id = ?", applicationID).First(&application).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Loan application not found"})
        return
    }

    c.JSON(http.StatusOK, api.LoanStatusResponse{
        Status:         "success",
        Message:        "Loan application status retrieved",
        ApplicationId:  fmt.Sprintf("%d", application.ID),  // 转换为字符串
        LoanStatus:     application.Status,
        Score:          int32(application.Score),           // 转换为 int32
        Comments:       "Loan application is under review",
    })    
}
