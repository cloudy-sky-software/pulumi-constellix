// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Analytics.Outputs
{

    [OutputType]
    public sealed class ListAnalyticsProperties
    {
        /// <summary>
        /// Analytics for your account
        /// </summary>
        public readonly Outputs.Analytics? Data;

        [OutputConstructor]
        private ListAnalyticsProperties(Outputs.Analytics? data)
        {
            Data = data;
        }
    }
}
