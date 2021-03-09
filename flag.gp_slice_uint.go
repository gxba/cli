///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Tue Mar 09 2021 11:43 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_uint_slice]
//
// Tool [github.com/gxlb/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : v3.1.0
// 
///////////////////////////////////////////////////////////////////

// MIT License
//
// Copyright (c) 2021 Ally Dale <vipally@gamil.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cli

import (
	"flag"
	"fmt"

	"cli/internal/impl"
	"cli/internal/util"
	"encoding/json"
	"strconv"
	"strings"
)

// UintSliceValue wraps []uint to satisfy flag.Value
type UintSliceValue struct {
	slice      []uint
	hasBeenSet bool
}

// NewUintSliceValue makes an *UintSliceValue with default values
func NewUintSliceValue(defaults ...uint) *UintSliceValue {
	return &UintSliceValue{
		slice:      append([]uint{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *UintSliceValue) clone() *UintSliceValue {
	n := &UintSliceValue{
		slice:      append([]uint{}, s.slice...),
		hasBeenSet: s.hasBeenSet,
	}
	return n
}

// AppendValues directly append values to the list of values
func (s *UintSliceValue) AppendValues(values ...uint) {
	s.setValues(false, values)
}

// SetValues directly overite values to the list of values
func (s *UintSliceValue) SetValues(values ...uint) {
	s.setValues(true, values)
}

func (s *UintSliceValue) setValues(overwrite bool, values []uint) {
	if !s.hasBeenSet || overwrite {
		s.Reset()
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, values...)
}

// Set parses the value and appends to the list of values
func (s *UintSliceValue) Set(value string) error {

	if strings.HasPrefix(value, impl.SerializedPrefix) {
		// Deserializing assumes overwrite
		_ = json.Unmarshal([]byte(strings.Replace(value, impl.SerializedPrefix, "", 1)), &s.slice)
		s.hasBeenSet = true
		return nil
	}

	//accept multi values for slice flags
	for _, val := range impl.FlagSplitMultiValues(value) {
		value := strings.TrimSpace(val)
		tmp, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return err
		}

		if !s.hasBeenSet {
			s.slice = []uint{}
			s.hasBeenSet = true
		}

		s.slice = append(s.slice, uint(tmp))
	}

	return nil
}

// Reset clean the last parsed values of this slice
func (s *UintSliceValue) Reset() {
	if s.slice == nil {
		s.slice = []uint{}
	} else {
		s.slice = s.slice[:0]
	}
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *UintSliceValue) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows UintSliceSlice to fulfill Serializer
func (s *UintSliceValue) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", impl.SerializedPrefix, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *UintSliceValue) Value() []uint {
	return s.slice
}

// Get returns the slice set by this flag
func (s *UintSliceValue) Get() interface{} {
	return *s
}

// UintSliceFlag define a value of type *UintSliceValue
type UintSliceFlag struct {
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
	Target      *UintSliceValue // Target value pointer outside
	Default     *UintSliceValue // Default value
	DefaultText string          // Default value display in help info
	Enums       []uint          // Enumeration of valid values
	Ranges      []uint          // {[min,max),[min,max),...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *UintSliceValue // target value pointer(maybe new(UintSliceValue) if Target not set)
	info   impl.FlagInfo   // parsed info of this flag
}

// init verify and init the flag info
func (f *UintSliceFlag) init(namegen *util.NameGenenerator) error {
	f.info.Flag = f
	f.info.EnvVars = append([]string{}, f.EnvVars...)
	f.info.Usage = f.Usage
	f.info.DefaultText = f.DefaultText
	f.info.Required = f.Required
	f.info.Hidden = f.Hidden
	f.info.FilePath = f.FilePath
	f.info.HasBeenSet = false
	f.info.Name = namegen.GetOrGenName(f.Name)
	f.info.NonameFlag = f.info.Name != f.Name
	f.info.LogicName = impl.FlagLogicName(f.Name, f.LogicName)
	f.info.ValueName = impl.FlagValueName(f.LogicName)
	impl.MergeNames(f.info.Name, f.Aliases, &f.info.Names) //use v.info.Name to enable auto-generated name

	//make the target pointer
	if f.Target != nil {
		f.target = f.Target
	} else {
		f.target = NewUintSliceValue()
	}

	if f.Name == "" && f.LogicName == "" { // Name & LogicName cannot both missing
		return fmt.Errorf("flag missing both Name & LogicName: %v", f)
	}
	if f.Name == "" && len(f.Aliases) > 0 { // Noname ones must without Aliases
		return fmt.Errorf("flag %s missing name, but has Aliases %v", f.info.LogicName, f.Aliases)
	}
	maxSliceLen := impl.MaxSliceLen
	if l := len(f.Enums); l > 0 { // Enums length check
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Enums too long: %d/%d", f.info.LogicName, l, maxSliceLen)
		}

		if l > 1 {
			var filter = make(map[uint]struct{})
			for _, v := range f.Enums {
				if i, ok := filter[v]; !ok {
					filter[v] = struct{}{}
				} else {
					return fmt.Errorf("flag %s.Enums error: duplicate %v at %d", f.info.LogicName, v, i)
				}
			}
		}
	}
	if l := len(f.Ranges); l > 0 { // Ranges length check and [min,max) pair check
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Ranges too long: %d/%d", f.info.LogicName, l, maxSliceLen)
		}
		if l%2 != 0 {
			return fmt.Errorf("flag %s.Ranges doesn't match [min,max) pairs: %d", f.info.LogicName, l)
		}

		for i := 0; i < l; i += 2 {
			min, max := f.Ranges[i], f.Ranges[i+1]
			if valid := min <= max; !valid {
				return fmt.Errorf("flag %s.Ranges doesn't match [min,max): (%v,%v)", f.info.LogicName, min, max)
			}
			for j := 0; j < i; j += 2 { //check range overlapping
				m, n := f.Ranges[j], f.Ranges[j+1]
				if m >= min && m < max || n >= min && n < max {
					return fmt.Errorf("flag %s.Ranges %d~[%v,%v) overlapping %d~[%v,%v) ", f.info.LogicName, i, min, max, j, m, n)
				}
			}
		}
	}
	if err := f.validateValues(f.Default); err != nil { // verify default values
		return fmt.Errorf("default value invalid: %s", err.Error())
	}
	if err := util.FiltNames(f.info.Names); err != nil { // verify name duplicate
		return fmt.Errorf("flag %s.Names error: %s", f.info.LogicName, err.Error())
	}
	if err := util.FiltNames(f.info.EnvVars); err != nil { // verify EnvVars duplicate
		return fmt.Errorf("flag %s.EnvVars error: %s", f.info.LogicName, err.Error())
	}
	return nil
}

