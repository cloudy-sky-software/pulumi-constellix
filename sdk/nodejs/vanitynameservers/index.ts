// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetVanityNameserverArgs, GetVanityNameserverOutputArgs } from "./getVanityNameserver";
export const getVanityNameserver: typeof import("./getVanityNameserver").getVanityNameserver = null as any;
export const getVanityNameserverOutput: typeof import("./getVanityNameserver").getVanityNameserverOutput = null as any;
utilities.lazyLoad(exports, ["getVanityNameserver","getVanityNameserverOutput"], () => require("./getVanityNameserver"));

export { ListVanityNameserversArgs } from "./listVanityNameservers";
export const listVanityNameservers: typeof import("./listVanityNameservers").listVanityNameservers = null as any;
export const listVanityNameserversOutput: typeof import("./listVanityNameservers").listVanityNameserversOutput = null as any;
utilities.lazyLoad(exports, ["listVanityNameservers","listVanityNameserversOutput"], () => require("./listVanityNameservers"));

export { VanityNameserverArgs } from "./vanityNameserver";
export type VanityNameserver = import("./vanityNameserver").VanityNameserver;
export const VanityNameserver: typeof import("./vanityNameserver").VanityNameserver = null as any;
utilities.lazyLoad(exports, ["VanityNameserver"], () => require("./vanityNameserver"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "constellix:vanitynameservers:VanityNameserver":
                return new VanityNameserver(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("constellix", "vanitynameservers", _module)
