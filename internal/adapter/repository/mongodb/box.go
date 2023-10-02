package repository

import (
	"bnt/bnt-box-service/internal/core/domain/entity"
	errorhelper "bnt/bnt-box-service/internal/core/helper/error-helper"
	logger "bnt/bnt-box-service/internal/core/helper/log-helper"
	"fmt"

	//ports "bnt/bnt-box-service/internal/port"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BoxInfra struct {
	BoxCollection *mongo.Collection
	IndentCollection *mongo.Collection
}

func NewBox(BoxCollection *mongo.Collection, IndentCollection *mongo.Collection) *BoxInfra {
	return &BoxInfra{BoxCollection, IndentCollection}
}


func (r *BoxInfra) Generate(boxes []entity.Box) (interface{}, error) {
	logger.LogEvent("INFO", "Persisting box configurations with reference")

	var documents []mongo.WriteModel
    for _, box := range boxes {	
		document := mongo.NewInsertOneModel()
		document.SetDocument(box)
		documents = append(documents, document)
	}

	bulkOption :=  options.BulkWriteOptions{}
	bulkOption.SetOrdered(true)


	_, err := r.BoxCollection.BulkWrite(context.TODO(), documents, &bulkOption)
	
	
	if err != nil {
		return nil, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error())
	}

	logger.LogEvent("INFO", "Persisting box configurations completed successfully...")
	return fmt.Sprint(len(documents)) + " documents successfully persisted", nil
}

func (r *BoxInfra) GetIndentByRef(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving indent configurations with reference: "+reference)
	indent := entity.Indent{}
	filter := bson.M{"reference": reference}
	err := r.IndentCollection.FindOne(context.TODO(), filter).Decode(&indent)
	if err != nil || indent == (entity.Indent{}) {
		return nil, err
	}
	logger.LogEvent("INFO", "Retrieving indent configurations with reference: "+reference+" completed successfully. ")
	return indent, nil
}
