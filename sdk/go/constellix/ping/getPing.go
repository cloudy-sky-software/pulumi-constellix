// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ping

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPing(ctx *pulumi.Context, args *GetPingArgs, opts ...pulumi.InvokeOption) (*GetPingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPingResult
	err := ctx.Invoke("constellix:ping:getPing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPingArgs struct {
}

type GetPingResult struct {
	Ip        *string `pulumi:"ip"`
	Timestamp *string `pulumi:"timestamp"`
	Version   *string `pulumi:"version"`
}

func GetPingOutput(ctx *pulumi.Context, args GetPingOutputArgs, opts ...pulumi.InvokeOption) GetPingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPingResultOutput, error) {
			args := v.(GetPingArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetPingResult
			secret, err := ctx.InvokePackageRaw("constellix:ping:getPing", args, &rv, "", opts...)
			if err != nil {
				return GetPingResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetPingResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetPingResultOutput), nil
			}
			return output, nil
		}).(GetPingResultOutput)
}

type GetPingOutputArgs struct {
}

func (GetPingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPingArgs)(nil)).Elem()
}

type GetPingResultOutput struct{ *pulumi.OutputState }

func (GetPingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPingResult)(nil)).Elem()
}

func (o GetPingResultOutput) ToGetPingResultOutput() GetPingResultOutput {
	return o
}

func (o GetPingResultOutput) ToGetPingResultOutputWithContext(ctx context.Context) GetPingResultOutput {
	return o
}

func (o GetPingResultOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPingResult) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o GetPingResultOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPingResult) *string { return v.Timestamp }).(pulumi.StringPtrOutput)
}

func (o GetPingResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPingResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPingResultOutput{})
}
