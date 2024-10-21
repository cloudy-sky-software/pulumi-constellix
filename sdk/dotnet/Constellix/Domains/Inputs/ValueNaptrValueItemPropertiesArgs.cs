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

    public sealed class ValueNaptrValueItemPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Any flags for this record
        /// </summary>
        [Input("flags")]
        public Input<string>? Flags { get; set; }

        /// <summary>
        /// The order of the record
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// The preference for the record
        /// </summary>
        [Input("preference")]
        public Input<int>? Preference { get; set; }

        /// <summary>
        /// A regular expression to use
        /// </summary>
        [Input("regularExpression")]
        public Input<string>? RegularExpression { get; set; }

        /// <summary>
        /// The replacement for the regular expression
        /// </summary>
        [Input("replacement")]
        public Input<string>? Replacement { get; set; }

        /// <summary>
        /// The service the record is used for
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public ValueNaptrValueItemPropertiesArgs()
        {
        }
        public static new ValueNaptrValueItemPropertiesArgs Empty => new ValueNaptrValueItemPropertiesArgs();
    }
}
