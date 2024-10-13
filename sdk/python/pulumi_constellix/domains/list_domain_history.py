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
    'ListDomainHistoryProperties',
    'AwaitableListDomainHistoryProperties',
    'list_domain_history',
    'list_domain_history_output',
]

@pulumi.output_type
class ListDomainHistoryProperties:
    def __init__(__self__, data=None, meta=None):
        if data and not isinstance(data, list):
            raise TypeError("Expected argument 'data' to be a list")
        pulumi.set(__self__, "data", data)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def data(self) -> Optional[Sequence['outputs.ListDomainHistoryPropertiesDataItem']]:
        """
        The domain history for this page
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def meta(self) -> Optional['outputs.ListMetadata']:
        """
        Metadata for list responses
        """
        return pulumi.get(self, "meta")


class AwaitableListDomainHistoryProperties(ListDomainHistoryProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDomainHistoryProperties(
            data=self.data,
            meta=self.meta)


def list_domain_history(domain_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDomainHistoryProperties:
    """
    Use this data source to access information about an existing resource.

    :param str domain_id: The ID of the domain object
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('constellix:domains:listDomainHistory', __args__, opts=opts, typ=ListDomainHistoryProperties).value

    return AwaitableListDomainHistoryProperties(
        data=pulumi.get(__ret__, 'data'),
        meta=pulumi.get(__ret__, 'meta'))
def list_domain_history_output(domain_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDomainHistoryProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str domain_id: The ID of the domain object
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('constellix:domains:listDomainHistory', __args__, opts=opts, typ=ListDomainHistoryProperties)
    return __ret__.apply(lambda __response__: ListDomainHistoryProperties(
        data=pulumi.get(__response__, 'data'),
        meta=pulumi.get(__response__, 'meta')))
