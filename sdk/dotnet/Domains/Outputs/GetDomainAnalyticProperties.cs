// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Domains.Outputs
{

    [OutputType]
    public sealed class GetDomainAnalyticProperties
    {
        /// <summary>
        /// Analytics for a specific domain
        /// </summary>
        public readonly Outputs.Domainanalytics? Data;

        [OutputConstructor]
        private GetDomainAnalyticProperties(Outputs.Domainanalytics? data)
        {
            Data = data;
        }
    }
}
