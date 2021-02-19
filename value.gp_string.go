///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sat Feb 20 2021 01:47:43]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/value.gp]
//   [github.com/gxlb/cli/value.gpg] [value_string]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

package cli

////////////////////////////////////////////////////////////////////////////////

// StringValue define a value of type string
type StringValue struct {
	Value       string   // The value from ENV of files
	Target      *string  // Target set the outer value pointer
	Default     string   // Default value
	DefaultText string   // Default value help info
	Enums       []string // Enumeration of valid values
	Ranges      []string // [min,max,min,max...] ranges of the valid values
	hasBeenSet  bool
}
