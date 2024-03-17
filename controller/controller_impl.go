package controller

import (
	"net/http"
	"strconv"

	"github.com/MochammadQemalFirza/assignment2/model/web"
	"github.com/MochammadQemalFirza/assignment2/service"
	"github.com/gin-gonic/gin"
)

type ControllerImpl struct {
	Service service.Service
}

// CreateOrder godoc
// @Tags orders
// @Description Create Order Data
// @ID create-new-order
// @Accept json
// @Produce json
// @Param RequestBody body web.CustItem true "request body json"
// @Success 201 {object} web.WebResponse
// @Router /orders [post]
func (controller *ControllerImpl) CreateOrdersItems(c *gin.Context) {
	
	payload := web.CustItem{}
	
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, web.BaseResponse{
			Message: "Invalid request body",
			Status : http.StatusBadRequest,
		})
		return
	}

	res, err := controller.Service.CreateOrdersItems(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BaseResponse{
			Message: "Failed to create order items",
			Status:  http.StatusInternalServerError,
		})
		return
	}


	c.JSON(http.StatusOK, web.WebResponse{
		Message: "successfully created order",
		Status: http.StatusOK,
		Result: res,
	})
	
}

// @Tags orders
// @Description Get Order with Item Data
// @ID get-orders-with-items
// @Produce json
// @Success 200 {object} web.WebResponse
// @Router /orders [get]
func(controller *ControllerImpl)GetAllOrdersItems(c *gin.Context){

	res, err:= controller.Service.GetAllOrdersItems()

	if err != nil{
		c.JSON(http.StatusInternalServerError, web.BaseResponse{
			Message: "Failed to get order items",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, web.BaseResponse{
			Message :"Data not found!",
			Status: http.StatusNotFound,
		})
		return 
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Message: "successfully get all orders items",
		Status: http.StatusOK,
		Result: res,
	})
}

// @Tags orders
// @Description Update Order Data By Id
// @ID update-order
// @Accept json
// @Produce json
// @Param orderId path int true "order's id"
// @Param RequestBody body web.CustItem true "request body json"
// @Success 200 {object} web.WebResponse
// @Router /orders/{orderId} [put]
func(controller *ControllerImpl)UpdateOrdersItems(c *gin.Context){

	orderIDStr := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, web.BaseResponse{
			Message: "invalid order ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	var payload web.CustItem
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, web.BaseResponse{
			Message: "invalid request body",
			Status:  http.StatusBadRequest,
		})
		return
	}

	res, err := controller.Service.UpdateOrdersItems(orderID, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BaseResponse{
			Message: "Failed to update order items",
			Status: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Message: "order items updated successfully",
		Status: http.StatusOK,
		Result: res,
	})
}

func(controller *ControllerImpl)DeleteOrdersItems(c *gin.Context){

	orderIDStr := c.Param("order_id")
	
	orderID, err := strconv.Atoi(orderIDStr)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, web.BaseResponse{
			Message: "invalid order ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = controller.Service.DeleteOrdersItemsByID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BaseResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, web.BaseResponse{
		Message :"Order deleted successfully",
		Status : http.StatusOK,
	})
}

func NewController(Service service.Service) Controller {
	return &ControllerImpl{Service: Service}
}