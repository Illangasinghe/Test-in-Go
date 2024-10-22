package data_helpers

// SKU represents the SKU structure of a product
type SKU struct {
	SKUId                 string    `json:"skuId"`
	Description           string    `json:"description"`
	UnitOfMeasure         string    `json:"unitOfMeasure"`
	HeightScalar          int       `json:"heightScalar"`
	HeightUnits           string    `json:"heightUnits"`
	WidthScalar           int       `json:"widthScalar"`
	WidthUnits            string    `json:"widthUnits"`
	DepthScalar           int       `json:"depthScalar"`
	DepthUnits            string    `json:"depthUnits"`
	WeightScalar          int       `json:"weightScalar"`
	WeightUnits           string    `json:"weightUnits"`
	VolumeScalar          int       `json:"volumeScalar"`
	VolumeUnits           string    `json:"volumeUnits"`
	NoOfUnits             int       `json:"noOfUnits"`
	SupplierID            string    `json:"supplierId"`
	SupplierReference     string    `json:"supplierReference"`
	MinimumLifeOnReceipt  int       `json:"minimumLifeOnReceipt"`
	MinimumLifeOnDespatch int       `json:"minimumLifeOnDespatch"`
	RetailPriceCents      int       `json:"retailPriceCents"`
	RetailPriceCurrency   string    `json:"retailPriceCurrency"`
	CountryOfOrigin       string    `json:"countryOfOrigin"`
	SKUBarcodes           []Barcode `json:"skuBarcodes"`
}

// Barcode represents the structure of a barcode
type Barcode struct {
	Barcode     string `json:"barcode"`
	BarcodeType string `json:"barcodeType"`
}

// Product represents the product message structure
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
	Secure                   bool      `json:"secure"`
	SKU                      []SKU     `json:"sku"`
}
