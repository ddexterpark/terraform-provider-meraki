package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Organization -
type Organization struct {
	ID        types.String  `tfsdk:"id"`
	Name      types.String  `tfsdk:"name"`
	Url       types.String  `tfsdk:"url"`
	Api       types.MapType `tfsdk:"api"`
	Licensing types.MapType `tfsdk:"licensing"`
}

// Organizations -
type Organizations struct {
	Organization []Organization `tfsdk:"organization"`
}
