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
    public sealed class GetMX
    {
        /// <summary>
        /// A simple version of a contact list when inclued with other resources
        /// </summary>
        public readonly Outputs.SimpleContactlist? Contacts;
        public readonly Outputs.SimpleDomain? Domain;
        /// <summary>
        /// Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        /// </summary>
        public readonly bool? GeoFailover;
        /// <summary>
        /// Geo Proximity Location
        /// </summary>
        public readonly Outputs.SimpleGeoproximity? Geoproximity;
        /// <summary>
        /// A unique ID for this domain record
        /// </summary>
        public readonly int? Id;
        public readonly Outputs.SimpleIpfilter? Ipfilter;
        /// <summary>
        /// If the requesting IP matches the IP filter, don't return a response
        /// </summary>
        public readonly bool? IpfilterDrop;
        /// <summary>
        /// The previous values of the record in the different modes
        /// </summary>
        public readonly Outputs.GetMXPropertiesLastValuesProperties? LastValues;
        /// <summary>
        /// Links for the domain record
        /// </summary>
        public readonly Outputs.TemplaterecordLinksProperties? Links;
        /// <summary>
        /// How the record should work
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Templates.GetMXPropertiesMode? Mode;
        /// <summary>
        /// The name of the record
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A note about the record. Max 512 characters.
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// The region for this record
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Templates.RecordRegion? Region;
        /// <summary>
        /// Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        /// </summary>
        public readonly bool? SkipLookup;
        public readonly Outputs.SimpleTemplate? Template;
        /// <summary>
        /// The time to live in seconds for this record. must be between 0 and 2147483647
        /// </summary>
        public readonly int? Ttl;
        /// <summary>
        /// The type of record
        /// </summary>
        public readonly CloudySkySoftware.Pulumi.Constellix.Templates.GetMXPropertiesType? Type;
        /// <summary>
        /// Standard record mode
        /// </summary>
        public readonly ImmutableArray<Outputs.ValueMxValueItemProperties> Value;

        [OutputConstructor]
        private GetMX(
            Outputs.SimpleContactlist? contacts,

            Outputs.SimpleDomain? domain,

            bool? enabled,

            bool? geoFailover,

            Outputs.SimpleGeoproximity? geoproximity,

            int? id,

            Outputs.SimpleIpfilter? ipfilter,

            bool? ipfilterDrop,

            Outputs.GetMXPropertiesLastValuesProperties? lastValues,

            Outputs.TemplaterecordLinksProperties? links,

            CloudySkySoftware.Pulumi.Constellix.Templates.GetMXPropertiesMode? mode,

            string? name,

            string? notes,

            CloudySkySoftware.Pulumi.Constellix.Templates.RecordRegion? region,

            bool? skipLookup,

            Outputs.SimpleTemplate? template,

            int? ttl,

            CloudySkySoftware.Pulumi.Constellix.Templates.GetMXPropertiesType? type,

            ImmutableArray<Outputs.ValueMxValueItemProperties> value)
        {
            Contacts = contacts;
            Domain = domain;
            Enabled = enabled;
            GeoFailover = geoFailover;
            Geoproximity = geoproximity;
            Id = id;
            Ipfilter = ipfilter;
            IpfilterDrop = ipfilterDrop;
            LastValues = lastValues;
            Links = links;
            Mode = mode;
            Name = name;
            Notes = notes;
            Region = region;
            SkipLookup = skipLookup;
            Template = template;
            Ttl = ttl;
            Type = type;
            Value = value;
        }
    }
}
