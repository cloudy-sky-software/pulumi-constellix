// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainRecord struct {
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
	// The name for the record
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A description of the record. It must be 512 characters or less.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// Optional region for this record. Will default to 'default'.
	Region DomainRecordPropertiesRegionPtrOutput `pulumi:"region"`
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup pulumi.BoolPtrOutput `pulumi:"skipLookup"`
	// How long DNS servers should cache the record for
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
}

// NewDomainRecord registers a new resource with the given unique name, arguments, and options.
func NewDomainRecord(ctx *pulumi.Context,
	name string, args *DomainRecordArgs, opts ...pulumi.ResourceOption) (*DomainRecord, error) {
	if args == nil {
		args = &DomainRecordArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainRecord
	err := ctx.RegisterResource("constellix:domains:DomainRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainRecord gets an existing DomainRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainRecordState, opts ...pulumi.ResourceOption) (*DomainRecord, error) {
	var resource DomainRecord
	err := ctx.ReadResource("constellix:domains:DomainRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainRecord resources.
type domainRecordState struct {
}

type DomainRecordState struct {
}

func (DomainRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRecordState)(nil)).Elem()
}

type domainRecordArgs struct {
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
	// The name for the record
	Name *string `pulumi:"name"`
	// A description of the record. It must be 512 characters or less.
	Notes *string `pulumi:"notes"`
	// Optional region for this record. Will default to 'default'.
	Region *DomainRecordPropertiesRegion `pulumi:"region"`
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup *bool `pulumi:"skipLookup"`
	// How long DNS servers should cache the record for
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a DomainRecord resource.
type DomainRecordArgs struct {
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
	// The name for the record
	Name pulumi.StringPtrInput
	// A description of the record. It must be 512 characters or less.
	Notes pulumi.StringPtrInput
	// Optional region for this record. Will default to 'default'.
	Region DomainRecordPropertiesRegionPtrInput
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup pulumi.BoolPtrInput
	// How long DNS servers should cache the record for
	Ttl pulumi.IntPtrInput
}

func (DomainRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRecordArgs)(nil)).Elem()
}

type DomainRecordInput interface {
	pulumi.Input

	ToDomainRecordOutput() DomainRecordOutput
	ToDomainRecordOutputWithContext(ctx context.Context) DomainRecordOutput
}

func (*DomainRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRecord)(nil)).Elem()
}

func (i *DomainRecord) ToDomainRecordOutput() DomainRecordOutput {
	return i.ToDomainRecordOutputWithContext(context.Background())
}

func (i *DomainRecord) ToDomainRecordOutputWithContext(ctx context.Context) DomainRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainRecordOutput)
}

type DomainRecordOutput struct{ *pulumi.OutputState }

func (DomainRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRecord)(nil)).Elem()
}

func (o DomainRecordOutput) ToDomainRecordOutput() DomainRecordOutput {
	return o
}

func (o DomainRecordOutput) ToDomainRecordOutputWithContext(ctx context.Context) DomainRecordOutput {
	return o
}

// Contact lists to be notified if a failover happens in a failover mode.
func (o DomainRecordOutput) Contacts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.IntArrayOutput { return v.Contacts }).(pulumi.IntArrayOutput)
}

func (o DomainRecordOutput) Data() DataPropertiesPtrOutput {
	return o.ApplyT(func(v *DomainRecord) DataPropertiesPtrOutput { return v.Data }).(DataPropertiesPtrOutput)
}

// Whether the record is enabled
func (o DomainRecordOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o DomainRecordOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.BoolPtrOutput { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
func (o DomainRecordOutput) Geoproximity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.IntPtrOutput { return v.Geoproximity }).(pulumi.IntPtrOutput)
}

// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
func (o DomainRecordOutput) Ipfilter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.IntPtrOutput { return v.Ipfilter }).(pulumi.IntPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o DomainRecordOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.BoolPtrOutput { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The name for the record
func (o DomainRecordOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A description of the record. It must be 512 characters or less.
func (o DomainRecordOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// Optional region for this record. Will default to 'default'.
func (o DomainRecordOutput) Region() DomainRecordPropertiesRegionPtrOutput {
	return o.ApplyT(func(v *DomainRecord) DomainRecordPropertiesRegionPtrOutput { return v.Region }).(DomainRecordPropertiesRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o DomainRecordOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.BoolPtrOutput { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

// How long DNS servers should cache the record for
func (o DomainRecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRecordInput)(nil)).Elem(), &DomainRecord{})
	pulumi.RegisterOutputType(DomainRecordOutput{})
}
