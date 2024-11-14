package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if !strings.HasPrefix(token, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        token = strings.TrimPrefix(token, "Bearer ")
        // 验证 token 的逻辑，可以使用JWT或其他方式
        if token != "valid_token" {  // 示例中直接对比固定字符串
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }
        c.Next()
    }
}
