package router

import (
	"github.com/haitwang-cloud/golang-echo-sample/controller"
	"github.com/haitwang-cloud/golang-echo-sample/utils/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
)

type bibleController struct {
	service *controller.BibleClient
	wrapper middlewares.Wrapper
}

func NewBibleController(wrapper middlewares.Wrapper) *bibleController {
	return &bibleController{
		service: controller.NewBibleClient(wrapper),
		wrapper: wrapper,
	}
}

// GetResult ShowBible godoc
// @Summary      Show an account
// @Description  Provides grabbing bible verses and passages
// @Tags         bible
// @Success      200  {object}  controller.BiBleResult
// @Failure      400  {object}  controller.HttpError
// @Failure      404  {object}  controller.HttpError
// @Failure      500  {object}  controller.HttpError
// @Router       /api/bible/result [get]
func (controller *bibleController) GetResult(c echo.Context) error {
	book, chapter, verse := "John", "3", "17"
	result, err := controller.service.GetResult(book, chapter, verse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
