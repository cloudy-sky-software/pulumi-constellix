// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTagResult
	err := ctx.Invoke("constellix:tags:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagArgs struct {
	// The ID of the Tag
	Id string `pulumi:"id"`
}

type LookupTagResult struct {
	// A tag is used to group resources together
	Data *TagType `pulumi:"data"`
}

func LookupTagOutput(ctx *pulumi.Context, args LookupTagOutputArgs, opts ...pulumi.InvokeOption) LookupTagResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagResultOutput, error) {
			args := v.(LookupTagArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupTagResult
			secret, err := ctx.InvokePackageRaw("constellix:tags:getTag", args, &rv, "", opts...)
			if err != nil {
				return LookupTagResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupTagResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupTagResultOutput), nil
			}
			return output, nil
		}).(LookupTagResultOutput)
}

type LookupTagOutputArgs struct {
	// The ID of the Tag
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagArgs)(nil)).Elem()
}

type LookupTagResultOutput struct{ *pulumi.OutputState }

func (LookupTagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagResult)(nil)).Elem()
}

func (o LookupTagResultOutput) ToLookupTagResultOutput() LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) ToLookupTagResultOutputWithContext(ctx context.Context) LookupTagResultOutput {
	return o
}

// A tag is used to group resources together
func (o LookupTagResultOutput) Data() TagTypePtrOutput {
	return o.ApplyT(func(v LookupTagResult) *TagType { return v.Data }).(TagTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagResultOutput{})
}
