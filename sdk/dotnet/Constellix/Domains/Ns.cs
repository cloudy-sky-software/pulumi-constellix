// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Domains
{
    [ConstellixResourceType("constellix:domains:Ns")]
    public partial class Ns : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Contact lists to be notified if a failover happens in a failover mode.
        /// </summary>
        [Output("contacts")]
        public Output<ImmutableArray<int>> Contacts { get; private set; } = null!;

        [Output("data")]
        public Output<Outputs.DataProperties?> Data { get; private set; } = null!;

        /// <summary>
        /// Whether the record is enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        /// </summary>
        [Output("geoFailover")]
        public Output<bool?> GeoFailover { get; private set; } = null!;

        /// <summary>
        /// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
        /// </summary>
        [Output("geoproximity")]
        public Output<int?> Geoproximity { get; private set; } = null!;

        /// <summary>
        /// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
        /// </summary>
        [Output("ipfilter")]
        public Output<int?> Ipfilter { get; private set; } = null!;

        /// <summary>
        /// If the requesting IP matches the IP filter, don't return a response
        /// </summary>
        [Output("ipfilterDrop")]
        public Output<bool?> IpfilterDrop { get; private set; } = null!;

        /// <summary>
        /// The current mode for this record
        /// </summary>
        [Output("mode")]
        public Output<CloudySkySoftware.Pulumi.Constellix.Domains.NsPropertiesMode?> Mode { get; private set; } = null!;

        /// <summary>
        /// The name for the record
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// A description of the record. It must be 512 characters or less.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// Optional region for this record. Will default to 'default'.
        /// </summary>
        [Output("region")]
        public Output<CloudySkySoftware.Pulumi.Constellix.Domains.RecordCreateDetailsRegion?> Region { get; private set; } = null!;

        /// <summary>
        /// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        /// </summary>
        [Output("skipLookup")]
        public Output<bool?> SkipLookup { get; private set; } = null!;

        /// <summary>
        /// How long DNS servers should cache the record for
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of record
        /// </summary>
        [Output("type")]
        public Output<CloudySkySoftware.Pulumi.Constellix.Domains.NsPropertiesType?> Type { get; private set; } = null!;

        /// <summary>
        /// Standard record mode
        /// </summary>
        [Output("value")]
        public Output<ImmutableArray<Outputs.ValueNsValueItemProperties>> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Ns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ns(string name, NsArgs? args = null, CustomResourceOptions? options = null)
            : base("constellix:domains:Ns", name, args ?? new NsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ns(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("constellix:domains:Ns", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Ns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ns Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Ns(name, id, options);
        }
    }

    public sealed class NsArgs : global::Pulumi.ResourceArgs
    {
        [Input("contacts")]
        private InputList<int>? _contacts;

        /// <summary>
        /// Contact lists to be notified if a failover happens in a failover mode.
        /// </summary>
        public InputList<int> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<int>());
            set => _contacts = value;
        }

        /// <summary>
        /// The ID of the domain object
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Whether the record is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        /// </summary>
        [Input("geoFailover")]
        public Input<bool>? GeoFailover { get; set; }

        /// <summary>
        /// The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
        /// </summary>
        [Input("geoproximity")]
        public Input<int>? Geoproximity { get; set; }

        /// <summary>
        /// The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
        /// </summary>
        [Input("ipfilter")]
        public Input<int>? Ipfilter { get; set; }

        /// <summary>
        /// If the requesting IP matches the IP filter, don't return a response
        /// </summary>
        [Input("ipfilterDrop")]
        public Input<bool>? IpfilterDrop { get; set; }

        /// <summary>
        /// The current mode for this record
        /// </summary>
        [Input("mode")]
        public Input<CloudySkySoftware.Pulumi.Constellix.Domains.NsPropertiesMode>? Mode { get; set; }

        /// <summary>
        /// The name for the record
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A description of the record. It must be 512 characters or less.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// Optional region for this record. Will default to 'default'.
        /// </summary>
        [Input("region")]
        public Input<CloudySkySoftware.Pulumi.Constellix.Domains.RecordCreateDetailsRegion>? Region { get; set; }

        /// <summary>
        /// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        /// </summary>
        [Input("skipLookup")]
        public Input<bool>? SkipLookup { get; set; }

        /// <summary>
        /// How long DNS servers should cache the record for
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of record
        /// </summary>
        [Input("type")]
        public Input<CloudySkySoftware.Pulumi.Constellix.Domains.NsPropertiesType>? Type { get; set; }

        [Input("value")]
        private InputList<Inputs.ValueNsValueItemPropertiesArgs>? _value;

        /// <summary>
        /// Standard record mode
        /// </summary>
        public InputList<Inputs.ValueNsValueItemPropertiesArgs> Value
        {
            get => _value ?? (_value = new InputList<Inputs.ValueNsValueItemPropertiesArgs>());
            set => _value = value;
        }

        public NsArgs()
        {
        }
        public static new NsArgs Empty => new NsArgs();
    }
}
