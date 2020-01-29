package feedback

import "github.com/Bezyee/Online-ticket-store/entity"

// CommentService specifies customer feedback related service
type CommentService interface {
	Comments() ([]entity.feedback, []error)
	feedback(id uint) (*entity.feedback, []error)
	UpdateComment(feedback *entity.feedback) (*entity.feedback, []error)
	DeleteComment(id uint) (*entity.feedback, []error)
	StoreComment(feedback *entity.feedback) (*entity.feedback, []error)
}
