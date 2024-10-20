// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getNS(args: GetNSArgs, opts?: pulumi.InvokeOptions): Promise<outputs.domains.GetNS> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:domains:getNS", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetNSArgs {
    /**
     * The ID of the domain object
     */
    domainId: string;
    /**
     * The ID of the record
     */
    id: string;
}
export function getNSOutput(args: GetNSOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.domains.GetNS> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:domains:getNS", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetNSOutputArgs {
    /**
     * The ID of the domain object
     */
    domainId: pulumi.Input<string>;
    /**
     * The ID of the record
     */
    id: pulumi.Input<string>;
}
