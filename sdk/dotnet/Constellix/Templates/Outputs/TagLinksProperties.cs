// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Templates.Outputs
{

    /// <summary>
    /// Links for tags
    /// </summary>
    [OutputType]
    public sealed class TagLinksProperties
    {
        public readonly string? Self;

        [OutputConstructor]
        private TagLinksProperties(string? self)
        {
            Self = self;
        }
    }
}
