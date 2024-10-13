// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Pools.Outputs
{

    [OutputType]
    public sealed class ListPoolsProperties
    {
        /// <summary>
        /// The pools for this page
        /// </summary>
        public readonly ImmutableArray<Outputs.Poolindex> Data;
        /// <summary>
        /// Metadata for list responses
        /// </summary>
        public readonly Outputs.ListMetadata? Meta;

        [OutputConstructor]
        private ListPoolsProperties(
            ImmutableArray<Outputs.Poolindex> data,

            Outputs.ListMetadata? meta)
        {
            Data = data;
            Meta = meta;
        }
    }
}
