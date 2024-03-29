package web

import "time"


type CustItem struct {
	OrderedAt time.Time `json:"orderedAt"`
	CustomerName string `json:"customerName"`
	Items []Items `json:"items"`
}

type Items struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int`json:"quantity"`
}

type WebResponse struct{
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Result  interface{} `json:"data"`
}

type BaseResponse struct{
	Message string      `json:"message"`
	Status  int         `json:"status"`
}
