//go:build tools

package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)

// ValueConversion - Extracts data from the GetPayload interface{} when the struct is not reveled
func ValueConversion(data interface{}, results interface{}) {
	jsonString, _ := json.Marshal(data)

	// convert json to struct
	err := json.Unmarshal(jsonString, &results)
	if err != nil {
		return
	}
}
