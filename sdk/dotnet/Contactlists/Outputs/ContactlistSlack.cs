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
    public sealed class ContactlistSlack
    {
        /// <summary>
        /// The channel to send the message to
        /// </summary>
        public readonly string? Channel;
        public readonly Outputs.ContactlistSlackContactlistProperties? Contactlist;
        public readonly int? Id;
        /// <summary>
        /// Links for the Slack webhook
        /// </summary>
        public readonly Outputs.ContactlistSlackLinksProperties? Links;
        /// <summary>
        /// The inbound webhook URL for Slack
        /// </summary>
        public readonly string? Webhook;

        [OutputConstructor]
        private ContactlistSlack(
            string? channel,

            Outputs.ContactlistSlackContactlistProperties? contactlist,

            int? id,

            Outputs.ContactlistSlackLinksProperties? links,

            string? webhook)
        {
            Channel = channel;
            Contactlist = contactlist;
            Id = id;
            Links = links;
            Webhook = webhook;
        }
    }
}
