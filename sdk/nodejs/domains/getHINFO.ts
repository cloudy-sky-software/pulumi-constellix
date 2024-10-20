// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getHINFO(args: GetHINFOArgs, opts?: pulumi.InvokeOptions): Promise<outputs.domains.GetHINFO> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:domains:getHINFO", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetHINFOArgs {
    /**
     * The ID of the domain object
     */
    domainId: string;
    /**
     * The ID of the record
     */
    id: string;
}
export function getHINFOOutput(args: GetHINFOOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.domains.GetHINFO> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:domains:getHINFO", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetHINFOOutputArgs {
    /**
     * The ID of the domain object
     */
    domainId: pulumi.Input<string>;
    /**
     * The ID of the record
     */
    id: pulumi.Input<string>;
}
