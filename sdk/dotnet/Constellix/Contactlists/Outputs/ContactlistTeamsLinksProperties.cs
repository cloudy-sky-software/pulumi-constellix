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
    /// Links for the MS Teams Webhook
    /// </summary>
    [OutputType]
    public sealed class ContactlistTeamsLinksProperties
    {
        public readonly string? Self;

        [OutputConstructor]
        private ContactlistTeamsLinksProperties(string? self)
        {
            Self = self;
        }
    }
}
