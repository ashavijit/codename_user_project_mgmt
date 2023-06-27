package configs

import (
	"context"
	"fmt"
	"log"
	"project-management/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)


type Database struct {
	client *mongo.Client
}

func MongoConnect() *Database {
	client , err := mongo.NewClient(options.Client().ApplyURI(EnvMONGO_URI()))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!" + EnvMONGO_URI())
	return &Database{client: client}
}

func CreateCollection(db *Database, collectionName string) *mongo.Collection {
	return db.client.Database("projectMngt").Collection(collectionName)
}

func (db *Database) CreateProject(input *model.NewProject) (*model.Project, error) {
	collection := CreateCollection(db, "project")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	project := &model.Project{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		OwnerID:     input.OwnerID,
		Name:        input.Name,
		Description: input.Description,
		Status:      model.StatusNotStarted,
	}

	return project, err
}

func (db *Database) CreateOwner(input *model.NewOwner) (*model.Owner, error) {
	collection := CreateCollection(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	owner := &model.Owner{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	return owner, err
}

func (db *Database) GetOwners() ([]*model.Owner, error) {
	collection := CreateCollection(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, nil)

	if err != nil {
		return nil, err
	}

	var owners []*model.Owner

	for cursor.Next(ctx) {
		var owner *model.Owner
		cursor.Decode(&owner)
		owners = append(owners, owner)
	}

	return owners, err
}

func (db *Database) GetProjects() ([]*model.Project, error) {
	collection := CreateCollection(db, "project")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, nil)

	if err != nil {
		return nil, err
	}

	var projects []*model.Project

	for cursor.Next(ctx) {
		var project *model.Project
		cursor.Decode(&project)
		projects = append(projects, project)
	}

	return projects, err
}


func (db *Database) SingleOwner(ID string) (*model.Owner, error) {
	collection := CreateCollection(db, "owner")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var owner *model.Owner
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&owner)

	return owner, err
}

func (db *Database) SingleProject(ID string) (*model.Project, error) {
	collection := CreateCollection(db, "project")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var project *model.Project
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&project)

	return project, err
}