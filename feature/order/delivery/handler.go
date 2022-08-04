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

type orderHandler struct {
	orderUsecase domain.OrderUseCase
}

func New(pu domain.OrderUseCase) domain.OrderHandler {
	return &orderHandler{
		orderUsecase: pu,
	}
}
func (oh *orderHandler) UpdateOrder() echo.HandlerFunc {
	return func(c echo.Context) error {

		qry := map[string]interface{}{}
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp OrderInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		if tmp.CartID != 0 {
			qry["cart"] = tmp.CartID
		}
		if tmp.UserID != 0 {
			qry["user_id"] = tmp.UserID
		}
		if tmp.ProductID != 0 {
			qry["product_id"] = tmp.ProductID
		}
		if tmp.Payment != "" {
			qry["payment"] = tmp.Payment
		}
		if tmp.Status != 0 {
			qry["status"] = tmp.Status
		}

		data, err := oh.orderUsecase.UpOrder(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success Update",
			"data":    FromDomain(data),
		})
	}
}

func (oh *orderHandler) InsertOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp OrderInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		fmt.Println(tmp)

		// var userid = common.ExtractData(c)
		userid, _ := common.ExtractData2(c)

		data, err := oh.orderUsecase.AddOrder(common.ExtractData(c), tmp.ToDomain())

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

func (oh *orderHandler) DeleteOrder() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := oh.orderUsecase.DelOrder(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete order",
		})
	}
}

func (oh *orderHandler) GetAllOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := oh.orderUsecase.GetAllO()
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all Order",
			"data":    data,
		})
	}
}

func (oh *orderHandler) GetOrderID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idOrder := c.Param("id")
		id, _ := strconv.Atoi(idOrder)
		data, err := oh.orderUsecase.GetSpecificOrder(id)
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data",
			"data":    data,
		})
	}
}
