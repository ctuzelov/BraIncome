package handler

import (
	"braincome/internal/helper"
	"braincome/internal/models"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) GetUsers(c *gin.Context) {
	if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err1 := strconv.Atoi(c.Query("page"))
	if err1 != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(c.Query("startIndex"))

	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
	groupStage := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
		{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
		{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}}}}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}}},
		},
	}

	result, err := helper.UserCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage})

	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
	}

	var allUsers []bson.M
	if err = result.All(ctx, &allUsers); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, allUsers[0])
}

func (h *Handler) GetUser(c *gin.Context) {
	userId := c.Param("user_id")

	if err := helper.MathUserTypeToUid(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var user models.User
	err := helper.UserCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
