// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupA(ctx *pulumi.Context, args *LookupAArgs, opts ...pulumi.InvokeOption) (*LookupAResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAResult
	err := ctx.Invoke("constellix:domains:getA", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAArgs struct {
	// The ID of the domain object
	DomainId string `pulumi:"domainId"`
	// The ID of the record
	Id string `pulumi:"id"`
}

type LookupAResult struct {
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
	LastValues *GetAPropertiesLastValuesProperties `pulumi:"lastValues"`
	// Links for the domain record
	Links *TemplaterecordLinksProperties `pulumi:"links"`
	// How the record should work
	Mode *GetAPropertiesMode `pulumi:"mode"`
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
	Type  *ValueAType  `pulumi:"type"`
	Value *ValueAValue `pulumi:"value"`
}

func LookupAOutput(ctx *pulumi.Context, args LookupAOutputArgs, opts ...pulumi.InvokeOption) LookupAResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAResultOutput, error) {
			args := v.(LookupAArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupAResult
			secret, err := ctx.InvokePackageRaw("constellix:domains:getA", args, &rv, "", opts...)
			if err != nil {
				return LookupAResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupAResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupAResultOutput), nil
			}
			return output, nil
		}).(LookupAResultOutput)
}

type LookupAOutputArgs struct {
	// The ID of the domain object
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The ID of the record
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAArgs)(nil)).Elem()
}

type LookupAResultOutput struct{ *pulumi.OutputState }

func (LookupAResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAResult)(nil)).Elem()
}

func (o LookupAResultOutput) ToLookupAResultOutput() LookupAResultOutput {
	return o
}

func (o LookupAResultOutput) ToLookupAResultOutputWithContext(ctx context.Context) LookupAResultOutput {
	return o
}

// A simple version of a contact list when inclued with other resources
func (o LookupAResultOutput) Contacts() SimpleContactlistPtrOutput {
	return o.ApplyT(func(v LookupAResult) *SimpleContactlist { return v.Contacts }).(SimpleContactlistPtrOutput)
}

func (o LookupAResultOutput) Domain() SimpleDomainPtrOutput {
	return o.ApplyT(func(v LookupAResult) *SimpleDomain { return v.Domain }).(SimpleDomainPtrOutput)
}

// Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
func (o LookupAResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o LookupAResultOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAResult) *bool { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// Geo Proximity Location
func (o LookupAResultOutput) Geoproximity() SimpleGeoproximityPtrOutput {
	return o.ApplyT(func(v LookupAResult) *SimpleGeoproximity { return v.Geoproximity }).(SimpleGeoproximityPtrOutput)
}

// A unique ID for this domain record
func (o LookupAResultOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAResult) *int { return v.Id }).(pulumi.IntPtrOutput)
}

func (o LookupAResultOutput) Ipfilter() SimpleIpfilterPtrOutput {
	return o.ApplyT(func(v LookupAResult) *SimpleIpfilter { return v.Ipfilter }).(SimpleIpfilterPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o LookupAResultOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAResult) *bool { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The previous values of the record in the different modes
func (o LookupAResultOutput) LastValues() GetAPropertiesLastValuesPropertiesPtrOutput {
	return o.ApplyT(func(v LookupAResult) *GetAPropertiesLastValuesProperties { return v.LastValues }).(GetAPropertiesLastValuesPropertiesPtrOutput)
}

// Links for the domain record
func (o LookupAResultOutput) Links() TemplaterecordLinksPropertiesPtrOutput {
	return o.ApplyT(func(v LookupAResult) *TemplaterecordLinksProperties { return v.Links }).(TemplaterecordLinksPropertiesPtrOutput)
}

// How the record should work
func (o LookupAResultOutput) Mode() GetAPropertiesModePtrOutput {
	return o.ApplyT(func(v LookupAResult) *GetAPropertiesMode { return v.Mode }).(GetAPropertiesModePtrOutput)
}

// The name of the record
func (o LookupAResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A note about the record. Max 512 characters.
func (o LookupAResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The region for this record
func (o LookupAResultOutput) Region() RecordRegionPtrOutput {
	return o.ApplyT(func(v LookupAResult) *RecordRegion { return v.Region }).(RecordRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o LookupAResultOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAResult) *bool { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

func (o LookupAResultOutput) Template() SimpleTemplatePtrOutput {
	return o.ApplyT(func(v LookupAResult) *SimpleTemplate { return v.Template }).(SimpleTemplatePtrOutput)
}

// The time to live in seconds for this record. must be between 0 and 2147483647
func (o LookupAResultOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAResult) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o LookupAResultOutput) Type() ValueATypePtrOutput {
	return o.ApplyT(func(v LookupAResult) *ValueAType { return v.Type }).(ValueATypePtrOutput)
}

func (o LookupAResultOutput) Value() ValueAValuePtrOutput {
	return o.ApplyT(func(v LookupAResult) *ValueAValue { return v.Value }).(ValueAValuePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAResultOutput{})
}
