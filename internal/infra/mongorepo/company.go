package mongorepo

import (
	"context"
	"log"
	"os"
	"time"
	"xm-challenge/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COMPANYCOLLECTION = "company"
)

var db *mongo.Database

type CompanyRepo struct {
	database *mongo.Database
}

func NewRepo(mongoURL string) (*CompanyRepo, error) {
	log.Println("Starting DB")
	database := os.Getenv("DATABASE")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}
	log.Println("Starting at " + mongoURL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if database == "" {
		database = "company-api"
	}
	db = client.Database(database)
	if err != nil {
		log.Panicln(err.Error())
		return nil, err

	}
	return &CompanyRepo{
		database: db,
	}, nil
}

func (cRepo *CompanyRepo) CreateOrganization(ctx context.Context, companyGot domain.Company) (companyCreated domain.Company, err error) {
	res, err := cRepo.database.Collection(COMPANYCOLLECTION).InsertOne(context.TODO(), companyGot)
	if err == nil {
		id := res.InsertedID.(string)
		companyGot.ID = id
	}
	return companyGot, err
}

func (cRepo *CompanyRepo) UpdateOrg(ctx context.Context, orgToUpdate domain.Company) (orgUpdated domain.Company, err error) {
	filter := bson.D{primitive.E{Key: "_id", Value: orgToUpdate.ID}}
	update := bson.M{"$set": bson.M{"name": orgToUpdate.Name, "description": orgToUpdate.Description, "employees": orgToUpdate.Employees, "registered": orgToUpdate.Registered, "company_type": orgToUpdate.CompanyType}}
	// resp, erro := db.Collection(COLLECTION).UpdateOne(context.TODO(), filter, update)
	resp, erro := cRepo.database.Collection(COMPANYCOLLECTION).UpdateOne(context.TODO(), filter, update)
	var b []byte
	resp.UnmarshalBSON(b)
	return orgToUpdate, erro
}

func (cRepo *CompanyRepo) DeleteOrg(ctx context.Context, id string) (orgsDeleted int64, err error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	dr, err := cRepo.database.Collection(COMPANYCOLLECTION).DeleteOne(context.TODO(), filter)
	return dr.DeletedCount, err
}

func (cRepo *CompanyRepo) GetOrgByName(ctx context.Context, name string) (counted int, orgsFound []*domain.Company, err error) {
	filter := bson.D{primitive.E{Key: "Name", Value: name}}
	cur, erro := cRepo.database.Collection(COMPANYCOLLECTION).Find(context.TODO(), filter)
	var results []*domain.Company
	for cur.Next(context.TODO()) {
		var elem domain.Company
		cur.Decode(&elem)
		results = append(results, &elem)
	}
	found := len(results)
	// Close the cursor once finished
	cur.Close(context.TODO())
	return found, results, erro
}
