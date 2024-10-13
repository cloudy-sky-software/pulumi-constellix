// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package contactlists

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListContactSlackWebhooks(ctx *pulumi.Context, args *ListContactSlackWebhooksArgs, opts ...pulumi.InvokeOption) (*ListContactSlackWebhooksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListContactSlackWebhooksResult
	err := ctx.Invoke("constellix:contactlists:listContactSlackWebhooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListContactSlackWebhooksArgs struct {
	// The ID of the Contact List
	Id string `pulumi:"id"`
}

type ListContactSlackWebhooksResult struct {
	// The webhooks for this page
	Data []ContactlistSlack `pulumi:"data"`
	// Metadata for list responses
	Meta *ListMetadata `pulumi:"meta"`
}

func ListContactSlackWebhooksOutput(ctx *pulumi.Context, args ListContactSlackWebhooksOutputArgs, opts ...pulumi.InvokeOption) ListContactSlackWebhooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListContactSlackWebhooksResultOutput, error) {
			args := v.(ListContactSlackWebhooksArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListContactSlackWebhooksResult
			secret, err := ctx.InvokePackageRaw("constellix:contactlists:listContactSlackWebhooks", args, &rv, "", opts...)
			if err != nil {
				return ListContactSlackWebhooksResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListContactSlackWebhooksResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListContactSlackWebhooksResultOutput), nil
			}
			return output, nil
		}).(ListContactSlackWebhooksResultOutput)
}

type ListContactSlackWebhooksOutputArgs struct {
	// The ID of the Contact List
	Id pulumi.StringInput `pulumi:"id"`
}

func (ListContactSlackWebhooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContactSlackWebhooksArgs)(nil)).Elem()
}

type ListContactSlackWebhooksResultOutput struct{ *pulumi.OutputState }

func (ListContactSlackWebhooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContactSlackWebhooksResult)(nil)).Elem()
}

func (o ListContactSlackWebhooksResultOutput) ToListContactSlackWebhooksResultOutput() ListContactSlackWebhooksResultOutput {
	return o
}

func (o ListContactSlackWebhooksResultOutput) ToListContactSlackWebhooksResultOutputWithContext(ctx context.Context) ListContactSlackWebhooksResultOutput {
	return o
}

// The webhooks for this page
func (o ListContactSlackWebhooksResultOutput) Data() ContactlistSlackArrayOutput {
	return o.ApplyT(func(v ListContactSlackWebhooksResult) []ContactlistSlack { return v.Data }).(ContactlistSlackArrayOutput)
}

// Metadata for list responses
func (o ListContactSlackWebhooksResultOutput) Meta() ListMetadataPtrOutput {
	return o.ApplyT(func(v ListContactSlackWebhooksResult) *ListMetadata { return v.Meta }).(ListMetadataPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListContactSlackWebhooksResultOutput{})
}
