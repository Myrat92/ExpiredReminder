package controllers

type FoodController struct {
	baseController
}

func (c *FoodController) Logging() {
	c.TplName = c.controllerName + "/logging.html"
}