// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listDomainHistory(args: ListDomainHistoryArgs, opts?: pulumi.InvokeOptions): Promise<outputs.domains.ListDomainHistoryProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:domains:listDomainHistory", {
        "domainId": args.domainId,
    }, opts);
}

export interface ListDomainHistoryArgs {
    /**
     * The ID of the domain object
     */
    domainId: string;
}
export function listDomainHistoryOutput(args: ListDomainHistoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.domains.ListDomainHistoryProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:domains:listDomainHistory", {
        "domainId": args.domainId,
    }, opts);
}

export interface ListDomainHistoryOutputArgs {
    /**
     * The ID of the domain object
     */
    domainId: pulumi.Input<string>;
}
