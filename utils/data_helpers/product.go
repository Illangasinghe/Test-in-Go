package data_helpers

// Root represents the root structure containing a list of products.
type Root struct {
	Products []Product `json:"products"`
}

// SKU represents the SKU structure of a product.
type SKU struct {
	SKUId                 string      `json:"skuId"`
	Description           string      `json:"description"`
	SKUUom                []SKUUom    `json:"skuUom"`
	SupplierID            string      `json:"supplierId"`
	SupplierReference     string      `json:"supplierReference"`
	MinimumLifeOnReceipt  int         `json:"minimumLifeOnReceipt"`
	MinimumLifeOnDespatch int         `json:"minimumLifeOnDespatch"`
	RetailPrice           RetailPrice `json:"retailPrice"`
	CountryOfOrigin       string      `json:"countryOfOrigin"`
	SKUBarcodes           []Barcode   `json:"skuBarcodes"`
}

// SKUUom represents the Unit of Measure structure for a SKU.
type SKUUom struct {
	UnitOfMeasure  string           `json:"unitOfMeasure"`
	Height         ScalarUnit       `json:"height"`
	Width          ScalarUnit       `json:"width"`
	Depth          ScalarUnit       `json:"depth"`
	Volume         ScalarUnit       `json:"volume"`
	Weight         ScalarUnit       `json:"weight"`
	UnitsPerParent []UnitsPerParent `json:"unitsPerParent"`
}

// ScalarUnit represents dimensions and other scalar measurements.
type ScalarUnit struct {
	Scalar int    `json:"scalar"`
	Units  string `json:"units"`
}

// UnitsPerParent represents how many units per parent SKU.
type UnitsPerParent struct {
	UnitOfMeasure string `json:"unitOfMeasure"`
	NoOfUnits     int    `json:"noOfUnits"`
}

// RetailPrice represents the price structure.
type RetailPrice struct {
	CentsValue int    `json:"centsValue"`
	Currency   string `json:"currency"`
}

// Barcode represents the structure of a barcode.
type Barcode struct {
	Barcode     string `json:"barcode"`
	BarcodeType string `json:"barcodeType"`
}

// Product represents the product message structure.
type Product struct {
	ProductCode              string    `json:"productCode"`
	LongDescription          string    `json:"longDescription"`
	ShortDescription         string    `json:"shortDescription"`
	ImageUrl                 string    `json:"imageUrl"`
	ProductClass             string    `json:"productClass"`
	ProductHierarchyId       string    `json:"productHierarchyId"`
	TemperatureClass         string    `json:"temperatureClass"`
	StorageArea              string    `json:"storageArea"`
	PickingCodeCheckRequired bool      `json:"pickingCodeCheckRequired"`
	PutawayCodeCheckRequired bool      `json:"putawayCodeCheckRequired"`
	BarcodeScanRequired      bool      `json:"barcodeScanRequired"`
	Barcodes                 []Barcode `json:"barcodes"`
	Sellable                 bool      `json:"sellable"`
	AgeRestriction           int       `json:"ageRestriction"`
	CanTaint                 bool      `json:"canTaint"`
	CanBeTainted             bool      `json:"canBeTainted"`
	Hazardous                string    `json:"hazardous"`
	Restricted               string    `json:"restricted"`
	FamilyGroup              string    `json:"familyGroup"`
	Secure                   bool      `json:"secure"`
	Catchweight              bool      `json:"catchweight"`
	Loose                    bool      `json:"loose"`
	PrePick                  bool      `json:"prePick"`
	InStoreBakery            bool      `json:"inStoreBakery"`
	SecurityTagged           bool      `json:"securityTagged"`
	GoodsNotReady            bool      `json:"goodsNotReady"`
	MadeToOrder              bool      `json:"madeToOrder"`
	Counter                  bool      `json:"counter"`
	Organic                  bool      `json:"organic"`
	VirtualStock             bool      `json:"virtualStock"`
	AlwaysBag                bool      `json:"alwaysBag"`
	SKU                      []SKU     `json:"sku"`
	SubstitutionFrom         string    `json:"substitutionFrom"`
	SubstitutionTo           string    `json:"substitutionTo"`
	SubstitutionMode         string    `json:"substitutionMode"`
}
