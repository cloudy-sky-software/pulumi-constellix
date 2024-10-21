// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cname struct {
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
	Mode CnamePropertiesModePtrOutput `pulumi:"mode"`
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
	Type  CnamePropertiesTypePtrOutput `pulumi:"type"`
	Value pulumi.AnyOutput             `pulumi:"value"`
}

// NewCname registers a new resource with the given unique name, arguments, and options.
func NewCname(ctx *pulumi.Context,
	name string, args *CnameArgs, opts ...pulumi.ResourceOption) (*Cname, error) {
	if args == nil {
		args = &CnameArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cname
	err := ctx.RegisterResource("constellix:domains:Cname", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCname gets an existing Cname resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCname(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CnameState, opts ...pulumi.ResourceOption) (*Cname, error) {
	var resource Cname
	err := ctx.ReadResource("constellix:domains:Cname", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cname resources.
type cnameState struct {
}

type CnameState struct {
}

func (CnameState) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameState)(nil)).Elem()
}

type cnameArgs struct {
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
	Mode *CnamePropertiesMode `pulumi:"mode"`
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
	Type  *CnamePropertiesType `pulumi:"type"`
	Value interface{}          `pulumi:"value"`
}

// The set of arguments for constructing a Cname resource.
type CnameArgs struct {
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
	Mode CnamePropertiesModePtrInput
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
	Type  CnamePropertiesTypePtrInput
	Value pulumi.Input
}

func (CnameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameArgs)(nil)).Elem()
}

type CnameInput interface {
	pulumi.Input

	ToCnameOutput() CnameOutput
	ToCnameOutputWithContext(ctx context.Context) CnameOutput
}

func (*Cname) ElementType() reflect.Type {
	return reflect.TypeOf((**Cname)(nil)).Elem()
}

func (i *Cname) ToCnameOutput() CnameOutput {
	return i.ToCnameOutputWithContext(context.Background())
}

func (i *Cname) ToCnameOutputWithContext(ctx context.Context) CnameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameOutput)
}

type CnameOutput struct{ *pulumi.OutputState }

func (CnameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cname)(nil)).Elem()
}

func (o CnameOutput) ToCnameOutput() CnameOutput {
	return o
}

func (o CnameOutput) ToCnameOutputWithContext(ctx context.Context) CnameOutput {
	return o
}

// Contact lists to be notified if a failover happens in a failover mode.
func (o CnameOutput) Contacts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Cname) pulumi.IntArrayOutput { return v.Contacts }).(pulumi.IntArrayOutput)
}

func (o CnameOutput) Data() DataPropertiesPtrOutput {
	return o.ApplyT(func(v *Cname) DataPropertiesPtrOutput { return v.Data }).(DataPropertiesPtrOutput)
}

// Whether the record is enabled
func (o CnameOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o CnameOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.BoolPtrOutput { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
func (o CnameOutput) Geoproximity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.IntPtrOutput { return v.Geoproximity }).(pulumi.IntPtrOutput)
}

// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
func (o CnameOutput) Ipfilter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.IntPtrOutput { return v.Ipfilter }).(pulumi.IntPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o CnameOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.BoolPtrOutput { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The current mode for this record
func (o CnameOutput) Mode() CnamePropertiesModePtrOutput {
	return o.ApplyT(func(v *Cname) CnamePropertiesModePtrOutput { return v.Mode }).(CnamePropertiesModePtrOutput)
}

// The name for the record
func (o CnameOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A description of the record. It must be 512 characters or less.
func (o CnameOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// Optional region for this record. Will default to 'default'.
func (o CnameOutput) Region() RecordCreateDetailsRegionPtrOutput {
	return o.ApplyT(func(v *Cname) RecordCreateDetailsRegionPtrOutput { return v.Region }).(RecordCreateDetailsRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o CnameOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.BoolPtrOutput { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

// How long DNS servers should cache the record for
func (o CnameOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cname) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o CnameOutput) Type() CnamePropertiesTypePtrOutput {
	return o.ApplyT(func(v *Cname) CnamePropertiesTypePtrOutput { return v.Type }).(CnamePropertiesTypePtrOutput)
}

func (o CnameOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v *Cname) pulumi.AnyOutput { return v.Value }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CnameInput)(nil)).Elem(), &Cname{})
	pulumi.RegisterOutputType(CnameOutput{})
}
