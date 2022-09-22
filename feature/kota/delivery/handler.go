package delivery

import (
	"log"
	"net/http"
	"strconv"
	"wilayah/domain"

	"github.com/labstack/echo/v4"
)

type kotaHandler struct {
	kotaUseCase domain.KotaUsecase
}

func New(ku domain.KotaUsecase) domain.KotaHandler {
	return &kotaHandler{
		kotaUseCase: ku,
	}
}

func (kh *kotaHandler) AddKota() echo.HandlerFunc {
	return func(c echo.Context) error {

		var tmp InserFormat
		errBind := c.Bind(&tmp)
		if errBind != nil {
			log.Println("cannot parse data", errBind)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "internal server error",
			})
		}
		_, err := kh.kotaUseCase.AddKota(tmp.ToModel())
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

func (kh *kotaHandler) GetKota() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := kh.kotaUseCase.GetAllKota()
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

func (kh *kotaHandler) UpdateDataKota() echo.HandlerFunc {
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

		data, err := kh.kotaUseCase.UpdateKota(id, updatedData.ToModel())
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

func (kh *kotaHandler) DeleteDataKota() echo.HandlerFunc {
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

		_, errDel := kh.kotaUseCase.DeleteKota(id)
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
