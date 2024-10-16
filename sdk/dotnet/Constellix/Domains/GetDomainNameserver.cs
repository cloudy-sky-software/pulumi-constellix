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
    public static class GetDomainNameserver
    {
        public static Task<Outputs.GetDomainNameserverProperties> InvokeAsync(GetDomainNameserverArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetDomainNameserverProperties>("constellix:domains:getDomainNameserver", args ?? new GetDomainNameserverArgs(), options.WithDefaults());

        public static Output<Outputs.GetDomainNameserverProperties> Invoke(GetDomainNameserverInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDomainNameserverProperties>("constellix:domains:getDomainNameserver", args ?? new GetDomainNameserverInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainNameserverArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDomainNameserverArgs()
        {
        }
        public static new GetDomainNameserverArgs Empty => new GetDomainNameserverArgs();
    }

    public sealed class GetDomainNameserverInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDomainNameserverInvokeArgs()
        {
        }
        public static new GetDomainNameserverInvokeArgs Empty => new GetDomainNameserverInvokeArgs();
    }
}
