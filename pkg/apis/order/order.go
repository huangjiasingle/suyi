package order

import (
	"encoding/json"

	"github.com/huangjiasingle/suyi/pkg/apis/base"
	"github.com/huangjiasingle/suyi/pkg/models"
	"github.com/huangjiasingle/suyi/pkg/tools/http/server"
	"github.com/huangjiasingle/suyi/pkg/tools/storage/db"
	"github.com/valyala/fasthttp"

	"k8s.io/klog/v2"
)

func RegistryOrderAPI(s *server.Server) {
	h := &OrderHandler{Handler: &base.Handler{}}
	s.Registry("/orders", fasthttp.MethodGet, h.List)
	s.Registry("/orders", fasthttp.MethodPost, h.Create)
	s.Registry("/orders/{id}", fasthttp.MethodDelete, h.Delete)
	s.Registry("/orders/{id}", fasthttp.MethodGet, h.Get)
	s.Registry("/orders", fasthttp.MethodPut, h.Update)
}

type OrderHandler struct {
	*base.Handler
}

func (h *OrderHandler) List(ctx *fasthttp.RequestCtx) {
	userName := string(ctx.QueryArgs().Peek("user_name"))
	status := string(ctx.QueryArgs().Peek("status"))
	pageSize := ctx.QueryArgs().GetUintOrZero("page_size")
	if pageSize == 0 {
		pageSize = 10
	}
	pageNum := ctx.QueryArgs().GetUintOrZero("page_num")
	orderList := []models.Order{}
	count, err := db.Engine.Where("user_name=? and status=?", userName, status).Limit(pageSize, pageSize*(pageNum-1)).FindAndCount(&orderList)
	if err != nil {
		h.WriteResponse(ctx, &base.Response{Code: fasthttp.StatusInternalServerError, Reason: err.Error()})
		return
	}

	h.WriteResponse(ctx, &base.Response{Code: fasthttp.StatusOK, Data: models.Pagination{
		PageSise: pageSize,
		PageNum:  pageNum,
		Total:    count,
		Data:     orderList,
	}})
}

func (h *OrderHandler) Create(ctx *fasthttp.RequestCtx) {
	var o models.Order
	if err := json.Unmarshal(ctx.PostBody(), &o); err != nil {
		h.WriteResponse(ctx, &base.Response{Code: fasthttp.StatusBadRequest, Reason: err.Error()})
		return
	}

	count, err := db.Engine.InsertOne(&o)
	if err != nil {
		h.WriteResponse(ctx, &base.Response{Code: fasthttp.StatusInternalServerError, Reason: err.Error()})
	}
	klog.Info(count)
	h.WriteResponse(ctx, &base.Response{Code: fasthttp.StatusOK, Message: "创建订单成功"})
}

func (h *OrderHandler) Delete(ctx *fasthttp.RequestCtx) {
}

func (h *OrderHandler) Get(ctx *fasthttp.RequestCtx) {
}

func (h *OrderHandler) Update(ctx *fasthttp.RequestCtx) {
}
