///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Mon Mar 08 2021 14:12 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_bool]
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
)

// BoolFlag define a value of type bool
type BoolFlag struct {
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
	Target      *bool  // Target value pointer outside
	Default     bool   // Default value
	DefaultText string // Default value display in help info

	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *bool         // target value pointer(maybe new(bool) if Target not set)
	info   impl.FlagInfo // parsed info of this flag
}

// init verify and init the flag info
func (f *BoolFlag) init(namegen *util.NameGenenerator) error {
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
		f.target = new(bool)
	}

	if f.Name == "" && f.LogicName == "" { // Name & LogicName cannot both missing
		return fmt.Errorf("flag missing both Name & LogicName: %v", f)
	}
	if f.Name == "" && len(f.Aliases) > 0 { // Noname ones must without Aliases
		return fmt.Errorf("flag %s missing name, but has Aliases %v", f.info.LogicName, f.Aliases)
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
func (f *BoolFlag) IsSet() bool {
	return f.info.HasBeenSet
}

//GetLogicName returns the logic name of the falg
func (f *BoolFlag) GetLogicName() string {
	return f.info.LogicName
}

//GetValueName returns the value name of the falg
func (f *BoolFlag) GetValueName() string {
	return f.info.ValueName
}

// Names returns the names of the flag
func (f *BoolFlag) Names() []string {
	return f.info.Names
}

// IsRequired returns whether or not the flag is required
func (f *BoolFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *BoolFlag) TakesValue() bool {
	return false
}

// GetUsage returns the usage string for the flag
func (f *BoolFlag) GetUsage() string {
	return f.info.Usage
}

// GetValue returns the flags value as string representation.
func (f *BoolFlag) GetValue() string {
	return ""
}

// Apply coordinate the value to flagset
func (f *BoolFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (f *BoolFlag) String() string {
	return ""
}

// ValidateValues verify if all values are valid
func (f *BoolFlag) ValidateValues() error {
	return f.validateValues(*f.target)
}

// Info returns parsed info of this flag, the returned object must READ-ONLY
func (v *BoolFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (f *BoolFlag) Reset() {
	*f.target = false
	f.info.HasBeenSet = false
}

// for default value verify
func (f *BoolFlag) validateValues(values bool) error {
	return f.validValue(values)
}

// check if value if valid for this flag
func (f *BoolFlag) validValue(value bool) error {
	return nil
}

// Bool looks up the value of a local BoolFlag
func (c *Context) Bool(name string) bool {
	if fs := lookupFlagSet(name, c); fs != nil {
		return lookupBool(name, fs)
	}
	return false
}

func lookupBool(name string, set *flag.FlagSet) bool {
	f := set.Lookup(name)
	if f != nil {

	}
	return false
}

var _ impl.Flag = (*BoolFlag)(nil) //for interface verification only
