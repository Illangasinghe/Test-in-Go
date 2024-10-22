package data_helpers

type ProductBuilder struct {
	product Product
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{
		product: Product{
			ProductCode:              "PRD-" + TestCode,
			LongDescription:          DefaultLongDescription,
			ShortDescription:         "desc" + TestCode,
			ImageUrl:                 DefaultImageUrl,
			ProductClass:             DefaultProductClass,
			ProductHierarchyId:       DefaultProductHierarchyID,
			TemperatureClass:         DefaultTemperatureClass,
			StorageArea:              DefaultStorageArea,
			PickingCodeCheckRequired: false,
			PutawayCodeCheckRequired: false,
			BarcodeScanRequired:      true,
			Barcodes: []Barcode{
				{Barcode: "421" + TestCode, BarcodeType: "EACH"},
			},
			Sellable: DefaultSellable,
			Secure:   DefaultSecure,
			SKU:      []SKU{},
		},
	}
}

func (b *ProductBuilder) WithLongDescription(longDescription string) *ProductBuilder {
	b.product.LongDescription = longDescription
	return b
}

func (b *ProductBuilder) WithoutOptionalFields() *ProductBuilder {
	b.product.LongDescription = ""
	b.product.ShortDescription = ""
	b.product.ImageUrl = ""
	return b
}

func (b *ProductBuilder) WithSKU(sku SKU) *ProductBuilder {
	b.product.SKU = append(b.product.SKU, sku)
	return b
}

// WithCalculatedTestCode calculates and sets the TestCode based on category, method, and testcase number
func (b *ProductBuilder) WithCalculatedTestCode(categoryNumber, methodNumber, testcaseNumber int) *ProductBuilder {
	testCode := GenerateTestCode(categoryNumber, methodNumber, testcaseNumber) // Updates the global TestCode variable
	b.product.ProductCode = "PRD-" + testCode
	b.product.ShortDescription = "desc" + testCode
	b.product.Barcodes[0].Barcode = "421" + testCode
	return b
}

func (b *ProductBuilder) Build() Product {
	return b.product
}

func GenerateSKU() SKU {
	return SKU{
		SKUId:                 "SKU-" + TestCode,
		Description:           "Product SKU",
		UnitOfMeasure:         "EACH",
		HeightScalar:          50,
		HeightUnits:           "mm",
		WidthScalar:           60,
		WidthUnits:            "mm",
		DepthScalar:           70,
		DepthUnits:            "mm",
		WeightScalar:          100,
		WeightUnits:           "g",
		VolumeScalar:          210,
		VolumeUnits:           "cc",
		NoOfUnits:             23,
		SupplierID:            "Gelato Inc",
		SupplierReference:     "stracciatella 500ml",
		MinimumLifeOnReceipt:  8,
		MinimumLifeOnDespatch: 7,
		RetailPriceCents:      123,
		RetailPriceCurrency:   "GBP",
		CountryOfOrigin:       "GBR",
		SKUBarcodes: []Barcode{
			{Barcode: "SKU" + TestCode, BarcodeType: "EACH"},
		},
	}
}

func GenerateProductWithMandatoryFields() Product {
	return NewProductBuilder().
		WithoutOptionalFields().
		WithSKU(GenerateSKU()).
		Build()
}

func GenerateProductWithMultipleSKUs() Product {
	sku1 := GenerateSKU()
	sku2 := GenerateSKU()
	return NewProductBuilder().
		WithSKU(sku1).
		WithSKU(sku2).
		Build()
}

// GenerateProductWithCustomFields allows flexible construction by passing custom SKUs or descriptions
func GenerateProductWithCustomFields(longDescription string, skus []SKU) Product {
	builder := NewProductBuilder()

	if longDescription != "" {
		builder.WithLongDescription(longDescription)
	}

	for _, sku := range skus {
		builder.WithSKU(sku)
	}

	return builder.Build()
}
