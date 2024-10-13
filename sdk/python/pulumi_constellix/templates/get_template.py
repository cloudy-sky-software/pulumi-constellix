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
    'GetTemplateProperties',
    'AwaitableGetTemplateProperties',
    'get_template',
    'get_template_output',
]

@pulumi.output_type
class GetTemplateProperties:
    def __init__(__self__, data=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def data(self) -> Optional['outputs.Template']:
        """
        A domain template
        """
        return pulumi.get(self, "data")


class AwaitableGetTemplateProperties(GetTemplateProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTemplateProperties(
            data=self.data)


def get_template(id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTemplateProperties:
    """
    Use this data source to access information about an existing resource.

    :param str id: The ID of the template object
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('constellix:templates:getTemplate', __args__, opts=opts, typ=GetTemplateProperties).value

    return AwaitableGetTemplateProperties(
        data=pulumi.get(__ret__, 'data'))
def get_template_output(id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTemplateProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The ID of the template object
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('constellix:templates:getTemplate', __args__, opts=opts, typ=GetTemplateProperties)
    return __ret__.apply(lambda __response__: GetTemplateProperties(
        data=pulumi.get(__response__, 'data')))
