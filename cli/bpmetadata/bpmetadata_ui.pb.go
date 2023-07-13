// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: bpmetadata_ui.proto

package bpmetadata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BlueprintUIInput is the structure for holding Input and Input Section (i.e. groups) specific metadata.
type BlueprintUIInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// variables is a map defining all inputs on the UI.
	// Gen: partial
	Variables map[string]*DisplayVariable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"variables,omitempty"` // @gotags: json:"variables,omitempty" yaml:"variables,omitempty"
	// Sections is a generic structure for grouping inputs together.
	// Gen: manually-authored
	Sections []*DisplaySection `protobuf:"bytes,2,rep,name=sections,proto3" json:"sections,omitempty" yaml:"sections,omitempty"` // @gotags: json:"sections,omitempty" yaml:"sections,omitempty"
}

func (x *BlueprintUIInput) Reset() {
	*x = BlueprintUIInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpmetadata_ui_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlueprintUIInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlueprintUIInput) ProtoMessage() {}

func (x *BlueprintUIInput) ProtoReflect() protoreflect.Message {
	mi := &file_bpmetadata_ui_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlueprintUIInput.ProtoReflect.Descriptor instead.
func (*BlueprintUIInput) Descriptor() ([]byte, []int) {
	return file_bpmetadata_ui_proto_rawDescGZIP(), []int{0}
}

func (x *BlueprintUIInput) GetVariables() map[string]*DisplayVariable {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *BlueprintUIInput) GetSections() []*DisplaySection {
	if x != nil {
		return x.Sections
	}
	return nil
}

// Additional display specific metadata pertaining to a particular
// input variable.
type DisplayVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The variable name from the corresponding standard metadata file.
	// Gen: auto-generated - the Terraform variable name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"` // @gotags: json:"name" yaml:"name"
	// Visible title for the variable on the UI. If not present,
	// Name will be used for the Title.
	// Gen: auto-generated - the Terraform variable converted to title case e.g.
	// variable "bucket_admins" will convert to "Bucket Admins" as the title.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" yaml:"title"` // @gotags: json:"title" yaml:"title"
	// A flag to hide or show the variable on the UI.
	// Gen: manually-authored
	Invisible bool `protobuf:"varint,3,opt,name=invisible,proto3" json:"invisible,omitempty" yaml:"invisible,omitempty"` // @gotags: json:"invisible,omitempty" yaml:"invisible,omitempty"
	// Variable tooltip.
	// Gen: manually-authored
	Tooltip string `protobuf:"bytes,4,opt,name=tooltip,proto3" json:"tooltip,omitempty" yaml:"tooltip,omitempty"` // @gotags: json:"tooltip,omitempty" yaml:"tooltip,omitempty"
	// Placeholder text (when there is no default).
	// Gen: manually-authored
	Placeholder string `protobuf:"bytes,5,opt,name=placeholder,proto3" json:"placeholder,omitempty" yaml:"placeholder,omitempty"` // @gotags: json:"placeholder,omitempty" yaml:"placeholder,omitempty"
	// Regex based validation rules for the variable.
	// Gen: manually-authored
	RegexValidation string `protobuf:"bytes,6,opt,name=regex_validation,json=regexValidation,proto3" json:"regexValidation,omitempty" yaml:"regexValidation,omitempty"` // @gotags: json:"regexValidation,omitempty" yaml:"regexValidation,omitempty"
	// Minimum no. of inputs for the input variable.
	// Gen: manually-authored
	MinItems int32 `protobuf:"varint,7,opt,name=min_items,json=minItems,proto3" json:"minItems,omitempty" yaml:"minItems,omitempty"` // @gotags: json:"minItems,omitempty" yaml:"minItems,omitempty"
	// Max no. of inputs for the input variable.
	// Gen: manually-authored
	MaxItems int32 `protobuf:"varint,8,opt,name=max_items,json=maxItems,proto3" json:"maxItems,omitempty" yaml:"maxItems,omitempty"` // @gotags: json:"maxItems,omitempty" yaml:"maxItems,omitempty"
	// Minimum length for string values.
	// Gen: manually-authored
	MinLength int32 `protobuf:"varint,9,opt,name=min_length,json=minLength,proto3" json:"minLength,omitempty" yaml:"minLength,omitempty"` // @gotags: json:"minLength,omitempty" yaml:"minLength,omitempty"
	// Max length for string values.
	// Gen: manually-authored
	MaxLength int32 `protobuf:"varint,10,opt,name=max_length,json=maxLength,proto3" json:"maxLength,omitempty" yaml:"maxLength,omitempty"` // @gotags: json:"maxLength,omitempty" yaml:"maxLength,omitempty"
	// Minimum value for numeric types.
	// Gen: manually-authored
	Min float32 `protobuf:"fixed32,11,opt,name=min,proto3" json:"min,omitempty" yaml:"min,omitempty"` // @gotags: json:"min,omitempty" yaml:"min,omitempty"
	// Max value for numeric types.
	// Gen: manually-authored
	Max float32 `protobuf:"fixed32,12,opt,name=max,proto3" json:"max,omitempty" yaml:"max,omitempty"` // @gotags: json:"max,omitempty" yaml:"max,omitempty"
	// The name of a section to which this variable belongs.
	// variables belong to the root section if this field is
	// not set.
	// Gen: manually-authored
	Section string `protobuf:"bytes,13,opt,name=section,proto3" json:"section,omitempty" yaml:"section,omitempty"` // @gotags: json:"section,omitempty" yaml:"section,omitempty"
	// UI extension associated with the input variable.
	// E.g. for rendering a GCE machine type selector:
	//
	// xGoogleProperty:
	//
	//	type: GCE_MACHINE_TYPE
	//	zoneProperty: myZone
	//	gceMachineType:
	//	  minCpu: 2
	//	  minRamGb:
	//
	// Gen: manually-authored
	XGoogleProperty *GooglePropertyExtension `protobuf:"bytes,14,opt,name=x_google_property,json=xGoogleProperty,proto3" json:"xGoogleProperty,omitempty" yaml:"xGoogleProperty,omitempty"` // @gotags: json:"xGoogleProperty,omitempty" yaml:"xGoogleProperty,omitempty"
	// Text describing the validation rules for the property. Typically shown
	// after an invalid input.
	// Optional. UTF-8 text. No markup. At most 128 characters.
	// Gen: manually-authored
	Validation string `protobuf:"bytes,15,opt,name=validation,proto3" json:"validation,omitempty" yaml:"validation,omitempty"` // @gotags: json:"validation,omitempty" yaml:"validation,omitempty"
}

