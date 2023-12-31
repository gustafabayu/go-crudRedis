package usecase

import (
	"errors"

	"github.com/gustafabayu/go-crudRedis/domain"
	"github.com/gustafabayu/go-crudRedis/model"
)

type novelUseCase struct {
	novelRepo domain.NovelRepo
}

// GetNovelById implements domain.NovelUseCase.
func (n *novelUseCase) GetNovelById(id int) (model.Novel, error) {
	res, err := n.novelRepo.GetNovelById(id)
	if err != nil {
		return model.Novel{}, errors.New("internal server error: " + err.Error())
	}
	return res, nil
}

// CreateNovel implements domain.NovelUseCase.
func (n *novelUseCase) CreateNovel(createNovel model.Novel) error {
	err := n.novelRepo.CreateNovel(createNovel)
	return errors.New("internal server error: " + err.Error())
}

func NewNovelUseCase(novelRepo domain.NovelRepo) domain.NovelUseCase {
	return &novelUseCase{
		novelRepo: novelRepo,
	}
}
