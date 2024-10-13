// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Domains.Outputs
{

    /// <summary>
    /// A simplied version of a rebranded nameserver using your own domain name.
    /// </summary>
    [OutputType]
    public sealed class SimpleVanitynameserver
    {
        /// <summary>
        /// A unique ID for the vanity nameserver
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Links relevant to this object
        /// </summary>
        public readonly Outputs.SimpleVanitynameserverLinksProperties? Links;

        [OutputConstructor]
        private SimpleVanitynameserver(
            int? id,

            Outputs.SimpleVanitynameserverLinksProperties? links)
        {
            Id = id;
            Links = links;
        }
    }
}
