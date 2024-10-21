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
    /// Failover record mode
    /// </summary>
    [OutputType]
    public sealed class ValueCnameRequestValueProperties
    {
        /// <summary>
        /// Whether this failover value is enabled or not
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The failover mode
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Domains.ValueCnameRequestValuePropertiesMode? Mode;
        public readonly ImmutableArray<Outputs.ValueCnameRequestValuePropertiesValuesItemProperties> Values;

        [OutputConstructor]
        private ValueCnameRequestValueProperties(
            bool? enabled,

            CloudySkySoftware.Pulumi.Constellix.Domains.ValueCnameRequestValuePropertiesMode? mode,

            ImmutableArray<Outputs.ValueCnameRequestValuePropertiesValuesItemProperties> values)
        {
            Enabled = enabled;
            Mode = mode;
            Values = values;
        }
    }
}
