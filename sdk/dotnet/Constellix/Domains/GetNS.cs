// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Domains
{
    public static class GetNS
    {
        public static Task<Outputs.GetNS> InvokeAsync(GetNSArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetNS>("constellix:domains:getNS", args ?? new GetNSArgs(), options.WithDefaults());

        public static Output<Outputs.GetNS> Invoke(GetNSInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetNS>("constellix:domains:getNS", args ?? new GetNSInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNSArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the record
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNSArgs()
        {
        }
        public static new GetNSArgs Empty => new GetNSArgs();
    }

    public sealed class GetNSInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the record
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNSInvokeArgs()
        {
        }
        public static new GetNSInvokeArgs Empty => new GetNSInvokeArgs();
    }
}
