package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Organizations -
type Organizations struct {
	Organizations []Organization `tfsdk:"organizations"`
}

// Organization -
type Organization struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Url       types.String `tfsdk:"url"`
	Cloud     types.String `tfsdk:"cloud"`
	Api       Api          `tfsdk:"api"`
	Licensing Licensing    `tfsdk:"licensing"`
}

type Api struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type Licensing struct {
	Model types.String `tfsdk:"model"`
}

// Organization

// Order -
type Order struct {
	ID          types.String `tfsdk:"id"`
	Items       []OrderItem  `tfsdk:"items"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

// OrderItem -
type OrderItem struct {
	Coffee   Coffee `tfsdk:"coffee"`
	Quantity int    `tfsdk:"quantity"`
}

// Coffee -
// This Coffee struct is for Order.Items[].Coffee which does not have an
// ingredients` field in the schema defined by plugin framework. Since the
// resource schema must match the struct exactly (extra field will return an
// error). This struct has Ingredients commented out.
type Coffee struct {
	ID          int          `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Teaser      types.String `tfsdk:"teaser"`
	Description types.String `tfsdk:"description"`
	Price       types.Number `tfsdk:"price"`
	Image       types.String `tfsdk:"image"`
	// Ingredients []Ingredient   `tfsdk:"ingredients"`
}
