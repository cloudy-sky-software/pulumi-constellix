// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAAAA(ctx *pulumi.Context, args *LookupAAAAArgs, opts ...pulumi.InvokeOption) (*LookupAAAAResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAAAAResult
	err := ctx.Invoke("constellix:domains:getAAAA", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAAAAArgs struct {
	// The ID of the domain object
	DomainId string `pulumi:"domainId"`
	// The ID of the record
	Id string `pulumi:"id"`
}

type LookupAAAAResult struct {
	// A simple version of a contact list when inclued with other resources
	Contacts *SimpleContactlist `pulumi:"contacts"`
	Domain   *SimpleDomain      `pulumi:"domain"`
	// Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
	Enabled *bool `pulumi:"enabled"`
	// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
	GeoFailover *bool `pulumi:"geoFailover"`
	// Geo Proximity Location
	Geoproximity *SimpleGeoproximity `pulumi:"geoproximity"`
	// A unique ID for this domain record
	Id       *int            `pulumi:"id"`
	Ipfilter *SimpleIpfilter `pulumi:"ipfilter"`
	// If the requesting IP matches the IP filter, don't return a response
	IpfilterDrop *bool `pulumi:"ipfilterDrop"`
	// The previous values of the record in the different modes
	LastValues *GetAAAAPropertiesLastValuesProperties `pulumi:"lastValues"`
	// Links for the domain record
	Links *TemplaterecordLinksProperties `pulumi:"links"`
	// How the record should work
	Mode *GetAAAAPropertiesMode `pulumi:"mode"`
	// The name of the record
	Name *string `pulumi:"name"`
	// A note about the record. Max 512 characters.
	Notes *string `pulumi:"notes"`
	// The region for this record
	Region *RecordRegion `pulumi:"region"`
	// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
	SkipLookup *bool           `pulumi:"skipLookup"`
	Template   *SimpleTemplate `pulumi:"template"`
	// The time to live in seconds for this record. must be between 0 and 2147483647
	Ttl *int `pulumi:"ttl"`
	// The type of record
	Type  *ValueAaaaType  `pulumi:"type"`
	Value *ValueAaaaValue `pulumi:"value"`
}

func LookupAAAAOutput(ctx *pulumi.Context, args LookupAAAAOutputArgs, opts ...pulumi.InvokeOption) LookupAAAAResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAAAAResultOutput, error) {
			args := v.(LookupAAAAArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupAAAAResult
			secret, err := ctx.InvokePackageRaw("constellix:domains:getAAAA", args, &rv, "", opts...)
			if err != nil {
				return LookupAAAAResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupAAAAResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupAAAAResultOutput), nil
			}
			return output, nil
		}).(LookupAAAAResultOutput)
}

type LookupAAAAOutputArgs struct {
	// The ID of the domain object
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The ID of the record
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAAAAOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAAAAArgs)(nil)).Elem()
}

type LookupAAAAResultOutput struct{ *pulumi.OutputState }

func (LookupAAAAResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAAAAResult)(nil)).Elem()
}

func (o LookupAAAAResultOutput) ToLookupAAAAResultOutput() LookupAAAAResultOutput {
	return o
}

func (o LookupAAAAResultOutput) ToLookupAAAAResultOutputWithContext(ctx context.Context) LookupAAAAResultOutput {
	return o
}

// A simple version of a contact list when inclued with other resources
func (o LookupAAAAResultOutput) Contacts() SimpleContactlistPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *SimpleContactlist { return v.Contacts }).(SimpleContactlistPtrOutput)
}

func (o LookupAAAAResultOutput) Domain() SimpleDomainPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *SimpleDomain { return v.Domain }).(SimpleDomainPtrOutput)
}

// Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
func (o LookupAAAAResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o LookupAAAAResultOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *bool { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// Geo Proximity Location
func (o LookupAAAAResultOutput) Geoproximity() SimpleGeoproximityPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *SimpleGeoproximity { return v.Geoproximity }).(SimpleGeoproximityPtrOutput)
}

// A unique ID for this domain record
func (o LookupAAAAResultOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *int { return v.Id }).(pulumi.IntPtrOutput)
}

func (o LookupAAAAResultOutput) Ipfilter() SimpleIpfilterPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *SimpleIpfilter { return v.Ipfilter }).(SimpleIpfilterPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o LookupAAAAResultOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *bool { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The previous values of the record in the different modes
func (o LookupAAAAResultOutput) LastValues() GetAAAAPropertiesLastValuesPropertiesPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *GetAAAAPropertiesLastValuesProperties { return v.LastValues }).(GetAAAAPropertiesLastValuesPropertiesPtrOutput)
}

// Links for the domain record
func (o LookupAAAAResultOutput) Links() TemplaterecordLinksPropertiesPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *TemplaterecordLinksProperties { return v.Links }).(TemplaterecordLinksPropertiesPtrOutput)
}

// How the record should work
func (o LookupAAAAResultOutput) Mode() GetAAAAPropertiesModePtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *GetAAAAPropertiesMode { return v.Mode }).(GetAAAAPropertiesModePtrOutput)
}

// The name of the record
func (o LookupAAAAResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A note about the record. Max 512 characters.
func (o LookupAAAAResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The region for this record
func (o LookupAAAAResultOutput) Region() RecordRegionPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *RecordRegion { return v.Region }).(RecordRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o LookupAAAAResultOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *bool { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

func (o LookupAAAAResultOutput) Template() SimpleTemplatePtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *SimpleTemplate { return v.Template }).(SimpleTemplatePtrOutput)
}

// The time to live in seconds for this record. must be between 0 and 2147483647
func (o LookupAAAAResultOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o LookupAAAAResultOutput) Type() ValueAaaaTypePtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *ValueAaaaType { return v.Type }).(ValueAaaaTypePtrOutput)
}

func (o LookupAAAAResultOutput) Value() ValueAaaaValuePtrOutput {
	return o.ApplyT(func(v LookupAAAAResult) *ValueAaaaValue { return v.Value }).(ValueAaaaValuePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAAAAResultOutput{})
}
