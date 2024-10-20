// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getPing(args?: GetPingArgs, opts?: pulumi.InvokeOptions): Promise<outputs.ping.Ping> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:ping:getPing", {
    }, opts);
}

export interface GetPingArgs {
}
export function getPingOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.ping.Ping> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:ping:getPing", {
    }, opts);
}

