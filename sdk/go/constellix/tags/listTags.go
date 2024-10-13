// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTags(ctx *pulumi.Context, args *ListTagsArgs, opts ...pulumi.InvokeOption) (*ListTagsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListTagsResult
	err := ctx.Invoke("constellix:tags:listTags", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTagsArgs struct {
}

type ListTagsResult struct {
	Data []TagType `pulumi:"data"`
	// Metadata for list responses
	Meta *ListMetadata `pulumi:"meta"`
}

func ListTagsOutput(ctx *pulumi.Context, args ListTagsOutputArgs, opts ...pulumi.InvokeOption) ListTagsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTagsResultOutput, error) {
			args := v.(ListTagsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListTagsResult
			secret, err := ctx.InvokePackageRaw("constellix:tags:listTags", args, &rv, "", opts...)
			if err != nil {
				return ListTagsResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListTagsResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListTagsResultOutput), nil
			}
			return output, nil
		}).(ListTagsResultOutput)
}

type ListTagsOutputArgs struct {
}

func (ListTagsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTagsArgs)(nil)).Elem()
}

type ListTagsResultOutput struct{ *pulumi.OutputState }

func (ListTagsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTagsResult)(nil)).Elem()
}

func (o ListTagsResultOutput) ToListTagsResultOutput() ListTagsResultOutput {
	return o
}

func (o ListTagsResultOutput) ToListTagsResultOutputWithContext(ctx context.Context) ListTagsResultOutput {
	return o
}

func (o ListTagsResultOutput) Data() TagTypeArrayOutput {
	return o.ApplyT(func(v ListTagsResult) []TagType { return v.Data }).(TagTypeArrayOutput)
}

// Metadata for list responses
func (o ListTagsResultOutput) Meta() ListMetadataPtrOutput {
	return o.ApplyT(func(v ListTagsResult) *ListMetadata { return v.Meta }).(ListMetadataPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTagsResultOutput{})
}
