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
from ._inputs import *

__all__ = ['VanityNameserverArgs', 'VanityNameserver']

@pulumi.input_type
class VanityNameserverArgs:
    def __init__(__self__, *,
                 default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nameserver_group: Optional[pulumi.Input['NameserverGroupPropertiesArgs']] = None,
                 nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VanityNameserver resource.
        :param pulumi.Input[bool] default: Is this the default nameserver for domains in the account
        :param pulumi.Input[str] name: A unique name for this vanity nameserver
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nameservers: The nameserver hostnames
        """
        if default is not None:
            pulumi.set(__self__, "default", default)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nameserver_group is not None:
            pulumi.set(__self__, "nameserver_group", nameserver_group)
        if nameservers is not None:
            pulumi.set(__self__, "nameservers", nameservers)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this the default nameserver for domains in the account
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for this vanity nameserver
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nameserverGroup")
    def nameserver_group(self) -> Optional[pulumi.Input['NameserverGroupPropertiesArgs']]:
        return pulumi.get(self, "nameserver_group")

    @nameserver_group.setter
    def nameserver_group(self, value: Optional[pulumi.Input['NameserverGroupPropertiesArgs']]):
        pulumi.set(self, "nameserver_group", value)

    @property
    @pulumi.getter
    def nameservers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The nameserver hostnames
        """
        return pulumi.get(self, "nameservers")

    @nameservers.setter
    def nameservers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nameservers", value)


class VanityNameserver(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nameserver_group: Optional[pulumi.Input[Union['NameserverGroupPropertiesArgs', 'NameserverGroupPropertiesArgsDict']]] = None,
                 nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        A rebranded nameserver using your own domain name.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] default: Is this the default nameserver for domains in the account
        :param pulumi.Input[str] name: A unique name for this vanity nameserver
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nameservers: The nameserver hostnames
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VanityNameserverArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A rebranded nameserver using your own domain name.

        :param str resource_name: The name of the resource.
        :param VanityNameserverArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VanityNameserverArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nameserver_group: Optional[pulumi.Input[Union['NameserverGroupPropertiesArgs', 'NameserverGroupPropertiesArgsDict']]] = None,
                 nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VanityNameserverArgs.__new__(VanityNameserverArgs)

            __props__.__dict__["default"] = default
            __props__.__dict__["name"] = name
            __props__.__dict__["nameserver_group"] = nameserver_group
            __props__.__dict__["nameservers"] = nameservers
            __props__.__dict__["data"] = None
        super(VanityNameserver, __self__).__init__(
            'constellix:vanitynameservers:VanityNameserver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VanityNameserver':
        """
        Get an existing VanityNameserver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VanityNameserverArgs.__new__(VanityNameserverArgs)

        __props__.__dict__["data"] = None
        __props__.__dict__["default"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["nameserver_group"] = None
        __props__.__dict__["nameservers"] = None
        return VanityNameserver(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional['outputs.DataProperties']]:
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def default(self) -> pulumi.Output[Optional[bool]]:
        """
        Is this the default nameserver for domains in the account
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        A unique name for this vanity nameserver
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameserverGroup")
    def nameserver_group(self) -> pulumi.Output[Optional['outputs.NameserverGroupProperties']]:
        return pulumi.get(self, "nameserver_group")

    @property
    @pulumi.getter
    def nameservers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The nameserver hostnames
        """
        return pulumi.get(self, "nameservers")

