package controllers

import (
	"assignment-project/database"
	"assignment-project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateStudents(ctx *gin.Context) {
	db := database.GetDB()

	var newStudents models.Students

	if err := ctx.BindJSON(&newStudents); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Students := models.Students{
		Name: newStudents.Name,
		Age:    newStudents.Age,
		Scores:        newStudents.Scores,
	}

	err := db.Create(&Students).Error

	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Students,
	})
}

func GetAllStudents(ctx *gin.Context) {
	db := database.GetDB()
	var results = []models.Students{}

	res := db.Preload("Scores").Find(&results)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Studentss"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

func GetStudentsById(ctx *gin.Context) {
	ID := ctx.Param("ID")
	db := database.GetDB()

	var results = models.Students{}

	id, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := db.Where(models.Students{ID: uint(id)}).Preload("Scores").First(&results)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

// gorm example with db transaction
func UpdateStudents(ctx *gin.Context) {
	ID := ctx.Param("ID")
	db := database.GetDB()

	var newStudents models.Students

	if err := ctx.BindJSON(&newStudents); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	Students := models.Students{
		Name: newStudents.Name,
		Age:    newStudents.Age,
		Scores:        newStudents.Scores,
	}

	// Fetch the existing Students from the database along with its associated Scores
	existingStudents := models.Students{}
	res := db.Preload("Scores").First(&existingStudents, uint(id))
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the Students"})
		return
	}

	// Start a transaction
	tx := db.Begin()

	// Update the fields of the fetched Students
	existingStudents.Name = Students.Name
	existingStudents.Age = Students.Age

	// Update the fields of the associated Scores
	for i, newItem := range Students.Scores {
		if i < len(existingStudents.Scores) {
			existingStudents.Scores[i].AssignmentTitle = newItem.AssignmentTitle
			existingStudents.Scores[i].Description = newItem.Description
			existingStudents.Scores[i].Score = newItem.Score
			// Save the Item update within the transaction
			if err := tx.Save(&existingStudents.Scores[i]).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update an item"})
				return
			}
		} else {
			// If the newItem is a new Item, add it to the existingStudents.Scores slice
			existingStudents.Scores = append(existingStudents.Scores, newItem)
			// Save the new Item within the transaction
			if err := tx.Create(&existingStudents.Scores[i]).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new item"})
				return
			}
		}
	}

	// Save the changes to the Students within the transaction
	if err := tx.Save(&existingStudents).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Students"})
		return
	}

	// Commit the transaction
	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    existingStudents,
	})
}

func DeleteStudents(ctx *gin.Context) {
	ID := ctx.Param("ID")
	db := database.GetDB()

	var Students models.Students

	id, err := strconv.Atoi(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the Students with the given ID
	res := db.Preload("Scores").First(&Students, id)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the Students"})
		return
	}

	// Start a transaction
	tx := db.Begin()

	// Delete the associated Scores within the transaction
	if err := tx.Delete(&Students.Scores).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the associated Scores"})
		return
	}

	// Delete the Students from the database within the transaction
	if err := tx.Delete(&Students).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the Students"})
		return
	}

	// Commit the transaction
	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
	})
}
