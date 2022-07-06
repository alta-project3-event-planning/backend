package bussiness

import (
	"fmt"
	"project3/eventapp/features/events"
	"testing"

	"github.com/stretchr/testify/assert"
)

//mock data success case
type mockEventDataSucces struct{}

func (mock mockEventDataSucces) SelectData() (data []events.Core, err error) {
	return []events.Core{
		{ID: 1, Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
		{ID: 1, Name: "festival 2", EventDetail: "detail festival 2", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
	}, nil
}

func (mock mockEventDataSucces) SelectDataByID(id int) (data events.Core, err error) {
	return events.Core{ID: 1, Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1}, nil
}

func (mock mockEventDataSucces) InsertData(data events.Core) (err error) {
	return nil
}

func (mock mockEventDataSucces) DeleteDataByID(id int, userid int) (err error) {
	return nil
}

func (mock mockEventDataSucces) UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error) {
	return nil
}

func (mock mockEventDataSucces) SelectDataByUserID(userId int) (data []events.Core, err error) {
	return []events.Core{
		{ID: 1, Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
		{ID: 1, Name: "festival 2", EventDetail: "detail festival 2", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
	}, nil
}

//mock data failed case
type mockEventDataFailed struct{}

func (mock mockEventDataFailed) SelectData() (data []events.Core, err error) {
	return nil, fmt.Errorf("Failed to select data")
}

func (mock mockEventDataFailed) SelectDataByID(id int) (data events.Core, err error) {
	return data, fmt.Errorf("Failed to select data")
}

func (mock mockEventDataFailed) InsertData(data events.Core) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockEventDataFailed) DeleteDataByID(id int, userid int) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockEventDataFailed) UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockEventDataFailed) SelectDataByUserID(userId int) (data []events.Core, err error) {
	return nil, fmt.Errorf("failed to select data ")
}

func TestGetAllEvent(t *testing.T) {
	t.Run("Test Get All Data Success", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		result, err := eventBusiness.GetAllEvent()
		assert.Nil(t, err)
		assert.Equal(t, "sepatu baru", result[0].Name)
	})

	t.Run("Test Get All Data Failed", func(t *testing.T) {

		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, err := eventBusiness.GetAllEvent()
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetEventByID(t *testing.T) {
	t.Run("Test Get event Data By ID Success", func(t *testing.T) {
		id := 1
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		result, err := eventBusiness.GetEventByID(id)
		assert.Nil(t, err)
		assert.Equal(t, "sepatu baru", result.Name)
	})

	t.Run("Test Get event Data By ID Failed", func(t *testing.T) {
		id := 3
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, err := eventBusiness.GetEventByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Name)
	})
}

func TestInsertEvent(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		newEvent := events.Core{
			Name: "sepatu baru", EventDetail: "ini sepatu baru", Price: 10000, Stock: 10, PhotoUrl: "example.com", Photo: "ini foto", UserID: 1,
		}
		err := eventBusiness.InsertEvent(newEvent)
		assert.Nil(t, err)
	})

	t.Run("Test Insert Data Failed", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		newEvent := events.Core{
			Name: "alta",
		}
		err := EventBusiness.InsertEvent(newEvent)
		assert.NotNil(t, err)
	})
}

func TestGetEventByUserID(t *testing.T) {
	t.Run("Test Get Event Data By ID User Success", func(t *testing.T) {
		id := 1
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		result, err := eventBusiness.GetEventByUserID(id)
		assert.Nil(t, err)
		assert.Equal(t, "sepatu baru", result[0].Name)
	})

	t.Run("Test Get Data By ID User Failed", func(t *testing.T) {
		id := 3
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, err := eventBusiness.GetEventByUserID(id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete Data", func(t *testing.T) {
		id := 1
		userid := 1
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		err := eventBusiness.DeleteEventByID(id, userid)
		assert.Nil(t, err)

	})
	t.Run("Test Delete Data Failed", func(t *testing.T) {
		id := 3
		userid := 1
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		err := eventBusiness.DeleteEventByID(id, userid)
		assert.NotNil(t, err)

	})
}

func TestUpdateEvent(t *testing.T) {
	t.Run("Test Update Data Success", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		id := 1
		userid := 1
		newEvent := events.Core{
			Name: "sepatu baru", EventDetail: "ini sepatu baru", Price: 10000, Stock: 10,
		}
		err := eventBusiness.UpdateEventByID(newEvent, id, userid)
		assert.Nil(t, err)
	})

	t.Run("Test Update Data Failed", func(t *testing.T) {
		id := 1
		userid := 0
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		newEvent := events.Core{
			Name: "septau",
		}
		err := eventBusiness.UpdateEventByID(newEvent, id, userid)
		assert.NotNil(t, err)
	})
}
