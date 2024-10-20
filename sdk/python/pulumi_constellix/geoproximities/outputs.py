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

__all__ = [
    'DataProperties',
    'DataPropertiesLinksProperties',
    'Geoproximity',
    'GeoproximityLinksProperties',
    'ListMetadata',
    'ListMetadataLinksProperties',
    'ListMetadataPaginationProperties',
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
class Geoproximity(dict):
    """
    Geo Proximity Location
    """
    def __init__(__self__, *,
                 city: Optional[int] = None,
                 country: Optional[str] = None,
                 id: Optional[int] = None,
                 latitude: Optional[float] = None,
                 links: Optional['outputs.GeoproximityLinksProperties'] = None,
                 longitude: Optional[float] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None):
        """
        Geo Proximity Location
        :param int city: ID of the city
        :param str country: 2 digit ISO country code
        :param int id: The unique ID for the Geo Proximity location
        :param float latitude: Latitude of the location
        :param 'GeoproximityLinksProperties' links: Links for domain objects
        :param float longitude: Longitude of the location
        :param str name: The name of the Geo Proximity location
        :param str region: Region, state or province code
        """
        if city is not None:
            pulumi.set(__self__, "city", city)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if latitude is not None:
            pulumi.set(__self__, "latitude", latitude)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if longitude is not None:
            pulumi.set(__self__, "longitude", longitude)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def city(self) -> Optional[int]:
        """
        ID of the city
        """
        return pulumi.get(self, "city")

    @property
    @pulumi.getter
    def country(self) -> Optional[str]:
        """
        2 digit ISO country code
        """
        return pulumi.get(self, "country")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        The unique ID for the Geo Proximity location
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def latitude(self) -> Optional[float]:
        """
        Latitude of the location
        """
        return pulumi.get(self, "latitude")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.GeoproximityLinksProperties']:
        """
        Links for domain objects
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def longitude(self) -> Optional[float]:
        """
        Longitude of the location
        """
        return pulumi.get(self, "longitude")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the Geo Proximity location
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        Region, state or province code
        """
        return pulumi.get(self, "region")


@pulumi.output_type
class GeoproximityLinksProperties(dict):
    """
    Links for domain objects
    """
    def __init__(__self__, *,
                 self: Optional[str] = None):
        """
        Links for domain objects
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


