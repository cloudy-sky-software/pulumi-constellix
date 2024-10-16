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
    'ListContactEmailProperties',
    'AwaitableListContactEmailProperties',
    'list_contact_email',
    'list_contact_email_output',
]

@pulumi.output_type
class ListContactEmailProperties:
    def __init__(__self__, data=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def data(self) -> Optional['outputs.ContactlistEmail']:
        return pulumi.get(self, "data")


class AwaitableListContactEmailProperties(ListContactEmailProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListContactEmailProperties(
            data=self.data)


def list_contact_email(contactlist_id: Optional[str] = None,
                       id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListContactEmailProperties:
    """
    Use this data source to access information about an existing resource.

    :param str contactlist_id: The ID of the Contact List
    :param str id: The ID of the email address
    """
    __args__ = dict()
    __args__['contactlistId'] = contactlist_id
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('constellix:contactlists:listContactEmail', __args__, opts=opts, typ=ListContactEmailProperties).value

    return AwaitableListContactEmailProperties(
        data=pulumi.get(__ret__, 'data'))
def list_contact_email_output(contactlist_id: Optional[pulumi.Input[str]] = None,
                              id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListContactEmailProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str contactlist_id: The ID of the Contact List
    :param str id: The ID of the email address
    """
    __args__ = dict()
    __args__['contactlistId'] = contactlist_id
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('constellix:contactlists:listContactEmail', __args__, opts=opts, typ=ListContactEmailProperties)
    return __ret__.apply(lambda __response__: ListContactEmailProperties(
        data=pulumi.get(__response__, 'data')))
