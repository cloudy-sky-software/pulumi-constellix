// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package templates

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTXT(ctx *pulumi.Context, args *LookupTXTArgs, opts ...pulumi.InvokeOption) (*LookupTXTResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTXTResult
	err := ctx.Invoke("constellix:templates:getTXT", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTXTArgs struct {
	// The ID of the record
	Id string `pulumi:"id"`
	// The ID of the template object
	TemplateId string `pulumi:"templateId"`
}

type LookupTXTResult struct {
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
	LastValues *GetTXTPropertiesLastValuesProperties `pulumi:"lastValues"`
	// Links for the domain record
	Links *TemplaterecordLinksProperties `pulumi:"links"`
	// How the record should work
	Mode *GetTXTPropertiesMode `pulumi:"mode"`
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
	Type *GetTXTPropertiesType `pulumi:"type"`
	// Standard record mode
	Value []ValueTxtValueItemProperties `pulumi:"value"`
}

func LookupTXTOutput(ctx *pulumi.Context, args LookupTXTOutputArgs, opts ...pulumi.InvokeOption) LookupTXTResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTXTResultOutput, error) {
			args := v.(LookupTXTArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupTXTResult
			secret, err := ctx.InvokePackageRaw("constellix:templates:getTXT", args, &rv, "", opts...)
			if err != nil {
				return LookupTXTResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupTXTResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupTXTResultOutput), nil
			}
			return output, nil
		}).(LookupTXTResultOutput)
}

type LookupTXTOutputArgs struct {
	// The ID of the record
	Id pulumi.StringInput `pulumi:"id"`
	// The ID of the template object
	TemplateId pulumi.StringInput `pulumi:"templateId"`
}

func (LookupTXTOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTXTArgs)(nil)).Elem()
}

type LookupTXTResultOutput struct{ *pulumi.OutputState }

func (LookupTXTResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTXTResult)(nil)).Elem()
}

func (o LookupTXTResultOutput) ToLookupTXTResultOutput() LookupTXTResultOutput {
	return o
}

func (o LookupTXTResultOutput) ToLookupTXTResultOutputWithContext(ctx context.Context) LookupTXTResultOutput {
	return o
}

// A simple version of a contact list when inclued with other resources
func (o LookupTXTResultOutput) Contacts() SimpleContactlistPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *SimpleContactlist { return v.Contacts }).(SimpleContactlistPtrOutput)
}

func (o LookupTXTResultOutput) Domain() SimpleDomainPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *SimpleDomain { return v.Domain }).(SimpleDomainPtrOutput)
}

// Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
func (o LookupTXTResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
func (o LookupTXTResultOutput) GeoFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *bool { return v.GeoFailover }).(pulumi.BoolPtrOutput)
}

// Geo Proximity Location
func (o LookupTXTResultOutput) Geoproximity() SimpleGeoproximityPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *SimpleGeoproximity { return v.Geoproximity }).(SimpleGeoproximityPtrOutput)
}

// A unique ID for this domain record
func (o LookupTXTResultOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *int { return v.Id }).(pulumi.IntPtrOutput)
}

func (o LookupTXTResultOutput) Ipfilter() SimpleIpfilterPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *SimpleIpfilter { return v.Ipfilter }).(SimpleIpfilterPtrOutput)
}

// If the requesting IP matches the IP filter, don't return a response
func (o LookupTXTResultOutput) IpfilterDrop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *bool { return v.IpfilterDrop }).(pulumi.BoolPtrOutput)
}

// The previous values of the record in the different modes
func (o LookupTXTResultOutput) LastValues() GetTXTPropertiesLastValuesPropertiesPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *GetTXTPropertiesLastValuesProperties { return v.LastValues }).(GetTXTPropertiesLastValuesPropertiesPtrOutput)
}

// Links for the domain record
func (o LookupTXTResultOutput) Links() TemplaterecordLinksPropertiesPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *TemplaterecordLinksProperties { return v.Links }).(TemplaterecordLinksPropertiesPtrOutput)
}

// How the record should work
func (o LookupTXTResultOutput) Mode() GetTXTPropertiesModePtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *GetTXTPropertiesMode { return v.Mode }).(GetTXTPropertiesModePtrOutput)
}

// The name of the record
func (o LookupTXTResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A note about the record. Max 512 characters.
func (o LookupTXTResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The region for this record
func (o LookupTXTResultOutput) Region() RecordRegionPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *RecordRegion { return v.Region }).(RecordRegionPtrOutput)
}

// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
func (o LookupTXTResultOutput) SkipLookup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *bool { return v.SkipLookup }).(pulumi.BoolPtrOutput)
}

func (o LookupTXTResultOutput) Template() SimpleTemplatePtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *SimpleTemplate { return v.Template }).(SimpleTemplatePtrOutput)
}

// The time to live in seconds for this record. must be between 0 and 2147483647
func (o LookupTXTResultOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of record
func (o LookupTXTResultOutput) Type() GetTXTPropertiesTypePtrOutput {
	return o.ApplyT(func(v LookupTXTResult) *GetTXTPropertiesType { return v.Type }).(GetTXTPropertiesTypePtrOutput)
}

// Standard record mode
func (o LookupTXTResultOutput) Value() ValueTxtValueItemPropertiesArrayOutput {
	return o.ApplyT(func(v LookupTXTResult) []ValueTxtValueItemProperties { return v.Value }).(ValueTxtValueItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTXTResultOutput{})
}
