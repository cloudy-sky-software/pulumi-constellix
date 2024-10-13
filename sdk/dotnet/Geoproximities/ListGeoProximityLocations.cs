// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CONSTELLIX.Geoproximities
{
    public static class ListGeoProximityLocations
    {
        public static Task<Outputs.ListGeoProximityLocationsProperties> InvokeAsync(ListGeoProximityLocationsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ListGeoProximityLocationsProperties>("constellix:geoproximities:listGeoProximityLocations", args ?? new ListGeoProximityLocationsArgs(), options.WithDefaults());

        public static Output<Outputs.ListGeoProximityLocationsProperties> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ListGeoProximityLocationsProperties>("constellix:geoproximities:listGeoProximityLocations", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListGeoProximityLocationsArgs : global::Pulumi.InvokeArgs
    {
        public ListGeoProximityLocationsArgs()
        {
        }
        public static new ListGeoProximityLocationsArgs Empty => new ListGeoProximityLocationsArgs();
    }
}
