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
    public static class ListDomains
    {
        public static Task<Outputs.ListDomainsProperties> InvokeAsync(ListDomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListDomainsProperties>("constellix:domains:listDomains", args ?? new ListDomainsArgs(), options.WithDefaults());

        public static Output<Outputs.ListDomainsProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListDomainsProperties>("constellix:domains:listDomains", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListDomainsArgs : global::Pulumi.InvokeArgs
    {
        public ListDomainsArgs()
        {
        }
        public static new ListDomainsArgs Empty => new ListDomainsArgs();
    }
}
