package delivery

import (
	"commerce-app/domain"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartUsecase domain.CartUseCase
}

func New(pu domain.CartUseCase) domain.CartHandler {
	return &cartHandler{
		cartUsecase: pu,
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
			"message": "success get all Product",
			"data":    data,
		})
	}
}

// func (ph *productHandler) InsertProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var tmp ProductInsertRequest
// 		err := c.Bind(&tmp)

// 		if err != nil {
// 			log.Println("Cannot parse data", err)
// 			c.JSON(http.StatusBadRequest, "error read input")
// 		}

// 		data, err := ph.productUsecase.AddProduct(tmp.ToDomain())
// 		if err != nil {
// 			log.Println("Cannot proces data", err)
// 			c.JSON(http.StatusInternalServerError, err)
// 		}

// 		return c.JSON(http.StatusCreated, map[string]interface{}{
// 			"message": "success create data",
// 			"data":    data,
// 		})
// 	}
// }

// func (ph *productHandler) UpdateProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		qry := map[string]interface{}{}
// 		cnv, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			log.Println("Cannot convert to int", err.Error())
// 			return c.JSON(http.StatusInternalServerError, "cannot convert id")
// 		}

// 		var tmp ProductInsertRequest
// 		res := c.Bind(&tmp)

// 		if res != nil {
// 			log.Println(res, "Cannot parse data")
// 			return c.JSON(http.StatusInternalServerError, "error read update")
// 		}

// 		if tmp.Name != "" {
// 			qry["name"] = tmp.Name
// 		}
// 		if tmp.Images != "" {
// 			qry["images"] = tmp.Images
// 		}
// 		if tmp.Description != "" {
// 			qry["description"] = tmp.Description
// 		}
// 		if tmp.Price != 0 {
// 			qry["price"] = tmp.Price
// 		}

// 		if tmp.Stock != 0 {
// 			qry["stock"] = tmp.Stock
// 		}
// 		data, err := ph.productUsecase.UpProduct(cnv, tmp.ToDomain())

// 		if err != nil {
// 			log.Println("Cannot update data", err)
// 			c.JSON(http.StatusInternalServerError, "cannot update")
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"ID":          data.ID,
// 			"Name":        data.Name,
// 			"Price":       data.Price,
// 			"Stock":       data.Stock,
// 			"Description": data.Description,
// 			"Images":      data.Images,
// 			"message ":    "Update Data Sukses",
// 		})
// 	}
// }

// func (ph *productHandler) DeleteProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		cnv, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			log.Println("Cannot convert to int", err.Error())
// 			return c.JSON(http.StatusInternalServerError, "cannot convert id")
// 		}

// 		data, err := ph.productUsecase.DelProduct(cnv)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "cannot delete user")
// 		}

// 		if !data {
// 			return c.JSON(http.StatusInternalServerError, "cannot delete")
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "success delete product",
// 		})
// 	}
// }

// func (ph *productHandler) GetProductID() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		idProduct := c.Param("id")
// 		id, _ := strconv.Atoi(idProduct)
// 		data, err := ph.productUsecase.GetSpecificProduct(id)

// 		if err != nil {
// 			log.Println("Cannot get data", err)
// 			return c.JSON(http.StatusBadRequest, "cannot read input")
// 		}
// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "success get data",
// 			"data":    data,
// 		})
// 	}
// }
