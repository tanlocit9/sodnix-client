package database

import "github.com/google/uuid"

// QueryOneById queries by primary key (id) into a model.
func QueryOneById[T any](id any, model *T) error {
	parsedID, err := uuid.Parse(id.(string))
	if err != nil {
		return err
	}
	result := DB.First(model, parsedID)
	return result.Error
}

// QueryMulti queries all records into the slice pointer (e.g. *[]User).
func QueryMulti[T any](model *T) error {
	result := DB.Find(model)
	return result.Error
}

// CreateOne inserts a new record into the database.
func CreateOne[T any](model *T) error {
	result := DB.Create(model)
	return result.Error
}

// UpdateOne updates the given model (based on primary key).
func UpdateOne[T any](model *T) error {
	result := DB.Save(model)
	return result.Error
}

// DeleteOne deletes the given model.
func DeleteOne[T any](model *T) error {
	result := DB.Delete(model)
	return result.Error
}
