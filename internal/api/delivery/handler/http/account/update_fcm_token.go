package account

import (
	"github.com/aasumitro/karlota/internal/api/domain"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// update fcm_token godoc
// @Schemes
// @Summary Update FCM TOKEN
// @Description Store Firebase Cloud Messaging Token.
// @Tags AccountHandler
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param fcm_token formData string true "firebase cloud messaging token"
// @Success 201 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE_ENTITY_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /v1/update/fcm [POST]
func (handler *accountHandler) updateFCMToken(context *gin.Context) {
	var form domain.UserFCMTokenForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewFormRequest(domain.UserFormErrorMessages).Validate(form, err)
		utils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	payload := context.MustGet("payload")
	email := payload.(map[string]interface{})["email"]
	if err := handler.service.Edit(&domain.User{
		Email:    email.(string),
		FCMToken: form.FCMToken,
	}); err != nil {
		utils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	utils.NewHttpRespond(context, http.StatusOK, "UPDATED")
}
