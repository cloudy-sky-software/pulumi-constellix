// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ptr struct {
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
	Mode PtrPropertiesModePtrOutput `pulumi:"mode"`
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
	Type PtrPropertiesTypePtrOutput `pulumi:"type"`
	// Standard record mode
	Value ValuePtrValueItemPropertiesArrayOutput `pulumi:"value"`
}

// NewPtr registers a new resource with the given unique name, arguments, and options.
func NewPtr(ctx *pulumi.Context,
	name string, args *PtrArgs, opts ...pulumi.ResourceOption) (*Ptr, error) {
	if args == nil {
		args = &PtrArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ptr
	err := ctx.RegisterResource("constellix:domains:Ptr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPtr gets an existing Ptr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPtr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PtrState, opts ...pulumi.ResourceOption) (*Ptr, error) {
	var resource Ptr
	err := ctx.ReadResource("constellix:domains:Ptr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ptr resources.
type ptrState struct {
}

type PtrState struct {
}

func (PtrState) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrState)(nil)).Elem()
}

type ptrArgs struct {
	// Contact lists to be notified if a failover happens in a failover mode.
	Contacts []int `pulumi:"contacts"`
	// The ID of the domain object
	DomainId *string `pulumi:"domainId"`
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
	Mode *PtrPropertiesMode `pulumi:"mode"`
	// The name for the record
	Name *string `pulumi:"name"`
	// A description of the record. It must be 512 characters or less.
	Notes *string `pulumi:"notes"`
	// Optional region for this record. Will default to 'default'.
	Region *RecordCreateDetailsRegion `pulumi:"region"`
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup *bool `pulumi:"skipLookup"`
	// How long DNS servers should cache the record for
	Ttl *int `pulumi:"ttl"`
	// The type of record
	Type *PtrPropertiesType `pulumi:"type"`
	// Standard record mode
	Value []ValuePtrValueItemProperties `pulumi:"value"`
}

// The set of arguments for constructing a Ptr resource.
type PtrArgs struct {
	// Contact lists to be notified if a failover happens in a failover mode.
	Contacts pulumi.IntArrayInput
	// The ID of the domain object
	DomainId pulumi.StringPtrInput
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
	Mode PtrPropertiesModePtrInput
	// The name for the record
	Name pulumi.StringPtrInput
	// A description of the record. It must be 512 characters or less.
	Notes pulumi.StringPtrInput
	// Optional region for this record. Will default to 'default'.
	Region RecordCreateDetailsRegionPtrInput
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup pulumi.BoolPtrInput
	// How long DNS servers should cache the record for
	Ttl pulumi.IntPtrInput
	// The type of record
	Type PtrPropertiesTypePtrInput
	// Standard record mode
	Value ValuePtrValueItemPropertiesArrayInput
}

func (PtrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrArgs)(nil)).Elem()
}

type PtrInput interface {
	pulumi.Input

	ToPtrOutput() PtrOutput
	ToPtrOutputWithContext(ctx context.Context) PtrOutput
}

func (*Ptr) ElementType() reflect.Type {
	return reflect.TypeOf((**Ptr)(nil)).Elem()
}

func (i *Ptr) ToPtrOutput() PtrOutput {
	return i.ToPtrOutputWithContext(context.Background())
}

func (i *Ptr) ToPtrOutputWithContext(ctx context.Context) PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrOutput)
}

type PtrOutput struct{ *pulumi.OutputState }

func (PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ptr)(nil)).Elem()
}

func (o PtrOutput) ToPtrOutput() PtrOutput {
	return o
}

func (o PtrOutput) ToPtrOutputWithContext(ctx context.Context) PtrOutput {
	return o
}

// Contact lists to be notified if a failover happens in a failover mode.
func (o PtrOutput) Contacts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Ptr) pulumi.IntArrayOutput { return v.Contacts }).(pulumi.IntArrayOutput)
}

func (o PtrOutput) Data() DataPropertiesPtrOutput {
	return o.ApplyT(func(v *Ptr) DataPropertiesPtrOutput { return v.Data }).(DataPropertiesPtrOutput)
}

// Whether the record is enabled
func (o PtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o PtrOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.BoolPtrOutput { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
func (o PtrOutput) Geoproximity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.IntPtrOutput { return v.Geoproximity }).(pulumi.IntPtrOutput)
}

// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
func (o PtrOutput) Ipfilter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.IntPtrOutput { return v.Ipfilter }).(pulumi.IntPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o PtrOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.BoolPtrOutput { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The current mode for this record
func (o PtrOutput) Mode() PtrPropertiesModePtrOutput {
	return o.ApplyT(func(v *Ptr) PtrPropertiesModePtrOutput { return v.Mode }).(PtrPropertiesModePtrOutput)
}

// The name for the record
func (o PtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A description of the record. It must be 512 characters or less.
func (o PtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// Optional region for this record. Will default to 'default'.
func (o PtrOutput) Region() RecordCreateDetailsRegionPtrOutput {
	return o.ApplyT(func(v *Ptr) RecordCreateDetailsRegionPtrOutput { return v.Region }).(RecordCreateDetailsRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o PtrOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.BoolPtrOutput { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

// How long DNS servers should cache the record for
func (o PtrOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ptr) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o PtrOutput) Type() PtrPropertiesTypePtrOutput {
	return o.ApplyT(func(v *Ptr) PtrPropertiesTypePtrOutput { return v.Type }).(PtrPropertiesTypePtrOutput)
}

// Standard record mode
func (o PtrOutput) Value() ValuePtrValueItemPropertiesArrayOutput {
	return o.ApplyT(func(v *Ptr) ValuePtrValueItemPropertiesArrayOutput { return v.Value }).(ValuePtrValueItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PtrInput)(nil)).Elem(), &Ptr{})
	pulumi.RegisterOutputType(PtrOutput{})
}
