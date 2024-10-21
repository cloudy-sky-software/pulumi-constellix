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

    [OutputType]
    public sealed class ListTemplateRecordsPropertiesDataItem
    {
        /// <summary>
        /// Links for the domain record
        /// </summary>
        public readonly Outputs.TemplaterecordLinksProperties? Links;
        public readonly Outputs.SimpleTemplate? Template;

        [OutputConstructor]
        private ListTemplateRecordsPropertiesDataItem(
            Outputs.TemplaterecordLinksProperties? links,

            Outputs.SimpleTemplate? template)
        {
            Links = links;
            Template = template;
        }
    }
}
