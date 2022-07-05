package presentation

import (
	"fmt"
	"net/http"
	"project3/eventapp/features/events"
	"time"

	_request_event "project3/eventapp/features/events/presentation/request"
	_response_event "project3/eventapp/features/events/presentation/response"
	"project3/eventapp/helper"
	"project3/eventapp/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventBusiness events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		eventBusiness: business,
	}
}

func (h *EventHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	name := c.QueryParam("name")
	city := c.QueryParam("city")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	result, err := h.eventBusiness.GetAllEvent(limitint, offsetint, name, city)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	respons := _response_event.FromCoreList(result)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get all events", respons))
}

func (h *EventHandler) GetDataById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.eventBusiness.GetEventByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get event"))
	}

	response := _response_event.FromCoreByID(result)

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success get event", response))
}

func (h *EventHandler) InsertData(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed insert data"))
	}

	event := _request_event.Event{}
	err_bind := c.Bind(&event)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseSuccessNoData("success insert event"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get file"))
	}

	extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
	}

	// check file size
	err_check_size := helper.CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
	}

	// memberikan nama file
	fileName := strconv.Itoa(userID_token) + "_" + event.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := helper.UploadImageToS3(fileName, fileData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to upload file"))
	}

	eventCore := _request_event.ToCore(event)
	eventCore.IDUser = userID_token
	eventCore.Url = url

	err := h.eventBusiness.InsertEvent(eventCore)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert data",
	})

}

func (h *EventHandler) DeleteData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}

	err := h.eventBusiness.DeleteEventByID(id, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete data" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data",
	})
}

func (h *EventHandler) UpdateData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	eventReq := _request_event.Event{}
	err_bind := c.Bind(&eventReq)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind data"))
	}

	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}

	eventCore := _request_event.ToCore(eventReq)

	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get file"))
		}

		extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
		}

		// check file size
		err_check_size := helper.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
		}

		// memberikan nama file
		fileName := strconv.Itoa(userID_token) + "_" + eventReq.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

		url, errUploadImg := helper.UploadImageToS3(fileName, fileData)

		if errUploadImg != nil {
			fmt.Println(errUploadImg)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to upload file"))
		}

		eventCore.Url = url
	}

	err := h.eventBusiness.UpdateEventByID(eventCore, id, userID_token)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed update data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("sucsess update event"))
}

func (h *EventHandler) GetEventByUser(c echo.Context) error {

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	id_user, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get user id"))
	}

	result, err := h.eventBusiness.GetEventByUserID(id_user, limitint, offsetint)

	respons := _response_event.FromCoreList(result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all my events"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success get all my events", respons))
}
