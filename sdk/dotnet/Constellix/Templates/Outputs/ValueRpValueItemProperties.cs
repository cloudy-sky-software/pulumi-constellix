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
    public sealed class ValueRpValueItemProperties
    {
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The email of the person responsible. Replace @ with .
        /// </summary>
        public readonly string? Mailbox;
        /// <summary>
        /// The name of a TXT record containing more information
        /// </summary>
        public readonly string? Txt;

        [OutputConstructor]
        private ValueRpValueItemProperties(
            bool? enabled,

            string? mailbox,

            string? txt)
        {
            Enabled = enabled;
            Mailbox = mailbox;
            Txt = txt;
        }
    }
}
