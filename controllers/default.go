package controllers

import (
	
	"mywebcalculator/mycalculator/calculator"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Layout="home.tpl"
	c.TplName = "home.tpl"
}

func (c *MainController) Post() {

	inputs:=c.Input()

	expression:=inputs.Get("expression")

	ca := calculator.NewCalculator(expression)

	result := ca.Run()

	c.Data["Expression"]=expression
	c.Data["Result"] = result

	c.TplName="layout.tpl"
}
