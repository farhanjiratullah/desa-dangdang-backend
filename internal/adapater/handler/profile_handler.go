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

type ProfileHandlerInterface interface {
	FetchByIDProfile(c echo.Context) error
	EditByIDProfile(c echo.Context) error
}

type profileHandler struct {
	profileService service.ProfileServiceInterface
}

// FetchByIDProfile implements ProfileHandlerInterface.
func (p *profileHandler) FetchByIDProfile(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respProfile response.ProfileResponse
	)

	// Fetch the profile by ID
	id, err := conv.StringToInt64(c.Param("id"))
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDProfile - 1: %v", err)
		respError.Meta.Message = "Invalid profile ID"
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := p.profileService.FetchByIDProfile(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDProfile - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	// Mapping the fetched data to response
	respProfile.ID = result.ID
	respProfile.Title = result.Title
	respProfile.Content = result.Content

	// Return the success response
	resp.Meta.Message = "Successfully fetched profile"
	resp.Meta.Status = true
	resp.Data = respProfile
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// EditByIDProfile implements ProfileHandlerInterface.
func (p *profileHandler) EditByIDProfile(c echo.Context) error {
	var (
		req       = request.ProfileRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDProfile - Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	// Fetch the profile by ID
	id, err := conv.StringToInt64(c.Param("id"))
	if err != nil {
		log.Errorf("[HANDLER] EditByIDProfile - 1: %v", err)
		respError.Meta.Message = "Invalid profile ID"
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDProfile - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	// Validate the input
	if err := c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDProfile - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	// Create the ProfileEntity
	reqEntity := entity.ProfileEntity{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}

	// Call service to update the profile
	err = p.profileService.EditByIDProfile(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDProfile - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Successfully updated profile"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewProfileHandler(c *echo.Echo, cfg *config.Config, profileService service.ProfileServiceInterface) ProfileHandlerInterface {
	profileHandler := &profileHandler{
		profileService: profileService,
	}

	mid := middleware.NewMiddleware(cfg)

	profileApp := c.Group("/profile")

	profileApp.GET("/:id", profileHandler.FetchByIDProfile)
	
	adminApp := profileApp.Group("/admin", mid.CheckToken())
	adminApp.PUT("/:id", profileHandler.EditByIDProfile)

	return profileHandler
}
