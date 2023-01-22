// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	Uid    int64 `json:"uid"`
	Oid    int64 `json:"oid"`
	Amount int64 `json:"amount"`
}

type CreateResponse struct {
	Id int64 `json:"id"`
}

type DetailRequest struct {
	Id int64 `json:"id"`
}

type DetailResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Oid    int64 `json:"oid"`
	Amount int64 `json:"amount"`
	Source int64 `json:"source"`
	Status int64 `json:"status"`
}

type CallbackRequest struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Oid    int64 `json:"oid"`
	Amount int64 `json:"amount"`
	Source int64 `json:"source"`
	Status int64 `json:"status"`
}

type CallbackResponse struct {
}
