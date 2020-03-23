package controllers

import (
	"../objects"
	"../services"
	"../helpers"

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

	group := router.Group("/shorten")
	group.GET("/:shortcode", handler.GetURLFromShortcode)
	group.POST("/", handler.PostShorten)
}

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