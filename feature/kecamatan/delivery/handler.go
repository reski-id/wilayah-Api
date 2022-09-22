package delivery

import (
	"log"
	"net/http"
	"strconv"
	"wilayah/domain"

	"github.com/labstack/echo/v4"
)

type kecHandler struct {
	kecUseCase domain.KecamatanUsecase
}

func New(ku domain.KecamatanUsecase) domain.KecamatanHandler {
	return &kecHandler{
		kecUseCase: ku,
	}
}

func (kh *kecHandler) AddKecamatan() echo.HandlerFunc {
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
		_, err := kh.kecUseCase.AddKecamatan(tmp.ToModel())
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

func (kh *kecHandler) GetKecamatan() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := kh.kecUseCase.GetAllKecamatan()
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

func (kh *kecHandler) UpdateDataKecamatan() echo.HandlerFunc {
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

		data, err := kh.kecUseCase.UpdateKecamatan(id, updatedData.ToModel())
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

func (kh *kecHandler) DeleteDataKecamatan() echo.HandlerFunc {
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

		_, errDel := kh.kecUseCase.DeleteKecamatan(id)
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
