// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package search

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSearchDomain(ctx *pulumi.Context, args *GetSearchDomainArgs, opts ...pulumi.InvokeOption) (*GetSearchDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSearchDomainResult
	err := ctx.Invoke("constellix:search:getSearchDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSearchDomainArgs struct {
}

type GetSearchDomainResult struct {
	// The search results for this page
	Data []Domainsearchresult `pulumi:"data"`
	// Metadata for list responses
	Meta *ListMetadata `pulumi:"meta"`
}

func GetSearchDomainOutput(ctx *pulumi.Context, args GetSearchDomainOutputArgs, opts ...pulumi.InvokeOption) GetSearchDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSearchDomainResultOutput, error) {
			args := v.(GetSearchDomainArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetSearchDomainResult
			secret, err := ctx.InvokePackageRaw("constellix:search:getSearchDomain", args, &rv, "", opts...)
			if err != nil {
				return GetSearchDomainResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetSearchDomainResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetSearchDomainResultOutput), nil
			}
			return output, nil
		}).(GetSearchDomainResultOutput)
}

type GetSearchDomainOutputArgs struct {
}

func (GetSearchDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSearchDomainArgs)(nil)).Elem()
}

type GetSearchDomainResultOutput struct{ *pulumi.OutputState }

func (GetSearchDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSearchDomainResult)(nil)).Elem()
}

func (o GetSearchDomainResultOutput) ToGetSearchDomainResultOutput() GetSearchDomainResultOutput {
	return o
}

func (o GetSearchDomainResultOutput) ToGetSearchDomainResultOutputWithContext(ctx context.Context) GetSearchDomainResultOutput {
	return o
}

// The search results for this page
func (o GetSearchDomainResultOutput) Data() DomainsearchresultArrayOutput {
	return o.ApplyT(func(v GetSearchDomainResult) []Domainsearchresult { return v.Data }).(DomainsearchresultArrayOutput)
}

// Metadata for list responses
func (o GetSearchDomainResultOutput) Meta() ListMetadataPtrOutput {
	return o.ApplyT(func(v GetSearchDomainResult) *ListMetadata { return v.Meta }).(ListMetadataPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSearchDomainResultOutput{})
}
