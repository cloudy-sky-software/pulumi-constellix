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
    /// Links for domain objects
    /// </summary>
    [OutputType]
    public sealed class DomainLinksProperties
    {
        public readonly string? Records;
        public readonly string? Self;

        [OutputConstructor]
        private DomainLinksProperties(
            string? records,

            string? self)
        {
            Records = records;
            Self = self;
        }
    }
}
