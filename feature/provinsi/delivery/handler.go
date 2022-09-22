package delivery

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"wilayah/domain"

	"github.com/labstack/echo/v4"
)

type provinsiHandler struct {
	provinsiUseCase domain.ProvinsiUsecase
}

func New(pu domain.ProvinsiUsecase) domain.ProvinsiHandler {
	return &provinsiHandler{
		provinsiUseCase: pu,
	}
}

func (ph *provinsiHandler) AddProvinsi() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("ok")
		// token := common.ExtractData(c)

		var tmp InserFormat
		errBind := c.Bind(&tmp)
		if errBind != nil {
			log.Println("cannot parse data", errBind)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}
		_, err := ph.provinsiUseCase.AddProvinsi(tmp.ToModel())
		if err != nil {
			log.Println("cannot proces data", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "succes",
		})
	}
}

func (ph *provinsiHandler) GetProvinsi() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ph.provinsiUseCase.GetAllProvinsi()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    FromModelToList(data),
		})
	}
}

func (ph *provinsiHandler) UpdateDataProvinsi() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Please enter data correctly",
			})
		}

		var updatedData InserFormat
		err = c.Bind(&updatedData)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		data, err := ph.provinsiUseCase.UpdateProvinsi(id, updatedData.ToModel())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success update data " + param,
			"data":    data,
		})
	}
}

func (ph *provinsiHandler) DeleteDataProvinsi() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Please enter data correctly",
			})
		}

		_, errDel := ph.provinsiUseCase.DeleteProvinsi(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success",
		})
	}
}
