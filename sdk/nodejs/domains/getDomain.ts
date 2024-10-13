// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getDomain(args: GetDomainArgs, opts?: pulumi.InvokeOptions): Promise<outputs.domains.GetDomainProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:domains:getDomain", {
        "id": args.id,
    }, opts);
}

export interface GetDomainArgs {
    /**
     * The ID of the domain object
     */
    id: string;
}
export function getDomainOutput(args: GetDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.domains.GetDomainProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:domains:getDomain", {
        "id": args.id,
    }, opts);
}

export interface GetDomainOutputArgs {
    /**
     * The ID of the domain object
     */
    id: pulumi.Input<string>;
}
