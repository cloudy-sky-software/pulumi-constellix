// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Optional region for this record. Will default to 'default'.
type DomainRecordPropertiesRegion string

const (
	DomainRecordPropertiesRegionDefault      = DomainRecordPropertiesRegion("default")
	DomainRecordPropertiesRegionEurope       = DomainRecordPropertiesRegion("europe")
	DomainRecordPropertiesRegionUsEast       = DomainRecordPropertiesRegion("us-east")
	DomainRecordPropertiesRegionUsWest       = DomainRecordPropertiesRegion("us-west")
	DomainRecordPropertiesRegionAsiaPacific  = DomainRecordPropertiesRegion("asia-pacific")
	DomainRecordPropertiesRegionOceania      = DomainRecordPropertiesRegion("oceania")
	DomainRecordPropertiesRegionSouthAmerica = DomainRecordPropertiesRegion("south-america")
)

func (DomainRecordPropertiesRegion) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainRecordPropertiesRegion)(nil)).Elem()
}

func (e DomainRecordPropertiesRegion) ToDomainRecordPropertiesRegionOutput() DomainRecordPropertiesRegionOutput {
	return pulumi.ToOutput(e).(DomainRecordPropertiesRegionOutput)
}

func (e DomainRecordPropertiesRegion) ToDomainRecordPropertiesRegionOutputWithContext(ctx context.Context) DomainRecordPropertiesRegionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DomainRecordPropertiesRegionOutput)
}

func (e DomainRecordPropertiesRegion) ToDomainRecordPropertiesRegionPtrOutput() DomainRecordPropertiesRegionPtrOutput {
	return e.ToDomainRecordPropertiesRegionPtrOutputWithContext(context.Background())
}

func (e DomainRecordPropertiesRegion) ToDomainRecordPropertiesRegionPtrOutputWithContext(ctx context.Context) DomainRecordPropertiesRegionPtrOutput {
	return DomainRecordPropertiesRegion(e).ToDomainRecordPropertiesRegionOutputWithContext(ctx).ToDomainRecordPropertiesRegionPtrOutputWithContext(ctx)
}

func (e DomainRecordPropertiesRegion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DomainRecordPropertiesRegion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DomainRecordPropertiesRegion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DomainRecordPropertiesRegion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DomainRecordPropertiesRegionOutput struct{ *pulumi.OutputState }

func (DomainRecordPropertiesRegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainRecordPropertiesRegion)(nil)).Elem()
}

func (o DomainRecordPropertiesRegionOutput) ToDomainRecordPropertiesRegionOutput() DomainRecordPropertiesRegionOutput {
	return o
}

func (o DomainRecordPropertiesRegionOutput) ToDomainRecordPropertiesRegionOutputWithContext(ctx context.Context) DomainRecordPropertiesRegionOutput {
	return o
}

func (o DomainRecordPropertiesRegionOutput) ToDomainRecordPropertiesRegionPtrOutput() DomainRecordPropertiesRegionPtrOutput {
	return o.ToDomainRecordPropertiesRegionPtrOutputWithContext(context.Background())
}

func (o DomainRecordPropertiesRegionOutput) ToDomainRecordPropertiesRegionPtrOutputWithContext(ctx context.Context) DomainRecordPropertiesRegionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainRecordPropertiesRegion) *DomainRecordPropertiesRegion {
		return &v
	}).(DomainRecordPropertiesRegionPtrOutput)
}

func (o DomainRecordPropertiesRegionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DomainRecordPropertiesRegionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainRecordPropertiesRegion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DomainRecordPropertiesRegionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainRecordPropertiesRegionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainRecordPropertiesRegion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DomainRecordPropertiesRegionPtrOutput struct{ *pulumi.OutputState }

func (DomainRecordPropertiesRegionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRecordPropertiesRegion)(nil)).Elem()
}

func (o DomainRecordPropertiesRegionPtrOutput) ToDomainRecordPropertiesRegionPtrOutput() DomainRecordPropertiesRegionPtrOutput {
	return o
}

func (o DomainRecordPropertiesRegionPtrOutput) ToDomainRecordPropertiesRegionPtrOutputWithContext(ctx context.Context) DomainRecordPropertiesRegionPtrOutput {
	return o
}

func (o DomainRecordPropertiesRegionPtrOutput) Elem() DomainRecordPropertiesRegionOutput {
	return o.ApplyT(func(v *DomainRecordPropertiesRegion) DomainRecordPropertiesRegion {
		if v != nil {
			return *v
		}
		var ret DomainRecordPropertiesRegion
		return ret
	}).(DomainRecordPropertiesRegionOutput)
}

func (o DomainRecordPropertiesRegionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainRecordPropertiesRegionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DomainRecordPropertiesRegion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DomainRecordPropertiesRegionInput is an input type that accepts values of the DomainRecordPropertiesRegion enum
// A concrete instance of `DomainRecordPropertiesRegionInput` can be one of the following:
//
//	DomainRecordPropertiesRegionDefault
//	DomainRecordPropertiesRegionEurope
//	DomainRecordPropertiesRegionUsEast
//	DomainRecordPropertiesRegionUsWest
//	DomainRecordPropertiesRegionAsiaPacific
//	DomainRecordPropertiesRegionOceania
//	DomainRecordPropertiesRegionSouthAmerica
type DomainRecordPropertiesRegionInput interface {
	pulumi.Input

	ToDomainRecordPropertiesRegionOutput() DomainRecordPropertiesRegionOutput
	ToDomainRecordPropertiesRegionOutputWithContext(context.Context) DomainRecordPropertiesRegionOutput
}

var domainRecordPropertiesRegionPtrType = reflect.TypeOf((**DomainRecordPropertiesRegion)(nil)).Elem()

type DomainRecordPropertiesRegionPtrInput interface {
	pulumi.Input

	ToDomainRecordPropertiesRegionPtrOutput() DomainRecordPropertiesRegionPtrOutput
	ToDomainRecordPropertiesRegionPtrOutputWithContext(context.Context) DomainRecordPropertiesRegionPtrOutput
}

type domainRecordPropertiesRegionPtr string

func DomainRecordPropertiesRegionPtr(v string) DomainRecordPropertiesRegionPtrInput {
	return (*domainRecordPropertiesRegionPtr)(&v)
}

func (*domainRecordPropertiesRegionPtr) ElementType() reflect.Type {
	return domainRecordPropertiesRegionPtrType
}

func (in *domainRecordPropertiesRegionPtr) ToDomainRecordPropertiesRegionPtrOutput() DomainRecordPropertiesRegionPtrOutput {
	return pulumi.ToOutput(in).(DomainRecordPropertiesRegionPtrOutput)
}

func (in *domainRecordPropertiesRegionPtr) ToDomainRecordPropertiesRegionPtrOutputWithContext(ctx context.Context) DomainRecordPropertiesRegionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DomainRecordPropertiesRegionPtrOutput)
}

type DomainStatus string

const (
	DomainStatusActive     = DomainStatus("ACTIVE")
	DomainStatusSuspended  = DomainStatus("SUSPENDED")
	DomainStatusTerminated = DomainStatus("TERMINATED")
)

type DomainStatusOutput struct{ *pulumi.OutputState }

func (DomainStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainStatus)(nil)).Elem()
}

func (o DomainStatusOutput) ToDomainStatusOutput() DomainStatusOutput {
	return o
}

func (o DomainStatusOutput) ToDomainStatusOutputWithContext(ctx context.Context) DomainStatusOutput {
	return o
}

func (o DomainStatusOutput) ToDomainStatusPtrOutput() DomainStatusPtrOutput {
	return o.ToDomainStatusPtrOutputWithContext(context.Background())
}

func (o DomainStatusOutput) ToDomainStatusPtrOutputWithContext(ctx context.Context) DomainStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainStatus) *DomainStatus {
		return &v
	}).(DomainStatusPtrOutput)
}

func (o DomainStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DomainStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DomainStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DomainStatusPtrOutput struct{ *pulumi.OutputState }

func (DomainStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainStatus)(nil)).Elem()
}

func (o DomainStatusPtrOutput) ToDomainStatusPtrOutput() DomainStatusPtrOutput {
	return o
}

func (o DomainStatusPtrOutput) ToDomainStatusPtrOutputWithContext(ctx context.Context) DomainStatusPtrOutput {
	return o
}

func (o DomainStatusPtrOutput) Elem() DomainStatusOutput {
	return o.ApplyT(func(v *DomainStatus) DomainStatus {
		if v != nil {
			return *v
		}
		var ret DomainStatus
		return ret
	}).(DomainStatusOutput)
}

func (o DomainStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DomainStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// The type of query the analytics are for
type DomainanalyticsQueriesItemPropertiesType string

const (
	DomainanalyticsQueriesItemPropertiesTypeGeoProximity = DomainanalyticsQueriesItemPropertiesType("geo_proximity")
	DomainanalyticsQueriesItemPropertiesTypeStandard     = DomainanalyticsQueriesItemPropertiesType("standard")
	DomainanalyticsQueriesItemPropertiesTypeGeoFilter    = DomainanalyticsQueriesItemPropertiesType("geo_filter")
)

type DomainanalyticsQueriesItemPropertiesTypeOutput struct{ *pulumi.OutputState }

func (DomainanalyticsQueriesItemPropertiesTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainanalyticsQueriesItemPropertiesType)(nil)).Elem()
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToDomainanalyticsQueriesItemPropertiesTypeOutput() DomainanalyticsQueriesItemPropertiesTypeOutput {
	return o
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToDomainanalyticsQueriesItemPropertiesTypeOutputWithContext(ctx context.Context) DomainanalyticsQueriesItemPropertiesTypeOutput {
	return o
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToDomainanalyticsQueriesItemPropertiesTypePtrOutput() DomainanalyticsQueriesItemPropertiesTypePtrOutput {
	return o.ToDomainanalyticsQueriesItemPropertiesTypePtrOutputWithContext(context.Background())
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToDomainanalyticsQueriesItemPropertiesTypePtrOutputWithContext(ctx context.Context) DomainanalyticsQueriesItemPropertiesTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainanalyticsQueriesItemPropertiesType) *DomainanalyticsQueriesItemPropertiesType {
		return &v
	}).(DomainanalyticsQueriesItemPropertiesTypePtrOutput)
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainanalyticsQueriesItemPropertiesType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainanalyticsQueriesItemPropertiesTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainanalyticsQueriesItemPropertiesType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DomainanalyticsQueriesItemPropertiesTypePtrOutput struct{ *pulumi.OutputState }

func (DomainanalyticsQueriesItemPropertiesTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainanalyticsQueriesItemPropertiesType)(nil)).Elem()
}

