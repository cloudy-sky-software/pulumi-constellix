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
    public sealed class ValueHinfoValueItemProperties
    {
        /// <summary>
        /// Text representing the CPU
        /// </summary>
        public readonly string? Cpu;
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Text representing the OS
        /// </summary>
        public readonly string? Os;

        [OutputConstructor]
        private ValueHinfoValueItemProperties(
            string? cpu,

            bool? enabled,

            string? os)
        {
            Cpu = cpu;
            Enabled = enabled;
            Os = os;
        }
    }
}
