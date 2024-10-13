// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Domains.Outputs
{

    /// <summary>
    /// The SOA details for the domain
    /// </summary>
    [OutputType]
    public sealed class SoaProperties
    {
        /// <summary>
        /// Email of the administrator for the domain. @ should be replaced with .
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// Number of seconds after which secondary nameservers should stop responding to queries, if the master does not respond
        /// </summary>
        public readonly int? Expire;
        /// <summary>
        /// How long NXDOMAIN responses should be cached for
        /// </summary>
        public readonly int? NegativeCache;
        /// <summary>
        /// Primary master nameserver for the domain
        /// </summary>
        public readonly string? PrimaryNameserver;
        /// <summary>
        /// The interval for secondary nameservers should query for the SOA record
        /// </summary>
        public readonly int? Refresh;
        /// <summary>
        /// The number of seconds after which secondary servers should retry to request the serial number if the master does not respond
        /// </summary>
        public readonly int? Retry;
        /// <summary>
        /// The Time To Live (TTL) in seconds for the SOA record
        /// </summary>
        public readonly int? Ttl;

        [OutputConstructor]
        private SoaProperties(
            string? email,

            int? expire,

            int? negativeCache,

            string? primaryNameserver,

            int? refresh,

            int? retry,

            int? ttl)
        {
            Email = email;
            Expire = expire;
            NegativeCache = negativeCache;
            PrimaryNameserver = primaryNameserver;
            Refresh = refresh;
            Retry = retry;
            Ttl = ttl;
        }
    }
}
