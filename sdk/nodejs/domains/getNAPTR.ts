// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getNAPTR(args: GetNAPTRArgs, opts?: pulumi.InvokeOptions): Promise<outputs.domains.GetNAPTR> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:domains:getNAPTR", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetNAPTRArgs {
    /**
     * The ID of the domain object
     */
    domainId: string;
    /**
     * The ID of the record
     */
    id: string;
}
export function getNAPTROutput(args: GetNAPTROutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.domains.GetNAPTR> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:domains:getNAPTR", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetNAPTROutputArgs {
    /**
     * The ID of the domain object
     */
    domainId: pulumi.Input<string>;
    /**
     * The ID of the record
     */
    id: pulumi.Input<string>;
}
