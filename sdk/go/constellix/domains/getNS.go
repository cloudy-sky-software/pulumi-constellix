// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNS(ctx *pulumi.Context, args *LookupNSArgs, opts ...pulumi.InvokeOption) (*LookupNSResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNSResult
	err := ctx.Invoke("constellix:domains:getNS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNSArgs struct {
	// The ID of the domain object
	DomainId string `pulumi:"domainId"`
	// The ID of the record
	Id string `pulumi:"id"`
}

type LookupNSResult struct {
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
	LastValues *GetNSPropertiesLastValuesProperties `pulumi:"lastValues"`
	// Links for the domain record
	Links *TemplaterecordLinksProperties `pulumi:"links"`
	// How the record should work
	Mode *GetNSPropertiesMode `pulumi:"mode"`
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
	Type *GetNSPropertiesType `pulumi:"type"`
	// Standard record mode
	Value []ValueNsValueItemProperties `pulumi:"value"`
}

func LookupNSOutput(ctx *pulumi.Context, args LookupNSOutputArgs, opts ...pulumi.InvokeOption) LookupNSResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNSResultOutput, error) {
			args := v.(LookupNSArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupNSResult
			secret, err := ctx.InvokePackageRaw("constellix:domains:getNS", args, &rv, "", opts...)
			if err != nil {
				return LookupNSResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupNSResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupNSResultOutput), nil
			}
			return output, nil
		}).(LookupNSResultOutput)
}

type LookupNSOutputArgs struct {
	// The ID of the domain object
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The ID of the record
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNSOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNSArgs)(nil)).Elem()
}

type LookupNSResultOutput struct{ *pulumi.OutputState }

func (LookupNSResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNSResult)(nil)).Elem()
}

func (o LookupNSResultOutput) ToLookupNSResultOutput() LookupNSResultOutput {
	return o
}

func (o LookupNSResultOutput) ToLookupNSResultOutputWithContext(ctx context.Context) LookupNSResultOutput {
	return o
}

// A simple version of a contact list when inclued with other resources
func (o LookupNSResultOutput) Contacts() SimpleContactlistPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *SimpleContactlist { return v.Contacts }).(SimpleContactlistPtrOutput)
}

func (o LookupNSResultOutput) Domain() SimpleDomainPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *SimpleDomain { return v.Domain }).(SimpleDomainPtrOutput)
}

// Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
func (o LookupNSResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o LookupNSResultOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *bool { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// Geo Proximity Location
func (o LookupNSResultOutput) Geoproximity() SimpleGeoproximityPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *SimpleGeoproximity { return v.Geoproximity }).(SimpleGeoproximityPtrOutput)
}

// A unique ID for this domain record
func (o LookupNSResultOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *int { return v.Id }).(pulumi.IntPtrOutput)
}

func (o LookupNSResultOutput) Ipfilter() SimpleIpfilterPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *SimpleIpfilter { return v.Ipfilter }).(SimpleIpfilterPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o LookupNSResultOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *bool { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The previous values of the record in the different modes
func (o LookupNSResultOutput) LastValues() GetNSPropertiesLastValuesPropertiesPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *GetNSPropertiesLastValuesProperties { return v.LastValues }).(GetNSPropertiesLastValuesPropertiesPtrOutput)
}

// Links for the domain record
func (o LookupNSResultOutput) Links() TemplaterecordLinksPropertiesPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *TemplaterecordLinksProperties { return v.Links }).(TemplaterecordLinksPropertiesPtrOutput)
}

// How the record should work
func (o LookupNSResultOutput) Mode() GetNSPropertiesModePtrOutput {
	return o.ApplyT(func(v LookupNSResult) *GetNSPropertiesMode { return v.Mode }).(GetNSPropertiesModePtrOutput)
}

// The name of the record
func (o LookupNSResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A note about the record. Max 512 characters.
func (o LookupNSResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The region for this record
func (o LookupNSResultOutput) Region() RecordRegionPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *RecordRegion { return v.Region }).(RecordRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o LookupNSResultOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *bool { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

func (o LookupNSResultOutput) Template() SimpleTemplatePtrOutput {
	return o.ApplyT(func(v LookupNSResult) *SimpleTemplate { return v.Template }).(SimpleTemplatePtrOutput)
}

// The time to live in seconds for this record. must be between 0 and 2147483647
func (o LookupNSResultOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNSResult) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o LookupNSResultOutput) Type() GetNSPropertiesTypePtrOutput {
	return o.ApplyT(func(v LookupNSResult) *GetNSPropertiesType { return v.Type }).(GetNSPropertiesTypePtrOutput)
}

// Standard record mode
func (o LookupNSResultOutput) Value() ValueNsValueItemPropertiesArrayOutput {
	return o.ApplyT(func(v LookupNSResult) []ValueNsValueItemProperties { return v.Value }).(ValueNsValueItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNSResultOutput{})
}
