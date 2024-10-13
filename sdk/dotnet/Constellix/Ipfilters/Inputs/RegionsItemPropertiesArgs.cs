// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Ipfilters.Inputs
{

    public sealed class RegionsItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional AS Number that this rule applies to. Values must be between 0 and 4,294,967,295
        /// </summary>
        [Input("asn")]
        public Input<int>? Asn { get; set; }

        /// <summary>
        /// The continent for this region
        /// </summary>
        [Input("continent")]
        public Input<CloudySkySoftware.Pulumi.Constellix.Ipfilters.RegionsItemPropertiesContinent>? Continent { get; set; }

        /// <summary>
        /// Optional 2 digit ISO code for the country - https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Optional 2 digit code for the region
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RegionsItemPropertiesArgs()
        {
        }
        public static new RegionsItemPropertiesArgs Empty => new RegionsItemPropertiesArgs();
    }
}
