// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package templates

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Http struct {
	pulumi.CustomResourceState

	// Contact lists to be notified if a failover happens in a failover mode.
	Contacts pulumi.IntArrayOutput   `pulumi:"contacts"`
	Data     DataPropertiesPtrOutput `pulumi:"data"`
	// Whether the record is enabled
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
	GeoFailover pulumi.BoolPtrOutput `pulumi:"geoFailover"`
	// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
	Geoproximity pulumi.IntPtrOutput `pulumi:"geoproximity"`
	// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
	Ipfilter pulumi.IntPtrOutput `pulumi:"ipfilter"`
	// If the requesting IP matches the IP filter, don't return a response
	IpfilterDrop pulumi.BoolPtrOutput `pulumi:"ipfilterDrop"`
	// The current mode for this record
	Mode HttpPropertiesModePtrOutput `pulumi:"mode"`
	// The name for the record
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A description of the record. It must be 512 characters or less.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// Optional region for this record. Will default to 'default'.
	Region RecordCreateDetailsRegionPtrOutput `pulumi:"region"`
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup pulumi.BoolPtrOutput `pulumi:"skipLookup"`
	// How long DNS servers should cache the record for
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The type of record
	Type  HttpPropertiesTypePtrOutput       `pulumi:"type"`
	Value ValueHttpValuePropertiesPtrOutput `pulumi:"value"`
}

// NewHttp registers a new resource with the given unique name, arguments, and options.
func NewHttp(ctx *pulumi.Context,
	name string, args *HttpArgs, opts ...pulumi.ResourceOption) (*Http, error) {
	if args == nil {
		args = &HttpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Http
	err := ctx.RegisterResource("constellix:templates:Http", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttp gets an existing Http resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpState, opts ...pulumi.ResourceOption) (*Http, error) {
	var resource Http
	err := ctx.ReadResource("constellix:templates:Http", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Http resources.
type httpState struct {
}

type HttpState struct {
}

func (HttpState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpState)(nil)).Elem()
}

type httpArgs struct {
	// Contact lists to be notified if a failover happens in a failover mode.
	Contacts []int `pulumi:"contacts"`
	// Whether the record is enabled
	Enabled *bool `pulumi:"enabled"`
	// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
	GeoFailover *bool `pulumi:"geoFailover"`
	// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
	Geoproximity *int `pulumi:"geoproximity"`
	// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
	Ipfilter *int `pulumi:"ipfilter"`
	// If the requesting IP matches the IP filter, don't return a response
	IpfilterDrop *bool `pulumi:"ipfilterDrop"`
	// The current mode for this record
	Mode *HttpPropertiesMode `pulumi:"mode"`
	// The name for the record
	Name *string `pulumi:"name"`
	// A description of the record. It must be 512 characters or less.
	Notes *string `pulumi:"notes"`
	// Optional region for this record. Will default to 'default'.
	Region *RecordCreateDetailsRegion `pulumi:"region"`
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup *bool `pulumi:"skipLookup"`
	// The ID of the template object
	TemplateId *string `pulumi:"templateId"`
	// How long DNS servers should cache the record for
	Ttl *int `pulumi:"ttl"`
	// The type of record
	Type  *HttpPropertiesType       `pulumi:"type"`
	Value *ValueHttpValueProperties `pulumi:"value"`
}

// The set of arguments for constructing a Http resource.
type HttpArgs struct {
	// Contact lists to be notified if a failover happens in a failover mode.
	Contacts pulumi.IntArrayInput
	// Whether the record is enabled
	Enabled pulumi.BoolPtrInput
	// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
	GeoFailover pulumi.BoolPtrInput
	// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
	Geoproximity pulumi.IntPtrInput
	// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
	Ipfilter pulumi.IntPtrInput
	// If the requesting IP matches the IP filter, don't return a response
	IpfilterDrop pulumi.BoolPtrInput
	// The current mode for this record
	Mode HttpPropertiesModePtrInput
	// The name for the record
	Name pulumi.StringPtrInput
	// A description of the record. It must be 512 characters or less.
	Notes pulumi.StringPtrInput
	// Optional region for this record. Will default to 'default'.
	Region RecordCreateDetailsRegionPtrInput
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup pulumi.BoolPtrInput
	// The ID of the template object
	TemplateId pulumi.StringPtrInput
	// How long DNS servers should cache the record for
	Ttl pulumi.IntPtrInput
	// The type of record
	Type  HttpPropertiesTypePtrInput
	Value ValueHttpValuePropertiesPtrInput
}

func (HttpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpArgs)(nil)).Elem()
}

type HttpInput interface {
	pulumi.Input

	ToHttpOutput() HttpOutput
	ToHttpOutputWithContext(ctx context.Context) HttpOutput
}

func (*Http) ElementType() reflect.Type {
	return reflect.TypeOf((**Http)(nil)).Elem()
}

func (i *Http) ToHttpOutput() HttpOutput {
	return i.ToHttpOutputWithContext(context.Background())
}

func (i *Http) ToHttpOutputWithContext(ctx context.Context) HttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpOutput)
}

type HttpOutput struct{ *pulumi.OutputState }

func (HttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Http)(nil)).Elem()
}

func (o HttpOutput) ToHttpOutput() HttpOutput {
	return o
}

func (o HttpOutput) ToHttpOutputWithContext(ctx context.Context) HttpOutput {
	return o
}

// Contact lists to be notified if a failover happens in a failover mode.
func (o HttpOutput) Contacts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Http) pulumi.IntArrayOutput { return v.Contacts }).(pulumi.IntArrayOutput)
}

func (o HttpOutput) Data() DataPropertiesPtrOutput {
	return o.ApplyT(func(v *Http) DataPropertiesPtrOutput { return v.Data }).(DataPropertiesPtrOutput)
}

// Whether the record is enabled
func (o HttpOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o HttpOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.BoolPtrOutput { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
func (o HttpOutput) Geoproximity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.IntPtrOutput { return v.Geoproximity }).(pulumi.IntPtrOutput)
}

// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
func (o HttpOutput) Ipfilter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.IntPtrOutput { return v.Ipfilter }).(pulumi.IntPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o HttpOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.BoolPtrOutput { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The current mode for this record
func (o HttpOutput) Mode() HttpPropertiesModePtrOutput {
	return o.ApplyT(func(v *Http) HttpPropertiesModePtrOutput { return v.Mode }).(HttpPropertiesModePtrOutput)
}

// The name for the record
func (o HttpOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A description of the record. It must be 512 characters or less.
func (o HttpOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// Optional region for this record. Will default to 'default'.
func (o HttpOutput) Region() RecordCreateDetailsRegionPtrOutput {
	return o.ApplyT(func(v *Http) RecordCreateDetailsRegionPtrOutput { return v.Region }).(RecordCreateDetailsRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o HttpOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.BoolPtrOutput { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

// How long DNS servers should cache the record for
func (o HttpOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Http) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o HttpOutput) Type() HttpPropertiesTypePtrOutput {
	return o.ApplyT(func(v *Http) HttpPropertiesTypePtrOutput { return v.Type }).(HttpPropertiesTypePtrOutput)
}

func (o HttpOutput) Value() ValueHttpValuePropertiesPtrOutput {
	return o.ApplyT(func(v *Http) ValueHttpValuePropertiesPtrOutput { return v.Value }).(ValueHttpValuePropertiesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpInput)(nil)).Elem(), &Http{})
	pulumi.RegisterOutputType(HttpOutput{})
}
