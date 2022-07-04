package bussiness

import (
	"errors"
	"project3/eventapp/features/products"
)

type productUseCase struct {
	productData products.Data
}

func NewProductBusiness(usrData products.Data) products.Business {
	return &productUseCase{
		productData: usrData,
	}
}

func (uc *productUseCase) GetAllProduct() (response []products.Core, err error) {
	resp, errData := uc.productData.SelectData()
	return resp, errData
}

func (uc *productUseCase) GetProductByID(id int) (response products.Core, err error) {
	response, err = uc.productData.SelectDataByID(id)
	return response, err
}

func (uc *productUseCase) InsertProduct(productRequest products.Core) error {
	if productRequest.Name == "" || productRequest.PhotoUrl == "" || productRequest.ProductDetail == "" || productRequest.Stock == 0 || productRequest.Price == 0 {
		return errors.New("all data must be filled")
	}

	err := uc.productData.InsertData(productRequest)
	return err
}

func (uc *productUseCase) DeleteProductByID(id int, userId int) (err error) {
	err = uc.productData.DeleteDataByID(id, userId)
	return err
}

func (uc *productUseCase) UpdateProductByID(productReq products.Core, id int, userId int) (err error) {
	updateMap := make(map[string]interface{})
	if productReq.Name != "" {
		updateMap["name"] = &productReq.Name
	}
	if productReq.ProductDetail != "" {
		updateMap["detail"] = &productReq.ProductDetail
	}
	if productReq.Stock != 0 {
		updateMap["stock"] = &productReq.Stock
	}

	err = uc.productData.UpdateDataByID(updateMap, id, userId)
	return err
}

func (uc *productUseCase) GetProductByUserID(id_user int) (response []products.Core, err error) {
	resp, errData := uc.productData.SelectDataByUserID(id_user)
	return resp, errData
}
