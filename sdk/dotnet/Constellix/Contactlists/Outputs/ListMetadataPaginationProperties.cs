// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Contactlists.Outputs
{

    /// <summary>
    /// Pagination details
    /// </summary>
    [OutputType]
    public sealed class ListMetadataPaginationProperties
    {
        /// <summary>
        /// The number of items in this page of the response
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// The current results page
        /// </summary>
        public readonly int? CurrentPage;
        /// <summary>
        /// The number of items per page
        /// </summary>
        public readonly int? PerPage;
        /// <summary>
        /// The total number of objects matching the query
        /// </summary>
        public readonly int? Total;
        /// <summary>
        /// The total number of pages
        /// </summary>
        public readonly int? TotalPages;

        [OutputConstructor]
        private ListMetadataPaginationProperties(
            int? count,

            int? currentPage,

            int? perPage,

            int? total,

            int? totalPages)
        {
            Count = count;
            CurrentPage = currentPage;
            PerPage = perPage;
            Total = total;
            TotalPages = totalPages;
        }
    }
}
