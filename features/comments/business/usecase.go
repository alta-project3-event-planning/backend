package business

import (
	"project3/eventapp/features/comments"
)

type commentUseCase struct {
	commentData comments.Data
}

func NewCommentBusiness(usrData comments.Data) comments.Business {
	return &commentUseCase{
		commentData: usrData,
	}
}

func (uc *commentUseCase) AddComment(data comments.Core) (row int, err error) {
	row, err = uc.commentData.Add(data)
	return row, err
}

func (uc *commentUseCase) GetCommentByIdEvent(limit, offset, eventId int) (response []comments.Core, err error) {
	response, err = uc.commentData.GetComment(limit, offset, eventId)
	return response, err
}
