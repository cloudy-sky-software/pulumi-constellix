// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Pools.Outputs
{

    /// <summary>
    /// Links for the pool
    /// </summary>
    [OutputType]
    public sealed class PoolindexLinksProperties
    {
        public readonly string? Self;

        [OutputConstructor]
        private PoolindexLinksProperties(string? self)
        {
            Self = self;
        }
    }
}
