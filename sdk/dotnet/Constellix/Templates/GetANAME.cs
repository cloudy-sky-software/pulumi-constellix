// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Templates
{
    public static class GetANAME
    {
        public static Task<Outputs.GetANAME> InvokeAsync(GetANAMEArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetANAME>("constellix:templates:getANAME", args ?? new GetANAMEArgs(), options.WithDefaults());

        public static Output<Outputs.GetANAME> Invoke(GetANAMEInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetANAME>("constellix:templates:getANAME", args ?? new GetANAMEInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetANAMEArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the record
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The ID of the template object
        /// </summary>
        [Input("templateId", required: true)]
        public string TemplateId { get; set; } = null!;

        public GetANAMEArgs()
        {
        }
        public static new GetANAMEArgs Empty => new GetANAMEArgs();
    }

    public sealed class GetANAMEInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the record
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The ID of the template object
        /// </summary>
        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        public GetANAMEInvokeArgs()
        {
        }
        public static new GetANAMEInvokeArgs Empty => new GetANAMEInvokeArgs();
    }
}