func (x *DisplayVariable) Reset() {
	*x = DisplayVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpmetadata_ui_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplayVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplayVariable) ProtoMessage() {}

func (x *DisplayVariable) ProtoReflect() protoreflect.Message {
	mi := &file_bpmetadata_ui_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplayVariable.ProtoReflect.Descriptor instead.
func (*DisplayVariable) Descriptor() ([]byte, []int) {
	return file_bpmetadata_ui_proto_rawDescGZIP(), []int{1}
}

func (x *DisplayVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DisplayVariable) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DisplayVariable) GetInvisible() bool {
	if x != nil {
		return x.Invisible
	}
	return false
}

func (x *DisplayVariable) GetTooltip() string {
	if x != nil {
		return x.Tooltip
	}
	return ""
}

func (x *DisplayVariable) GetPlaceholder() string {
	if x != nil {
		return x.Placeholder
	}
	return ""
}

func (x *DisplayVariable) GetRegexValidation() string {
	if x != nil {
		return x.RegexValidation
	}
	return ""
}

func (x *DisplayVariable) GetMinItems() int32 {
	if x != nil {
		return x.MinItems
	}
	return 0
}

func (x *DisplayVariable) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *DisplayVariable) GetMinLength() int32 {
	if x != nil {
		return x.MinLength
	}
	return 0
}

func (x *DisplayVariable) GetMaxLength() int32 {
	if x != nil {
		return x.MaxLength
	}
	return 0
}

func (x *DisplayVariable) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *DisplayVariable) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *DisplayVariable) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *DisplayVariable) GetXGoogleProperty() *GooglePropertyExtension {
	if x != nil {
		return x.XGoogleProperty
	}
	return nil
}

func (x *DisplayVariable) GetValidation() string {
	if x != nil {
		return x.Validation
	}
	return ""
}

