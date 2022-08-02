package delivery

import (
	"commerce-app/domain"
	"commerce-app/feature/common"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUsecase domain.ProductUseCase
}

func New(pu domain.ProductUseCase) domain.ProductHandler {
	return &productHandler{
		productUsecase: pu,
	}
}

func (ph *productHandler) InsertProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp ProductInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		fmt.Println(tmp)

		var userid = common.ExtractData(c)
		data, err := ph.productUsecase.AddProduct(common.ExtractData(c), tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Println(userid)

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}

func (ph *productHandler) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		qry := map[string]interface{}{}
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp ProductInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		if tmp.Name != "" {
			qry["name"] = tmp.Name
		}
		if tmp.Images != "" {
			qry["images"] = tmp.Images
		}
		if tmp.Description != "" {
			qry["description"] = tmp.Description
		}
		if tmp.Price != 0 {
			qry["price"] = tmp.Price
		}

		if tmp.Stock != 0 {
			qry["stock"] = tmp.Stock
		}
		data, err := ph.productUsecase.UpdateProduct(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "success update data",
			"id":          data.ID,
			"name":        data.Name,
			"price":       data.Price,
			"stock":       data.Stock,
			"description": data.Description,
			"images":      data.Images,
		})
	}
}

func (ph *productHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := ph.productUsecase.DeleteProduct(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}

		if data == 0 {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete news",
		})
	}
}

func (ph *productHandler) GetAllProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ph.productUsecase.GetAll()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all news",
			"users":   data,
		})
	}
}

func (ph *productHandler) GetProductID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idNews := c.Param("id")
		id, _ := strconv.Atoi(idNews)
		data, err := ph.productUsecase.GetProduct(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get product",
			"users":   data,
		})
	}
}
