package main

// TODO: set up tests for the web server and database

import (
	"net/http"
	"time"

	//"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/plexlad/gardi/server/lib"
)

const (
	CollectionInstances = "instances"
	CollectionSchemas   = "schemas"
)

type NewSchemaRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewInstanceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SchemaID    string `json:"schema_id"`
}

func main() {
	db := NewJsonDB("./data")

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	u := router.Group("/:user")
	schemas := u.Group("/schemas")
	instances := u.Group("/instances")

	schemas.GET("/:id", func(c echo.Context) error {
		user := c.Param("user")
		id := c.Param("id")

		var schema lib.Schema
		err := db.Get(CollectionSchemas, user, id, &schema)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, schema)
	})

	schemas.POST("/new", func(c echo.Context) error {
		user := c.Param("user")

		var req NewSchemaRequest

		// TODO: fix issues with USER and ID vs invalid JSON
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		schemaID := uuid.New().String()
		schema := lib.Schema{
			ID:          schemaID,
			Version:     1,
			UserVersion: 1,
			Name:        req.Name,
			Description: req.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := db.Set(CollectionSchemas, user, schemaID, schema)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, schema)
	})

	schemas.POST("/save", func(c echo.Context) error {
		user := c.Param("user")

		var req lib.Schema

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		err := db.Set(CollectionSchemas, user, req.ID, req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.String(http.StatusOK, "schema saved")
	})

	schemas.GET("", func(c echo.Context) error {
		user := c.Param("user")

		schemaIDs, err := db.List(CollectionSchemas, user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, schemaIDs)
	})

	instances.GET("/:id", func(c echo.Context) error {
		user := c.Param("user")
		id := c.Param("id")

		var instance lib.Instance
		err := db.Get(CollectionInstances, user, id, &instance)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, instance)
	})

	instances.POST("/new", func(c echo.Context) error {
		user := c.Param("user")

		var req NewInstanceRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		var schema lib.Schema
		err := db.Get(CollectionSchemas, user, req.SchemaID, &schema)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}

		instanceID := uuid.New().String()
		instance := lib.Instance{
			ID:          instanceID,
			SchemaID:    req.SchemaID,
			Name:        req.Name,
			Description: req.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err = db.Set(CollectionInstances, user, instanceID, instance)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, instance)
	})

	instances.POST("/save", func(c echo.Context) error {
		user := c.Param("user")

		var req lib.Instance

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		err := db.Set(CollectionInstances, user, req.ID, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.String(http.StatusOK, "instance saved")
	})

	instances.GET("", func(c echo.Context) error {
		user := c.Param("user")

		instanceIDs, err := db.List(CollectionInstances, user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, instanceIDs)
	})

	// run server with error checking
	router.Logger.Fatal(router.Start(":5499"))
}

// TODO: Use this function to make code nice to look at
func httpError(c echo.Context, code int, err error) error {
	return c.JSON(code, echo.Map{"error": err.Error()})
}
