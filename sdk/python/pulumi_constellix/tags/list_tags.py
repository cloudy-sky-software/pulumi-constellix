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

__all__ = [
    'ListTagsProperties',
    'AwaitableListTagsProperties',
    'list_tags',
    'list_tags_output',
]

@pulumi.output_type
class ListTagsProperties:
    def __init__(__self__, data=None, meta=None):
        if data and not isinstance(data, list):
            raise TypeError("Expected argument 'data' to be a list")
        pulumi.set(__self__, "data", data)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def data(self) -> Optional[Sequence['outputs.Tag']]:
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def meta(self) -> Optional['outputs.ListMetadata']:
        """
        Metadata for list responses
        """
        return pulumi.get(self, "meta")


class AwaitableListTagsProperties(ListTagsProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListTagsProperties(
            data=self.data,
            meta=self.meta)


def list_tags(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListTagsProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('constellix:tags:listTags', __args__, opts=opts, typ=ListTagsProperties).value

    return AwaitableListTagsProperties(
        data=pulumi.get(__ret__, 'data'),
        meta=pulumi.get(__ret__, 'meta'))
def list_tags_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListTagsProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('constellix:tags:listTags', __args__, opts=opts, typ=ListTagsProperties)
    return __ret__.apply(lambda __response__: ListTagsProperties(
        data=pulumi.get(__response__, 'data'),
        meta=pulumi.get(__response__, 'meta')))
