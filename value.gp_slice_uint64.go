///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Sat Feb 20 2021 14:45 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/value.gp]
//   [github.com/gxlb/cli/value.gpg] [value_uint64_slice]
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
	"strconv"
	//"encoding/json"
	//"strings"
)

var _ = (*strconv.NumError)(nil) //avoid compile error

// Uint64SliceSlice wraps []uint64 to satisfy flag.Value
type Uint64SliceSlice struct {
	slice      []uint64
	hasBeenSet bool
}

// NewUint64SliceSlice makes an *Uint64SliceSlice with default values
func NewUint64SliceSlice(defaults ...uint64) *Uint64SliceSlice {
	return &Uint64SliceSlice{slice: append([]uint64{}, defaults...)}
}

// clone allocate a copy of self object
func (s *Uint64SliceSlice) clone() *Uint64SliceSlice {
	n := &Uint64SliceSlice{
		slice:      make([]uint64, len(s.slice)),
		hasBeenSet: s.hasBeenSet,
	}
	copy(n.slice, s.slice)
	return n
}

// TODO: Consistently have specific Set function for Int64 and Float64 ?
// Append directly adds an integer to the list of values
func (s *Uint64SliceSlice) Append(value ...uint64) {
	if !s.hasBeenSet {
		s.slice = []uint64{}
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, value...)
}

// Set parses the value into an integer and appends it to the list of values
func (s *Uint64SliceSlice) Set(value string) error {
	if !s.hasBeenSet {
		s.slice = []uint64{}
		s.hasBeenSet = true
	}

	// if strings.HasPrefix(value, slPfx) {
	// 	// Deserializing assumes overwrite
	// 	_ = json.Unmarshal([]byte(strings.Replace(value, slPfx, "", 1)), &s.slice)
	// 	s.hasBeenSet = true
	// 	return nil
	// }

	tmp, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return err
	}

	s.slice = append(s.slice, uint64(tmp))

	return nil
}

// String returns a readable representation of this value (for usage defaults)
func (s *Uint64SliceSlice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows Uint64SliceSlice to fulfill Serializer
func (s *Uint64SliceSlice) Serialize() string {
	//TODO:
	// jsonBytes, _ := json.Marshal(s.slice)
	// return fmt.Sprintf("%s%s", slPfx, string(jsonBytes))
	return ""
}

// Value returns the slice of ints set by this flag
func (s *Uint64SliceSlice) Value() []uint64 {
	return s.slice
}

// Get returns the slice set by this flag
func (s *Uint64SliceSlice) Get() interface{} {
	return *s
}

// Uint64SliceValue define a value of type GOGPElemType
type Uint64SliceValue struct {
	Target      *Uint64SliceSlice // Target value pointer outside
	Default     *Uint64SliceSlice // Default value
	DefaultText string            // Default value in help info
	Enums       []uint64          // Enumeration of valid values
	Ranges      []uint64          // {[min,max),[min,max),[min...)} ranges of valid values
	value       *Uint64SliceSlice // The value from ENV of files
	hasBeenSet  bool              // if the value was set
	flag        *Flag             // pointer of owner flag
}

// Init verify and init the value by ower flag
func (v *Uint64SliceValue) Init(f *Flag) error {
	v.flag = f
	if l := len(v.Enums); l > maxSliceLen {
		return fmt.Errorf("flag %s.Enums too long: %d/%d", v.flag.logicName, l, maxSliceLen)
	}
	if l := len(v.Ranges); l > maxSliceLen {
		return fmt.Errorf("flag %s.Ranges too long: %d/%d", v.flag.logicName, l, maxSliceLen)
	}
	if err := v.validateValues(v.Default); err != nil {
		return fmt.Errorf("default value invalid: %s", err.Error())
	}
	return nil
}

// IsSet check if value was set
func (v *Uint64SliceValue) IsSet() bool {
	return v.value.hasBeenSet
}

// Apply coordinate the value to flagset
func (v *Uint64SliceValue) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (v *Uint64SliceValue) String() string {
	return ""
}

// ValidateValues verify if all values was valid
func (v *Uint64SliceValue) ValidateValues() error {
	return v.validateValues(v.value)
}

// for default value verify
func (v *Uint64SliceValue) validateValues(values *Uint64SliceSlice) error {
	for _, val := range values.slice {
		if err := v.validValue(val); err != nil {
			return err
		}
	}
	return nil
}

// check if value if valid for this flag
func (v *Uint64SliceValue) validValue(value uint64) error {
	f := v.flag
	if len(v.Enums) > 0 {
		found := false
		for _, v := range v.Enums {
			if value == v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Enums: %v", f.logicName, value, v.Enums)
		}
	}
	if len(v.Ranges) > 0 {
		found := false
		for i := 0; i < len(v.Ranges); i++ {
			min := v.Ranges[i]
			max := min
			if i++; i < len(v.Ranges) {
				max = v.Ranges[i]
			}
			if value >= min && value < max {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Ranges: %v", f.logicName, value, v.Enums)
		}
	}
	return nil
}

var _ Value = (*Uint64SliceValue)(nil) //for interface verification only
