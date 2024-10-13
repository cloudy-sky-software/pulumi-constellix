// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDomainSnapshots(ctx *pulumi.Context, args *ListDomainSnapshotsArgs, opts ...pulumi.InvokeOption) (*ListDomainSnapshotsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDomainSnapshotsResult
	err := ctx.Invoke("constellix:domains:listDomainSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainSnapshotsArgs struct {
	// The ID of the domain object
	DomainId string `pulumi:"domainId"`
}

type ListDomainSnapshotsResult struct {
	// The domain snapshots for this page
	Data []ListDomainSnapshotsPropertiesDataItem `pulumi:"data"`
	// Metadata for list responses
	Meta *ListMetadata `pulumi:"meta"`
}

func ListDomainSnapshotsOutput(ctx *pulumi.Context, args ListDomainSnapshotsOutputArgs, opts ...pulumi.InvokeOption) ListDomainSnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDomainSnapshotsResultOutput, error) {
			args := v.(ListDomainSnapshotsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListDomainSnapshotsResult
			secret, err := ctx.InvokePackageRaw("constellix:domains:listDomainSnapshots", args, &rv, "", opts...)
			if err != nil {
				return ListDomainSnapshotsResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListDomainSnapshotsResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListDomainSnapshotsResultOutput), nil
			}
			return output, nil
		}).(ListDomainSnapshotsResultOutput)
}

type ListDomainSnapshotsOutputArgs struct {
	// The ID of the domain object
	DomainId pulumi.StringInput `pulumi:"domainId"`
}

func (ListDomainSnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainSnapshotsArgs)(nil)).Elem()
}

type ListDomainSnapshotsResultOutput struct{ *pulumi.OutputState }

func (ListDomainSnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainSnapshotsResult)(nil)).Elem()
}

func (o ListDomainSnapshotsResultOutput) ToListDomainSnapshotsResultOutput() ListDomainSnapshotsResultOutput {
	return o
}

func (o ListDomainSnapshotsResultOutput) ToListDomainSnapshotsResultOutputWithContext(ctx context.Context) ListDomainSnapshotsResultOutput {
	return o
}

// The domain snapshots for this page
func (o ListDomainSnapshotsResultOutput) Data() ListDomainSnapshotsPropertiesDataItemArrayOutput {
	return o.ApplyT(func(v ListDomainSnapshotsResult) []ListDomainSnapshotsPropertiesDataItem { return v.Data }).(ListDomainSnapshotsPropertiesDataItemArrayOutput)
}

// Metadata for list responses
func (o ListDomainSnapshotsResultOutput) Meta() ListMetadataPtrOutput {
	return o.ApplyT(func(v ListDomainSnapshotsResult) *ListMetadata { return v.Meta }).(ListMetadataPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDomainSnapshotsResultOutput{})
}
