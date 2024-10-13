// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Pools
{
    /// <summary>
    /// The type of pool
    /// </summary>
    [EnumType]
    public readonly struct PoolType : IEquatable<PoolType>
    {
        private readonly string _value;

        private PoolType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PoolType A { get; } = new PoolType("A");
        public static PoolType Aaaa { get; } = new PoolType("AAAA");
        public static PoolType Cname { get; } = new PoolType("CNAME");

        public static bool operator ==(PoolType left, PoolType right) => left.Equals(right);
        public static bool operator !=(PoolType left, PoolType right) => !left.Equals(right);

        public static explicit operator string(PoolType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PoolType other && Equals(other);
        public bool Equals(PoolType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The failover/check policy for this value
    /// </summary>
    [EnumType]
    public readonly struct PoolValuesItemPropertiesPolicy : IEquatable<PoolValuesItemPropertiesPolicy>
    {
        private readonly string _value;

        private PoolValuesItemPropertiesPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PoolValuesItemPropertiesPolicy FollowSonar { get; } = new PoolValuesItemPropertiesPolicy("follow_sonar");
        public static PoolValuesItemPropertiesPolicy AlwaysOff { get; } = new PoolValuesItemPropertiesPolicy("always_off");
        public static PoolValuesItemPropertiesPolicy AlwaysOn { get; } = new PoolValuesItemPropertiesPolicy("always_on");
        public static PoolValuesItemPropertiesPolicy OffOnFailure { get; } = new PoolValuesItemPropertiesPolicy("off_on_failure");

        public static bool operator ==(PoolValuesItemPropertiesPolicy left, PoolValuesItemPropertiesPolicy right) => left.Equals(right);
        public static bool operator !=(PoolValuesItemPropertiesPolicy left, PoolValuesItemPropertiesPolicy right) => !left.Equals(right);

        public static explicit operator string(PoolValuesItemPropertiesPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PoolValuesItemPropertiesPolicy other && Equals(other);
        public bool Equals(PoolValuesItemPropertiesPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of pool
    /// </summary>
    [EnumType]
    public readonly struct PoolindexType : IEquatable<PoolindexType>
    {
        private readonly string _value;

        private PoolindexType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PoolindexType A { get; } = new PoolindexType("A");
        public static PoolindexType Aaaa { get; } = new PoolindexType("AAAA");
        public static PoolindexType Cname { get; } = new PoolindexType("CNAME");

        public static bool operator ==(PoolindexType left, PoolindexType right) => left.Equals(right);
        public static bool operator !=(PoolindexType left, PoolindexType right) => !left.Equals(right);

        public static explicit operator string(PoolindexType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PoolindexType other && Equals(other);
        public bool Equals(PoolindexType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Percentage of how much is the response time allowed to deviate?
    /// </summary>
    public enum PoolitoConfigPropertiesDeviationAllowance
    {
        PoolitoConfigPropertiesDeviationAllowance_d_float64_10_ = 10,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_20_ = 20,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_30_ = 30,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_40_ = 40,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_50_ = 50,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_60_ = 60,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_70_ = 70,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_80_ = 80,
        PoolitoConfigPropertiesDeviationAllowance_d_float64_90_ = 90,
    }

    [EnumType]
    public readonly struct PoolitoConfigPropertiesHandicapFactor : IEquatable<PoolitoConfigPropertiesHandicapFactor>
    {
        private readonly string _value;

        private PoolitoConfigPropertiesHandicapFactor(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PoolitoConfigPropertiesHandicapFactor None { get; } = new PoolitoConfigPropertiesHandicapFactor("none");
        public static PoolitoConfigPropertiesHandicapFactor Percent { get; } = new PoolitoConfigPropertiesHandicapFactor("percent");
        public static PoolitoConfigPropertiesHandicapFactor Speed { get; } = new PoolitoConfigPropertiesHandicapFactor("speed");

        public static bool operator ==(PoolitoConfigPropertiesHandicapFactor left, PoolitoConfigPropertiesHandicapFactor right) => left.Equals(right);
        public static bool operator !=(PoolitoConfigPropertiesHandicapFactor left, PoolitoConfigPropertiesHandicapFactor right) => !left.Equals(right);

        public static explicit operator string(PoolitoConfigPropertiesHandicapFactor value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PoolitoConfigPropertiesHandicapFactor other && Equals(other);
        public bool Equals(PoolitoConfigPropertiesHandicapFactor other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Where monitoring should be performed from
    /// </summary>
    [EnumType]
    public readonly struct PoolitoConfigPropertiesMonitoringRegion : IEquatable<PoolitoConfigPropertiesMonitoringRegion>
    {
        private readonly string _value;

        private PoolitoConfigPropertiesMonitoringRegion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PoolitoConfigPropertiesMonitoringRegion World { get; } = new PoolitoConfigPropertiesMonitoringRegion("world");
        public static PoolitoConfigPropertiesMonitoringRegion Asiapac { get; } = new PoolitoConfigPropertiesMonitoringRegion("asiapac");
        public static PoolitoConfigPropertiesMonitoringRegion Europe { get; } = new PoolitoConfigPropertiesMonitoringRegion("europe");
        public static PoolitoConfigPropertiesMonitoringRegion Nacentral { get; } = new PoolitoConfigPropertiesMonitoringRegion("nacentral");
        public static PoolitoConfigPropertiesMonitoringRegion Naeast { get; } = new PoolitoConfigPropertiesMonitoringRegion("naeast");
        public static PoolitoConfigPropertiesMonitoringRegion Nawest { get; } = new PoolitoConfigPropertiesMonitoringRegion("nawest");
        public static PoolitoConfigPropertiesMonitoringRegion Oceania { get; } = new PoolitoConfigPropertiesMonitoringRegion("oceania");
        public static PoolitoConfigPropertiesMonitoringRegion Southamerica { get; } = new PoolitoConfigPropertiesMonitoringRegion("southamerica");

        public static bool operator ==(PoolitoConfigPropertiesMonitoringRegion left, PoolitoConfigPropertiesMonitoringRegion right) => left.Equals(right);
        public static bool operator !=(PoolitoConfigPropertiesMonitoringRegion left, PoolitoConfigPropertiesMonitoringRegion right) => !left.Equals(right);

        public static explicit operator string(PoolitoConfigPropertiesMonitoringRegion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PoolitoConfigPropertiesMonitoringRegion other && Equals(other);
        public bool Equals(PoolitoConfigPropertiesMonitoringRegion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The number of seconds between each check
    /// </summary>
    public enum PoolitoConfigPropertiesPeriod
    {
        PoolitoConfigPropertiesPeriod_d_float64_30_ = 30,
        PoolitoConfigPropertiesPeriod_d_float64_60_ = 60,
        PoolitoConfigPropertiesPeriod_d_float64_120_ = 120,
        PoolitoConfigPropertiesPeriod_d_float64_180_ = 180,
        PoolitoConfigPropertiesPeriod_d_float64_240_ = 240,
        PoolitoConfigPropertiesPeriod_d_float64_300_ = 300,
    }

    [EnumType]
    public readonly struct SimpleDomainStatus : IEquatable<SimpleDomainStatus>
    {
        private readonly string _value;

        private SimpleDomainStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SimpleDomainStatus Active { get; } = new SimpleDomainStatus("ACTIVE");
        public static SimpleDomainStatus Suspended { get; } = new SimpleDomainStatus("SUSPENDED");
        public static SimpleDomainStatus Terminated { get; } = new SimpleDomainStatus("TERMINATED");

        public static bool operator ==(SimpleDomainStatus left, SimpleDomainStatus right) => left.Equals(right);
        public static bool operator !=(SimpleDomainStatus left, SimpleDomainStatus right) => !left.Equals(right);

        public static explicit operator string(SimpleDomainStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SimpleDomainStatus other && Equals(other);
        public bool Equals(SimpleDomainStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of pool, either A, AAAA or CNAME
    /// </summary>
    [EnumType]
    public readonly struct Type : IEquatable<Type>
    {
        private readonly string _value;

        private Type(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Type A { get; } = new Type("A");
        public static Type Aaaa { get; } = new Type("AAAA");
        public static Type Cname { get; } = new Type("CNAME");

        public static bool operator ==(Type left, Type right) => left.Equals(right);
        public static bool operator !=(Type left, Type right) => !left.Equals(right);

        public static explicit operator string(Type value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Type other && Equals(other);
        public bool Equals(Type other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The failover/check policy for this value
    /// </summary>
    [EnumType]
    public readonly struct ValuesItemPropertiesPolicy : IEquatable<ValuesItemPropertiesPolicy>
    {
        private readonly string _value;

        private ValuesItemPropertiesPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ValuesItemPropertiesPolicy FollowSonar { get; } = new ValuesItemPropertiesPolicy("follow_sonar");
        public static ValuesItemPropertiesPolicy AlwaysOff { get; } = new ValuesItemPropertiesPolicy("always_off");
        public static ValuesItemPropertiesPolicy AlwaysOn { get; } = new ValuesItemPropertiesPolicy("always_on");
        public static ValuesItemPropertiesPolicy OffOnFailure { get; } = new ValuesItemPropertiesPolicy("off_on_failure");

        public static bool operator ==(ValuesItemPropertiesPolicy left, ValuesItemPropertiesPolicy right) => left.Equals(right);
        public static bool operator !=(ValuesItemPropertiesPolicy left, ValuesItemPropertiesPolicy right) => !left.Equals(right);

        public static explicit operator string(ValuesItemPropertiesPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ValuesItemPropertiesPolicy other && Equals(other);
        public bool Equals(ValuesItemPropertiesPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
