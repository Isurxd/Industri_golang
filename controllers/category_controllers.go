package controllers

import (
	"crud-api/configs"
	"crud-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Read All Categorys
func ReadAllCategorys(c echo.Context) (err error) {
	var responses []models.CategoryResponse

	// Buat koneksi ke database
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const readAllCategorysQuery = `
	SELECT
		id, Category_name
	FROM
		Categories
	`

	rows, err := db.QueryContext(c.Request().Context(), readAllCategorysQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading all Categorys!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	for rows.Next() {
		var response models.CategoryResponse

		err = rows.Scan(
			&response.ID,
			&response.CategoryName,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Error reading all Categorys!",
				"page":    nil,
				"data":    nil,
				"error":   err.Error(),
			})
		}

		responses = append(responses, response)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success reading all Categorys!",
		"page":    nil,
		"data":    responses,
		"error":   nil,
	})
}

// Read Detail Category
func ReadDetailCategorys(c echo.Context) (err error) {
	var response models.CategoryResponse

	// Buat koneksi ke database
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error parsing parameter to integer!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	const readDetailCategoryQuery = `
	SELECT
		Categorys.id, Category_name, Categorys.price, categories.category_name, Categorys.description
	FROM
		Categorys
		INNER JOIN categories ON Category.category_id = category_id = categories.id
	WHERE
		Categorys.id = ?
		
	`
	// LEFT JOIN categories ON Categorys.category_id = categories.id

	row := db.QueryRowContext(c.Request().Context(), readDetailCategoryQuery, id)

	err = row.Scan(
		&response.ID,
		&response.CategoryName,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading detail Category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successss all Categorys!",
		"page":    nil,
		"data":    response,
		"error":   nil,
	})
}

// Create Category
func CreateCategorys(c echo.Context) (err error) {
	var request models.CategoryRequest

	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{

			"message": "Error binding request",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	db, err := configs.ConnectDatabase()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{

			"message": "Error binding request",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const createCategoryQuery = `
	INSERT INTO categories
	( category_name)
	VALUES
	(? )
	`

	fmt.Println(request)

	_, err = db.ExecContext(c.Request().Context(), createCategoryQuery,
		request.CategoryName,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error creating data Category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Succes creating data Category!",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})

}

// Update Category
func UpdateCategorys(c echo.Context) (err error) {
	var request models.CategoryRequest

	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{

			"message": "Error binding request",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	db, err := configs.ConnectDatabase()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{

			"message": "Error binding request",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const updateCategoryQuery = `
	UPDATE Categories set category_name = ?,
	    where id = ?
	`

	fmt.Println(request)

	_, err = db.ExecContext(c.Request().Context(), updateCategoryQuery,
		request.CategoryName,
		request.ID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error update data Category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Succes update data Category!",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})

}

// Delete Category

func DeleteCategorys(c echo.Context) (err error) {
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Succes update data Category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Failded  converting id",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	const deleteCategoryQuery = `
	DELETE
	FROM 
	Categories 
	where
	id = ?`

	_, err = db.ExecContext(c.Request().Context(), deleteCategoryQuery, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Failded  delete Category",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Delete Category",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})
}
