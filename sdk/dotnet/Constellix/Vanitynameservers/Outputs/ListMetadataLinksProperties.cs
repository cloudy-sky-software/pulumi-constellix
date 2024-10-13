// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Vanitynameservers.Outputs
{

    /// <summary>
    /// Relevant links for this list
    /// </summary>
    [OutputType]
    public sealed class ListMetadataLinksProperties
    {
        public readonly string? First;
        public readonly string? Last;
        public readonly string? Next;
        public readonly string? Previous;
        public readonly string? Self;

        [OutputConstructor]
        private ListMetadataLinksProperties(
            string? first,

            string? last,

            string? next,

            string? previous,

            string? self)
        {
            First = first;
            Last = last;
            Next = next;
            Previous = previous;
            Self = self;
        }
    }
}
