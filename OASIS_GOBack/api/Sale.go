// api/users.go

package api

// type User struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// func GetUsers(context *gin.Context) {
// 	context.JSON(http.StatusOK, gin.H{"data": users})
// }

// func CreateUser(context *gin.Context) {
// 	var user User
// 	if err := context.BindJSON(&user); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
// 		return
// 	}
// 	users = append(users, user)
// 	context.Status(http.StatusCreated)
// 	context.JSON(http.StatusOK, gin.H{"data": users})
// }
