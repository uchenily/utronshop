package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gernest/utron/controller"
	"utronshop.io/models"
)

//Product is a controller for product
type Product struct {
	controller.BaseController
	Routes []string
}

// Detail 商品详情页面
func (p *Product) Detail() {
	cookie, err := p.Ctx.Request().Cookie("username")
	if err != nil {
		log.Println(">>INFO>> product detail page, user not login")
	} else {
		p.Ctx.Data["User"] = models.User{Name: cookie.Value}
	}
	// 获取商品ID
	// s := p.Ctx.Request().RequestURI
	// id, _ := strconv.Atoi(s[len("/detail/"):])
	number, _ := strconv.Atoi(p.Ctx.Params["number"])
	log.Println("product id: ", number)
	product := &models.Product{}
	p.Ctx.DB.Where("number = ?", number).First(&product)
	p.Ctx.Data["Product"] = product

	curdir, _ := os.Getwd()
	BriefByteData, err := ioutil.ReadFile(curdir + product.Brief)
	DetailByteData, err := ioutil.ReadFile(curdir + product.Detail)
	fmt.Println(curdir + product.Brief)
	BriefData := string(BriefByteData)
	DetailData := string(DetailByteData)
	p.Ctx.Data["Product_Desc"] = BriefData
	p.Ctx.Data["Product_Detail"] = DetailData

	p.Ctx.Template = "details"
	p.HTML(http.StatusOK)
}

//Search 搜索页面
func (p *Product) Search() {
	cookie, err := p.Ctx.Request().Cookie("username")
	if err != nil {
		log.Println(">>INFO>> product detail page, user not login")
	} else {
		p.Ctx.Data["User"] = models.User{Name: cookie.Value}
	}

	keyword := p.Ctx.Request().FormValue("keyword")
	if keyword != "" {
		products := &[]models.Product{} // 数组
		p.Ctx.DB.Where("type LIKE ? or name LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Limit(5).Find(&products)
		p.Ctx.Data["Products"] = products
	}
	p.Ctx.Template = "search"
	p.HTML(http.StatusOK)
}

//NewProduct returns a new  product controller
func NewProduct() controller.Controller {
	return &Product{
		Routes: []string{
			"get;/detail/{number:[0-9]+};Detail",
			"get;/search;Search",
		},
	}
}
