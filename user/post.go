package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerPostUser(router *gin.Engine, db *sql.DB) {
	router.POST("/users", func(c *gin.Context) {
		var newUser User
		var lastInsertID int64

		// Call BindJSON to bind the received JSON to
		// newUser.
		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		login := newUser.Login
		password := newUser.Password
		email := newUser.Email

		// Add the new user to the DB.

		err := db.QueryRow("INSERT INTO u_user(login, password, email) VALUES($1, $2, $3) returning id;", login, password, email).Scan(&lastInsertID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		newUser.ID = lastInsertID
		c.IndentedJSON(http.StatusCreated, newUser)
	})

}

func registerPostRole(router *gin.Engine, db *sql.DB) {
	router.POST("/roles", func(c *gin.Context) {
		var newRole Role
		var lastInsertID int64

		// Call BindJSON to bind the received JSON to
		// newRole.
		if err := c.BindJSON(&newRole); err != nil {
			return
		}

		name := newRole.Name
		authItemName := newRole.AuthItemName

		// Add the new role to the DB.

		err := db.QueryRow("INSERT INTO u_role(name, auth_item_name) VALUES($1, $2) returning id;", name, authItemName).Scan(&lastInsertID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		newRole.ID = lastInsertID
		c.IndentedJSON(http.StatusCreated, newRole)
	})

}

func registerPostLinkingRoleToUser(router *gin.Engine, db *sql.DB) {
	router.POST("users/roles", func(c *gin.Context) {
		var newLink Link

		if err := c.BindJSON(&newLink); err != nil {
			return
		}

		uUserID := newLink.UUserID
		uRoleID := newLink.UROleID

		// Add the new link to the DB.

		_, err := db.Exec("INSERT INTO u_user_role(u_user_id, u_role_id) VALUES($1, $2);", uUserID, uRoleID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.IndentedJSON(http.StatusCreated, newLink)
	})

}
