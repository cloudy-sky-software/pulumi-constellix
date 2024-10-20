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

    [OutputType]
    public sealed class ValueAaaaValue
    {
        /// <summary>
        /// Whether this failover value is enabled or not
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The failover mode
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Domains.ValueAaaaValuePropertiesMode? Mode;
        public readonly ImmutableArray<Outputs.ValueAaaaValuePropertiesValuesItem> Values;

        [OutputConstructor]
        private ValueAaaaValue(
            bool? enabled,

            CloudySkySoftware.Pulumi.Constellix.Domains.ValueAaaaValuePropertiesMode? mode,

            ImmutableArray<Outputs.ValueAaaaValuePropertiesValuesItem> values)
        {
            Enabled = enabled;
            Mode = mode;
            Values = values;
        }
    }
}
