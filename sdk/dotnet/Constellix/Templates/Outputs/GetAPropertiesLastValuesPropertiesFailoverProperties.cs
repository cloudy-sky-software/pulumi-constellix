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
    /// The previous values in failover mode
    /// </summary>
    [OutputType]
    public sealed class GetAPropertiesLastValuesPropertiesFailoverProperties
    {
        /// <summary>
        /// The failover mode
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Templates.GetAPropertiesLastValuesPropertiesFailoverPropertiesMode? Mode;
        public readonly ImmutableArray<Outputs.GetAPropertiesLastValuesPropertiesFailoverPropertiesValuesItem> Values;

        [OutputConstructor]
        private GetAPropertiesLastValuesPropertiesFailoverProperties(
            CloudySkySoftware.Pulumi.Constellix.Templates.GetAPropertiesLastValuesPropertiesFailoverPropertiesMode? mode,

            ImmutableArray<Outputs.GetAPropertiesLastValuesPropertiesFailoverPropertiesValuesItem> values)
        {
            Mode = mode;
            Values = values;
        }
    }
}
