// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Contactlists.Outputs
{

    [OutputType]
    public sealed class ListContactsProperties
    {
        public readonly ImmutableArray<Outputs.Contactlistwithoutemails> Data;
        /// <summary>
        /// Metadata for list responses
        /// </summary>
        public readonly Outputs.ListMetadata? Meta;

        [OutputConstructor]
        private ListContactsProperties(
            ImmutableArray<Outputs.Contactlistwithoutemails> data,

            Outputs.ListMetadata? meta)
        {
            Data = data;
            Meta = meta;
        }
    }
}
