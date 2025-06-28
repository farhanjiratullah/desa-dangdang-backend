package handler

import (
	"desadangdang/config"
	"desadangdang/internal/adapater/handler/request"
	"desadangdang/internal/adapater/handler/response"
	"desadangdang/internal/core/domain/entity"
	"desadangdang/internal/core/service"
	"desadangdang/utils/conv"
	"desadangdang/utils/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type StatisticHandlerInterface interface {
	CreateStatistic(c echo.Context) error
	FetchAllStatistic(c echo.Context) error
	FetchByIDStatistic(c echo.Context) error
	EditByIDStatistic(c echo.Context) error
	DeleteByIDStatistic(c echo.Context) error
}

type statisticHandler struct {
	statisticService service.StatisticServiceInterface
}

// CreateStatistic implements StatisticHandlerInterface.
func (s *statisticHandler) CreateStatistic(c echo.Context) error {
	var (
		req       = request.StatisticRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateStatistic - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateStatistic - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err := c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateStatistic - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.StatisticEntity{
		Name:  req.Name,
		Total: req.Total,
		Icon:  req.Icon,
	}

	err := s.statisticService.CreateStatistic(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateStatistic - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create statistic"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDStatistic implements StatisticHandlerInterface.
func (s *statisticHandler) DeleteByIDStatistic(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDStatistic - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idStat := c.Param("id")
	id, err := conv.StringToInt64(idStat)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDStatistic - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = s.statisticService.DeleteByIDStatistic(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDStatistic - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete statistic"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// EditByIDStatistic implements StatisticHandlerInterface.
func (s *statisticHandler) EditByIDStatistic(c echo.Context) error {
	var (
		req       = request.StatisticRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDStatistic - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idStat := c.Param("id")
	id, err := conv.StringToInt64(idStat)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDStatistic - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDStatistic - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err := c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDStatistic - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.StatisticEntity{
		ID:    id,
		Name:  req.Name,
		Total: req.Total,
		Icon:  req.Icon,
	}

	err = s.statisticService.EditByIDStatistic(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDStatistic - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit statistic"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllStatistic implements StatisticHandlerInterface.
func (s *statisticHandler) FetchAllStatistic(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respStat  = []response.StatisticResponse{}
	)

	results, err := s.statisticService.FetchAllStatistic(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllStatistic - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respStat = append(respStat, response.StatisticResponse{
			ID:    val.ID,
			Name:  val.Name,
			Total: val.Total,
			Icon:  val.Icon,
		})
	}

	resp.Meta.Message = "Success fetch all statistic"
	resp.Meta.Status = true
	resp.Data = respStat
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDStatistic implements StatisticHandlerInterface.
func (s *statisticHandler) FetchByIDStatistic(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respStat  = response.StatisticResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDStatistic - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idStat := c.Param("id")
	id, err := conv.StringToInt64(idStat)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDStatistic - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := s.statisticService.FetchByIDStatistic(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDStatistic - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respStat.ID = result.ID
	respStat.Name = result.Name
	respStat.Total = result.Total
	respStat.Icon = result.Icon
	resp.Meta.Message = "Success fetch statistic by ID"
	resp.Meta.Status = true
	resp.Data = respStat
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewStatisticHandler(c *echo.Echo, cfg *config.Config, statisticService service.StatisticServiceInterface) StatisticHandlerInterface {
	statHandler := &statisticHandler{
		statisticService: statisticService,
	}

	mid := middleware.NewMiddleware(cfg)

	statApp := c.Group("/statistics")

	statApp.GET("", statHandler.FetchAllStatistic)

	adminApp := statApp.Group("/admin", mid.CheckToken())
	adminApp.GET("/:id", statHandler.FetchByIDStatistic)
	adminApp.POST("", statHandler.CreateStatistic)
	adminApp.PUT("/:id", statHandler.EditByIDStatistic)
	adminApp.DELETE("/:id", statHandler.DeleteByIDStatistic)

	return statHandler
}