func (o DomainanalyticsQueriesItemPropertiesTypePtrOutput) ToDomainanalyticsQueriesItemPropertiesTypePtrOutput() DomainanalyticsQueriesItemPropertiesTypePtrOutput {
	return o
}

func (o DomainanalyticsQueriesItemPropertiesTypePtrOutput) ToDomainanalyticsQueriesItemPropertiesTypePtrOutputWithContext(ctx context.Context) DomainanalyticsQueriesItemPropertiesTypePtrOutput {
	return o
}

func (o DomainanalyticsQueriesItemPropertiesTypePtrOutput) Elem() DomainanalyticsQueriesItemPropertiesTypeOutput {
	return o.ApplyT(func(v *DomainanalyticsQueriesItemPropertiesType) DomainanalyticsQueriesItemPropertiesType {
		if v != nil {
			return *v
		}
		var ret DomainanalyticsQueriesItemPropertiesType
		return ret
	}).(DomainanalyticsQueriesItemPropertiesTypeOutput)
}

func (o DomainanalyticsQueriesItemPropertiesTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainanalyticsQueriesItemPropertiesTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DomainanalyticsQueriesItemPropertiesType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SimpleDomainStatus string

const (
	SimpleDomainStatusActive     = SimpleDomainStatus("ACTIVE")
	SimpleDomainStatusSuspended  = SimpleDomainStatus("SUSPENDED")
	SimpleDomainStatusTerminated = SimpleDomainStatus("TERMINATED")
)

type SimpleDomainStatusOutput struct{ *pulumi.OutputState }

func (SimpleDomainStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleDomainStatus)(nil)).Elem()
}

func (o SimpleDomainStatusOutput) ToSimpleDomainStatusOutput() SimpleDomainStatusOutput {
	return o
}

func (o SimpleDomainStatusOutput) ToSimpleDomainStatusOutputWithContext(ctx context.Context) SimpleDomainStatusOutput {
	return o
}

func (o SimpleDomainStatusOutput) ToSimpleDomainStatusPtrOutput() SimpleDomainStatusPtrOutput {
	return o.ToSimpleDomainStatusPtrOutputWithContext(context.Background())
}

func (o SimpleDomainStatusOutput) ToSimpleDomainStatusPtrOutputWithContext(ctx context.Context) SimpleDomainStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SimpleDomainStatus) *SimpleDomainStatus {
		return &v
	}).(SimpleDomainStatusPtrOutput)
}

func (o SimpleDomainStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SimpleDomainStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SimpleDomainStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SimpleDomainStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SimpleDomainStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SimpleDomainStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SimpleDomainStatusPtrOutput struct{ *pulumi.OutputState }

func (SimpleDomainStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimpleDomainStatus)(nil)).Elem()
}

func (o SimpleDomainStatusPtrOutput) ToSimpleDomainStatusPtrOutput() SimpleDomainStatusPtrOutput {
	return o
}

func (o SimpleDomainStatusPtrOutput) ToSimpleDomainStatusPtrOutputWithContext(ctx context.Context) SimpleDomainStatusPtrOutput {
	return o
}

func (o SimpleDomainStatusPtrOutput) Elem() SimpleDomainStatusOutput {
	return o.ApplyT(func(v *SimpleDomainStatus) SimpleDomainStatus {
		if v != nil {
			return *v
		}
		var ret SimpleDomainStatus
		return ret
	}).(SimpleDomainStatusOutput)
}

func (o SimpleDomainStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SimpleDomainStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SimpleDomainStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRecordPropertiesRegionInput)(nil)).Elem(), DomainRecordPropertiesRegion("default"))
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRecordPropertiesRegionPtrInput)(nil)).Elem(), DomainRecordPropertiesRegion("default"))
	pulumi.RegisterOutputType(DomainRecordPropertiesRegionOutput{})
	pulumi.RegisterOutputType(DomainRecordPropertiesRegionPtrOutput{})
	pulumi.RegisterOutputType(DomainStatusOutput{})
	pulumi.RegisterOutputType(DomainStatusPtrOutput{})
	pulumi.RegisterOutputType(DomainanalyticsQueriesItemPropertiesTypeOutput{})
	pulumi.RegisterOutputType(DomainanalyticsQueriesItemPropertiesTypePtrOutput{})
	pulumi.RegisterOutputType(SimpleDomainStatusOutput{})
	pulumi.RegisterOutputType(SimpleDomainStatusPtrOutput{})
}
