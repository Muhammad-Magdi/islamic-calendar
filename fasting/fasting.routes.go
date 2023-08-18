package fasting

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FastingRouter struct {
	fastingService *FastingService
}

func NewFastingRouter(fastingService *FastingService) FastingRouter {
	return FastingRouter{fastingService}
}

func (r FastingRouter) GetCurrentMonthFastingDays(c *gin.Context) {
	layoutISO := "2006-01-02"

	params := FastingDaysQueryParams{}
	err := c.BindQuery(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"message": "TODO", "error": err.Error()})
		return
	}

	fastingStrategy := NewFastingStrategy(params.Strategy)
	r.fastingService.SetFastingStrategy(fastingStrategy)
	fastingDays := r.fastingService.GetCurrentMonthFastingDays()

	formattedDates := make(map[FastingDay]string)
	for key, value := range fastingDays {
		formattedDates[key] = value.Format(layoutISO)
	}

	c.JSON(http.StatusOK, formattedDates)
}
