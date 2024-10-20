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
    /// Failover record mode
    /// </summary>
    [OutputType]
    public sealed class ValueAnameRequestValueProperties
    {
        /// <summary>
        /// Whether this failover value is enabled or not
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The failover mode
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Templates.ValueAnameRequestValuePropertiesMode? Mode;
        public readonly ImmutableArray<Outputs.ValueAnameRequestValuePropertiesValuesItemProperties> Values;

        [OutputConstructor]
        private ValueAnameRequestValueProperties(
            bool? enabled,

            CloudySkySoftware.Pulumi.Constellix.Templates.ValueAnameRequestValuePropertiesMode? mode,

            ImmutableArray<Outputs.ValueAnameRequestValuePropertiesValuesItemProperties> values)
        {
            Enabled = enabled;
            Mode = mode;
            Values = values;
        }
    }
}
