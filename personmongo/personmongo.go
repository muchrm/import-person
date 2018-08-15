package personmongo

import (
	"context"

	"github.com/muchrm/import-person/config"
	"github.com/muchrm/import-person/person"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func getConnection() (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), config.GetMongoHost(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.GetMongoDB())
	return db, nil
}
func AddPerson(person person.Person) error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	coll := db.Collection("people")
	collHistoryEducation := db.Collection("historyeducations")
	collHistoryWork := db.Collection("historyworks")
	result, err := coll.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("officerCode", int32(person.OfficerCode)),
			bson.EC.String("officerName", person.OfficerName),
			bson.EC.String("officerSurname", person.OfficerSurname),
			bson.EC.String("officerNameEng", person.OfficerNameEng),
			bson.EC.String("officerSurnameEng", person.OfficerSurnameEng),
			bson.EC.String("officerPosition", person.OfficerPosition),
			bson.EC.String("officerType", person.OfficerType),
			bson.EC.String("officerLogin", person.OfficerLogin),
			bson.EC.String("officerStatus", person.OfficerStatus),
			bson.EC.String("officerEmail", person.Email),
		))
	if err != nil {
		return err
	}
	for _, historyeducation := range person.HistoryEducations {
		collHistoryEducation.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.String("level", historyeducation.LevelName),
				bson.EC.String("major", historyeducation.MajorName),
				bson.EC.String("place", historyeducation.PlaceName),
				bson.EC.String("country", historyeducation.CountryName),
				bson.EC.Int32("graduatedYear", int32(historyeducation.EndYear)),
				bson.EC.ObjectID("personId", result.InsertedID.(objectid.ObjectID)),
			))
	}
	for _, historyWork := range person.HistoryWorks {
		collHistoryWork.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.DateTime("dateStart", historyWork.StartDate.UnixNano()/1e6),
				bson.EC.DateTime("dateEnd", historyWork.EndDate.UnixNano()/1e6),
				bson.EC.String("position", historyWork.Position),
				bson.EC.String("place", historyWork.Workplace),
				bson.EC.Boolean("dateless", historyWork.DateLess),
				bson.EC.ObjectID("personId", result.InsertedID.(objectid.ObjectID)),
			))
	}
	return nil
}
