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

    /// <summary>
    /// Links for the domain record
    /// </summary>
    [OutputType]
    public sealed class DomainrecordPropertiesLinksProperties
    {
        public readonly string? Self;

        [OutputConstructor]
        private DomainrecordPropertiesLinksProperties(string? self)
        {
            Self = self;
        }
    }
}
