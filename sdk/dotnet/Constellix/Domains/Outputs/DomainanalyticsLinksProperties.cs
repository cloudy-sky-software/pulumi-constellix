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
    public sealed class DomainanalyticsLinksProperties
    {
        /// <summary>
        /// The URL for the domain
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// The URL for these analytics
        /// </summary>
        public readonly string? Self;

        [OutputConstructor]
        private DomainanalyticsLinksProperties(
            string? domain,

            string? self)
        {
            Domain = domain;
            Self = self;
        }
    }
}
