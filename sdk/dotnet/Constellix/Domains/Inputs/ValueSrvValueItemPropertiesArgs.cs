// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Domains.Inputs
{

    public sealed class ValueSrvValueItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The hostname for the service
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The port the service runs on
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// A priority for this record
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// A weight for this record
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ValueSrvValueItemPropertiesArgs()
        {
        }
        public static new ValueSrvValueItemPropertiesArgs Empty => new ValueSrvValueItemPropertiesArgs();
    }
}
