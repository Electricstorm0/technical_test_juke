package controller

import (
	"net/http"
	"strconv"
	"technical_test/src/model"
	"technical_test/src/service"

	"github.com/gin-gonic/gin"
)

// CreateEmployeeController godoc
// @Summary Create new employee
// @Description Add a new employee to the database
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Employee Data"
// @Success 201 {object} model.Employee
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /employees [post]
func CreateEmployeeController(c *gin.Context) {
	var emp model.Employee
	if err:= c.BindJSON(&emp)
	err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "invalid request body",
			"status": "fail",
	})
	return
	}
	id,err:= service.CreateEmployeeServices(&emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"data": id,
		"message": "add employee successfully",
		"status": "success",
	})

}


// GetAllController godoc
// @Summary Get all employees
// @Description Retrieve all employee data
// @Tags employees
// @Produce json
// @Success 200 {object} model.Employee
// @Failure 500 {object} map[string]interface{}
// @Router /employees [get]
func GetAllController(c *gin.Context) {
	employee,err:= service.GetAllService()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": employee,
		"status": "success",
	})

}


// GetByIdController godoc
// @Summary Get employee by ID
// @Description Retrieve a single employee by ID
// @Tags employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.Employee
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /employees/{id} [get]
func GetByIdController(c *gin.Context) {
	id,err:=strconv.Atoi(c.Param("id"))
	if err != nil {
	c.JSON(http.StatusBadRequest,gin.H{
		"message": "invalid Id",
		"status": "fail",
	})
	return
	}
	
	employee,err:= service.GetByIdService(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": employee,
		"status": "success",
	})
}


// UpdateController godoc
// @Summary Update employee
// @Description Update employee data by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID" example(1)
// @Param employee body model.UpdatedEmployee true "Updated Employee Data"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "invalid id or body"
// @Failure 500 {object} map[string]interface{} "server error"
// @Router /employees/{id} [put]
func UpdateController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Id",
			"status":  "fail",
		})
		return
	}

	var emp model.Employee
	if err := c.BindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"status":  "fail",
		})
		return
	}

	_, err = service.UpdateServices(int64(id), emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "employee updated successfully",
		"status":  "success",
	})
}



// DeletedController godoc
// @Summary Delete employee
// @Description Delete employee by ID
// @Tags employees
// @Produce json
// @Param id path int true "Employee ID" example(1)
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "invalid id"
// @Failure 500 {object} map[string]interface{} "server error"
// @Router /employees/{id} [delete]
func DeletedController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Id",
			"status":  "fail",
		})
		return
	}

	_, err = service.DeleteServices(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "employee deleted successfully",
		"status":  "success",
	})
}

