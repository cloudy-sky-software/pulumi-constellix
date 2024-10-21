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

    [OutputType]
    public sealed class ValueNsValueItemProperties
    {
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The hostname for the nameserver
        /// </summary>
        public readonly string? Host;

        [OutputConstructor]
        private ValueNsValueItemProperties(
            bool? enabled,

            string? host)
        {
            Enabled = enabled;
            Host = host;
        }
    }
}
