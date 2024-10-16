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

__all__ = [
    'DataProperties',
    'DataPropertiesLinksProperties',
    'ListMetadata',
    'ListMetadataLinksProperties',
    'ListMetadataPaginationProperties',
    'ListTemplateRecordsPropertiesDataItem',
    'SimpleTemplate',
    'SimpleTemplateLinksProperties',
    'Template',
    'TemplateLinksProperties',
    'TemplaterecordPropertiesLinksProperties',
]

@pulumi.output_type
class DataProperties(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 links: Optional['outputs.DataPropertiesLinksProperties'] = None):
        """
        :param int id: The ID of the object
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if links is not None:
            pulumi.set(__self__, "links", links)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        The ID of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.DataPropertiesLinksProperties']:
        return pulumi.get(self, "links")


@pulumi.output_type
class DataPropertiesLinksProperties(dict):
    def __init__(__self__, *,
                 self: Optional[str] = None):
        """
        :param str self: The URL for the new object
        """
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def self(self) -> Optional[str]:
        """
        The URL for the new object
        """
        return pulumi.get(self, "self")


@pulumi.output_type
class ListMetadata(dict):
    """
    Metadata for list responses
    """
    def __init__(__self__, *,
                 links: Optional['outputs.ListMetadataLinksProperties'] = None,
                 pagination: Optional['outputs.ListMetadataPaginationProperties'] = None):
        """
        Metadata for list responses
        :param 'ListMetadataLinksProperties' links: Relevant links for this list
        :param 'ListMetadataPaginationProperties' pagination: Pagination details
        """
        if links is not None:
            pulumi.set(__self__, "links", links)
        if pagination is not None:
            pulumi.set(__self__, "pagination", pagination)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.ListMetadataLinksProperties']:
        """
        Relevant links for this list
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def pagination(self) -> Optional['outputs.ListMetadataPaginationProperties']:
        """
        Pagination details
        """
        return pulumi.get(self, "pagination")


