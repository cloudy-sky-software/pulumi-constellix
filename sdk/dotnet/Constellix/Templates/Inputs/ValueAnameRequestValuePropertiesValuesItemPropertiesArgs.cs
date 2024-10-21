// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Templates.Inputs
{

    public sealed class ValueAnameRequestValuePropertiesValuesItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this entry is considered active or not
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Whether the failover entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The sort order of the entry. Lower order entries are preferred over higher order entries
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// The ID in Sonar to use for checking if the record should be used
        /// </summary>
        [Input("sonarCheckId")]
        public Input<int>? SonarCheckId { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ValueAnameRequestValuePropertiesValuesItemPropertiesArgs()
        {
        }
        public static new ValueAnameRequestValuePropertiesValuesItemPropertiesArgs Empty => new ValueAnameRequestValuePropertiesValuesItemPropertiesArgs();
    }
}
