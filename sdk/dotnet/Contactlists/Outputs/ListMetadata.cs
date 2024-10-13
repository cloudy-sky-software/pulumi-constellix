// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Contactlists.Outputs
{

    /// <summary>
    /// Metadata for list responses
    /// </summary>
    [OutputType]
    public sealed class ListMetadata
    {
        /// <summary>
        /// Relevant links for this list
        /// </summary>
        public readonly Outputs.ListMetadataLinksProperties? Links;
        /// <summary>
        /// Pagination details
        /// </summary>
        public readonly Outputs.ListMetadataPaginationProperties? Pagination;

        [OutputConstructor]
        private ListMetadata(
            Outputs.ListMetadataLinksProperties? links,

            Outputs.ListMetadataPaginationProperties? pagination)
        {
            Links = links;
            Pagination = pagination;
        }
    }
}
