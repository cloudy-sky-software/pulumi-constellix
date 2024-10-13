// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Domains.Outputs
{

    [OutputType]
    public sealed class SimpleDomain
    {
        public readonly string? CreatedAt;
        /// <summary>
        /// Is the domain enabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Is GeoIP functionality enabled for the domain
        /// </summary>
        public readonly bool? Geoip;
        /// <summary>
        /// Is Global Traffic Director enabled for the domain
        /// </summary>
        public readonly bool? Gtd;
        /// <summary>
        /// A unique numeric ID for this domain
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Links for domain objects
        /// </summary>
        public readonly Outputs.SimpleDomainLinksProperties? Links;
        /// <summary>
        /// The name of the domain
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A note for the domain
        /// </summary>
        public readonly string? Note;
        public readonly CloudySkySoftware.Pulumi.Constellix.Domains.SimpleDomainStatus? Status;
        /// <summary>
        /// An array of tags for this domain.
        /// </summary>
        public readonly ImmutableArray<Outputs.Tag> Tags;
        public readonly string? UpdatedAt;
        /// <summary>
        /// The version of the domain resource
        /// </summary>
        public readonly int? Version;

        [OutputConstructor]
        private SimpleDomain(
            string? createdAt,

            bool? enabled,

            bool? geoip,

            bool? gtd,

            int? id,

            Outputs.SimpleDomainLinksProperties? links,

            string? name,

            string? note,

            CloudySkySoftware.Pulumi.Constellix.Domains.SimpleDomainStatus? status,

            ImmutableArray<Outputs.Tag> tags,

            string? updatedAt,

            int? version)
        {
            CreatedAt = createdAt;
            Enabled = enabled;
            Geoip = geoip;
            Gtd = gtd;
            Id = id;
            Links = links;
            Name = name;
            Note = note;
            Status = status;
            Tags = tags;
            UpdatedAt = updatedAt;
            Version = version;
        }
    }
}
