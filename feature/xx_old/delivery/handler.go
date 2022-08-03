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

type cartHandler struct {
	cartUsecase domain.CartUseCase
}

func New(cu domain.CartUseCase) domain.CartHandler {
	return &cartHandler{
		cartUsecase: cu,
	}
}

func (ch *cartHandler) InsertCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp CartInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		fmt.Println(tmp)

		var userid = common.ExtractData(c)
		data, err := ch.cartUsecase.AddCart(common.ExtractData(c), tmp.ToDomain())

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

func (ch *cartHandler) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {

		qry := map[string]interface{}{}
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp CartInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		if tmp.Quantity != 0 {
			qry["quantity"] = tmp.Quantity
		}

		data, err := ch.cartUsecase.UpCart(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success update data",
			"ID":       data.ID,
			"Quantity": data.Quantity,
		})
	}
}

func (ch *cartHandler) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := ch.cartUsecase.DelCart(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete Cart")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete Cart",
		})
	}
}

func (ch *cartHandler) GetAllCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ch.cartUsecase.GetAllC()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all Cart",
			"users":   data,
		})
	}
}

func (ch *cartHandler) GetCartID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idNews := c.Param("id")
		id, _ := strconv.Atoi(idNews)
		data, err := ch.cartUsecase.GetSpecificCart(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my cart",
			"data":    data,
		})
	}
}