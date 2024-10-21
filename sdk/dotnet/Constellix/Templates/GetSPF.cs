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
    public static class GetSPF
    {
        public static Task<Outputs.GetSPF> InvokeAsync(GetSPFArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetSPF>("constellix:templates:getSPF", args ?? new GetSPFArgs(), options.WithDefaults());

        public static Output<Outputs.GetSPF> Invoke(GetSPFInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetSPF>("constellix:templates:getSPF", args ?? new GetSPFInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSPFArgs : global::Pulumi.InvokeArgs
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

        public GetSPFArgs()
        {
        }
        public static new GetSPFArgs Empty => new GetSPFArgs();
    }

    public sealed class GetSPFInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSPFInvokeArgs()
        {
        }
        public static new GetSPFInvokeArgs Empty => new GetSPFInvokeArgs();
    }
}
