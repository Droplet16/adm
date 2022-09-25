package user

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerGetUsers(router *gin.Engine, db *sql.DB) {
	router.GET("/users", func(c *gin.Context) {
		var u UserFilter
		var users []User

		if err := c.ShouldBind(&u); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		var orderBy string

		if u.OrderBy != nil && checkOrderByColumn(*u.OrderBy) {
			orderBy = *u.OrderBy
		} else {
			orderBy = "login"
		}

		query := fmt.Sprintf(`SELECT * from u_user
						  WHERE ($1::varchar IS NULL OR login = $1::varchar)
							AND ($2::varchar IS NULL OR password = $2::varchar)
							AND ($3::varchar IS NULL OR email = $3::varchar)
						  ORDER BY %s
						  LIMIT $4 OFFSET $5`, orderBy)

		rows, err := db.Query(query, u.Login, u.Password, u.Email, u.Limit, u.Offset)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var usr User
			if err := rows.Scan(&usr.ID, &usr.Login, &usr.Password, &usr.Email); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
			users = append(users, usr)
		}
		if err := rows.Err(); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.IndentedJSON(http.StatusOK, users)

	})
}

func registerGetUserByID(router *gin.Engine, db *sql.DB) {
	router.GET("/users/:id", func(c *gin.Context) {
		var usr User
		id := c.Param("id")

		row := db.QueryRow("SELECT * FROM u_user WHERE id = $1", id)
		if err := row.Scan(&usr.ID, &usr.Login, &usr.Password, &usr.Email); err != nil {
			if err == sql.ErrNoRows {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "getUserByID " + id + ": no such user"})
				return
			}
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, usr)
	})
}

func registerGetRolesByUserID(router *gin.Engine, db *sql.DB) {
	router.GET("/users/:id/roles", func(c *gin.Context) {
		var roles []Role
		userID := c.Param("id")

		rows, err := db.Query(`SELECT r.id, r.name, r.auth_item_name
		                         FROM u_user_role ur
		                         LEFT JOIN u_role r ON r.id = ur.u_role_id
							     LEFT JOIN u_user u ON u.id = ur.u_user_id
                                WHERE u.id = $1`, userID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var role Role
			if err := rows.Scan(&role.ID, &role.Name, &role.AuthItemName); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
			roles = append(roles, role)
		}
		if err := rows.Err(); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.IndentedJSON(http.StatusOK, roles)

	})
}

func registerGetRoles(router *gin.Engine, db *sql.DB) {
	router.GET("/roles", func(c *gin.Context) {
		var r RoleFilter
		var roles []Role

		if err := c.ShouldBind(&r); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		var orderBy string

		if r.OrderBy != nil && checkOrderByColumnRoles(*r.OrderBy) {
			orderBy = *r.OrderBy
		} else {
			orderBy = "name"
		}

		query := fmt.Sprintf(`SELECT * from u_role
						  WHERE ($1::varchar IS NULL OR name = $1::varchar)
							AND ($2::varchar IS NULL OR auth_item_name = $2::varchar)
						  ORDER BY %s
						  LIMIT $3 OFFSET $4`, orderBy)

		rows, err := db.Query(query, r.Name, r.AuthItemName, r.Limit, r.Offset)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var role Role
			if err := rows.Scan(&role.ID, &role.Name, &role.AuthItemName); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
			roles = append(roles, role)
		}
		if err := rows.Err(); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.IndentedJSON(http.StatusOK, roles)

	})
}