// IsSet check if value was set
func (f *UintSliceFlag) IsSet() bool {
	return f.target.hasBeenSet
}

//GetLogicName returns the logic name of the falg
func (f *UintSliceFlag) GetLogicName() string {
	return f.info.LogicName
}

//GetValueName returns the value name of the falg
func (f *UintSliceFlag) GetValueName() string {
	return f.info.ValueName
}

// Names returns the names of the flag
func (f *UintSliceFlag) Names() []string {
	return f.info.Names
}

// IsRequired returns whether or not the flag is required
func (f *UintSliceFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *UintSliceFlag) TakesValue() bool {
	return false
}

// GetUsage returns the usage string for the flag
func (f *UintSliceFlag) GetUsage() string {
	return f.info.Usage
}

// GetValue returns the flags value as string representation.
func (f *UintSliceFlag) GetValue() string {
	return ""
}

// Apply coordinate the value to flagset
func (f *UintSliceFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (f *UintSliceFlag) String() string {
	return ""
}

// ValidateValues verify if all values are valid
func (f *UintSliceFlag) ValidateValues() error {
	return f.validateValues(f.target)
}

// Info returns parsed info of this flag, the returned object must READ-ONLY
func (v *UintSliceFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (f *UintSliceFlag) Reset() {
	f.target.Reset()
	f.info.HasBeenSet = false
}

// for default value verify
func (f *UintSliceFlag) validateValues(values *UintSliceValue) error {
	for _, val := range values.slice {
		if err := f.validValue(val); err != nil {
			return err
		}
	}
	return nil
}

// check if value if valid for this flag
func (f *UintSliceFlag) validValue(value uint) error {
	if len(f.Enums) > 0 {
		found := false
		for _, v := range f.Enums {
			if value == v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Enums: %v", f.info.LogicName, value, f.Enums)
		}
	}
	if len(f.Ranges) > 0 {
		found := false
		for i := 0; i < len(f.Ranges); i += 2 {
			min, max := f.Ranges[i], f.Ranges[i+1]
			if value >= min && value < max {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Ranges: %v", f.info.LogicName, value, f.Enums)
		}
	}
	return nil
}

// UintSlice looks up the value of a local UintSliceFlag
func (c *Context) UintSlice(name string) []uint {
	if fs := lookupFlagSet(name, c); fs != nil {
		return lookupUintSlice(name, fs)
	}
	return nil
}

func lookupUintSlice(name string, set *flag.FlagSet) []uint {
	f := set.Lookup(name)
	if f != nil {
		if slice, ok := f.Value.(*UintSliceValue); ok {
			return slice.Value()
		}
	}
	return nil
}

var _ impl.Flag = (*UintSliceFlag)(nil) //for interface verification only
var _ = (*strconv.NumError)(nil)        //avoid compile error