@pulumi.output_type
class ListMetadataLinksProperties(dict):
    """
    Relevant links for this list
    """
    def __init__(__self__, *,
                 first: Optional[str] = None,
                 last: Optional[str] = None,
                 next: Optional[str] = None,
                 previous: Optional[str] = None,
                 self: Optional[str] = None):
        """
        Relevant links for this list
        """
        if first is not None:
            pulumi.set(__self__, "first", first)
        if last is not None:
            pulumi.set(__self__, "last", last)
        if next is not None:
            pulumi.set(__self__, "next", next)
        if previous is not None:
            pulumi.set(__self__, "previous", previous)
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def first(self) -> Optional[str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def previous(self) -> Optional[str]:
        return pulumi.get(self, "previous")

    @property
    @pulumi.getter
    def self(self) -> Optional[str]:
        return pulumi.get(self, "self")


@pulumi.output_type
class ListMetadataPaginationProperties(dict):
    """
    Pagination details
    """
    def __init__(__self__, *,
                 count: Optional[int] = None,
                 current_page: Optional[int] = None,
                 per_page: Optional[int] = None,
                 total: Optional[int] = None,
                 total_pages: Optional[int] = None):
        """
        Pagination details
        :param int count: The number of items in this page of the response
        :param int current_page: The current results page
        :param int per_page: The number of items per page
        :param int total: The total number of objects matching the query
        :param int total_pages: The total number of pages
        """
        if count is not None:
            pulumi.set(__self__, "count", count)
        if current_page is not None:
            pulumi.set(__self__, "current_page", current_page)
        if per_page is not None:
            pulumi.set(__self__, "per_page", per_page)
        if total is not None:
            pulumi.set(__self__, "total", total)
        if total_pages is not None:
            pulumi.set(__self__, "total_pages", total_pages)

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        """
        The number of items in this page of the response
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="currentPage")
    def current_page(self) -> Optional[int]:
        """
        The current results page
        """
        return pulumi.get(self, "current_page")

    @property
    @pulumi.getter(name="perPage")
    def per_page(self) -> Optional[int]:
        """
        The number of items per page
        """
        return pulumi.get(self, "per_page")

    @property
    @pulumi.getter
    def total(self) -> Optional[int]:
        """
        The total number of objects matching the query
        """
        return pulumi.get(self, "total")

    @property
    @pulumi.getter(name="totalPages")
    def total_pages(self) -> Optional[int]:
        """
        The total number of pages
        """
        return pulumi.get(self, "total_pages")


@pulumi.output_type
class ListTemplateRecordsPropertiesDataItem(dict):
    def __init__(__self__, *,
                 links: Optional['outputs.TemplaterecordPropertiesLinksProperties'] = None,
                 template: Optional['outputs.SimpleTemplate'] = None):
        """
        :param 'TemplaterecordPropertiesLinksProperties' links: Links for the domain record
        """
        if links is not None:
            pulumi.set(__self__, "links", links)
        if template is not None:
            pulumi.set(__self__, "template", template)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.TemplaterecordPropertiesLinksProperties']:
        """
        Links for the domain record
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def template(self) -> Optional['outputs.SimpleTemplate']:
        return pulumi.get(self, "template")


@pulumi.output_type
class SimpleTemplate(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 links: Optional['outputs.SimpleTemplateLinksProperties'] = None,
                 name: Optional[str] = None,
                 version: Optional[int] = None):
        """
        :param int id: The unique ID for this template
        :param 'SimpleTemplateLinksProperties' links: Links for the template
        :param str name: The name for the template
        :param int version: The version of the template resource
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        The unique ID for this template
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.SimpleTemplateLinksProperties']:
        """
        Links for the template
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name for the template
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        """
        The version of the template resource
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SimpleTemplateLinksProperties(dict):
    """
    Links for the template
    """
    def __init__(__self__, *,
                 records: Optional[str] = None,
                 self: Optional[str] = None):
        """
        Links for the template
        """
        if records is not None:
            pulumi.set(__self__, "records", records)
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def records(self) -> Optional[str]:
        return pulumi.get(self, "records")

    @property
    @pulumi.getter
    def self(self) -> Optional[str]:
        return pulumi.get(self, "self")


@pulumi.output_type
class Template(dict):
    """
    A domain template
    """
    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 geoip: Optional[bool] = None,
                 gtd: Optional[bool] = None,
                 id: Optional[int] = None,
                 links: Optional['outputs.TemplateLinksProperties'] = None,
                 name: Optional[str] = None,
                 updated_at: Optional[str] = None,
                 version: Optional[int] = None):
        """
        A domain template
        :param bool geoip: Is GeoIP functionality enabled for the template
        :param bool gtd: Is Global Traffic Director enabled for the template
        :param int id: The unique ID for this template
        :param 'TemplateLinksProperties' links: Links for the template
        :param str name: The name for the template
        :param int version: The version of the template resource
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if geoip is not None:
            pulumi.set(__self__, "geoip", geoip)
        if gtd is not None:
            pulumi.set(__self__, "gtd", gtd)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def geoip(self) -> Optional[bool]:
        """
        Is GeoIP functionality enabled for the template
        """
        return pulumi.get(self, "geoip")

    @property
    @pulumi.getter
    def gtd(self) -> Optional[bool]:
        """
        Is Global Traffic Director enabled for the template
        """
        return pulumi.get(self, "gtd")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        The unique ID for this template
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.TemplateLinksProperties']:
        """
        Links for the template
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name for the template
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[str]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        """
        The version of the template resource
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class TemplateLinksProperties(dict):
    """
    Links for the template
    """
    def __init__(__self__, *,
                 records: Optional[str] = None,
                 self: Optional[str] = None):
        """
        Links for the template
        """
        if records is not None:
            pulumi.set(__self__, "records", records)
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def records(self) -> Optional[str]:
        return pulumi.get(self, "records")

    @property
    @pulumi.getter
    def self(self) -> Optional[str]:
        return pulumi.get(self, "self")


@pulumi.output_type
class TemplaterecordPropertiesLinksProperties(dict):
    """
    Links for the domain record
    """
    def __init__(__self__, *,
                 self: Optional[str] = None):
        """
        Links for the domain record
        """
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def self(self) -> Optional[str]:
        return pulumi.get(self, "self")


