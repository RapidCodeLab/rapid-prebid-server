package interfaces

// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/master/AdCOM%20v1.0%20FINAL.md#list--category-taxonomies-
const (
	EntityTypeBanner EntityType = iota + 1
	EntityTypeNative
)

const (
	InventoryTypeSite InventoryType = iota + 1
	InventoryTypeApp
)

type (
	EntityType    int
	InventoryType int

	InventoryStorager interface {
		EntityStorager
		ReadAllInventories() ([]Inventory, error)
		CreateInventory(Inventory) error
		ReadInventory(ID string) (Inventory, error)
		UpdateInventory(Inventory) error
		DeleteInventory(ID string) error
	}

	EntityStorager interface {
		EntityReadStorager
		ReadAllEntities() ([]Entity, error)
		CreateEntity(Entity) error
		UpdateEntity(Entity) error
		DeleteEntity(ID string) error
	}

	EntityReadStorager interface {
		ReadEntity(ID string) (Entity, error)
	}

	Entity struct {
		ID                              string        `json:"id"`
		Type                            EntityType    `json:"type"`
		InventoryID                     string        `json:"inventory_id"`   // OpenRTB Site.id or App.id
		InventoryType                   InventoryType `json:"inventory_type"` // OpenRTB Site or App object
		IABCategories                   []string      `json:"iab_categories"`
		BlockedAdvertisierIABCategories []string      `json:"blocked_advertisier_iab_categories"`
		IABCategoriesTaxonomy           int           `json:"iab_categories_taxonomy"`
		PlacementCount                  int           `json:"placement_count"` // for Native
		Width                           int           `json:"width"`           // will be use as wmin for Native
		Height                          int           `json:"height"`          // will be use as hmin for Native
	}

	Inventory struct {
		ID                              string        `json:"id,required"`
		Name                            string        `json:"name"`
		InventoryType                   InventoryType `json:"inventory_type"` // OpenRTB Site or App object
		IABCategories                   []string      `json:"iab_categories"`
		BlockedAdvertisierIABCategories []string      `json:"blocked_advertiser_iab_categories"`
		IABCategoriesTaxonomy           int           `json:"iab_categories_taxonomy"`
	}
)
