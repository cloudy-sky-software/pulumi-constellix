// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetTemplateArgs, GetTemplateOutputArgs } from "./getTemplate";
export const getTemplate: typeof import("./getTemplate").getTemplate = null as any;
export const getTemplateOutput: typeof import("./getTemplate").getTemplateOutput = null as any;
utilities.lazyLoad(exports, ["getTemplate","getTemplateOutput"], () => require("./getTemplate"));

export { GetTemplateRecordArgs, GetTemplateRecordOutputArgs } from "./getTemplateRecord";
export const getTemplateRecord: typeof import("./getTemplateRecord").getTemplateRecord = null as any;
export const getTemplateRecordOutput: typeof import("./getTemplateRecord").getTemplateRecordOutput = null as any;
utilities.lazyLoad(exports, ["getTemplateRecord","getTemplateRecordOutput"], () => require("./getTemplateRecord"));

export { ListTemplateRecordsArgs, ListTemplateRecordsOutputArgs } from "./listTemplateRecords";
export const listTemplateRecords: typeof import("./listTemplateRecords").listTemplateRecords = null as any;
export const listTemplateRecordsOutput: typeof import("./listTemplateRecords").listTemplateRecordsOutput = null as any;
utilities.lazyLoad(exports, ["listTemplateRecords","listTemplateRecordsOutput"], () => require("./listTemplateRecords"));

export { ListTemplatesArgs } from "./listTemplates";
export const listTemplates: typeof import("./listTemplates").listTemplates = null as any;
export const listTemplatesOutput: typeof import("./listTemplates").listTemplatesOutput = null as any;
utilities.lazyLoad(exports, ["listTemplates","listTemplatesOutput"], () => require("./listTemplates"));

export { TemplateArgs } from "./template";
export type Template = import("./template").Template;
export const Template: typeof import("./template").Template = null as any;
utilities.lazyLoad(exports, ["Template"], () => require("./template"));

export { TemplateRecordArgs } from "./templateRecord";
export type TemplateRecord = import("./templateRecord").TemplateRecord;
export const TemplateRecord: typeof import("./templateRecord").TemplateRecord = null as any;
utilities.lazyLoad(exports, ["TemplateRecord"], () => require("./templateRecord"));


// Export enums:
export * from "../types/enums/templates";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "constellix:templates:Template":
                return new Template(name, <any>undefined, { urn })
            case "constellix:templates:TemplateRecord":
                return new TemplateRecord(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("constellix", "templates", _module)
