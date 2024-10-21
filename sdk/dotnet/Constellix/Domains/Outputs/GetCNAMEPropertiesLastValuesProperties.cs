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
    /// The previous values of the record in the different modes
    /// </summary>
    [OutputType]
    public sealed class GetCNAMEPropertiesLastValuesProperties
    {
        /// <summary>
        /// The previous values in failover mode
        /// </summary>
        public readonly Outputs.GetCNAMEPropertiesLastValuesPropertiesFailoverProperties? Failover;
        /// <summary>
        /// The previous values in Pools mode
        /// </summary>
        public readonly ImmutableArray<Outputs.SimplePool> Pools;
        /// <summary>
        /// The previous values in standard mode
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCNAMEPropertiesLastValuesPropertiesStandardItemProperties> Standard;

        [OutputConstructor]
        private GetCNAMEPropertiesLastValuesProperties(
            Outputs.GetCNAMEPropertiesLastValuesPropertiesFailoverProperties? failover,

            ImmutableArray<Outputs.SimplePool> pools,

            ImmutableArray<Outputs.GetCNAMEPropertiesLastValuesPropertiesStandardItemProperties> standard)
        {
            Failover = failover;
            Pools = pools;
            Standard = standard;
        }
    }
}
