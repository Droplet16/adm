package user

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type User struct {
	ID       int64  `json:"id" form:"id"`
	Login    string `json:"login" form:"login"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

type UserFilter struct {
	ID       *int64  `form:"id"`
	Login    *string `form:"login"`
	Password *string `form:"password"`
	Email    *string `form:"email"`
	OrderBy  *string `form:"orderby"`
	Limit    *int64  `form:"limit"`
	Offset   *int64  `form:"offset"`
}

type Role struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	AuthItemName string `json:"auth_item_name"`
}

type Link struct {
	UUserID int64 `json:"u_user_id"`
	UROleID int64 `json:"u_role_id"`
}

func RegisterUserHandlers(router *gin.Engine, db *sql.DB) {

	registerGetUsers(router, db)
	registerGetUserByID(router, db)
	registerUpdateUser(router, db)
	registerPostUser(router, db)
	registerDeleteUser(router, db)
	registerGetRolesByUserID(router, db)
	registerPostRole(router, db)
	registerDeleteRole(router, db)
	registerPostLinkingRoleToUser(router, db)

	router.Run("localhost:8080")

}

func checkOrderByColumn(x string) bool {
	cols := []string{"id", "login", "password", "email"}
	for _, n := range cols {
		if x == n {
			return true
		}
	}
	return false
}

/*
func getUsers(c *gin.Context) {
	var users []User

	rows, err := db.Query("SELECT * FROM u_user")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.ID, &usr.Login, &usr.Password, &usr.Email, &usr.URoleID); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, users)
}
*/
