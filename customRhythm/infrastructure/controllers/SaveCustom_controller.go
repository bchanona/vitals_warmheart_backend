package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/customRhythm/application/usecase"
	"github.com/bchanona/vitals_warmheart_backend/customRhythm/domain"
	"github.com/gin-gonic/gin"
)

type SaveCustomController struct {
	useCase *usecase.SaveCustomRhytmUc
}

func NewSaveCustomController(useCase *usecase.SaveCustomRhytmUc) *SaveCustomController{
	return &SaveCustomController{useCase: useCase}
}
func (controller *SaveCustomController) Execute(ctx *gin.Context){
	var customRhythm domain.CustomRhythmModel

	userIdStr, exists := ctx.Get("user_id") 
	
	if !exists {
		ctx.JSON(400, gin.H{"error": "id_user no encontrado"})
		return
	}
	if err := ctx.ShouldBindJSON(&customRhythm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	customRhythm.UserId = userIdStr.(int)

	if err := controller.useCase.Execute(customRhythm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering custom rhythm"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Custom rhythm registered successfully"})

}