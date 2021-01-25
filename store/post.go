package store

import (
	"github.com/jinzhu/gorm"
	"github.com/sOM2H/golang_trainee_backend/model"
)

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{
		db: db,
	}
}

func (ps *PostStore) ShowPost(id uint) (*model.Post, error) {
	var p model.Post
	if err := ps.db.First(&p, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (as *PostStore) CreatePost(a *model.Post) error {
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (as *PostStore) UpdatePost(a *model.Post) error {
	tx := as.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (as *PostStore) DeletePost(a *model.Post) error {
	return as.db.Delete(a).Error
}

func (as *PostStore) ListPost(offset, limit int) ([]model.Post, int, error) {
	var (
		posts []model.Post
		count int
	)

	as.db.Order("created_at desc").Find(&posts)
	as.db.Model(&posts).Count(&count)

	return posts, count, nil
}

func (us *PostStore) GetPostByID(id uint) (*model.Post, error) {
	var p model.Post
	if err := us.db.First(&p, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}
