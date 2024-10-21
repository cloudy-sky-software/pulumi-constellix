// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Domains.Outputs
{

    [OutputType]
    public sealed class ValueCertValueItemProperties
    {
        /// <summary>
        /// An integer representing the algorithm
        /// </summary>
        public readonly int? Algorithm;
        /// <summary>
        /// A base 64 encoded string containing the certificate information
        /// </summary>
        public readonly string? Certificate;
        /// <summary>
        /// An integer representing the type of certificate
        /// </summary>
        public readonly int? CertificateType;
        /// <summary>
        /// Whether the entry is enabled or not. Disabled entries will not be included in a response
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// An integer representing the key tag
        /// </summary>
        public readonly int? KeyTag;

        [OutputConstructor]
        private ValueCertValueItemProperties(
            int? algorithm,

            string? certificate,

            int? certificateType,

            bool? enabled,

            int? keyTag)
        {
            Algorithm = algorithm;
            Certificate = certificate;
            CertificateType = certificateType;
            Enabled = enabled;
            KeyTag = keyTag;
        }
    }
}
