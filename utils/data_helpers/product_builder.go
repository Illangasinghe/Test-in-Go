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
			Sellable:         DefaultSellable,
			Secure:           DefaultSecure,
			SKU:              []SKU{},
			AgeRestriction:   0,
			CanTaint:         false,
			CanBeTainted:     false,
			Hazardous:        "string",
			Restricted:       "string",
			FamilyGroup:      "Family Group 1",
			Catchweight:      true,
			Loose:            true,
			PrePick:          false,
			InStoreBakery:    false,
			SecurityTagged:   false,
			GoodsNotReady:    false,
			MadeToOrder:      false,
			Counter:          false,
			Organic:          true,
			VirtualStock:     false,
			AlwaysBag:        false,
			SubstitutionFrom: "DEFAULT",
			SubstitutionTo:   "DEFAULT",
			SubstitutionMode: "DEFAULT",
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

func (b *ProductBuilder) WithTestCode() *ProductBuilder {
	b.product.ProductCode = "PRD-" + TestCode
	b.product.ShortDescription = "desc" + TestCode
	b.product.Barcodes[0].Barcode = "421" + TestCode
	return b
}

func (b *ProductBuilder) WithShortDescription(description string) *ProductBuilder {
	b.product.ShortDescription = description
	return b
}

func (b *ProductBuilder) Build() Product {
	return b.product
}

func GenerateSKU() SKU {
	return SKU{
		SKUId:       "SKU-" + TestCode,
		Description: "Product SKU",
		SKUUom: []SKUUom{
			{
				UnitOfMeasure: "EACH",
				Height: ScalarUnit{
					Scalar: 50,
					Units:  "MM",
				},
				Width: ScalarUnit{
					Scalar: 60,
					Units:  "MM",
				},
				Depth: ScalarUnit{
					Scalar: 70,
					Units:  "MM",
				},
				Volume: ScalarUnit{
					Scalar: 210,
					Units:  "CC",
				},
				Weight: ScalarUnit{
					Scalar: 100,
					Units:  "G",
				},
				UnitsPerParent: []UnitsPerParent{
					{
						UnitOfMeasure: "EACH",
						NoOfUnits:     23,
					},
				},
			},
		},
		SupplierID:            "Gelato Inc",
		SupplierReference:     "stracciatella 500ml",
		MinimumLifeOnReceipt:  8,
		MinimumLifeOnDespatch: 7,
		RetailPrice: RetailPrice{
			CentsValue: 123,
			Currency:   "GBP",
		},
		CountryOfOrigin: "GBR",
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

// GenerateProductsRoot generates a Root structure containing multiple products.
func GenerateProductsRoot() Root {
	product1 := GenerateProductWithMandatoryFields()
	product2 := GenerateProductWithMultipleSKUs()

	return Root{
		Products: []Product{product1, product2},
	}
}
