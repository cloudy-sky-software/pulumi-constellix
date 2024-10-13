// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'constellix:domains:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * Contactlists to be notified if the domain is updated
     */
    public readonly contacts!: pulumi.Output<number[] | undefined>;
    public /*out*/ readonly data!: pulumi.Output<outputs.domains.DataProperties | undefined>;
    /**
     * Is the domain enabled
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Is GeoIP functionality enabled for the domain
     */
    public readonly geoip!: pulumi.Output<boolean | undefined>;
    /**
     * Is Global Traffic Director enabled for the domain
     */
    public readonly gtd!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the domain
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A note for the domain
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * The SOA details for the domain
     */
    public readonly soa!: pulumi.Output<outputs.domains.SoaProperties | undefined>;
    /**
     * The numeric IDs of tags you want to apply to this domain
     */
    public readonly tags!: pulumi.Output<number[] | undefined>;
    /**
     * The template to use for creating this domain. It will be linked to this template so any changes made to the template will apply to this domain.
     */
    public readonly template!: pulumi.Output<number | undefined>;
    /**
     * The vanity nameserver to use for this domain.
     */
    public readonly vanityNameserver!: pulumi.Output<number | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DomainArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["contacts"] = args ? args.contacts : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["geoip"] = args ? args.geoip : undefined;
            resourceInputs["gtd"] = args ? args.gtd : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["soa"] = args ? args.soa : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["vanityNameserver"] = args ? args.vanityNameserver : undefined;
            resourceInputs["data"] = undefined /*out*/;
        } else {
            resourceInputs["contacts"] = undefined /*out*/;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["geoip"] = undefined /*out*/;
            resourceInputs["gtd"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["note"] = undefined /*out*/;
            resourceInputs["soa"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["template"] = undefined /*out*/;
            resourceInputs["vanityNameserver"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * Contactlists to be notified if the domain is updated
     */
    contacts?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Is the domain enabled
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Is GeoIP functionality enabled for the domain
     */
    geoip?: pulumi.Input<boolean>;
    /**
     * Is Global Traffic Director enabled for the domain
     */
    gtd?: pulumi.Input<boolean>;
    /**
     * The name of the domain
     */
    name?: pulumi.Input<string>;
    /**
     * A note for the domain
     */
    note?: pulumi.Input<string>;
    /**
     * The SOA details for the domain
     */
    soa?: pulumi.Input<inputs.domains.SoaPropertiesArgs>;
    /**
     * The numeric IDs of tags you want to apply to this domain
     */
    tags?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The template to use for creating this domain. It will be linked to this template so any changes made to the template will apply to this domain.
     */
    template?: pulumi.Input<number>;
    /**
     * The vanity nameserver to use for this domain.
     */
    vanityNameserver?: pulumi.Input<number>;
}