// A logical group of variables. [Section][]s may also be grouped into
// sub-sections.
type DisplaySection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the section, referenced by DisplayVariable.Section
	// Section names must be unique.
	// Gen: manually-authored
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"` // @gotags: json:"name" yaml:"name"
	// Section title.
	// If not provided, name will be used instead.
	// Gen: manually-authored
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"` // @gotags: json:"title,omitempty" yaml:"title,omitempty"
	// Section tooltip.
	// Gen: manually-authored
	Tooltip string `protobuf:"bytes,3,opt,name=tooltip,proto3" json:"tooltip,omitempty" yaml:"tooltip,omitempty"` // @gotags: json:"tooltip,omitempty" yaml:"tooltip,omitempty"
	// Section subtext.
	// Gen: manually-authored
	Subtext string `protobuf:"bytes,4,opt,name=subtext,proto3" json:"subtext,omitempty" yaml:"subtext,omitempty"` // @gotags: json:"subtext,omitempty" yaml:"subtext,omitempty"
	// The name of the parent section (if parent is not the root section).
	// Gen: manually-authored
	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty" yaml:"parent,omitempty"` // @gotags: json:"parent,omitempty" yaml:"parent,omitempty"
}

func (x *DisplaySection) Reset() {
	*x = DisplaySection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpmetadata_ui_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplaySection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplaySection) ProtoMessage() {}

func (x *DisplaySection) ProtoReflect() protoreflect.Message {
	mi := &file_bpmetadata_ui_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplaySection.ProtoReflect.Descriptor instead.
func (*DisplaySection) Descriptor() ([]byte, []int) {
	return file_bpmetadata_ui_proto_rawDescGZIP(), []int{2}
}

func (x *DisplaySection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DisplaySection) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DisplaySection) GetTooltip() string {
	if x != nil {
		return x.Tooltip
	}
	return ""
}

func (x *DisplaySection) GetSubtext() string {
	if x != nil {
		return x.Subtext
	}
	return ""
}

func (x *DisplaySection) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

type BlueprintUIOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Short message to be displayed while the blueprint is deploying.
	// At most 128 characters.
	// Gen: manually-authored
	OutputMessage string `protobuf:"bytes,1,opt,name=output_message,json=outputMessage,proto3" json:"outputMessage,omitempty" yaml:"outputMessage,omitempty"` // @gotags: json:"outputMessage,omitempty" yaml:"outputMessage,omitempty"
	// List of suggested actions to take.
	// Gen: manually-authored
	SuggestedActions []*UIActionItem `protobuf:"bytes,2,rep,name=suggested_actions,json=suggestedActions,proto3" json:"suggestedActions,omitempty" yaml:"suggestedActions,omitempty"` // @gotags: json:"suggestedActions,omitempty" yaml:"suggestedActions,omitempty"
	// outputs is a map defining a subset of Terraform outputs on the UI
	// that may need additional UI configuration.
	// Gen: manually-authored
	Outputs map[string]*DisplayOutput `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"outputs,omitempty"` // @gotags: json:"outputs,omitempty" yaml:"outputs,omitempty"
}

func (x *BlueprintUIOutput) Reset() {
	*x = BlueprintUIOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpmetadata_ui_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlueprintUIOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlueprintUIOutput) ProtoMessage() {}

func (x *BlueprintUIOutput) ProtoReflect() protoreflect.Message {
	mi := &file_bpmetadata_ui_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlueprintUIOutput.ProtoReflect.Descriptor instead.
func (*BlueprintUIOutput) Descriptor() ([]byte, []int) {
	return file_bpmetadata_ui_proto_rawDescGZIP(), []int{3}
}

func (x *BlueprintUIOutput) GetOutputMessage() string {
	if x != nil {
		return x.OutputMessage
	}
	return ""
}

func (x *BlueprintUIOutput) GetSuggestedActions() []*UIActionItem {
	if x != nil {
		return x.SuggestedActions
	}
	return nil
}

func (x *BlueprintUIOutput) GetOutputs() map[string]*DisplayOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

// An item appearing in a list of required or suggested steps.
type UIActionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Summary heading for the item.
	// Required. Accepts string expressions. At most 64 characters.
	// Gen: manually-authored
	Heading string `protobuf:"bytes,1,opt,name=heading,proto3" json:"heading" yaml:"heading"` // @gotags: json:"heading" yaml:"heading"
	// Longer description of the item.
	// At least one description or snippet is required.
	// Accepts string expressions. HTML <code>&lt;a href&gt;</code>
	// tags only. At most 512 characters.
	// Gen: manually-authored
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description,omitempty"` // @gotags: json:"description,omitempty" yaml:"description,omitempty"
	// Fixed-width formatted code snippet.
	// At least one description or snippet is required.
	// Accepts string expressions. UTF-8 text. At most 512 characters.
	// Gen: manually-authored
	Snippet string `protobuf:"bytes,3,opt,name=snippet,proto3" json:"snippet,omitempty" yaml:"snippet,omitempty"` // @gotags: json:"snippet,omitempty" yaml:"snippet,omitempty"
	// If present, this expression determines whether the item is shown.
	// Should be in the form of a Boolean expression e.g. outputs().hasExternalIP
	// where `externalIP` is the output.
	// Gen: manually-authored
	ShowIf string `protobuf:"bytes,4,opt,name=show_if,json=showIf,proto3" json:"showIf,omitempty" yaml:"showIf,omitempty"` // @gotags: json:"showIf,omitempty" yaml:"showIf,omitempty"
}

