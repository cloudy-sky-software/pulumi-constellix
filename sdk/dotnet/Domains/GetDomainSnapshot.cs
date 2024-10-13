// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Domains
{
    public static class GetDomainSnapshot
    {
        public static Task<Outputs.GetDomainSnapshotProperties> InvokeAsync(GetDomainSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetDomainSnapshotProperties>("constellix:domains:getDomainSnapshot", args ?? new GetDomainSnapshotArgs(), options.WithDefaults());

        public static Output<Outputs.GetDomainSnapshotProperties> Invoke(GetDomainSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetDomainSnapshotProperties>("constellix:domains:getDomainSnapshot", args ?? new GetDomainSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainSnapshotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// The version number of the snapshot
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetDomainSnapshotArgs()
        {
        }
        public static new GetDomainSnapshotArgs Empty => new GetDomainSnapshotArgs();
    }

    public sealed class GetDomainSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// The version number of the snapshot
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GetDomainSnapshotInvokeArgs()
        {
        }
        public static new GetDomainSnapshotInvokeArgs Empty => new GetDomainSnapshotInvokeArgs();
    }
}
