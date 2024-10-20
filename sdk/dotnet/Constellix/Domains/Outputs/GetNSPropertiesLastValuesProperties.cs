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
    public sealed class GetNSPropertiesLastValuesProperties
    {
        /// <summary>
        /// The previous values in standard mode
        /// </summary>
        public readonly ImmutableArray<Outputs.ValueNsPropertiesValueItems> Standard;

        [OutputConstructor]
        private GetNSPropertiesLastValuesProperties(ImmutableArray<Outputs.ValueNsPropertiesValueItems> standard)
        {
            Standard = standard;
        }
    }
}
