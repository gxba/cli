//#GOGP_IGNORE_BEGIN
///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Fri Mar 19 2021 16:15 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp.go]
//   [github.com/gxlb/cli/internal/gp/flag.gpg] [GOGP_REVERSE_flag]
//
// Tool [github.com/gxlb/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : v4.0.0
// 
///////////////////////////////////////////////////////////////////
//#GOGP_IGNORE_END

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

<PACKAGE>

import (
	"flag"
	"fmt"

	"cli/internal/impl"
	"cli/internal/util"

	//#GOGP_IFDEF <IF_IS_SLICE_TYPE>
	"encoding/json"
	"strconv"
	"strings"
	//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>
)

//#GOGP_REQUIRE(github.com/gxlb/gogp/lib/fakedef,_)

//#GOGP_IFDEF <IF_IS_SLICE_TYPE>||<IF_IS_TIMESTAMP>
//    #GOGP_MAP(GOGP_IfIsPointerFlagValue, yes)
//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>||<IF_IS_TIMESTAMP>

//#GOGP_SWITCH
// #GOGP_CASE <IF_IS_SLICE_TYPE>
//    #GOGP_REPLACE(GOGP_ReplaceDefaultValueType, *<GLOBAL_NAME_PREFIX>Value)
// #GOGP_ENDCASE
// #GOGP_CASE <IF_IS_TIMESTAMP>
//    #GOGP_REPLACE(GOGP_ReplaceDefaultValueType, *impl.TimestampValue)
// #GOGP_ENDCASE
// #GOGP_DEFAULT
//   #GOGP_REPLACE(GOGP_ReplaceDefaultValueType, <VALUE_TYPE>)
// #GOGP_ENDCASE
//#GOGP_ENDSWITCH

//#GOGP_IFDEF GOGP_IfIsPointerFlagValue
//    #GOGP_REPLACE(GOGP_ReplaceDefaultValueType, <VALUE_TYPE>)
//    #GOGP_REPLACE(GOGP_RepaceZeroValue, nil)
//#GOGP_ENDIF //GOGP_IfIsPointerFlagValue

//#GOGP_IFDEF <IF_IS_TIMESTAMP>

//#GOGP_REPLACE(<REP_SINGLE_VALUE>, values)
//#GOGP_REPLACE(*<REP_ELEM_TYPE>, <VALUE_TYPE>)
//#GOGP_REPLACE(<REP_SLICE_VALUE>, f.target)

//#GOGP_ENDIF //<IF_IS_TIMESTAMP>

//#GOGP_IFDEF <IF_IS_SLICE_TYPE>

// <GLOBAL_NAME_PREFIX>Value wraps []<VALUE_TYPE> to satisfy flag.Value
type <GLOBAL_NAME_PREFIX>Value struct {
	slice      []<VALUE_TYPE>
	hasBeenSet bool
}

