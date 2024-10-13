// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Pools
{
    public static class GetPool
    {
        public static Task<Outputs.GetPoolProperties> InvokeAsync(GetPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetPoolProperties>("constellix:pools:getPool", args ?? new GetPoolArgs(), options.WithDefaults());

        public static Output<Outputs.GetPoolProperties> Invoke(GetPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetPoolProperties>("constellix:pools:getPool", args ?? new GetPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the pool
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The type of pool
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetPoolArgs()
        {
        }
        public static new GetPoolArgs Empty => new GetPoolArgs();
    }

    public sealed class GetPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the pool
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The type of pool
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetPoolInvokeArgs()
        {
        }
        public static new GetPoolInvokeArgs Empty => new GetPoolInvokeArgs();
    }
}
