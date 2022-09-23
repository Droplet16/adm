package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerDeleteUser(router *gin.Engine, db *sql.DB) {
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM u_user where id = $1", id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

	})
}

func registerDeleteRole(router *gin.Engine, db *sql.DB) {
	router.DELETE("/roles/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM u_role where id = $1", id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

	})
}
