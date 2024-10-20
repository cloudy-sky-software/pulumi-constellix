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
    /// Lists of email addresses used for notifications and messages about domains and records.
    /// </summary>
    [OutputType]
    public sealed class Contactlistwithoutemails
    {
        /// <summary>
        /// The number of emails in this contact list
        /// </summary>
        public readonly int? EmailCount;
        /// <summary>
        /// Unique ID for the contact list
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Links for the object
        /// </summary>
        public readonly Outputs.ContactlistwithoutemailsLinksProperties? Links;
        /// <summary>
        /// A name for this contact list
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private Contactlistwithoutemails(
            int? emailCount,

            int? id,

            Outputs.ContactlistwithoutemailsLinksProperties? links,

            string? name)
        {
            EmailCount = emailCount;
            Id = id;
            Links = links;
            Name = name;
        }
    }
}
