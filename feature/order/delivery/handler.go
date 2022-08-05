package delivery

import (
	"commerce-app/domain"
	"commerce-app/feature/common"
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

		// var userid = common.ExtractData(c)
		userid, _ := common.ExtractData2(c)

		data, err := oh.orderUsecase.AddOrder(userid, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

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
		_, role := common.ExtractData2(c)

		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can view all order history",
			})
		}

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
		_, role := common.ExtractData2(c)

		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can view order by id",
			})
		}
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
func (oh *orderHandler) GetMYOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		// idNews := c.Param("id")
		userid, _ := common.ExtractData2(c)
		data, err := oh.orderUsecase.GetmyOrder(userid)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my order",
			"data":    data,
		})
	}
}
