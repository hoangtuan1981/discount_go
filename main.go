package main

import (
	"Discount/controllers"
	"Discount/middlewares"
	"net/http"

	//"Discount/middleware"
	// Import the database package
	//Discount is declar in `go.mod` file
	"github.com/gin-gonic/gin"
)

func main() {
	//database.ConnectDatabase()
	//database.DB.AutoMigrate(&database.User{})

	r := gin.Default()

	// Endpoint đăng ký người dùng
	r.POST("/register", register)

	// Endpoint đăng nhập
	r.POST("/login", login)

	//r.GET("/users", controllers.GetUsers)
	// //TODO Endpoint yêu cầu xác thực
	// authenticated := r.Group("/api")
	// authenticated.Use(authMiddleware())
	// {
	// 	authenticated.GET("/users", getUsers)
	// }
	r.GET("/users", middlewares.RoleAuthorization("nhanvien"), controllers.GetUsers)

	r.Run(":50917")
}

// Hàm đăng ký người dùng
func register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể mã hóa mật khẩu"})
	// 	return
	// }

	// user := database.User{
	// 	Username: input.Username,
	// 	Password: string(hashedPassword),
	// }

	// if err := database.DB.Create(&user).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Tên người dùng đã tồn tại"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công"})
}

// Hàm đăng nhập
func login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var user database.User
	// if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Tên người dùng hoặc mật khẩu không đúng"})
	// 	return
	// }

	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Tên người dùng hoặc mật khẩu không đúng"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Đăng nhập thành công"})
}

// Middleware xác thực
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Giả sử xác thực bằng token (có thể thay đổi theo yêu cầu)
		token := c.GetHeader("Authorization")
		if token != "valid-token" { // Thay bằng logic xác thực thực tế
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Không có quyền truy cập"})
			c.Abort()
			return
		}
		c.Next()
	}
}
