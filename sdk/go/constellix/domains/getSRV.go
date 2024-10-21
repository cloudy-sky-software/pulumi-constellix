// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSRV(ctx *pulumi.Context, args *LookupSRVArgs, opts ...pulumi.InvokeOption) (*LookupSRVResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSRVResult
	err := ctx.Invoke("constellix:domains:getSRV", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSRVArgs struct {
	// The ID of the domain object
	DomainId string `pulumi:"domainId"`
	// The ID of the record
	Id string `pulumi:"id"`
}

type LookupSRVResult struct {
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
	LastValues *GetSRVPropertiesLastValuesProperties `pulumi:"lastValues"`
	// Links for the domain record
	Links *TemplaterecordLinksProperties `pulumi:"links"`
	// How the record should work
	Mode *GetSRVPropertiesMode `pulumi:"mode"`
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
	Type *GetSRVPropertiesType `pulumi:"type"`
	// Standard record mode
	Value []ValueSrvValueItemProperties `pulumi:"value"`
}

func LookupSRVOutput(ctx *pulumi.Context, args LookupSRVOutputArgs, opts ...pulumi.InvokeOption) LookupSRVResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSRVResultOutput, error) {
			args := v.(LookupSRVArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupSRVResult
			secret, err := ctx.InvokePackageRaw("constellix:domains:getSRV", args, &rv, "", opts...)
			if err != nil {
				return LookupSRVResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupSRVResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupSRVResultOutput), nil
			}
			return output, nil
		}).(LookupSRVResultOutput)
}

type LookupSRVOutputArgs struct {
	// The ID of the domain object
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The ID of the record
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSRVOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSRVArgs)(nil)).Elem()
}

type LookupSRVResultOutput struct{ *pulumi.OutputState }

func (LookupSRVResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSRVResult)(nil)).Elem()
}

func (o LookupSRVResultOutput) ToLookupSRVResultOutput() LookupSRVResultOutput {
	return o
}

func (o LookupSRVResultOutput) ToLookupSRVResultOutputWithContext(ctx context.Context) LookupSRVResultOutput {
	return o
}

// A simple version of a contact list when inclued with other resources
func (o LookupSRVResultOutput) Contacts() SimpleContactlistPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *SimpleContactlist { return v.Contacts }).(SimpleContactlistPtrOutput)
}

func (o LookupSRVResultOutput) Domain() SimpleDomainPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *SimpleDomain { return v.Domain }).(SimpleDomainPtrOutput)
}

// Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
func (o LookupSRVResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o LookupSRVResultOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *bool { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// Geo Proximity Location
func (o LookupSRVResultOutput) Geoproximity() SimpleGeoproximityPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *SimpleGeoproximity { return v.Geoproximity }).(SimpleGeoproximityPtrOutput)
}

// A unique ID for this domain record
func (o LookupSRVResultOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *int { return v.Id }).(pulumi.IntPtrOutput)
}

func (o LookupSRVResultOutput) Ipfilter() SimpleIpfilterPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *SimpleIpfilter { return v.Ipfilter }).(SimpleIpfilterPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o LookupSRVResultOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *bool { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The previous values of the record in the different modes
func (o LookupSRVResultOutput) LastValues() GetSRVPropertiesLastValuesPropertiesPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *GetSRVPropertiesLastValuesProperties { return v.LastValues }).(GetSRVPropertiesLastValuesPropertiesPtrOutput)
}

// Links for the domain record
func (o LookupSRVResultOutput) Links() TemplaterecordLinksPropertiesPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *TemplaterecordLinksProperties { return v.Links }).(TemplaterecordLinksPropertiesPtrOutput)
}

// How the record should work
func (o LookupSRVResultOutput) Mode() GetSRVPropertiesModePtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *GetSRVPropertiesMode { return v.Mode }).(GetSRVPropertiesModePtrOutput)
}

// The name of the record
func (o LookupSRVResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A note about the record. Max 512 characters.
func (o LookupSRVResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The region for this record
func (o LookupSRVResultOutput) Region() RecordRegionPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *RecordRegion { return v.Region }).(RecordRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o LookupSRVResultOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *bool { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

func (o LookupSRVResultOutput) Template() SimpleTemplatePtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *SimpleTemplate { return v.Template }).(SimpleTemplatePtrOutput)
}

// The time to live in seconds for this record. must be between 0 and 2147483647
func (o LookupSRVResultOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o LookupSRVResultOutput) Type() GetSRVPropertiesTypePtrOutput {
	return o.ApplyT(func(v LookupSRVResult) *GetSRVPropertiesType { return v.Type }).(GetSRVPropertiesTypePtrOutput)
}

// Standard record mode
func (o LookupSRVResultOutput) Value() ValueSrvValueItemPropertiesArrayOutput {
	return o.ApplyT(func(v LookupSRVResult) []ValueSrvValueItemProperties { return v.Value }).(ValueSrvValueItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSRVResultOutput{})
}
