package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type rwhandler struct {
	context.Context
	*mongo.Collection
}

type Record struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Value string             `json:"value" bson:"value"`
}

var (
	port = flag.Int("p", 8080, "Listening port")
	rwh  *rwhandler
)

func init() {
	flag.Parse()

	rwh = &rwhandler{}

	rwh.Context = context.Background()
	client, err := mongo.Connect(rwh.Context, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	quitOnError(err, "Failed to connect to mongodb")
	err = client.Ping(context.TODO(), readpref.Primary())
	quitOnError(err, "Failed to connect to mongodb")

	rwh.Collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("AmayTestLog")

	log.Println("Connected to mongoDB")
}

func main() {
	router := gin.Default()

	router.GET("/read")

	router.Run(fmt.Sprintf(":%v", *port))
}

func (rwh *rwhandler) ReadHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Argument id can't be empty/null"})
	}

	var res Record
	cur := rwh.Collection.FindOne(rwh.Context, bson.M{"id": id})
	if cur != nil {
		cur.Decode(&res)
	}

	c.JSON(http.StatusOK, res)
	return
}

func (rwh *rwhandler) WriteHandler(c *gin.Context) {
	var record Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	_, err := rwh.Collection.InsertOne(rwh.Context, record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": record.ID})
	return
}

func quitOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func quitOnErrorf(err error, msg string, a ...interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", fmt.Sprintf(msg, a...), err.Error())
	}
}
