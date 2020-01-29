package repository

import (
	"github.com/Bezyee/Online-ticket-store/feedback"
	"github.com/Bezyee/Online-ticket-store/entity"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type CommentGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo returns new object of CommentGormRepo
func NewCommentGormRepo(db *gorm.DB) feedback.CommentRepository {
	return &CommentGormRepo{conn: db}
}

// Comments returns all customer comments stored in the database
func (cmntRepo *CommentGormRepo) Comments() ([]entity.feedback, []error) {
	cmnts := []entity.feedback{}
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// feedback retrieves a customer feedback from the database by its id
func (cmntRepo *CommentGormRepo) feedback(id uint) (*entity.feedback, []error) {
	cmnt := entity.feedback{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}

// UpdateComment updates a given customer feedback in the database
func (cmntRepo *CommentGormRepo) UpdateComment(feedback *entity.feedback) (*entity.feedback, []error) {
	cmnt := feedback
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteComment deletes a given customer feedback from the database
func (cmntRepo *CommentGormRepo) DeleteComment(id uint) (*entity.feedback, []error) {
	cmnt, errs := cmntRepo.feedback(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreComment stores a given customer feedback in the database
func (cmntRepo *CommentGormRepo) StoreComment(feedback *entity.feedback) (*entity.feedback, []error) {
	cmnt := feedback
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
