// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Tag struct {
	pulumi.CustomResourceState

	Data DataPropertiesPtrOutput `pulumi:"data"`
	// A name for this tag
	Name pulumi.StringPtrOutput `pulumi:"name"`
}

// NewTag registers a new resource with the given unique name, arguments, and options.
func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil {
		args = &TagArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Tag
	err := ctx.RegisterResource("constellix:tags:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTag gets an existing Tag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("constellix:tags:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tag resources.
type tagState struct {
}

type TagState struct {
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	// A name for this tag
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Tag resource.
type TagArgs struct {
	// A name for this tag
	Name pulumi.StringPtrInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}

type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(ctx context.Context) TagOutput
}

func (*Tag) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (i *Tag) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i *Tag) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

func (o TagOutput) Data() DataPropertiesPtrOutput {
	return o.ApplyT(func(v *Tag) DataPropertiesPtrOutput { return v.Data }).(DataPropertiesPtrOutput)
}

// A name for this tag
func (o TagOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagInput)(nil)).Elem(), &Tag{})
	pulumi.RegisterOutputType(TagOutput{})
}
