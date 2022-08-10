//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package filter

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchConfig) DeepCopyInto(out *MatchConfig) {
	*out = *in
	if in.And != nil {
		in, out := &in.And, &out.And
		*out = make([]MatchExpr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(MatchExpr)
		(*in).DeepCopyInto(*out)
	}
	if in.Regexp != nil {
		in, out := &in.Regexp, &out.Regexp
		*out = new(RegexpMatchExpr)
		(*in).DeepCopyInto(*out)
	}
	if in.Or != nil {
		in, out := &in.Or, &out.Or
		*out = make([]MatchExpr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchConfig.
func (in *MatchConfig) DeepCopy() *MatchConfig {
	if in == nil {
		return nil
	}
	out := new(MatchConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchExpr) DeepCopyInto(out *MatchExpr) {
	*out = *in
	if in.And != nil {
		in, out := &in.And, &out.And
		*out = make([]MatchExpr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(MatchExpr)
		(*in).DeepCopyInto(*out)
	}
	if in.Regexp != nil {
		in, out := &in.Regexp, &out.Regexp
		*out = new(RegexpMatchExpr)
		(*in).DeepCopyInto(*out)
	}
	if in.Or != nil {
		in, out := &in.Or, &out.Or
		*out = make([]MatchExpr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpr.
func (in *MatchExpr) DeepCopy() *MatchExpr {
	if in == nil {
		return nil
	}
	out := new(MatchExpr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParserConfig) DeepCopyInto(out *ParserConfig) {
	*out = *in
	if in.Regexp != nil {
		in, out := &in.Regexp, &out.Regexp
		*out = new(RegexpParser)
		(*in).DeepCopyInto(*out)
	}
	if in.SyslogParser != nil {
		in, out := &in.SyslogParser, &out.SyslogParser
		*out = new(SyslogParser)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParserConfig.
func (in *ParserConfig) DeepCopy() *ParserConfig {
	if in == nil {
		return nil
	}
	out := new(ParserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegexpMatchExpr) DeepCopyInto(out *RegexpMatchExpr) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexpMatchExpr.
func (in *RegexpMatchExpr) DeepCopy() *RegexpMatchExpr {
	if in == nil {
		return nil
	}
	out := new(RegexpMatchExpr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegexpParser) DeepCopyInto(out *RegexpParser) {
	*out = *in
	if in.Patterns != nil {
		in, out := &in.Patterns, &out.Patterns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexpParser.
func (in *RegexpParser) DeepCopy() *RegexpParser {
	if in == nil {
		return nil
	}
	out := new(RegexpParser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenameConfig) DeepCopyInto(out *RenameConfig) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(MatchExpr)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenameConfig.
func (in *RenameConfig) DeepCopy() *RenameConfig {
	if in == nil {
		return nil
	}
	out := new(RenameConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RewriteConfig) DeepCopyInto(out *RewriteConfig) {
	*out = *in
	if in.Rename != nil {
		in, out := &in.Rename, &out.Rename
		*out = new(RenameConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = new(SetConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Substitute != nil {
		in, out := &in.Substitute, &out.Substitute
		*out = new(SubstituteConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Unset != nil {
		in, out := &in.Unset, &out.Unset
		*out = new(UnsetConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RewriteConfig.
func (in *RewriteConfig) DeepCopy() *RewriteConfig {
	if in == nil {
		return nil
	}
	out := new(RewriteConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetConfig) DeepCopyInto(out *SetConfig) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(MatchExpr)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetConfig.
func (in *SetConfig) DeepCopy() *SetConfig {
	if in == nil {
		return nil
	}
	out := new(SetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubstituteConfig) DeepCopyInto(out *SubstituteConfig) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(MatchExpr)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubstituteConfig.
func (in *SubstituteConfig) DeepCopy() *SubstituteConfig {
	if in == nil {
		return nil
	}
	out := new(SubstituteConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyslogParser) DeepCopyInto(out *SyslogParser) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyslogParser.
func (in *SyslogParser) DeepCopy() *SyslogParser {
	if in == nil {
		return nil
	}
	out := new(SyslogParser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnsetConfig) DeepCopyInto(out *UnsetConfig) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(MatchExpr)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnsetConfig.
func (in *UnsetConfig) DeepCopy() *UnsetConfig {
	if in == nil {
		return nil
	}
	out := new(UnsetConfig)
	in.DeepCopyInto(out)
	return out
}
