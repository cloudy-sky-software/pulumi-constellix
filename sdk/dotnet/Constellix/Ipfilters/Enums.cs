// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Ipfilters
{
    [EnumType]
    public readonly struct ContinentsItem : IEquatable<ContinentsItem>
    {
        private readonly string _value;

        private ContinentsItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContinentsItem World { get; } = new ContinentsItem("world");
        public static ContinentsItem Af { get; } = new ContinentsItem("AF");
        public static ContinentsItem An { get; } = new ContinentsItem("AN");
        public static ContinentsItem As { get; } = new ContinentsItem("AS");
        public static ContinentsItem Eu { get; } = new ContinentsItem("EU");
        public static ContinentsItem Na { get; } = new ContinentsItem("NA");
        public static ContinentsItem Oc { get; } = new ContinentsItem("OC");
        public static ContinentsItem Sa { get; } = new ContinentsItem("SA");

        public static bool operator ==(ContinentsItem left, ContinentsItem right) => left.Equals(right);
        public static bool operator !=(ContinentsItem left, ContinentsItem right) => !left.Equals(right);

        public static explicit operator string(ContinentsItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContinentsItem other && Equals(other);
        public bool Equals(ContinentsItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct IpfilterContinentsItem : IEquatable<IpfilterContinentsItem>
    {
        private readonly string _value;

        private IpfilterContinentsItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IpfilterContinentsItem Af { get; } = new IpfilterContinentsItem("AF");
        public static IpfilterContinentsItem An { get; } = new IpfilterContinentsItem("AN");
        public static IpfilterContinentsItem As { get; } = new IpfilterContinentsItem("AS");
        public static IpfilterContinentsItem Eu { get; } = new IpfilterContinentsItem("EU");
        public static IpfilterContinentsItem Na { get; } = new IpfilterContinentsItem("NA");
        public static IpfilterContinentsItem Oc { get; } = new IpfilterContinentsItem("OC");
        public static IpfilterContinentsItem Sa { get; } = new IpfilterContinentsItem("SA");

        public static bool operator ==(IpfilterContinentsItem left, IpfilterContinentsItem right) => left.Equals(right);
        public static bool operator !=(IpfilterContinentsItem left, IpfilterContinentsItem right) => !left.Equals(right);

        public static explicit operator string(IpfilterContinentsItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IpfilterContinentsItem other && Equals(other);
        public bool Equals(IpfilterContinentsItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The continent for this region
    /// </summary>
    [EnumType]
    public readonly struct RegionsItemPropertiesContinent : IEquatable<RegionsItemPropertiesContinent>
    {
        private readonly string _value;

        private RegionsItemPropertiesContinent(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RegionsItemPropertiesContinent World { get; } = new RegionsItemPropertiesContinent("world");
        public static RegionsItemPropertiesContinent Af { get; } = new RegionsItemPropertiesContinent("AF");
        public static RegionsItemPropertiesContinent An { get; } = new RegionsItemPropertiesContinent("AN");
        public static RegionsItemPropertiesContinent As { get; } = new RegionsItemPropertiesContinent("AS");
        public static RegionsItemPropertiesContinent Eu { get; } = new RegionsItemPropertiesContinent("EU");
        public static RegionsItemPropertiesContinent Na { get; } = new RegionsItemPropertiesContinent("NA");
        public static RegionsItemPropertiesContinent Oc { get; } = new RegionsItemPropertiesContinent("OC");
        public static RegionsItemPropertiesContinent Sa { get; } = new RegionsItemPropertiesContinent("SA");

        public static bool operator ==(RegionsItemPropertiesContinent left, RegionsItemPropertiesContinent right) => left.Equals(right);
        public static bool operator !=(RegionsItemPropertiesContinent left, RegionsItemPropertiesContinent right) => !left.Equals(right);

        public static explicit operator string(RegionsItemPropertiesContinent value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RegionsItemPropertiesContinent other && Equals(other);
        public bool Equals(RegionsItemPropertiesContinent other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
