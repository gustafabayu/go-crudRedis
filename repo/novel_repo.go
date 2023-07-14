package repo

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gustafabayu/go-crudRedis/domain"
	"github.com/gustafabayu/go-crudRedis/model"
	"gorm.io/gorm"
)

type novelRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

// GetNovelById implements domain.NovelRepo.
func (n *novelRepo) GetNovelById(id int) (model.Novel, error) {
	var novels model.Novel
	var ctx = context.Background()
	result, err := n.rdb.Get(ctx, "novel"+strconv.Itoa(id)).Result()
	if err != nil && err != redis.Nil {
		return novels, err
	}

	if len(result) > 0 {
		err := json.Unmarshal([]byte(result), &novels)
		return novels, err
	}

	err = n.db.Model(model.Novel{}).Select("id", "name", "description", "author").Where("id=?", id).Find(&novels).Error
	if err != nil {
		return novels, err
	}

	jsonBytes, err := json.Marshal(novels)
	if err != nil {
		return novels, err
	}
	jsonString := string(jsonBytes)

	err = n.rdb.Set(ctx, "novel"+strconv.Itoa(id), jsonString, 24*time.Hour).Err()
	if err != nil {
		return novels, err
	}

	return novels, nil
}

// CreateNovel implements domain.NovelRepo.
func (n *novelRepo) CreateNovel(createNovel model.Novel) error {
	if err := n.db.Create(&createNovel).Error; err != nil {
		return errors.New("internal server error: cannot create novel")
	}
	return nil
}

func NewNovelRepo(db *gorm.DB, rdb *redis.Client) domain.NovelRepo {
	return &novelRepo{
		db:  db,
		rdb: rdb,
	}
}
