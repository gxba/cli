///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Tue Mar 02 2021 22:15 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_int64_slice]
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

// Int64SliceSlice wraps []int64 to satisfy flag.Value
type Int64Slice struct {
	slice      []int64
	hasBeenSet bool
}

// NewInt64Slice makes an *NewInt64Slice with default values
func NewInt64Slice(defaults ...int64) *Int64Slice {
	return &Int64Slice{
		slice:      append([]int64{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *Int64Slice) clone() *Int64Slice {
	n := &Int64Slice{
		slice:      append([]int64{}, s.slice...),
		hasBeenSet: s.hasBeenSet,
	}
	return n
}

// AppendValues directly append values to the list of values
func (s *Int64Slice) AppendValues(values ...int64) {
	s.setValues(false, values)
}

// SetValues directly overite values to the list of values
func (s *Int64Slice) SetValues(values ...int64) {
	s.setValues(true, values)
}

func (s *Int64Slice) setValues(overwrite bool, values []int64) {
	if !s.hasBeenSet || overwrite {
		s.Reset()
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, values...)
}

// Set parses the value and appends to the list of values
func (s *Int64Slice) Set(value string) error {

	if strings.HasPrefix(value, impl.SerializedPrefix) {
		// Deserializing assumes overwrite
		_ = json.Unmarshal([]byte(strings.Replace(value, impl.SerializedPrefix, "", 1)), &s.slice)
		s.hasBeenSet = true
		return nil
	}

	//accept multi values for slice flags
	for _, val := range impl.FlagSplitMultiValues(value) {
		value := strings.TrimSpace(val)
		tmp, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return err
		}

		if !s.hasBeenSet {
			s.slice = []int64{}
			s.hasBeenSet = true
		}

		s.slice = append(s.slice, int64(tmp))
	}

	return nil
}

// Reset clean the last parsed values of this slice
func (s *Int64Slice) Reset() {
	if s.slice == nil {
		s.slice = []int64{}
	} else {
		s.slice = s.slice[:0]
	}
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *Int64Slice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows Int64SliceSlice to fulfill Serializer
func (s *Int64Slice) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", impl.SerializedPrefix, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *Int64Slice) Value() []int64 {
	return s.slice
}

// Get returns the slice set by this flag
func (s *Int64Slice) Get() interface{} {
	return *s
}

// Int64SliceFlag define a value of type *Int64Slice
type Int64SliceFlag struct {
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
	Target      *Int64Slice // Target value pointer outside
	Default     *Int64Slice // Default value
	DefaultText string      // Default value display in help info
	Enums       []int64     // Enumeration of valid values
	Ranges      []int64     // {[min,max),[min,max),...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *Int64Slice   // target value pointer(maybe new(*Int64Slice) if Target not set)
	info   impl.FlagInfo // parsed info of this flag
}

// init verify and init the flag info
func (f *Int64SliceFlag) init(namegen *util.NameGenenerator) error {
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
		f.target = NewInt64Slice()
	}

	maxSliceLen := impl.MaxSliceLen
	if f.Name == "" && f.LogicName == "" { // Name & LogicName cannot both missing
		return fmt.Errorf("flag missing both Name & LogicName: %v", f)
	}
	if f.Name == "" && len(f.Aliases) > 0 { // Noname ones must without Aliases
		return fmt.Errorf("flag %s missing name, but has Aliases %v", f.info.LogicName, f.Aliases)
	}
	if l := len(f.Enums); l > 0 { // Enums length check
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Enums too long: %d/%d", f.info.LogicName, l, maxSliceLen)
		}

		if l > 1 {
			var filter = make(map[int64]struct{})
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
func (f *Int64SliceFlag) IsSet() bool {
	return f.target.hasBeenSet
}

//GetLogicName returns the logic name of the falg
func (f *Int64SliceFlag) GetLogicName() string {
	return f.info.LogicName
}

//GetValueName returns the value name of the falg
func (f *Int64SliceFlag) GetValueName() string {
	return f.info.ValueName
}

// Names returns the names of the flag
func (f *Int64SliceFlag) Names() []string {
	return f.info.Names
}

// IsRequired returns whether or not the flag is required
func (f *Int64SliceFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *Int64SliceFlag) TakesValue() bool {
	return false
}

// GetUsage returns the usage string for the flag
func (f *Int64SliceFlag) GetUsage() string {
	return f.info.Usage
}

// GetValue returns the flags value as string representation.
func (f *Int64SliceFlag) GetValue() string {
	return ""
}

// Apply coordinate the value to flagset
func (f *Int64SliceFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (f *Int64SliceFlag) String() string {
	return ""
}

// ValidateValues verify if all values are valid
func (f *Int64SliceFlag) ValidateValues() error {
	return f.validateValues(f.target)
}

// Info returns parsed info of this flag, the returned object must READ-ONLY
func (v *Int64SliceFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (f *Int64SliceFlag) Reset() {
	f.target.Reset()
	f.info.HasBeenSet = false
}

// for default value verify
func (f *Int64SliceFlag) validateValues(values *Int64Slice) error {
	for _, val := range values.slice {
		if err := f.validValue(val); err != nil {
			return err
		}
	}
	return nil
}

// check if value if valid for this flag
func (f *Int64SliceFlag) validValue(value int64) error {
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

// // Int64Slice looks up the value of a local Int64SliceFlag
// func (c *Context) Int64Slice(name string) []int64 {
// 	if fs := lookupFlagSet(name, c); fs != nil {
// 		return lookupInt64Slice(name, fs)
// 	}
// 	return nil
// }

// func lookupInt64Slice(name string, set *flag.FlagSet) []int64 {
// 	f := set.Lookup(name)
// 	if f != nil {
// 		if slice, ok := f.Value.(*Int64Slice); ok {
// 			return slice.Value()
// 		}
// 	}
// 	return nil
// }

var _ impl.Flag = (*Int64SliceFlag)(nil) //for interface verification only
var _ = (*strconv.NumError)(nil)         //avoid compile error
