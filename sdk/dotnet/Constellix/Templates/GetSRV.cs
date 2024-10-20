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
    public static class GetSRV
    {
        public static Task<Outputs.GetSRV> InvokeAsync(GetSRVArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetSRV>("constellix:templates:getSRV", args ?? new GetSRVArgs(), options.WithDefaults());

        public static Output<Outputs.GetSRV> Invoke(GetSRVInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetSRV>("constellix:templates:getSRV", args ?? new GetSRVInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSRVArgs : global::Pulumi.InvokeArgs
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

        public GetSRVArgs()
        {
        }
        public static new GetSRVArgs Empty => new GetSRVArgs();
    }

    public sealed class GetSRVInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSRVInvokeArgs()
        {
        }
        public static new GetSRVInvokeArgs Empty => new GetSRVInvokeArgs();
    }
}
