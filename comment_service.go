package service

import (
	"github.com/Bezyee/Online-ticket-store/feedback"
	"github.com/Bezyee/Online-ticket-store/entity"
)

// CommentService implements menu.CommentService interface
type CommentService struct {
	commentRepo feedback.CommentRepository
}

// NewCommentService returns a new CommentService object
func NewCommentService(commRepo feedback.CommentRepository) feedback.CommentService {
	return &CommentService{commentRepo: commRepo}
}

// Comments returns all stored comments
func (cs *CommentService) Comments() ([]entity.feedback, []error) {
	cmnts, errs := cs.commentRepo.Comments()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// feedback retrieves stored feedback by its id
func (cs *CommentService) feedback(id uint) (*entity.feedback, []error) {
	cmnt, errs := cs.commentRepo.feedback(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// UpdateComment updates a given feedback
func (cs *CommentService) UpdateComment(feedback *entity.feedback) (*entity.feedback, []error) {
	cmnt, errs := cs.commentRepo.UpdateComment(feedback)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteComment deletes a given feedback
func (cs *CommentService) DeleteComment(id uint) (*entity.feedback, []error) {
	cmnt, errs := cs.commentRepo.DeleteComment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreComment stores a given feedback
func (cs *CommentService) StoreComment(feedback *entity.feedback) (*entity.feedback, []error) {
	cmnt, errs := cs.commentRepo.StoreComment(feedback)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
