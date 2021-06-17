package service

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetInjury(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchText := ""
	if arg := c.Query("searchText"); arg != "" {
		searchText = arg
	}
	injuryParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchText": searchText,
	}
	err, info, total := models.GetAllInjury(injuryParam)
	//fmt.Println(info);
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"OK")
}

