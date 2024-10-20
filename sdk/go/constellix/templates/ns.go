// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package templates

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ns struct {
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
	Mode NsPropertiesModePtrOutput `pulumi:"mode"`
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
	Type NsPropertiesTypePtrOutput `pulumi:"type"`
	// Standard record mode
	Value ValueNsValueItemPropertiesArrayOutput `pulumi:"value"`
}

// NewNs registers a new resource with the given unique name, arguments, and options.
func NewNs(ctx *pulumi.Context,
	name string, args *NsArgs, opts ...pulumi.ResourceOption) (*Ns, error) {
	if args == nil {
		args = &NsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ns
	err := ctx.RegisterResource("constellix:templates:Ns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNs gets an existing Ns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsState, opts ...pulumi.ResourceOption) (*Ns, error) {
	var resource Ns
	err := ctx.ReadResource("constellix:templates:Ns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ns resources.
type nsState struct {
}

type NsState struct {
}

func (NsState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsState)(nil)).Elem()
}

type nsArgs struct {
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
	Mode *NsPropertiesMode `pulumi:"mode"`
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
	Type *NsPropertiesType `pulumi:"type"`
	// Standard record mode
	Value []ValueNsValueItemProperties `pulumi:"value"`
}

// The set of arguments for constructing a Ns resource.
type NsArgs struct {
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
	Mode NsPropertiesModePtrInput
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
	Type NsPropertiesTypePtrInput
	// Standard record mode
	Value ValueNsValueItemPropertiesArrayInput
}

func (NsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsArgs)(nil)).Elem()
}

type NsInput interface {
	pulumi.Input

	ToNsOutput() NsOutput
	ToNsOutputWithContext(ctx context.Context) NsOutput
}

func (*Ns) ElementType() reflect.Type {
	return reflect.TypeOf((**Ns)(nil)).Elem()
}

func (i *Ns) ToNsOutput() NsOutput {
	return i.ToNsOutputWithContext(context.Background())
}

func (i *Ns) ToNsOutputWithContext(ctx context.Context) NsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsOutput)
}

type NsOutput struct{ *pulumi.OutputState }

func (NsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ns)(nil)).Elem()
}

func (o NsOutput) ToNsOutput() NsOutput {
	return o
}

func (o NsOutput) ToNsOutputWithContext(ctx context.Context) NsOutput {
	return o
}

// Contact lists to be notified if a failover happens in a failover mode.
func (o NsOutput) Contacts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Ns) pulumi.IntArrayOutput { return v.Contacts }).(pulumi.IntArrayOutput)
}

func (o NsOutput) Data() DataPropertiesPtrOutput {
	return o.ApplyT(func(v *Ns) DataPropertiesPtrOutput { return v.Data }).(DataPropertiesPtrOutput)
}

// Whether the record is enabled
func (o NsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o NsOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.BoolPtrOutput { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
func (o NsOutput) Geoproximity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.IntPtrOutput { return v.Geoproximity }).(pulumi.IntPtrOutput)
}

// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
func (o NsOutput) Ipfilter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.IntPtrOutput { return v.Ipfilter }).(pulumi.IntPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o NsOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.BoolPtrOutput { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The current mode for this record
func (o NsOutput) Mode() NsPropertiesModePtrOutput {
	return o.ApplyT(func(v *Ns) NsPropertiesModePtrOutput { return v.Mode }).(NsPropertiesModePtrOutput)
}

// The name for the record
func (o NsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A description of the record. It must be 512 characters or less.
func (o NsOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// Optional region for this record. Will default to 'default'.
func (o NsOutput) Region() RecordCreateDetailsRegionPtrOutput {
	return o.ApplyT(func(v *Ns) RecordCreateDetailsRegionPtrOutput { return v.Region }).(RecordCreateDetailsRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o NsOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.BoolPtrOutput { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

// How long DNS servers should cache the record for
func (o NsOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ns) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o NsOutput) Type() NsPropertiesTypePtrOutput {
	return o.ApplyT(func(v *Ns) NsPropertiesTypePtrOutput { return v.Type }).(NsPropertiesTypePtrOutput)
}

// Standard record mode
func (o NsOutput) Value() ValueNsValueItemPropertiesArrayOutput {
	return o.ApplyT(func(v *Ns) ValueNsValueItemPropertiesArrayOutput { return v.Value }).(ValueNsValueItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsInput)(nil)).Elem(), &Ns{})
	pulumi.RegisterOutputType(NsOutput{})
}
