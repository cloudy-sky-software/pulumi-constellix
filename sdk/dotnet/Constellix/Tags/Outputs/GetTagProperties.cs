// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Tags.Outputs
{

    [OutputType]
    public sealed class GetTagProperties
    {
        /// <summary>
        /// A tag is used to group resources together
        /// </summary>
        public readonly Outputs.Tag? Data;

        [OutputConstructor]
        private GetTagProperties(Outputs.Tag? data)
        {
            Data = data;
        }
    }
}
