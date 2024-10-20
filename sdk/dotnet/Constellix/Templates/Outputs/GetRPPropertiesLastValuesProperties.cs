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
    public sealed class GetRPPropertiesLastValuesProperties
    {
        /// <summary>
        /// The previous values in standard mode
        /// </summary>
        public readonly ImmutableArray<Outputs.ValueRpPropertiesValueItems> Standard;

        [OutputConstructor]
        private GetRPPropertiesLastValuesProperties(ImmutableArray<Outputs.ValueRpPropertiesValueItems> standard)
        {
            Standard = standard;
        }
    }
}
