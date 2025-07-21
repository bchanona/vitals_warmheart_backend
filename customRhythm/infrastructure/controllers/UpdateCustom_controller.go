package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/customRhythm/application/usecase"
	"github.com/bchanona/vitals_warmheart_backend/customRhythm/domain"
	"github.com/gin-gonic/gin"
)

type UpdateCustomRhytmController struct {
	useCase *usecase.UpdateCustomRhytmUc
}

func NewUpdateCustomRhytmController(useCase *usecase.UpdateCustomRhytmUc) *UpdateCustomRhytmController {
	return &UpdateCustomRhytmController{useCase: useCase}
}
func (controller *UpdateCustomRhytmController) Execute(ctx *gin.Context) {
	var customRhythm domain.CustomRhythmModel

	idUser, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(400, gin.H{"error": "User ID not found "})
		return
	}

	if err := ctx.ShouldBindJSON(&customRhythm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := controller.useCase.Execute(idUser.(int), customRhythm)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating custom rhythm"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Custom rhythm updated successfully"})

}
