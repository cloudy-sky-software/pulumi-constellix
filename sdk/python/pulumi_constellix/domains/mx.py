# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['MxArgs', 'Mx']

@pulumi.input_type
class MxArgs:
    def __init__(__self__, *,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 geo_failover: Optional[pulumi.Input[bool]] = None,
                 geoproximity: Optional[pulumi.Input[int]] = None,
                 ipfilter: Optional[pulumi.Input[int]] = None,
                 ipfilter_drop: Optional[pulumi.Input[bool]] = None,
                 mode: Optional[pulumi.Input['MxPropertiesMode']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['RecordCreateDetailsRegion']] = None,
                 skip_lookup: Optional[pulumi.Input[bool]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input['MxPropertiesType']] = None,
                 value: Optional[pulumi.Input[Sequence[pulumi.Input['ValueMxValueItemPropertiesArgs']]]] = None):
        """
        The set of arguments for constructing a Mx resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] contacts: Contact lists to be notified if a failover happens in a failover mode.
        :param pulumi.Input[str] domain_id: The ID of the domain object
        :param pulumi.Input[bool] enabled: Whether the record is enabled
        :param pulumi.Input[bool] geo_failover: Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        :param pulumi.Input[int] geoproximity: The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
        :param pulumi.Input[int] ipfilter: The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
        :param pulumi.Input[bool] ipfilter_drop: If the requesting IP matches the IP filter, don't return a response
        :param pulumi.Input['MxPropertiesMode'] mode: The current mode for this record
        :param pulumi.Input[str] name: The name for the record
        :param pulumi.Input[str] notes: A description of the record. It must be 512 characters or less.
        :param pulumi.Input['RecordCreateDetailsRegion'] region: Optional region for this record. Will default to 'default'.
        :param pulumi.Input[bool] skip_lookup: Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        :param pulumi.Input[int] ttl: How long DNS servers should cache the record for
        :param pulumi.Input['MxPropertiesType'] type: The type of record
        :param pulumi.Input[Sequence[pulumi.Input['ValueMxValueItemPropertiesArgs']]] value: Standard record mode
        """
        if contacts is not None:
            pulumi.set(__self__, "contacts", contacts)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if geo_failover is not None:
            pulumi.set(__self__, "geo_failover", geo_failover)
        if geoproximity is not None:
            pulumi.set(__self__, "geoproximity", geoproximity)
        if ipfilter is not None:
            pulumi.set(__self__, "ipfilter", ipfilter)
        if ipfilter_drop is not None:
            pulumi.set(__self__, "ipfilter_drop", ipfilter_drop)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if skip_lookup is not None:
            pulumi.set(__self__, "skip_lookup", skip_lookup)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def contacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Contact lists to be notified if a failover happens in a failover mode.
        """
        return pulumi.get(self, "contacts")

    @contacts.setter
    def contacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "contacts", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the domain object
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the record is enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="geoFailover")
    def geo_failover(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        """
        return pulumi.get(self, "geo_failover")

    @geo_failover.setter
    def geo_failover(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "geo_failover", value)

    @property
    @pulumi.getter
    def geoproximity(self) -> Optional[pulumi.Input[int]]:
        """
        The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
        """
        return pulumi.get(self, "geoproximity")

    @geoproximity.setter
    def geoproximity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "geoproximity", value)

    @property
    @pulumi.getter
    def ipfilter(self) -> Optional[pulumi.Input[int]]:
        """
        The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
        """
        return pulumi.get(self, "ipfilter")

    @ipfilter.setter
    def ipfilter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipfilter", value)

    @property
    @pulumi.getter(name="ipfilterDrop")
    def ipfilter_drop(self) -> Optional[pulumi.Input[bool]]:
        """
        If the requesting IP matches the IP filter, don't return a response
        """
        return pulumi.get(self, "ipfilter_drop")

    @ipfilter_drop.setter
    def ipfilter_drop(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ipfilter_drop", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input['MxPropertiesMode']]:
        """
        The current mode for this record
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input['MxPropertiesMode']]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the record
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the record. It must be 512 characters or less.
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input['RecordCreateDetailsRegion']]:
        """
        Optional region for this record. Will default to 'default'.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input['RecordCreateDetailsRegion']]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="skipLookup")
    def skip_lookup(self) -> Optional[pulumi.Input[bool]]:
        """
        Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        """
        return pulumi.get(self, "skip_lookup")

    @skip_lookup.setter
    def skip_lookup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_lookup", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        How long DNS servers should cache the record for
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['MxPropertiesType']]:
        """
        The type of record
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['MxPropertiesType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ValueMxValueItemPropertiesArgs']]]]:
        """
        Standard record mode
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ValueMxValueItemPropertiesArgs']]]]):
        pulumi.set(self, "value", value)


