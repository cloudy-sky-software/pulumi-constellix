// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace CloudySkySoftware.Pulumi.Constellix
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("constellix");

        private static readonly __Value<string?> _apiKey = new __Value<string?>(() => __config.Get("apiKey"));
        /// <summary>
        /// The Constellix API key
        /// </summary>
        public static string? ApiKey
        {
            get => _apiKey.Get();
            set => _apiKey.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey"));
        /// <summary>
        /// The Constellix secret key
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

    }
}
