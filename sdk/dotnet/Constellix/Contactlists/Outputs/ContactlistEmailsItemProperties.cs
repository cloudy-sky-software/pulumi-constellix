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
    /// An object for a single email in the contact list
    /// </summary>
    [OutputType]
    public sealed class ContactlistEmailsItemProperties
    {
        public readonly string? Address;
        /// <summary>
        /// Has the email been verified or not
        /// </summary>
        public readonly bool? Verified;

        [OutputConstructor]
        private ContactlistEmailsItemProperties(
            string? address,

            bool? verified)
        {
            Address = address;
            Verified = verified;
        }
    }
}
