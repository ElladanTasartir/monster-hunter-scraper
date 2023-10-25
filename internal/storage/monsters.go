package storage

import (
	"context"
	"github.com/ElladanTasartir/monster-hunter-scraper/internal/scraper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monster struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Image string             `bson:"image_url" json:"image_url"`
}

const (
	monstersCollection = "monsters"
)

func (s *Storage) CreateMonsters(ctx context.Context, monsters []scraper.Monster) ([]Monster, error) {
	var storageMonsters []Monster
	for _, monster := range monsters {
		storageMonsters = append(storageMonsters, Monster{
			ID:    primitive.NewObjectID(),
			Name:  monster.Name,
			Image: monster.Image,
		})
	}

	dataToInsert := make([]interface{}, 0)
	for _, monster := range storageMonsters {
		dataToInsert = append(dataToInsert, bson.D{
			{"_id", monster.ID},
			{"name", monster.Name},
			{"image_url", monster.Image},
		})
	}

	_, err := s.database.Collection(monstersCollection).InsertMany(ctx, dataToInsert)
	if err != nil {
		return nil, err
	}

	return storageMonsters, nil
}

func (s *Storage) FindMonsters(ctx context.Context) ([]Monster, error) {
	monsters := make([]Monster, 0)
	cursor, err := s.database.Collection(monstersCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &monsters); err != nil {
		return nil, err
	}

	return monsters, nil
}
