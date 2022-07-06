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

func (uc *commentUseCase) GetCommentByIdEvent(limit, offset, eventId int) (response []comments.Core, total int64, err error) {
	page := limit * (offset - 1)
	response, total, err = uc.commentData.GetComment(limit, page, eventId)
	total = total/int64(limit) + 1
	return response, total, err
}
