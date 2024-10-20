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
    /// The previous values in failover mode
    /// </summary>
    [OutputType]
    public sealed class GetANAMEPropertiesLastValuesPropertiesFailoverProperties
    {
        /// <summary>
        /// The failover mode
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Domains.GetANAMEPropertiesLastValuesPropertiesFailoverPropertiesMode? Mode;
        public readonly ImmutableArray<Outputs.GetANAMEPropertiesLastValuesPropertiesFailoverPropertiesValuesItem> Values;

        [OutputConstructor]
        private GetANAMEPropertiesLastValuesPropertiesFailoverProperties(
            CloudySkySoftware.Pulumi.Constellix.Domains.GetANAMEPropertiesLastValuesPropertiesFailoverPropertiesMode? mode,

            ImmutableArray<Outputs.GetANAMEPropertiesLastValuesPropertiesFailoverPropertiesValuesItem> values)
        {
            Mode = mode;
            Values = values;
        }
    }
}
