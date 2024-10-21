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

    [OutputType]
    public sealed class ValueNaptrValueItemProperties
    {
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Any flags for this record
        /// </summary>
        public readonly string? Flags;
        /// <summary>
        /// The order of the record
        /// </summary>
        public readonly int? Order;
        /// <summary>
        /// The preference for the record
        /// </summary>
        public readonly int? Preference;
        /// <summary>
        /// A regular expression to use
        /// </summary>
        public readonly string? RegularExpression;
        /// <summary>
        /// The replacement for the regular expression
        /// </summary>
        public readonly string? Replacement;
        /// <summary>
        /// The service the record is used for
        /// </summary>
        public readonly string? Service;

        [OutputConstructor]
        private ValueNaptrValueItemProperties(
            bool? enabled,

            string? flags,

            int? order,

            int? preference,

            string? regularExpression,

            string? replacement,

            string? service)
        {
            Enabled = enabled;
            Flags = flags;
            Order = order;
            Preference = preference;
            RegularExpression = regularExpression;
            Replacement = replacement;
            Service = service;
        }
    }
}
