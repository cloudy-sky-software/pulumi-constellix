// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Pulumi.Constellix.Templates.Inputs
{

    /// <summary>
    /// Failover record mode
    /// </summary>
    public sealed class ValueARequestValuePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The failover mode
        /// </summary>
        [Input("mode")]
        public Input<CloudySkySoftware.Pulumi.Constellix.Templates.ValueARequestValuePropertiesMode>? Mode { get; set; }

        [Input("values")]
        private InputList<Inputs.ValueARequestValuePropertiesValuesItemPropertiesArgs>? _values;
        public InputList<Inputs.ValueARequestValuePropertiesValuesItemPropertiesArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.ValueARequestValuePropertiesValuesItemPropertiesArgs>());
            set => _values = value;
        }

        public ValueARequestValuePropertiesArgs()
        {
        }
        public static new ValueARequestValuePropertiesArgs Empty => new ValueARequestValuePropertiesArgs();
    }
}