func (x *UIActionItem) Reset() {
	*x = UIActionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpmetadata_ui_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UIActionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UIActionItem) ProtoMessage() {}

func (x *UIActionItem) ProtoReflect() protoreflect.Message {
	mi := &file_bpmetadata_ui_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UIActionItem.ProtoReflect.Descriptor instead.
func (*UIActionItem) Descriptor() ([]byte, []int) {
	return file_bpmetadata_ui_proto_rawDescGZIP(), []int{4}
}

func (x *UIActionItem) GetHeading() string {
	if x != nil {
		return x.Heading
	}
	return ""
}

func (x *UIActionItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UIActionItem) GetSnippet() string {
	if x != nil {
		return x.Snippet
	}
	return ""
}

func (x *UIActionItem) GetShowIf() string {
	if x != nil {
		return x.ShowIf
	}
	return ""
}

// Additional display specific metadata pertaining to a particular
// Terraform output.
type DisplayOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// open_in_new_tab defines if the Output action should be opened
	// in a new tab.
	// Gen: manually-authored
	OpenInNewTab bool `protobuf:"varint,1,opt,name=open_in_new_tab,json=openInNewTab,proto3" json:"openInNewTab,omitempty" yaml:"openInNewTab,omitempty"` // @gotags: json:"openInNewTab,omitempty" yaml:"openInNewTab,omitempty"
	// show_in_notification defines if the Output should shown in
	// notification for the deployment.
	// Gen: manually-authored
	ShowInNotification bool `protobuf:"varint,2,opt,name=show_in_notification,json=showInNotification,proto3" json:"showInNotification,omitempty" yaml:"showInNotification,omitempty"` // @gotags: json:"showInNotification,omitempty" yaml:"showInNotification,omitempty"
}

func (x *DisplayOutput) Reset() {
	*x = DisplayOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpmetadata_ui_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplayOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplayOutput) ProtoMessage() {}

func (x *DisplayOutput) ProtoReflect() protoreflect.Message {
	mi := &file_bpmetadata_ui_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplayOutput.ProtoReflect.Descriptor instead.
func (*DisplayOutput) Descriptor() ([]byte, []int) {
	return file_bpmetadata_ui_proto_rawDescGZIP(), []int{5}
}

func (x *DisplayOutput) GetOpenInNewTab() bool {
	if x != nil {
		return x.OpenInNewTab
	}
	return false
}

func (x *DisplayOutput) GetShowInNotification() bool {
	if x != nil {
		return x.ShowInNotification
	}
	return false
}

var File_bpmetadata_ui_proto protoreflect.FileDescriptor

var file_bpmetadata_ui_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x70, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62, 0x70, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x62, 0x70, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x75, 0x69, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac,
	0x02, 0x0a, 0x10, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x55, 0x49, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x5d, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62, 0x70, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x55, 0x49, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x12, 0x4a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62, 0x70, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x6d,
	0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62, 0x70, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfb, 0x03,
	0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6e, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x6e, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x6f,
	0x6c, 0x74, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6f, 0x6c,
	0x74, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69,
	0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d,
	0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x63, 0x0a, 0x11, 0x78, 0x5f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62, 0x70, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x78, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x0e,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x6f, 0x6c,
	0x74, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6f, 0x6c, 0x74,
	0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x22, 0xda, 0x02, 0x0a, 0x11, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x55, 0x49, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x59, 0x0a, 0x11, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x62, 0x70, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x49,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x10, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x58, 0x0a, 0x07,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x62, 0x70, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42,
	0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x55, 0x49, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x1a, 0x69, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x62, 0x70,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x7d, 0x0a, 0x0c, 0x55, 0x49, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x77, 0x5f,
	0x69, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x66,
	0x22, 0x68, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x25, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x6e, 0x65, 0x77,
	0x5f, 0x74, 0x61, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e,
	0x49, 0x6e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x62, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x68, 0x6f, 0x77,
	0x5f, 0x69, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x74, 0x6f,
	0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x62, 0x70, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bpmetadata_ui_proto_rawDescOnce sync.Once
	file_bpmetadata_ui_proto_rawDescData = file_bpmetadata_ui_proto_rawDesc
)