class Mx(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 geo_failover: Optional[pulumi.Input[bool]] = None,
                 geoproximity: Optional[pulumi.Input[int]] = None,
                 ipfilter: Optional[pulumi.Input[int]] = None,
                 ipfilter_drop: Optional[pulumi.Input[bool]] = None,
                 mode: Optional[pulumi.Input['MxPropertiesMode']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['RecordCreateDetailsRegion']] = None,
                 skip_lookup: Optional[pulumi.Input[bool]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input['MxPropertiesType']] = None,
                 value: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ValueMxValueItemPropertiesArgs', 'ValueMxValueItemPropertiesArgsDict']]]]] = None,
                 __props__=None):
        """
        Create a Mx resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] contacts: Contact lists to be notified if a failover happens in a failover mode.
        :param pulumi.Input[str] domain_id: The ID of the domain object
        :param pulumi.Input[bool] enabled: Whether the record is enabled
        :param pulumi.Input[bool] geo_failover: Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        :param pulumi.Input[int] geoproximity: The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
        :param pulumi.Input[int] ipfilter: The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
        :param pulumi.Input[bool] ipfilter_drop: If the requesting IP matches the IP filter, don't return a response
        :param pulumi.Input['MxPropertiesMode'] mode: The current mode for this record
        :param pulumi.Input[str] name: The name for the record
        :param pulumi.Input[str] notes: A description of the record. It must be 512 characters or less.
        :param pulumi.Input['RecordCreateDetailsRegion'] region: Optional region for this record. Will default to 'default'.
        :param pulumi.Input[bool] skip_lookup: Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        :param pulumi.Input[int] ttl: How long DNS servers should cache the record for
        :param pulumi.Input['MxPropertiesType'] type: The type of record
        :param pulumi.Input[Sequence[pulumi.Input[Union['ValueMxValueItemPropertiesArgs', 'ValueMxValueItemPropertiesArgsDict']]]] value: Standard record mode
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MxArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Mx resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MxArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MxArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 geo_failover: Optional[pulumi.Input[bool]] = None,
                 geoproximity: Optional[pulumi.Input[int]] = None,
                 ipfilter: Optional[pulumi.Input[int]] = None,
                 ipfilter_drop: Optional[pulumi.Input[bool]] = None,
                 mode: Optional[pulumi.Input['MxPropertiesMode']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['RecordCreateDetailsRegion']] = None,
                 skip_lookup: Optional[pulumi.Input[bool]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input['MxPropertiesType']] = None,
                 value: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ValueMxValueItemPropertiesArgs', 'ValueMxValueItemPropertiesArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MxArgs.__new__(MxArgs)

            __props__.__dict__["contacts"] = contacts
            __props__.__dict__["domain_id"] = domain_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["geo_failover"] = geo_failover
            __props__.__dict__["geoproximity"] = geoproximity
            __props__.__dict__["ipfilter"] = ipfilter
            __props__.__dict__["ipfilter_drop"] = ipfilter_drop
            __props__.__dict__["mode"] = mode
            __props__.__dict__["name"] = name
            __props__.__dict__["notes"] = notes
            __props__.__dict__["region"] = region
            __props__.__dict__["skip_lookup"] = skip_lookup
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["type"] = type
            __props__.__dict__["value"] = value
            __props__.__dict__["data"] = None
        super(Mx, __self__).__init__(
            'constellix:domains:Mx',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Mx':
        """
        Get an existing Mx resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MxArgs.__new__(MxArgs)

        __props__.__dict__["contacts"] = None
        __props__.__dict__["data"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["geo_failover"] = None
        __props__.__dict__["geoproximity"] = None
        __props__.__dict__["ipfilter"] = None
        __props__.__dict__["ipfilter_drop"] = None
        __props__.__dict__["mode"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notes"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["skip_lookup"] = None
        __props__.__dict__["ttl"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["value"] = None
        return Mx(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def contacts(self) -> pulumi.Output[Optional[Sequence[int]]]:
        """
        Contact lists to be notified if a failover happens in a failover mode.
        """
        return pulumi.get(self, "contacts")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional['outputs.DataProperties']]:
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the record is enabled
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="geoFailover")
    def geo_failover(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        """
        return pulumi.get(self, "geo_failover")

    @property
    @pulumi.getter
    def geoproximity(self) -> pulumi.Output[Optional[int]]:
        """
        The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
        """
        return pulumi.get(self, "geoproximity")

    @property
    @pulumi.getter
    def ipfilter(self) -> pulumi.Output[Optional[int]]:
        """
        The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
        """
        return pulumi.get(self, "ipfilter")

    @property
    @pulumi.getter(name="ipfilterDrop")
    def ipfilter_drop(self) -> pulumi.Output[Optional[bool]]:
        """
        If the requesting IP matches the IP filter, don't return a response
        """
        return pulumi.get(self, "ipfilter_drop")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional['MxPropertiesMode']]:
        """
        The current mode for this record
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name for the record
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the record. It must be 512 characters or less.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional['RecordCreateDetailsRegion']]:
        """
        Optional region for this record. Will default to 'default'.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="skipLookup")
    def skip_lookup(self) -> pulumi.Output[Optional[bool]]:
        """
        Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        """
        return pulumi.get(self, "skip_lookup")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        How long DNS servers should cache the record for
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional['MxPropertiesType']]:
        """
        The type of record
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[Sequence['outputs.ValueMxValueItemProperties']]]:
        """
        Standard record mode
        """
        return pulumi.get(self, "value")

