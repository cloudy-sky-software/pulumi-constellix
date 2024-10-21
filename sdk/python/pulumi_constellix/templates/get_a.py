# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetA',
    'AwaitableGetA',
    'get_a',
    'get_a_output',
]

@pulumi.output_type
class GetA:
    def __init__(__self__, contacts=None, domain=None, enabled=None, geo_failover=None, geoproximity=None, id=None, ipfilter=None, ipfilter_drop=None, last_values=None, links=None, mode=None, name=None, notes=None, region=None, skip_lookup=None, template=None, ttl=None, type=None, value=None):
        if contacts and not isinstance(contacts, dict):
            raise TypeError("Expected argument 'contacts' to be a dict")
        pulumi.set(__self__, "contacts", contacts)
        if domain and not isinstance(domain, dict):
            raise TypeError("Expected argument 'domain' to be a dict")
        pulumi.set(__self__, "domain", domain)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if geo_failover and not isinstance(geo_failover, bool):
            raise TypeError("Expected argument 'geo_failover' to be a bool")
        pulumi.set(__self__, "geo_failover", geo_failover)
        if geoproximity and not isinstance(geoproximity, dict):
            raise TypeError("Expected argument 'geoproximity' to be a dict")
        pulumi.set(__self__, "geoproximity", geoproximity)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if ipfilter and not isinstance(ipfilter, dict):
            raise TypeError("Expected argument 'ipfilter' to be a dict")
        pulumi.set(__self__, "ipfilter", ipfilter)
        if ipfilter_drop and not isinstance(ipfilter_drop, bool):
            raise TypeError("Expected argument 'ipfilter_drop' to be a bool")
        pulumi.set(__self__, "ipfilter_drop", ipfilter_drop)
        if last_values and not isinstance(last_values, dict):
            raise TypeError("Expected argument 'last_values' to be a dict")
        pulumi.set(__self__, "last_values", last_values)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if skip_lookup and not isinstance(skip_lookup, bool):
            raise TypeError("Expected argument 'skip_lookup' to be a bool")
        pulumi.set(__self__, "skip_lookup", skip_lookup)
        if template and not isinstance(template, dict):
            raise TypeError("Expected argument 'template' to be a dict")
        pulumi.set(__self__, "template", template)
        if ttl and not isinstance(ttl, int):
            raise TypeError("Expected argument 'ttl' to be a int")
        pulumi.set(__self__, "ttl", ttl)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, dict):
            raise TypeError("Expected argument 'value' to be a dict")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def contacts(self) -> Optional['outputs.SimpleContactlist']:
        """
        A simple version of a contact list when inclued with other resources
        """
        return pulumi.get(self, "contacts")

    @property
    @pulumi.getter
    def domain(self) -> Optional['outputs.SimpleDomain']:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="geoFailover")
    def geo_failover(self) -> Optional[bool]:
        """
        Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
        """
        return pulumi.get(self, "geo_failover")

    @property
    @pulumi.getter
    def geoproximity(self) -> Optional['outputs.SimpleGeoproximity']:
        """
        Geo Proximity Location
        """
        return pulumi.get(self, "geoproximity")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        A unique ID for this domain record
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ipfilter(self) -> Optional['outputs.SimpleIpfilter']:
        return pulumi.get(self, "ipfilter")

    @property
    @pulumi.getter(name="ipfilterDrop")
    def ipfilter_drop(self) -> Optional[bool]:
        """
        If the requesting IP matches the IP filter, don't return a response
        """
        return pulumi.get(self, "ipfilter_drop")

    @property
    @pulumi.getter(name="lastValues")
    def last_values(self) -> Optional['outputs.GetAPropertiesLastValuesProperties']:
        """
        The previous values of the record in the different modes
        """
        return pulumi.get(self, "last_values")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.TemplaterecordLinksProperties']:
        """
        Links for the domain record
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def mode(self) -> Optional['GetAPropertiesMode']:
        """
        How the record should work
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the record
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        """
        A note about the record. Max 512 characters.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def region(self) -> Optional['RecordRegion']:
        """
        The region for this record
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="skipLookup")
    def skip_lookup(self) -> Optional[bool]:
        """
        Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
        """
        return pulumi.get(self, "skip_lookup")

    @property
    @pulumi.getter
    def template(self) -> Optional['outputs.SimpleTemplate']:
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def ttl(self) -> Optional[int]:
        """
        The time to live in seconds for this record. must be between 0 and 2147483647
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> Optional['ValueAType']:
        """
        The type of record
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional['outputs.ValueAValue']:
        return pulumi.get(self, "value")


