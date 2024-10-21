// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Rp extends pulumi.CustomResource {
    /**
     * Get an existing Rp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Rp {
        return new Rp(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'constellix:templates:Rp';

    /**
     * Returns true if the given object is an instance of Rp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rp.__pulumiType;
    }

    /**
     * Contact lists to be notified if a failover happens in a failover mode.
     */
    public readonly contacts!: pulumi.Output<number[] | undefined>;
    public /*out*/ readonly data!: pulumi.Output<outputs.templates.DataProperties | undefined>;
    /**
     * Whether the record is enabled
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
     */
    public readonly geoFailover!: pulumi.Output<boolean | undefined>;
    /**
     * The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
     */
    public readonly geoproximity!: pulumi.Output<number | undefined>;
    /**
     * The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
     */
    public readonly ipfilter!: pulumi.Output<number | undefined>;
    /**
     * If the requesting IP matches the IP filter, don't return a response
     */
    public readonly ipfilterDrop!: pulumi.Output<boolean | undefined>;
    /**
     * The current mode for this record
     */
    public readonly mode!: pulumi.Output<enums.templates.RpPropertiesMode | undefined>;
    /**
     * The name for the record
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A description of the record. It must be 512 characters or less.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * Optional region for this record. Will default to 'default'.
     */
    public readonly region!: pulumi.Output<enums.templates.RecordCreateDetailsRegion | undefined>;
    /**
     * Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
     */
    public readonly skipLookup!: pulumi.Output<boolean | undefined>;
    /**
     * How long DNS servers should cache the record for
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The type of record
     */
    public readonly type!: pulumi.Output<enums.templates.RpPropertiesType | undefined>;
    /**
     * Standard record mode
     */
    public readonly value!: pulumi.Output<outputs.templates.ValueRpValueItemProperties[] | undefined>;

    /**
     * Create a Rp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RpArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["contacts"] = args ? args.contacts : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["geoFailover"] = args ? args.geoFailover : undefined;
            resourceInputs["geoproximity"] = args ? args.geoproximity : undefined;
            resourceInputs["ipfilter"] = args ? args.ipfilter : undefined;
            resourceInputs["ipfilterDrop"] = args ? args.ipfilterDrop : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["skipLookup"] = args ? args.skipLookup : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["data"] = undefined /*out*/;
        } else {
            resourceInputs["contacts"] = undefined /*out*/;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["geoFailover"] = undefined /*out*/;
            resourceInputs["geoproximity"] = undefined /*out*/;
            resourceInputs["ipfilter"] = undefined /*out*/;
            resourceInputs["ipfilterDrop"] = undefined /*out*/;
            resourceInputs["mode"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notes"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["skipLookup"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Rp resource.
 */
export interface RpArgs {
    /**
     * Contact lists to be notified if a failover happens in a failover mode.
     */
    contacts?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Whether the record is enabled
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
     */
    geoFailover?: pulumi.Input<boolean>;
    /**
     * The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
     */
    geoproximity?: pulumi.Input<number>;
    /**
     * The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
     */
    ipfilter?: pulumi.Input<number>;
    /**
     * If the requesting IP matches the IP filter, don't return a response
     */
    ipfilterDrop?: pulumi.Input<boolean>;
    /**
     * The current mode for this record
     */
    mode?: pulumi.Input<enums.templates.RpPropertiesMode>;
    /**
     * The name for the record
     */
    name?: pulumi.Input<string>;
    /**
     * A description of the record. It must be 512 characters or less.
     */
    notes?: pulumi.Input<string>;
    /**
     * Optional region for this record. Will default to 'default'.
     */
    region?: pulumi.Input<enums.templates.RecordCreateDetailsRegion>;
    /**
     * Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
     */
    skipLookup?: pulumi.Input<boolean>;
    /**
     * The ID of the template object
     */
    templateId?: pulumi.Input<string>;
    /**
     * How long DNS servers should cache the record for
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of record
     */
    type?: pulumi.Input<enums.templates.RpPropertiesType>;
    /**
     * Standard record mode
     */
    value?: pulumi.Input<pulumi.Input<inputs.templates.ValueRpValueItemPropertiesArgs>[]>;
}
