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
    /// Links for the email address
    /// </summary>
    [OutputType]
    public sealed class ContactlistEmailLinksProperties
    {
        public readonly string? Self;

        [OutputConstructor]
        private ContactlistEmailLinksProperties(string? self)
        {
            Self = self;
        }
    }
}
