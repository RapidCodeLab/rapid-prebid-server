package inventoryapi_handler

import "github.com/buaazp/fasthttprouter"

func (h *Handler) LoadRoutes(r *fasthttprouter.Router) {
	r.GET("/inventory-api/inventories", h.HandleReadAllInventories)
	r.POST("/inventory-api/inventory/create", h.HandleCreateInventory)
	r.GET("/inventory-api/inventory/read/:id", h.HandleReadInventory)
	r.POST("/inventory-api/inventory/update", h.HandleUpdateInventory)
	r.GET("/inventory-api/inventory/delete/:id", h.HandleDeleteInventory)
	//
	r.GET("/inventory-api/entities", h.HandleReadAllEntities)
	r.POST("/inventory-api/entity/create", h.HandleCreateEntity)
	r.GET("/inventory-api/entity/read/:id", h.HandleReadEntity)
	r.POST("/inventory-api/entity/update", h.HandleUpdateEntity)
	r.GET("/inventory-api/entity/delete/:id", h.HandleDeleteEntity)
}
