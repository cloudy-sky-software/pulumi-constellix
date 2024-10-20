// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getGeoProximityLocation(args: GetGeoProximityLocationArgs, opts?: pulumi.InvokeOptions): Promise<outputs.geoproximities.GetGeoProximityLocationProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:geoproximities:getGeoProximityLocation", {
        "id": args.id,
    }, opts);
}

export interface GetGeoProximityLocationArgs {
    /**
     * The ID of the Geo Proximity
     */
    id: string;
}
export function getGeoProximityLocationOutput(args: GetGeoProximityLocationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.geoproximities.GetGeoProximityLocationProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:geoproximities:getGeoProximityLocation", {
        "id": args.id,
    }, opts);
}

export interface GetGeoProximityLocationOutputArgs {
    /**
     * The ID of the Geo Proximity
     */
    id: pulumi.Input<string>;
}
