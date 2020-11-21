package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/service/category"
	"github.com/kekaiwang/go-blog/utils/tools"
)

func GetCategoryList(ctx *gin.Context) {
	fmt.Println(ctx.Param("link"))
	fmt.Println(ctx.Query("type"))

	ctx.Header("Content-type", "text/html; charset=utf-8")

	var (
		link    = ctx.Param("link")
		req     category.GetCategoryReq
		page    int64
		err     error
		limit   = int64(10)
		pageStr = ctx.Query("page")
	)

	if link == "" {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			ctx.HTML(http.StatusOK, "error.html", nil)
			return
		}
	}

	if page == 0 {
		page = 1
	}
	req.Offset = (page - 1) * limit
	req.Limit = limit
	fmt.Println(req)
	req.Link = link

	data, err := req.GetCategoryList()
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"Title": "Kekai Wang",
		})
		return
	}

	meta := category.Meta{
		Name: "分类",
		Type: ctx.Query("type"),
	}

	ctx.HTML(http.StatusOK, "category.html", gin.H{
		"data":          data.Data,
		"Title":         data.Name,
		"total":         data.Total,
		"total_page":    tools.NewTotalPage(data.Total, limit),
		"current_page":  page,
		"next_page":     page + 1,
		"previous_page": page - 1,
		"meta":          meta,
	})
}
