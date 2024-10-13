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

    [OutputType]
    public sealed class ContactlistSlackContactlistProperties
    {
        /// <summary>
        /// Unique ID for the contact list
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Links for the contact list
        /// </summary>
        public readonly Outputs.ContactlistSlackContactlistPropertiesLinksProperties? Links;

        [OutputConstructor]
        private ContactlistSlackContactlistProperties(
            int? id,

            Outputs.ContactlistSlackContactlistPropertiesLinksProperties? links)
        {
            Id = id;
            Links = links;
        }
    }
}
