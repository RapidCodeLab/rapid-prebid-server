package inventoryapi_handler

import "github.com/buaazp/fasthttprouter"

func (h *Handler) LoadRoutes(r *fasthttprouter.Router) {
	r.GET("", h.HandleReadAllInventories)
	r.POST("", h.HandleCreateInventory)
	r.GET("", h.HandleReadInventory)
	r.POST("", h.HandleUpdateInventory)
	r.GET("", h.HandleDeleteInventory)
	//
	r.GET("", h.HandleReadAllEntities)
	r.POST("", h.HandleCreateEntity)
	r.GET("", h.HandleReadEntity)
	r.POST("", h.HandleUpdateEntity)
	r.GET("", h.HandleDeleteEntity)
}
