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

    public sealed class ValueHinfoValueItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text representing the CPU
        /// </summary>
        [Input("cpu")]
        public Input<string>? Cpu { get; set; }

        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Text representing the OS
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        public ValueHinfoValueItemPropertiesArgs()
        {
        }
        public static new ValueHinfoValueItemPropertiesArgs Empty => new ValueHinfoValueItemPropertiesArgs();
    }
}
