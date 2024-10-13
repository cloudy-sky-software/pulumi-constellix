// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Tags
{
    public static class GetTag
    {
        public static Task<Outputs.GetTagProperties> InvokeAsync(GetTagArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetTagProperties>("constellix:tags:getTag", args ?? new GetTagArgs(), options.WithDefaults());

        public static Output<Outputs.GetTagProperties> Invoke(GetTagInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetTagProperties>("constellix:tags:getTag", args ?? new GetTagInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Tag
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTagArgs()
        {
        }
        public static new GetTagArgs Empty => new GetTagArgs();
    }

    public sealed class GetTagInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Tag
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTagInvokeArgs()
        {
        }
        public static new GetTagInvokeArgs Empty => new GetTagInvokeArgs();
    }
}
