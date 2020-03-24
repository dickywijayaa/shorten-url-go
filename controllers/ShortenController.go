package controllers

import (
	"github.com/dickywijayaa/shorten-url-go/objects"
	"github.com/dickywijayaa/shorten-url-go/services"
	"github.com/dickywijayaa/shorten-url-go/helpers"

	"os"
	"regexp"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ShortenController struct {
	service services.ShortenService
	responseHelper helpers.ResponseHelper
	helper helpers.Helper
}

func ShortenControllerHandler(router *gin.Engine) {
	handler := ShortenController{
		service: services.ShortenServiceHandler(),
		responseHelper: helpers.ResponseHelperHandler(),
		helper: helpers.HelperHandler(),
	}

	group := router.Group("/shorten", gin.BasicAuth(gin.Accounts{
		os.Getenv("API_USERNAME"): os.Getenv("API_PASSWORD"),
	}))

	group.GET("/:shortcode", handler.GetURLFromShortcode)
	group.POST("/", handler.PostShorten)
}

// godoc
// @Summary Return Redirect to URL based on Shortcode
// @Description Return Redirect to URL based on Shortcode stored in Database
// @Tags Shorten
// @Accept  json
// @Produce  json
// @Param shortcode path string true "shortcode"
// @Success 200 {object} objects.Response
// @Failure 400 {object} objects.Response
// @Failure 401 {string} string "Unauthorized"
// @Router /shorten/{shortcode} [get]
// @Security BasicAuth
func (c *ShortenController) GetURLFromShortcode(ctx *gin.Context) {
	code := ctx.Param("shortcode")
	if code == "" {
		response := c.responseHelper.BadRequestResponse("res", "failed validation")
		ctx.JSON(response.Code, response)
		return
	}

	res, err := c.service.FetchURLByCode(code)

	if err != nil {
		response := c.responseHelper.SuccessResponse(code, err.Error())
		ctx.JSON(response.Code, response)
		return
	}

	// redirect to url
	ctx.Redirect(http.StatusMovedPermanently, res)
	return
}

// godoc
// @Summary Store a shorten code for an URL
// @Description Store a shorten code for an URL in database
// @Tags Shorten
// @Accept  json
// @Produce  json
// @Param payload body objects.ShortenRequest true "Add shorten"
// @Success 200 {object} objects.Response
// @Failure 400 {object} objects.Response
// @Failure 401 {string} string "Unauthorized"
// @Router /shorten [post]
// @Security BasicAuth
func (c *ShortenController) PostShorten(ctx *gin.Context) {
	var payload objects.ShortenRequest 
	ctx.BindJSON(&payload)

	if payload.URL == "" {
		response := c.responseHelper.BadRequestResponse(payload, "URL is required")
		ctx.JSON(response.Code, response)
		return
	}

	if payload.Shortcode == "" {
		payload.Shortcode = c.helper.GenerateRandomString(6)
	}

	// shortcode must be 6 characters
	codePattern := regexp.MustCompile(`^[0-9a-zA-Z_]{6}$`)
	if !codePattern.MatchString(payload.Shortcode) {
		response := c.responseHelper.BadRequestResponse(payload, "invalid shortcode pattern")
		ctx.JSON(response.Code, response)
		return
	}

	res, err := c.service.StoreShortenURL(payload.URL, payload.Shortcode)

	if err == nil {
		response := c.responseHelper.SuccessResponse(res, "success insert to database")
		ctx.JSON(response.Code, response)
		return
	}

	response := c.responseHelper.FailedResponse(payload, err.Error())
	ctx.JSON(response.Code, response)
	return
}