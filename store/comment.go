package store

import (
	"github.com/jinzhu/gorm"
	"github.com/sOM2H/golang_trainee_backend/model"
)

type CommentStore struct {
	db *gorm.DB
}

func NewCommentStore(db *gorm.DB) *CommentStore {
	return &CommentStore{
		db: db,
	}
}

func (us *CommentStore) Show(id uint) (*model.Comment, error) {
	var c model.Comment
	if err := us.db.First(&c, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (as *CommentStore) CreateComment(a *model.Comment) error {
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (as *CommentStore) UpdateComment(a *model.Comment) error {
	tx := as.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (cs *CommentStore) DeleteComment(c *model.Comment) error {
	return cs.db.Delete(c).Error
}

func (as *CommentStore) ListComment(offset, limit int) ([]model.Comment, int, error) {
	var (
		comments []model.Comment
		count    int
	)

	as.db.Order("created_at desc").Find(&comments)
	as.db.Model(&comments).Count(&count)

	return comments, count, nil
}

func (us *CommentStore) GetCommentByID(id uint) (*model.Comment, error) {
	var c model.Comment
	if err := us.db.First(&c, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}
