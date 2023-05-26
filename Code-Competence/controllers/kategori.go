package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"Code-Competence/models"
	
)

type KategoriController interface {
	GetKategorisController(c echo.Context) error
	GetKategoriController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type kategoriController struct {
	KategoriS services.KategoriService
}

func NewKategoriController(KategoriS services.KategoriService) KategoriController {
	return &kategoriController{
		KategoriS: KategoriS,
	}
}

func (k *kategoriController) GetKategorisController(c echo.Context) error {
	Kategoris, err := k.KategoriS.GetKategorisService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Kategoris,
		Message: "Get all Kategoris success",
		Status:  true,
	})
}

func (k *kategoriController) GetKategoriController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Kategori *models.Kategori

	Kategori, err = k.KategoriS.GetKategoriService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Kategori,
		Message: "Get Kategori success",
		Status:  true,
	})
}

func (k *kategoriController) CreateController(c echo.Context) error {
	var Kategori *models.Kategori

	err := c.Bind(&Kategori)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Kategori, err = k.KategoriS.CreateService(*Kategori)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Kategori,
		Message: "Create Kategori success",
		Status:  true,
	})
}

func (k *kategoriController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Kategori *models.Kategori

	err = c.Bind(&Kategori)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Kategori, err = k.KategoriS.UpdateService(id, *Kategori)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Kategori,
		Message: "Update Kategori success",
		Status:  true,
	})
}

func (k *kategoriController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = k.KategoriS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Kategori success",
		Status:  true,
	})
}