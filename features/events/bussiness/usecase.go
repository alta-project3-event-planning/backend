package bussiness

import (
	"errors"
	"project3/eventapp/features/events"
)

type eventUseCase struct {
	eventData events.Data
}

func NewEventBusiness(usrData events.Data) events.Business {
	return &eventUseCase{
		eventData: usrData,
	}
}

func (uc *eventUseCase) GetAllEvent() (response []events.Core, err error) {
	resp, errData := uc.eventData.SelectData()
	return resp, errData
}

func (uc *eventUseCase) GetEventByID(id int) (response events.Core, err error) {
	response, err = uc.eventData.SelectDataByID(id)
	return response, err
}

func (uc *eventUseCase) InsertEvent(eventRequest events.Core) error {
	if eventRequest.Name == "" || eventRequest.Url == "" || eventRequest.EventDetail == "" || eventRequest.City == "" || eventRequest.Location == "" || eventRequest.Performers == "" || eventRequest.HostedBy == "" {
		return errors.New("all data must be filled")
	}

	err := uc.eventData.InsertData(eventRequest)
	return err
}

func (uc *eventUseCase) DeleteEventByID(id int, userId int) (err error) {
	err = uc.eventData.DeleteDataByID(id, userId)
	return err
}

func (uc *eventUseCase) UpdateEventByID(eventReq events.Core, id int, userId int) (err error) {
	updateMap := make(map[string]interface{})
	if eventReq.Name != "" {
		updateMap["name"] = &eventReq.Name
	}
	if eventReq.EventDetail != "" {
		updateMap["detail"] = &eventReq.EventDetail
	}
	if eventReq.City != "" {
		updateMap["city"] = &eventReq.City
	}
	if eventReq.Location != "" {
		updateMap["location"] = &eventReq.Location
	}
	if eventReq.Performers != "" {
		updateMap["performers"] = &eventReq.Performers
	}
	if eventReq.HostedBy != "" {
		updateMap["hostedby"] = &eventReq.HostedBy
	}

	err = uc.eventData.UpdateDataByID(updateMap, id, userId)
	return err
}

func (uc *eventUseCase) GetEventByUserID(id_user int) (response []events.Core, err error) {
	resp, errData := uc.eventData.SelectDataByUserID(id_user)
	return resp, errData
}