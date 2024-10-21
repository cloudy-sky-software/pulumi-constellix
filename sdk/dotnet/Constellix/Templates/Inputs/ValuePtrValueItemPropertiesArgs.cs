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

    public sealed class ValuePtrValueItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The hostname for the IP address
        /// </summary>
        [Input("system")]
        public Input<string>? System { get; set; }

        public ValuePtrValueItemPropertiesArgs()
        {
        }
        public static new ValuePtrValueItemPropertiesArgs Empty => new ValuePtrValueItemPropertiesArgs();
    }
}
