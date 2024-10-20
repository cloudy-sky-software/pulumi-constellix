// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Templates.Outputs
{

    /// <summary>
    /// The previous values of the record in the different modes
    /// </summary>
    [OutputType]
    public sealed class GetAPropertiesLastValuesProperties
    {
        /// <summary>
        /// The previous values in failover mode
        /// </summary>
        public readonly Outputs.GetAPropertiesLastValuesPropertiesFailoverProperties? Failover;
        /// <summary>
        /// The previous values in Pools mode
        /// </summary>
        public readonly ImmutableArray<Outputs.SimplePool> Pools;
        /// <summary>
        /// The previous values in Round-Robin Failover mode
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAPropertiesLastValuesPropertiesRoundRobinFailoverItem> RoundRobinFailover;
        /// <summary>
        /// The previous values in standard mode
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAPropertiesLastValuesPropertiesStandardItemProperties> Standard;

        [OutputConstructor]
        private GetAPropertiesLastValuesProperties(
            Outputs.GetAPropertiesLastValuesPropertiesFailoverProperties? failover,

            ImmutableArray<Outputs.SimplePool> pools,

            ImmutableArray<Outputs.GetAPropertiesLastValuesPropertiesRoundRobinFailoverItem> roundRobinFailover,

            ImmutableArray<Outputs.GetAPropertiesLastValuesPropertiesStandardItemProperties> standard)
        {
            Failover = failover;
            Pools = pools;
            RoundRobinFailover = roundRobinFailover;
            Standard = standard;
        }
    }
}
