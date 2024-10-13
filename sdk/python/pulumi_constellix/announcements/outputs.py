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
    'Announcement',
    'AnnouncementLinksProperties',
    'ListMetadata',
    'ListMetadataLinksProperties',
    'ListMetadataPaginationProperties',
]

@pulumi.output_type
class Announcement(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 link: Optional[str] = None,
                 links: Optional['outputs.AnnouncementLinksProperties'] = None,
                 title: Optional[str] = None,
                 type: Optional['AnnouncementType'] = None):
        """
        :param int id: A numeric ID for the Announcement
        :param str link: A link for the announcement
        :param 'AnnouncementLinksProperties' links: Links for announcements
        :param str title: The announcement
        :param 'AnnouncementType' type: The type of Announcement
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if link is not None:
            pulumi.set(__self__, "link", link)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        A numeric ID for the Announcement
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def link(self) -> Optional[str]:
        """
        A link for the announcement
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.AnnouncementLinksProperties']:
        """
        Links for announcements
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        """
        The announcement
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> Optional['AnnouncementType']:
        """
        The type of Announcement
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class AnnouncementLinksProperties(dict):
    """
    Links for announcements
    """
    def __init__(__self__, *,
                 self: Optional[str] = None):
        """
        Links for announcements
        """
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def self(self) -> Optional[str]:
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


