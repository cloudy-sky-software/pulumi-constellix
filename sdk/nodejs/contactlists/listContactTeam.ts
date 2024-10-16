// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listContactTeam(args: ListContactTeamArgs, opts?: pulumi.InvokeOptions): Promise<outputs.contactlists.ListContactTeamProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("constellix:contactlists:listContactTeam", {
        "contactlistId": args.contactlistId,
        "id": args.id,
    }, opts);
}

export interface ListContactTeamArgs {
    /**
     * The ID of the Contact List
     */
    contactlistId: string;
    /**
     * The ID of the MS Teams Webhook
     */
    id: string;
}
export function listContactTeamOutput(args: ListContactTeamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.contactlists.ListContactTeamProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("constellix:contactlists:listContactTeam", {
        "contactlistId": args.contactlistId,
        "id": args.id,
    }, opts);
}

export interface ListContactTeamOutputArgs {
    /**
     * The ID of the Contact List
     */
    contactlistId: pulumi.Input<string>;
    /**
     * The ID of the MS Teams Webhook
     */
    id: pulumi.Input<string>;
}