class AwaitableGetA(GetA):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetA(
            contacts=self.contacts,
            domain=self.domain,
            enabled=self.enabled,
            geo_failover=self.geo_failover,
            geoproximity=self.geoproximity,
            id=self.id,
            ipfilter=self.ipfilter,
            ipfilter_drop=self.ipfilter_drop,
            last_values=self.last_values,
            links=self.links,
            mode=self.mode,
            name=self.name,
            notes=self.notes,
            region=self.region,
            skip_lookup=self.skip_lookup,
            template=self.template,
            ttl=self.ttl,
            type=self.type,
            value=self.value)


def get_a(id: Optional[str] = None,
          template_id: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetA:
    """
    Use this data source to access information about an existing resource.

    :param str id: The ID of the record
    :param str template_id: The ID of the template object
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['templateId'] = template_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('constellix:templates:getA', __args__, opts=opts, typ=GetA).value

    return AwaitableGetA(
        contacts=pulumi.get(__ret__, 'contacts'),
        domain=pulumi.get(__ret__, 'domain'),
        enabled=pulumi.get(__ret__, 'enabled'),
        geo_failover=pulumi.get(__ret__, 'geo_failover'),
        geoproximity=pulumi.get(__ret__, 'geoproximity'),
        id=pulumi.get(__ret__, 'id'),
        ipfilter=pulumi.get(__ret__, 'ipfilter'),
        ipfilter_drop=pulumi.get(__ret__, 'ipfilter_drop'),
        last_values=pulumi.get(__ret__, 'last_values'),
        links=pulumi.get(__ret__, 'links'),
        mode=pulumi.get(__ret__, 'mode'),
        name=pulumi.get(__ret__, 'name'),
        notes=pulumi.get(__ret__, 'notes'),
        region=pulumi.get(__ret__, 'region'),
        skip_lookup=pulumi.get(__ret__, 'skip_lookup'),
        template=pulumi.get(__ret__, 'template'),
        ttl=pulumi.get(__ret__, 'ttl'),
        type=pulumi.get(__ret__, 'type'),
        value=pulumi.get(__ret__, 'value'))
def get_a_output(id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetA]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The ID of the record
    :param str template_id: The ID of the template object
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['templateId'] = template_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('constellix:templates:getA', __args__, opts=opts, typ=GetA)
    return __ret__.apply(lambda __response__: GetA(
        contacts=pulumi.get(__response__, 'contacts'),
        domain=pulumi.get(__response__, 'domain'),
        enabled=pulumi.get(__response__, 'enabled'),
        geo_failover=pulumi.get(__response__, 'geo_failover'),
        geoproximity=pulumi.get(__response__, 'geoproximity'),
        id=pulumi.get(__response__, 'id'),
        ipfilter=pulumi.get(__response__, 'ipfilter'),
        ipfilter_drop=pulumi.get(__response__, 'ipfilter_drop'),
        last_values=pulumi.get(__response__, 'last_values'),
        links=pulumi.get(__response__, 'links'),
        mode=pulumi.get(__response__, 'mode'),
        name=pulumi.get(__response__, 'name'),
        notes=pulumi.get(__response__, 'notes'),
        region=pulumi.get(__response__, 'region'),
        skip_lookup=pulumi.get(__response__, 'skip_lookup'),
        template=pulumi.get(__response__, 'template'),
        ttl=pulumi.get(__response__, 'ttl'),
        type=pulumi.get(__response__, 'type'),
        value=pulumi.get(__response__, 'value')))
