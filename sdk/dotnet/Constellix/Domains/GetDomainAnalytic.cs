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
    public static class GetDomainAnalytic
    {
        public static Task<Outputs.GetDomainAnalyticProperties> InvokeAsync(GetDomainAnalyticArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetDomainAnalyticProperties>("constellix:domains:getDomainAnalytic", args ?? new GetDomainAnalyticArgs(), options.WithDefaults());

        public static Output<Outputs.GetDomainAnalyticProperties> Invoke(GetDomainAnalyticInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDomainAnalyticProperties>("constellix:domains:getDomainAnalytic", args ?? new GetDomainAnalyticInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainAnalyticArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDomainAnalyticArgs()
        {
        }
        public static new GetDomainAnalyticArgs Empty => new GetDomainAnalyticArgs();
    }

    public sealed class GetDomainAnalyticInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDomainAnalyticInvokeArgs()
        {
        }
        public static new GetDomainAnalyticInvokeArgs Empty => new GetDomainAnalyticInvokeArgs();
    }
}
