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
    public sealed class ListContactSlackWebhookProperties
    {
        public readonly Outputs.ContactlistSlack? Data;

        [OutputConstructor]
        private ListContactSlackWebhookProperties(Outputs.ContactlistSlack? data)
        {
            Data = data;
        }
    }
}
