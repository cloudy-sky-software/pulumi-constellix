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
    public static class GetAAAA
    {
        public static Task<Outputs.GetAAAA> InvokeAsync(GetAAAAArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetAAAA>("constellix:templates:getAAAA", args ?? new GetAAAAArgs(), options.WithDefaults());

        public static Output<Outputs.GetAAAA> Invoke(GetAAAAInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetAAAA>("constellix:templates:getAAAA", args ?? new GetAAAAInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAAAAArgs : global::Pulumi.InvokeArgs
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

        public GetAAAAArgs()
        {
        }
        public static new GetAAAAArgs Empty => new GetAAAAArgs();
    }

    public sealed class GetAAAAInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAAAAInvokeArgs()
        {
        }
        public static new GetAAAAInvokeArgs Empty => new GetAAAAInvokeArgs();
    }
}
