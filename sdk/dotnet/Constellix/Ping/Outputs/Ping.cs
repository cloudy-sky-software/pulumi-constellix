// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Ping.Outputs
{

    [OutputType]
    public sealed class Ping
    {
        public readonly string? Ip;
        public readonly string? Timestamp;
        public readonly string? Version;

        [OutputConstructor]
        private Ping(
            string? ip,

            string? timestamp,

            string? version)
        {
            Ip = ip;
            Timestamp = timestamp;
            Version = version;
        }
    }
}
