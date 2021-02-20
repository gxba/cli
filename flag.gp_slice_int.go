///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Sat Feb 20 2021 23:58 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_int_slice]
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
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var _ = (*strconv.NumError)(nil) //avoid compile error

// IntSliceSlice wraps []int to satisfy flag.Value
type IntSliceSlice struct {
	slice      []int
	hasBeenSet bool
}

// NewIntSliceSlice makes an *IntSliceSlice with default values
func NewIntSliceSlice(defaults ...int) *IntSliceSlice {
	return &IntSliceSlice{
		slice:      append([]int{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *IntSliceSlice) clone() *IntSliceSlice {
	n := &IntSliceSlice{
		slice:      make([]int, len(s.slice)),
		hasBeenSet: s.hasBeenSet,
	}
	copy(n.slice, s.slice)
	return n
}

// TODO: Consistently have specific Set function for Int64 and Float64 ?
// Append directly adds values to the list of values
func (s *IntSliceSlice) Append(value ...int) {
	if !s.hasBeenSet {
		s.slice = []int{}
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, value...)
}

// Set parses the value into an integer and appends it to the list of values
func (s *IntSliceSlice) Set(value string) error {

	if strings.HasPrefix(value, slPfx) {
		// Deserializing assumes overwrite
		_ = json.Unmarshal([]byte(strings.Replace(value, slPfx, "", 1)), &s.slice)
		s.hasBeenSet = true
		return nil
	}

	//accept multi values for slice flags
	for _, val := range flagSplitMultiValues(value) {
		value := strings.TrimSpace(val)
		tmp, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return err
		}

		if !s.hasBeenSet {
			s.slice = []int{}
			s.hasBeenSet = true
		}

		s.slice = append(s.slice, int(tmp))
	}

	return nil
}

// Reset clean the last parsed value of this slice
func (s *IntSliceSlice) Reset() {
	s.slice = s.slice[:0]
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *IntSliceSlice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows IntSliceSlice to fulfill Serializer
func (s *IntSliceSlice) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", slPfx, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *IntSliceSlice) Value() []int {
	return s.slice
}

// Get returns the slice set by this flag
func (s *IntSliceSlice) Get() interface{} {
	return *s
}

// IntSliceFlag define a value of type *IntSliceSlice
type IntSliceFlag struct {
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
	Target      *IntSliceSlice // Target value pointer outside
	Default     *IntSliceSlice // Default value
	DefaultText string         // Default value in help info
	Enums       []int          // Enumeration of valid values
	Ranges      []int          // {[min,max),[min,)...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *IntSliceSlice // target value pointer(maybe new(*IntSliceSlice) if Target not set)
	info   FlagInfo       // parsed info of this flag
}

// Init verify and init the value by ower flag
func (v *IntSliceFlag) Init() error {
	v.info.Flag = v
	v.info.EnvVars = v.EnvVars
	v.info.Usage = v.Usage
	v.info.DefaultText = v.DefaultText
	v.info.Required = v.Required
	v.info.Hidden = v.Hidden
	v.info.FilePath = v.FilePath
	v.info.LogicName = logicName(v.LogicName)
	v.info.Name = v.Name //TODO: deal with noname
	mergeNames(v.Name, v.Aliases, &v.info.Names)

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
		v.target = NewIntSliceSlice()
	}
	return nil
}

// IsSet check if value was set
func (v *IntSliceFlag) IsSet() bool {
	return v.target.hasBeenSet
}

// Apply coordinate the value to flagset
func (v *IntSliceFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (v *IntSliceFlag) String() string {
	return ""
}

// ValidateValues verify if all values was valid
func (v *IntSliceFlag) ValidateValues() error {
	return v.validateValues(v.target)
}

// Info returns parsed info of this flag
func (v *IntSliceFlag) Info() *FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (v *IntSliceFlag) Reset() {
	v.target.Reset()
	v.info.HasBeenSet = false
}

// for default value verify
func (v *IntSliceFlag) validateValues(values *IntSliceSlice) error {
	for _, val := range values.slice {
		if err := v.validValue(val); err != nil {
			return err
		}
	}
	return nil
}

// check if value if valid for this flag
func (v *IntSliceFlag) validValue(value int) error {
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
