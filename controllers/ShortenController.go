package controllers

import (
	"../objects"
	"../services"
	"../helpers"

	"os"
	"regexp"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v10"
)

var validate *validator.Validate

type ShortenController struct {
	service services.ShortenService
	helper helpers.ResponseHelper
}

func ShortenControllerHandler(router *gin.Engine) {
	handler := ShortenController{
		service: services.ShortenServiceHandler(),
		helper: helpers.ResponseHelperHandler(),
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
func (h *ShortenController) GetURLFromShortcode(ctx *gin.Context) {
	code := ctx.Param("shortcode")
	if code == "" {
		response := h.helper.BadRequestResponse("res", "failed validation")
		ctx.JSON(response.Code, response)
		return
	}

	res, err := h.service.FetchURLByCode(code)

	if err != nil {
		response := h.helper.SuccessResponse(code, err.Error())
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
func (h *ShortenController) PostShorten(ctx *gin.Context) {
	var payload objects.ShortenRequest 
	ctx.BindJSON(&payload)

	validate = validator.New()
	err := validate.Struct(payload)
	if err != nil {
		response := h.helper.BadRequestResponse(payload, err.Error())
		ctx.JSON(response.Code, response)
		return
	}

	// shortcode must be 6 characters
	codePattern := regexp.MustCompile(`^[0-9a-zA-Z_]{6}$`)
	if !codePattern.MatchString(payload.Shortcode) {
		response := h.helper.BadRequestResponse(payload, "invalid shortcode pattern")
		ctx.JSON(response.Code, response)
		return
	}

	res, err := h.service.StoreShortenURL(payload.URL, payload.Shortcode)

	if res == true && err == nil {
		response := h.helper.SuccessResponse(res, "success insert to database")
		ctx.JSON(response.Code, response)
		return
	}

	response := h.helper.FailedResponse(payload, err.Error())
	ctx.JSON(response.Code, response)
	return
}