package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerUpdateUser(router *gin.Engine, db *sql.DB) {
	router.PUT("/users/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")

		// Call BindJSON to bind the received JSON to
		// user.
		if err := c.BindJSON(&user); err != nil {
			return
		}

		login := user.Login
		password := user.Password
		email := user.Email

		//row := db.QueryRow("UPDATE u_user SET login = $2, password = $3, email = $4, u_role_id = $5 where id = $1;", id, login, password, email, uroleid)
		_, err := db.Exec("UPDATE u_user SET login = $2, password = $3, email = $4 where id = $1;", id, login, password, email)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		newId, err := strconv.Atoi(id)
		user.ID = int64(newId)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.IndentedJSON(http.StatusCreated, user)
	})
}
