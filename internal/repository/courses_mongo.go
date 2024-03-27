package repository

import (
	"braincome/internal/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CoursesMongo struct {
	db *mongo.Client
}

func NewCoursesMongo(db *mongo.Client) *CoursesMongo {
	return &CoursesMongo{db: db}
}

func (c *CoursesMongo) Insert(course models.Course) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := c.db.Database("cluster0").Collection("courses")

	defer cancel()

	course.ID = primitive.NewObjectID()
	fmt.Println(course.ID)

	// Insert the user document into the MongoDB collection
	_, err := collection.InsertOne(ctx, course)
	if err != nil {
		return err
	}

	return nil
}

func (c *CoursesMongo) Update(course models.Course) error {
	return nil
}

func (c *CoursesMongo) GetAll() ([]models.Course, error) {
	collection := c.db.Database("cluster0").Collection("courses")

	var courses []models.Course
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoursesMongo) GetSeveral(limit int, offset int) ([]models.Course, error) {
	collection := c.db.Database("cluster0").Collection("courses")

	var courses []models.Course
	sortOptions := options.Find().SetSort(bson.D{{Key: "rating", Value: offset}})
	limitOptions := options.Find().SetLimit(int64(limit))

	cursor, err := collection.Find(context.Background(), bson.M{}, sortOptions, limitOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoursesMongo) GetById(id primitive.ObjectID) (models.Course, error) {
	collection := c.db.Database("cluster0").Collection("courses")

	var course models.Course
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&course)
	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func (c *CoursesMongo) GetByInstructor(instructor string) ([]models.Course, error) {
	collection := c.db.Database("cluster0").Collection("courses")

	var courses []models.Course
	filter := bson.M{"instructor.name": instructor}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoursesMongo) GetByRating() ([]models.Course, error) {
	// collection := c.db.Database("cluster0").Collection("courses")

	// var courses []models.Course
	// cursor, err := collection.Find(context.Background(), bson.M{}).Sort("rating").Limit(10)
	// if err != nil {
	// 	return nil, err
	// }
	// defer cursor.Close(context.Background())

	// if err := cursor.All(context.Background(), &courses); err != nil {
	// 	return nil, err
	// }

	return []models.Course{}, nil
}

func (c *CoursesMongo) GetByEnrolled() ([]models.Course, error) {
	// collection := c.db.Database("cluster0").Collection("courses")

	// var courses []models.Course
	// cursor, err := collection.Find(context.Background(), bson.M{}).Sort("enrolled").Limit(10)
	// if err != nil {
	// 	return nil, err
	// }
	// defer cursor.Close(context.Background())

	// if err := cursor.All(context.Background(), &courses); err != nil {
	// 	return nil, err
	// }

	return []models.Course{}, nil
}

func (c *CoursesMongo) GetByCategory(category string) ([]models.Course, error) {
	collection := c.db.Database("cluster0").Collection("courses")

	var courses []models.Course
	filter := bson.M{"categories": category}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoursesMongo) GetByTitle(title string) ([]models.Course, error) {
	collection := c.db.Database("cluster0").Collection("courses")

	var courses []models.Course
	filter := bson.M{"title": primitive.Regex{Pattern: title, Options: "i"}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoursesMongo) GetByLanguage(language string) ([]models.Course, error) {
	collection := c.db.Database("cluster0").Collection("courses")

	var courses []models.Course
	filter := bson.M{"language": language}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &courses); err != nil {
		return nil, err
	}

	return courses, nil
}