func file_bpmetadata_ui_proto_rawDescGZIP() []byte {
	file_bpmetadata_ui_proto_rawDescOnce.Do(func() {
		file_bpmetadata_ui_proto_rawDescData = protoimpl.X.CompressGZIP(file_bpmetadata_ui_proto_rawDescData)
	})
	return file_bpmetadata_ui_proto_rawDescData
}

var file_bpmetadata_ui_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bpmetadata_ui_proto_goTypes = []interface{}{
	(*BlueprintUIInput)(nil),        // 0: google.cloud.config.bpmetadata.BlueprintUIInput
	(*DisplayVariable)(nil),         // 1: google.cloud.config.bpmetadata.DisplayVariable
	(*DisplaySection)(nil),          // 2: google.cloud.config.bpmetadata.DisplaySection
	(*BlueprintUIOutput)(nil),       // 3: google.cloud.config.bpmetadata.BlueprintUIOutput
	(*UIActionItem)(nil),            // 4: google.cloud.config.bpmetadata.UIActionItem
	(*DisplayOutput)(nil),           // 5: google.cloud.config.bpmetadata.DisplayOutput
	nil,                             // 6: google.cloud.config.bpmetadata.BlueprintUIInput.VariablesEntry
	nil,                             // 7: google.cloud.config.bpmetadata.BlueprintUIOutput.OutputsEntry
	(*GooglePropertyExtension)(nil), // 8: google.cloud.config.bpmetadata.GooglePropertyExtension
}
var file_bpmetadata_ui_proto_depIdxs = []int32{
	6, // 0: google.cloud.config.bpmetadata.BlueprintUIInput.variables:type_name -> google.cloud.config.bpmetadata.BlueprintUIInput.VariablesEntry
	2, // 1: google.cloud.config.bpmetadata.BlueprintUIInput.sections:type_name -> google.cloud.config.bpmetadata.DisplaySection
	8, // 2: google.cloud.config.bpmetadata.DisplayVariable.x_google_property:type_name -> google.cloud.config.bpmetadata.GooglePropertyExtension
	4, // 3: google.cloud.config.bpmetadata.BlueprintUIOutput.suggested_actions:type_name -> google.cloud.config.bpmetadata.UIActionItem
	7, // 4: google.cloud.config.bpmetadata.BlueprintUIOutput.outputs:type_name -> google.cloud.config.bpmetadata.BlueprintUIOutput.OutputsEntry
	1, // 5: google.cloud.config.bpmetadata.BlueprintUIInput.VariablesEntry.value:type_name -> google.cloud.config.bpmetadata.DisplayVariable
	5, // 6: google.cloud.config.bpmetadata.BlueprintUIOutput.OutputsEntry.value:type_name -> google.cloud.config.bpmetadata.DisplayOutput
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_bpmetadata_ui_proto_init() }
func file_bpmetadata_ui_proto_init() {
	if File_bpmetadata_ui_proto != nil {
		return
	}
	file_bpmetadata_ui_ext_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bpmetadata_ui_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlueprintUIInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bpmetadata_ui_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplayVariable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bpmetadata_ui_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplaySection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bpmetadata_ui_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlueprintUIOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bpmetadata_ui_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UIActionItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bpmetadata_ui_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplayOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bpmetadata_ui_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bpmetadata_ui_proto_goTypes,
		DependencyIndexes: file_bpmetadata_ui_proto_depIdxs,
		MessageInfos:      file_bpmetadata_ui_proto_msgTypes,
	}.Build()
	File_bpmetadata_ui_proto = out.File
	file_bpmetadata_ui_proto_rawDesc = nil
	file_bpmetadata_ui_proto_goTypes = nil
	file_bpmetadata_ui_proto_depIdxs = nil
}
