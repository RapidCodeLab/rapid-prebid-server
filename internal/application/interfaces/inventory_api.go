package interfaces

// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/master/AdCOM%20v1.0%20FINAL.md#list--category-taxonomies-
type (
	EntityType    int
	InventoryType int
)

const (
	EntityTypeBanner EntityType = iota + 1
	EntityTypeNative
)

const (
	InventoryTypeSite InventoryType = iota + 1
	InventoryTypeApp
)

type InventoryEntity struct {
	EntityID                     string
	EntityType                   EntityType
	InventoryID                  string // OpenRTB Site.id or App.id
	InventoryType                int    // OpenRTB Site or App object
	EntityCategories             []string
	BlockedAdvertisierCategories []string
	CategoriesTaxonomy           int
	PlacementCount               int // for Native
	Width                        int // will be use as wmin for Native
	Height                       int // will be use as hmin for Native
}

type Inventory struct{}

type InventoryAPI interface {
	EntityAPI
}

type EntityAPI interface {
	Create(InventoryEntity) error
	Read(EntityID string) (InventoryEntity, error)
	Update(InventoryEntity) error
	Delete(InventoryEntity) error
}
