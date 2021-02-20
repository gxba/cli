///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Sat Feb 20 2021 17:45 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_string]
//
// Tool [github.com/gxlb/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

package cli

import (
	"flag"
	"fmt"
)

// StringFlag define a value of type string
type StringFlag struct {
	//
	//name related area
	//
	LogicName string   // logic name of the flag
	Name      string   // name of the flag
	Aliases   []string // aliases of the flag
	Usage     string   // usage string
	Required  bool     // if required
	Hidden    bool     // hidden this flag
	EnvVars   []string // environment values
	FilePath  string   // file path
	//
	//value related area
	//
	Target      *string  // Target value pointer outside
	Default     string   // Default value
	DefaultText string   // Default value in help info
	Enums       []string // Enumeration of valid values
	Ranges      []string // {[min,max),[min,)...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *string  // target value pointer(maybe new(string) if Target not set)
	info   FlagInfo // parsed info of this flag
}

// Init verify and init the value by ower flag
func (v *StringFlag) Init() error {
	v.info.Flag = v
	if l := len(v.Enums); l > maxSliceLen {
		return fmt.Errorf("flag %s.Enums too long: %d/%d", v.info.LogicName, l, maxSliceLen)
	}
	if l := len(v.Ranges); l > 0 {
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Ranges too long: %d/%d", v.info.LogicName, l, maxSliceLen)
		}
		if l%2 != 0 {
			return fmt.Errorf("flag %s.Ranges doesn't match [min,max) pairs: %d", v.info.LogicName, l)
		}
		for i := 0; i < l; i += 2 {
			min, max := v.Ranges[i], v.Ranges[i+1]
			if valid := min <= max; !valid {
				return fmt.Errorf("flag %s.Ranges doesn't match [min,max): (%d,%d)", v.info.LogicName, min, max)
			}
		}
	}
	if v.Name == "" && v.LogicName == "" {
		return fmt.Errorf("flag missing both Name & LogicName: %v", v)
	}

	if err := v.validateValues(v.Default); err != nil {
		return fmt.Errorf("default value invalid: %s", err.Error())
	}
	if v.Target != nil {
		v.target = v.Target
	} else {
		v.target = new(string)
	}
	return nil
}

// IsSet check if value was set
func (v *StringFlag) IsSet() bool {
	return v.info.HasBeenSet
}

// Apply coordinate the value to flagset
func (v *StringFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (v *StringFlag) String() string {
	return ""
}

// ValidateValues verify if all values was valid
func (v *StringFlag) ValidateValues() error {
	return v.validateValues(*v.target)
}

// Info returns parsed info of this flag
func (v *StringFlag) Info() *FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (v *StringFlag) Reset() {
	var t string
	*v.target = t
	v.info.HasBeenSet = false
}

// for default value verify
func (v *StringFlag) validateValues(values string) error {
	return v.validValue(values)
}

// check if value if valid for this flag
func (v *StringFlag) validValue(value string) error {
	f := &v.info
	if len(v.Enums) > 0 {
		found := false
		for _, v := range v.Enums {
			if value == v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Enums: %v", f.LogicName, value, v.Enums)
		}
	}
	if len(v.Ranges) > 0 {
		found := false
		for i := 0; i < len(v.Ranges); i += 2 {
			min, max := v.Ranges[i], v.Ranges[i+1]
			if value >= min && value < max {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Ranges: %v", f.LogicName, value, v.Enums)
		}
	}
	return nil
}

var _ Flag = (*StringFlag)(nil) //for interface verification only
