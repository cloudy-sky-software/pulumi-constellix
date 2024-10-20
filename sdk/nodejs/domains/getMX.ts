// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getMX(args: GetMXArgs, opts?: pulumi.InvokeOptions): Promise<outputs.domains.GetMX> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:domains:getMX", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetMXArgs {
    /**
     * The ID of the domain object
     */
    domainId: string;
    /**
     * The ID of the record
     */
    id: string;
}
export function getMXOutput(args: GetMXOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.domains.GetMX> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:domains:getMX", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetMXOutputArgs {
    /**
     * The ID of the domain object
     */
    domainId: pulumi.Input<string>;
    /**
     * The ID of the record
     */
    id: pulumi.Input<string>;
}
