// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Domains.Outputs
{

    [OutputType]
    public sealed class Domainsnapshot
    {
        public readonly Outputs.SimpleDomain? Domain;
        /// <summary>
        /// The name of the domain
        /// </summary>
        public readonly string? Name;
        public readonly string? UpdatedAt;
        /// <summary>
        /// The version of the domain resource
        /// </summary>
        public readonly int? Version;

        [OutputConstructor]
        private Domainsnapshot(
            Outputs.SimpleDomain? domain,

            string? name,

            string? updatedAt,

            int? version)
        {
            Domain = domain;
            Name = name;
            UpdatedAt = updatedAt;
            Version = version;
        }
    }
}