// New<GLOBAL_NAME_PREFIX>Value makes an *<GLOBAL_NAME_PREFIX>Value with default values
func New<GLOBAL_NAME_PREFIX>Value(defaults ...<VALUE_TYPE>) *<GLOBAL_NAME_PREFIX>Value {
	return &<GLOBAL_NAME_PREFIX>Value{
		slice:      append([]<VALUE_TYPE>{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *<GLOBAL_NAME_PREFIX>Value) clone() *<GLOBAL_NAME_PREFIX>Value {
	n := &<GLOBAL_NAME_PREFIX>Value{
		slice:      append([]<VALUE_TYPE>{}, s.slice...),
		hasBeenSet: s.hasBeenSet,
	}
	return n
}

// AppendValues directly append values to the list of values
func (s *<GLOBAL_NAME_PREFIX>Value) AppendValues(values ...<VALUE_TYPE>) {
	s.setValues(false, values)
}

// SetValues directly overite values to the list of values
func (s *<GLOBAL_NAME_PREFIX>Value) SetValues(values ...<VALUE_TYPE>) {
	s.setValues(true, values)
}

func (s *<GLOBAL_NAME_PREFIX>Value) setValues(overwrite bool, values []<VALUE_TYPE>) {
	if !s.hasBeenSet || overwrite {
		s.Reset()
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, values...)
}

// Set parses the value and appends to the list of values
func (s *<GLOBAL_NAME_PREFIX>Value) Set(value string) error {

	if strings.HasPrefix(value, impl.SerializedPrefix) {
		// Deserializing assumes overwrite
		_ = json.Unmarshal([]byte(strings.Replace(value, impl.SerializedPrefix, "", 1)), &s.slice)
		s.hasBeenSet = true
		return nil
	}

	//accept multi values for slice flags
	for _, val := range impl.FlagSplitMultiValues(value) {
		value := strings.TrimSpace(val)
		tmp, err := <REP_PARSE_STRING>(value)
		if err != nil {
			return err
		}

		if !s.hasBeenSet {
			s.slice = []<VALUE_TYPE>{}
			s.hasBeenSet = true
		}

		s.slice = append(s.slice, <VALUE_TYPE>(tmp))
	}

	return nil
}

// Reset clean the last parsed values of this slice
func (s *<GLOBAL_NAME_PREFIX>Value) Reset() {
	if s.slice == nil {
		s.slice = []<VALUE_TYPE>{}
	} else {
		s.slice = s.slice[:0]
	}
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *<GLOBAL_NAME_PREFIX>Value) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows <GLOBAL_NAME_PREFIX>Slice to fulfill Serializer
func (s *<GLOBAL_NAME_PREFIX>Value) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", impl.SerializedPrefix, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *<GLOBAL_NAME_PREFIX>Value) Value() []<VALUE_TYPE> {
	return s.slice
}

// Get returns the slice set by this flag
func (s *<GLOBAL_NAME_PREFIX>Value) Get() interface{} {
	return *s
}

//#GOGP_REPLACE(*<REP_ELEM_TYPE>, *<GLOBAL_NAME_PREFIX>Value)
//#GOGP_REPLACE(<REP_ELEM_TYPE>, *<GLOBAL_NAME_PREFIX>Value)
//#GOGP_REPLACE(<REP_PARSE_STRING>(value), <PARSE_STRING>)
//#GOGP_REPLACE(<REP_SLICE_VALUE>, f.target)
//#GOGP_REPLACE(<REP_RAW_ELEM_TYPE>, <GLOBAL_NAME_PREFIX>Value)
//#GOGP_REPLACE(<REP_VALUE_TYPE>, []<VALUE_TYPE>)
//#GOGP_REPLACE(<REP_SINGLE_VALUE>, values)

//#GOGP_ELSE // <IF_IS_SLICE_TYPE>

//#GOGP_REPLACE(<REP_SINGLE_VALUE>, values)
//#GOGP_REPLACE(<REP_ELEM_TYPE>, <VALUE_TYPE>)
//#GOGP_REPLACE(<REP_RAW_ELEM_TYPE>, <VALUE_TYPE>)
//#GOGP_REPLACE(<REP_VALUE_TYPE>, <VALUE_TYPE>)

//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>

// <GLOBAL_NAME_PREFIX>Flag define a value of type <REP_ELEM_TYPE>
type <GLOBAL_NAME_PREFIX>Flag struct {
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
	Target      GOGP_ReplaceTargetValueType  // Target value pointer outside
	Default     GOGP_ReplaceDefaultValueType // Default value
	DefaultText string                       // Default value display in help info
	//#GOGP_IFDEF <IF_IS_TIMESTAMP>
	Layout string // Layout of the time format, default **2006-01-02T15:04:05 MST**
	//#GOGP_ENDIF //<IF_IS_TIMESTAMP>

	//#GOGP_IFDEF !<IF_NO_COMPARE>
	Enums  []<VALUE_TYPE> // Enumeration of valid values
	Ranges []<VALUE_TYPE> // {[min,max),[min,max),...} ranges of valid values
	//#GOGP_ENDIF //<IF_NO_COMPARE>

	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target GOGP_ReplaceTargetValueType // target value pointer(maybe new(<REP_RAW_ELEM_TYPE>) if Target not set)
	info   impl.FlagInfo               // parsed info of this flag
}

// init verify and init the flag info
func (f *<GLOBAL_NAME_PREFIX>Flag) init(namegen *util.NameGenenerator) error {
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
		//#GOGP_IFDEF <IF_IS_SLICE_TYPE>
		f.target = New<GLOBAL_NAME_PREFIX>Value()
		//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>

		//#GOGP_IFDEF <IF_IS_TIMESTAMP>
		//#GOGP_COMMENT f.target = impl.NewEmptyTimestampValue(f.Layout)
		//#GOGP_ELSE
		f.target = new(<REP_RAW_ELEM_TYPE>)
		//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>
	}

	if f.Name == "" && f.LogicName == "" { // Name & LogicName cannot both missing
		return fmt.Errorf("flag missing both Name & LogicName: %v", f)
	}
	if f.Name == "" && len(f.Aliases) > 0 { // Noname ones must without Aliases
		return fmt.Errorf("flag %s missing name, but has Aliases %v", f.info.LogicName, f.Aliases)
	}
	//#GOGP_IFDEF !<IF_NO_COMPARE>
	maxSliceLen := impl.MaxSliceLen
	if l := len(f.Enums); l > 0 { // Enums length check
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Enums too long: %d/%d", f.info.LogicName, l, maxSliceLen)
		}

		if l > 1 {
			var filter = make(map[<VALUE_TYPE>]struct{})
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
	//#GOGP_ENDIF //<IF_NO_COMPARE>
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
func (f *<GLOBAL_NAME_PREFIX>Flag) IsSet() bool {
	//#GOGP_IFDEF <IF_IS_SLICE_TYPE>
	return f.target.hasBeenSet
	//#GOGP_ELSE
	return f.info.HasBeenSet
	//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>
}

//GetLogicName returns the logic name of the falg
func (f *<GLOBAL_NAME_PREFIX>Flag) GetLogicName() string {
	return f.info.LogicName
}

//GetValueName returns the value name of the falg
func (f *<GLOBAL_NAME_PREFIX>Flag) GetValueName() string {
	return f.info.ValueName
}

// Names returns the names of the flag
func (f *<GLOBAL_NAME_PREFIX>Flag) Names() []string {
	return f.info.Names
}

// IsRequired returns whether or not the flag is required
func (f *<GLOBAL_NAME_PREFIX>Flag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *<GLOBAL_NAME_PREFIX>Flag) TakesValue() bool {
	return false
}

// GetUsage returns the usage string for the flag
func (f *<GLOBAL_NAME_PREFIX>Flag) GetUsage() string {
	return f.info.Usage
}

// GetValue returns the flags value as string representation.
func (f *<GLOBAL_NAME_PREFIX>Flag) GetValue() string {
	return ""
}

// Apply coordinate the value to flagset
func (f *<GLOBAL_NAME_PREFIX>Flag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (f *<GLOBAL_NAME_PREFIX>Flag) String() string {
	return ""
}

// ValidateValues verify if all values are valid
func (f *<GLOBAL_NAME_PREFIX>Flag) ValidateValues() error {
	//#GOGP_IFDEF <IF_IS_SLICE_TYPE>||<IF_IS_TIMESTAMP>
	return f.validateValues(<REP_SLICE_VALUE>)
	//#GOGP_ELSE
	return f.validateValues(*f.target)
	//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>
}

// Info returns parsed info of this flag, the returned object must READ-ONLY
func (v *<GLOBAL_NAME_PREFIX>Flag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (f *<GLOBAL_NAME_PREFIX>Flag) Reset() {
	//#GOGP_IFDEF <IF_IS_SLICE_TYPE>||<IF_IS_TIMESTAMP>
	f.target.Reset()
	//#GOGP_ELSE
	//#GOGP_COMMENT *f.target = <ZREO_VALUE>
	//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>
	f.info.HasBeenSet = false
}

// for default value verify
func (f *<GLOBAL_NAME_PREFIX>Flag) validateValues(values <REP_ELEM_TYPE>) error {
	//#GOGP_IFDEF <IF_IS_SLICE_TYPE>
	for _, val := range values.slice {
		if err := f.validValue(val); err != nil {
			return err
		}
	}
	return nil
	//#GOGP_ELSE
	return f.validValue(<REP_SINGLE_VALUE>)
	//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>
}

// check if value if valid for this flag
func (f *<GLOBAL_NAME_PREFIX>Flag) validValue(value <VALUE_TYPE>) error {
	//#GOGP_IFDEF <IF_NO_COMPARE>
	return nil
	//#GOGP_ELSE
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
	//#GOGP_ENDIF //<IF_NO_COMPARE>
}

// <GLOBAL_NAME_PREFIX> looks up the value of a local <GLOBAL_NAME_PREFIX>Flag
func (c *Context) <GLOBAL_NAME_PREFIX>(name string) <REP_VALUE_TYPE> {
	if fs := lookupFlagSet(name, c); fs != nil {
		return lookup<GLOBAL_NAME_PREFIX>(name, fs)
	}
	return <ZREO_VALUE>
}

func lookup<GLOBAL_NAME_PREFIX>(name string, set *flag.FlagSet) <REP_VALUE_TYPE> {
	f := set.Lookup(name)
	if f != nil {
		//#GOGP_IFDEF <IF_IS_SLICE_TYPE>
		if slice, ok := f.Value.(*<GLOBAL_NAME_PREFIX>Value); ok {
			return slice.Value()
		}
		//#GOGP_ELSE //<IF_IS_SLICE_TYPE>
		//TODO:
		//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>
	}
	return <ZREO_VALUE>
}

var _ impl.Flag = (*<GLOBAL_NAME_PREFIX>Flag)(nil) //for interface verification only
//#GOGP_IFDEF <IF_IS_SLICE_TYPE>
var _ = (*strconv.NumError)(nil) //avoid compile error
//#GOGP_ENDIF //<IF_IS_SLICE_TYPE>

