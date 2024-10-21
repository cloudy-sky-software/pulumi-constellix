// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getCERT(args: GetCERTArgs, opts?: pulumi.InvokeOptions): Promise<outputs.templates.GetCERT> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:templates:getCERT", {
        "id": args.id,
        "templateId": args.templateId,
    }, opts);
}

export interface GetCERTArgs {
    /**
     * The ID of the record
     */
    id: string;
    /**
     * The ID of the template object
     */
    templateId: string;
}
export function getCERTOutput(args: GetCERTOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.templates.GetCERT> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:templates:getCERT", {
        "id": args.id,
        "templateId": args.templateId,
    }, opts);
}

export interface GetCERTOutputArgs {
    /**
     * The ID of the record
     */
    id: pulumi.Input<string>;
    /**
     * The ID of the template object
     */
    templateId: pulumi.Input<string>;
}
