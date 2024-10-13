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

__all__ = ['ContactListSlackWebhookArgs', 'ContactListSlackWebhook']

@pulumi.input_type
class ContactListSlackWebhookArgs:
    def __init__(__self__, *,
                 channel: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 webhook: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ContactListSlackWebhook resource.
        :param pulumi.Input[str] channel: Slack channel to send notifications to
        :param pulumi.Input[str] id: The ID of the Contact List
        :param pulumi.Input[str] webhook: Incoming Webhook URL
        """
        if channel is not None:
            pulumi.set(__self__, "channel", channel)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if webhook is not None:
            pulumi.set(__self__, "webhook", webhook)

    @property
    @pulumi.getter
    def channel(self) -> Optional[pulumi.Input[str]]:
        """
        Slack channel to send notifications to
        """
        return pulumi.get(self, "channel")

    @channel.setter
    def channel(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "channel", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Contact List
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def webhook(self) -> Optional[pulumi.Input[str]]:
        """
        Incoming Webhook URL
        """
        return pulumi.get(self, "webhook")

    @webhook.setter
    def webhook(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webhook", value)


class ContactListSlackWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 webhook: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ContactListSlackWebhook resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] channel: Slack channel to send notifications to
        :param pulumi.Input[str] id: The ID of the Contact List
        :param pulumi.Input[str] webhook: Incoming Webhook URL
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContactListSlackWebhookArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ContactListSlackWebhook resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ContactListSlackWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContactListSlackWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 webhook: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContactListSlackWebhookArgs.__new__(ContactListSlackWebhookArgs)

            __props__.__dict__["channel"] = channel
            __props__.__dict__["id"] = id
            __props__.__dict__["webhook"] = webhook
            __props__.__dict__["data"] = None
        super(ContactListSlackWebhook, __self__).__init__(
            'constellix:contactlists:ContactListSlackWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ContactListSlackWebhook':
        """
        Get an existing ContactListSlackWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ContactListSlackWebhookArgs.__new__(ContactListSlackWebhookArgs)

        __props__.__dict__["channel"] = None
        __props__.__dict__["data"] = None
        __props__.__dict__["webhook"] = None
        return ContactListSlackWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def channel(self) -> pulumi.Output[Optional[str]]:
        """
        Slack channel to send notifications to
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional['outputs.DataProperties']]:
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def webhook(self) -> pulumi.Output[Optional[str]]:
        """
        Incoming Webhook URL
        """
        return pulumi.get(self, "webhook")

