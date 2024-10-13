// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Templates
{
    [ConstellixResourceType("constellix:templates:Template")]
    public partial class Template : global::Pulumi.CustomResource
    {
        [Output("data")]
        public Output<Outputs.DataProperties?> Data { get; private set; } = null!;

        /// <summary>
        /// Is GeoIP functionality enabled for the template
        /// </summary>
        [Output("geoip")]
        public Output<bool?> Geoip { get; private set; } = null!;

        /// <summary>
        /// Is Global Traffic Director enabled for the template
        /// </summary>
        [Output("gtd")]
        public Output<bool?> Gtd { get; private set; } = null!;

        /// <summary>
        /// The name of the template
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("constellix:templates:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("constellix:templates:Template", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-constellix",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Template(name, id, options);
        }
    }

    public sealed class TemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is GeoIP functionality enabled for the template
        /// </summary>
        [Input("geoip")]
        public Input<bool>? Geoip { get; set; }

        /// <summary>
        /// Is Global Traffic Director enabled for the template
        /// </summary>
        [Input("gtd")]
        public Input<bool>? Gtd { get; set; }

        /// <summary>
        /// The name of the template
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TemplateArgs()
        {
        }
        public static new TemplateArgs Empty => new TemplateArgs();
    }
}
