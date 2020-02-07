package main

import (
	_ "classOne/routers"
	"github.com/astaxie/beego"
	_ "classOne/models"
)

func main() {
	beego.AddFuncMap("prePage", HandlePrePage)
	beego.AddFuncMap("nextPage", HandleNextPage)
	beego.Run()

}

func HandlePrePage(pageIndex int) int {
	if pageIndex <= 0 {
		pageIndex = 1
		return pageIndex
	}
	return pageIndex-1
}

func HandleNextPage(pageIndex, pageCount int) int {
	if pageIndex >= pageCount{
		pageIndex = pageCount
		return pageIndex
	}
	return pageIndex+1;
}