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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PostHandlerInterface interface {
	CreatePost(c echo.Context) error
	FetchAllPosts(c echo.Context) error
	FetchByIDPost(c echo.Context) error
	FetchBySlugPost(c echo.Context) error
	EditByIDPost(c echo.Context) error
	DeleteByIDPost(c echo.Context) error
}

type postHandler struct {
	postService service.PostServiceInterface
}

// CreatePost implements PostHandlerInterface.
func (p *postHandler) CreatePost(c echo.Context) error {
	var (
		req       = request.PostRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreatePost - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreatePost - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err := c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreatePost - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringPublishedAt, err := time.Parse("2006-01-02", req.PublishedAt)
	if err != nil {
		log.Errorf("[HANDLER] CreatePost - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PostEntity{
		Title:        req.Title,
		Slug:         req.Slug,
		Author:       req.Author,
		FeaturedImage: req.FeaturedImage,
		Content:      req.Content,
		PublishedAt:  stringPublishedAt,
	}

	err = p.postService.CreatePost(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreatePost - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create post"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDPost implements PostHandlerInterface.
func (p *postHandler) DeleteByIDPost(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDPost - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPost := c.Param("id")
	id, err := conv.StringToInt64(idPost)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPost - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = p.postService.DeleteByIDPost(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPost - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete post"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// EditByIDPost implements PostHandlerInterface.
func (p *postHandler) EditByIDPost(c echo.Context) error {
	var (
		req       = request.PostRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDPost - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPost := c.Param("id")
	id, err := conv.StringToInt64(idPost)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPost - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDPost - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err := c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDPost - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringPublishedAt, err := time.Parse("2006-01-02", req.PublishedAt)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPost - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PostEntity{
		ID:           id,
		Title:        req.Title,
		Slug:         req.Slug,
		Author:       req.Author,
		FeaturedImage: req.FeaturedImage,
		Content:      req.Content,
		PublishedAt:  stringPublishedAt,
	}

	err = p.postService.EditByIDPost(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPost - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit post"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllPosts implements PostHandlerInterface.
func (p *postHandler) FetchAllPosts(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respPosts = []response.PostResponse{}
	)

	results, err := p.postService.FetchAllPosts(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllPosts - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respPosts = append(respPosts, response.PostResponse{
			ID:           val.ID,
			Title:        val.Title,
			Slug:         val.Slug,
			Author:       val.Author,
			FeaturedImage: val.FeaturedImage,
			Content:      val.Content,
			PublishedAt:  val.PublishedAt.Format("02 Jan 2006 15:04:05"),
		})
	}

	resp.Meta.Message = "Success fetch all posts"
	resp.Meta.Status = true
	resp.Data = respPosts
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDPost implements PostHandlerInterface.
func (p *postHandler) FetchByIDPost(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respPost  = response.PostResponse{}
	)

	idPost := c.Param("id")
	id, err := conv.StringToInt64(idPost)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPost - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := p.postService.FetchByIDPost(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPost - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPost.ID = result.ID
	respPost.Title = result.Title
	respPost.Slug = result.Slug
	respPost.Author = result.Author
	respPost.FeaturedImage = result.FeaturedImage
	respPost.Content = result.Content
	respPost.PublishedAt = result.PublishedAt.Format("02 Jan 2006 15:04:05")
	resp.Meta.Message = "Success fetch post by ID"
	resp.Meta.Status = true
	resp.Data = respPost
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func (p *postHandler) FetchBySlugPost(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respPost  = response.PostResponse{}
	)

	slug := c.Param("slug")

	result, err := p.postService.FetchBySlugPost(ctx, slug)
	if err != nil {
		log.Errorf("[HANDLER] FetchBySlugPost - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPost.ID = result.ID
	respPost.Title = result.Title
	respPost.Slug = result.Slug
	respPost.Author = result.Author
	respPost.FeaturedImage = result.FeaturedImage
	respPost.Content = result.Content
	respPost.PublishedAt = result.PublishedAt.Format("02 Jan 2006 15:04:05")
	resp.Meta.Message = "Success fetch post by slug"
	resp.Meta.Status = true
	resp.Data = respPost
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewPostHandler(c *echo.Echo, cfg *config.Config, postService service.PostServiceInterface) PostHandlerInterface {
	postHandler := &postHandler{
		postService: postService,
	}

	mid := middleware.NewMiddleware(cfg)

	postApp := c.Group("/posts")

	postApp.GET("", postHandler.FetchAllPosts)
	postApp.GET("/:id", postHandler.FetchByIDPost)
	postApp.GET("/slug/:slug", postHandler.FetchBySlugPost)

	adminApp := postApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", postHandler.CreatePost)
	adminApp.PUT("/:id", postHandler.EditByIDPost)
	adminApp.DELETE("/:id", postHandler.DeleteByIDPost)

	return postHandler
}
