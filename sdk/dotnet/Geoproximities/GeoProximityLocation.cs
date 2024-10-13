// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Geoproximities
{
    /// <summary>
    /// Geo Proximity Location
    /// </summary>
    [CONSTELLIXResourceType("constellix:geoproximities:GeoProximityLocation")]
    public partial class GeoProximityLocation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The numeric ID for a city
        /// </summary>
        [Output("city")]
        public Output<int?> City { get; private set; } = null!;

        /// <summary>
        /// 2 digit ISO country code
        /// </summary>
        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        [Output("data")]
        public Output<Outputs.DataProperties?> Data { get; private set; } = null!;

        /// <summary>
        /// Latitude of the location
        /// </summary>
        [Output("latitude")]
        public Output<double?> Latitude { get; private set; } = null!;

        /// <summary>
        /// Longitude of the location
        /// </summary>
        [Output("longitude")]
        public Output<double?> Longitude { get; private set; } = null!;

        /// <summary>
        /// The name of the Geo Proximity location
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Region, state or province code
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;


        /// <summary>
        /// Create a GeoProximityLocation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeoProximityLocation(string name, GeoProximityLocationArgs? args = null, CustomResourceOptions? options = null)
            : base("constellix:geoproximities:GeoProximityLocation", name, args ?? new GeoProximityLocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeoProximityLocation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("constellix:geoproximities:GeoProximityLocation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GeoProximityLocation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeoProximityLocation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GeoProximityLocation(name, id, options);
        }
    }

    public sealed class GeoProximityLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The numeric ID for a city
        /// </summary>
        [Input("city")]
        public Input<int>? City { get; set; }

        /// <summary>
        /// 2 digit ISO country code
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Latitude of the location
        /// </summary>
        [Input("latitude")]
        public Input<double>? Latitude { get; set; }

        /// <summary>
        /// Longitude of the location
        /// </summary>
        [Input("longitude")]
        public Input<double>? Longitude { get; set; }

        /// <summary>
        /// The name of the Geo Proximity location
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Region, state or province code
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GeoProximityLocationArgs()
        {
        }
        public static new GeoProximityLocationArgs Empty => new GeoProximityLocationArgs();
    }
}
