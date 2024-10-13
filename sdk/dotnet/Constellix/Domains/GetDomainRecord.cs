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
    public static class GetDomainRecord
    {
        public static Task<Outputs.GetDomainRecord> InvokeAsync(GetDomainRecordArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetDomainRecord>("constellix:domains:getDomainRecord", args ?? new GetDomainRecordArgs(), options.WithDefaults());

        public static Output<Outputs.GetDomainRecord> Invoke(GetDomainRecordInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDomainRecord>("constellix:domains:getDomainRecord", args ?? new GetDomainRecordInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainRecordArgs : global::Pulumi.InvokeArgs
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

        public GetDomainRecordArgs()
        {
        }
        public static new GetDomainRecordArgs Empty => new GetDomainRecordArgs();
    }

    public sealed class GetDomainRecordInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetDomainRecordInvokeArgs()
        {
        }
        public static new GetDomainRecordInvokeArgs Empty => new GetDomainRecordInvokeArgs();
    }
}
