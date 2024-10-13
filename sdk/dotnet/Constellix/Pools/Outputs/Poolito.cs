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

    [OutputType]
    public sealed class Poolito
    {
        /// <summary>
        /// The Ito configuration
        /// </summary>
        public readonly Outputs.PoolitoConfigProperties? Config;
        /// <summary>
        /// Is Ito enabled for this pool?
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private Poolito(
            Outputs.PoolitoConfigProperties? config,

            bool? enabled)
        {
            Config = config;
            Enabled = enabled;
        }
    }
}
