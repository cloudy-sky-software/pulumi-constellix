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
    'Analytics',
    'AnalyticsIntervalProperties',
    'AnalyticsLinksProperties',
    'AnalyticsStatsProperties',
    'AnalyticsValuesItemProperties',
]

@pulumi.output_type
class Analytics(dict):
    """
    Analytics for your account
    """
    def __init__(__self__, *,
                 end: Optional[str] = None,
                 interval: Optional['outputs.AnalyticsIntervalProperties'] = None,
                 links: Optional['outputs.AnalyticsLinksProperties'] = None,
                 start: Optional[str] = None,
                 stats: Optional['outputs.AnalyticsStatsProperties'] = None,
                 values: Optional[Sequence['outputs.AnalyticsValuesItemProperties']] = None):
        """
        Analytics for your account
        :param str end: The end date for the analytics
        :param 'AnalyticsIntervalProperties' interval: Details about the interval between time periods in the analytics
        :param str start: The start date for the analytics
        :param 'AnalyticsStatsProperties' stats: Some statistics for these analytics
        :param Sequence['AnalyticsValuesItemProperties'] values: Query counts for all dates within the requested range
        """
        if end is not None:
            pulumi.set(__self__, "end", end)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if start is not None:
            pulumi.set(__self__, "start", start)
        if stats is not None:
            pulumi.set(__self__, "stats", stats)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def end(self) -> Optional[str]:
        """
        The end date for the analytics
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def interval(self) -> Optional['outputs.AnalyticsIntervalProperties']:
        """
        Details about the interval between time periods in the analytics
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.AnalyticsLinksProperties']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def start(self) -> Optional[str]:
        """
        The start date for the analytics
        """
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def stats(self) -> Optional['outputs.AnalyticsStatsProperties']:
        """
        Some statistics for these analytics
        """
        return pulumi.get(self, "stats")

    @property
    @pulumi.getter
    def values(self) -> Optional[Sequence['outputs.AnalyticsValuesItemProperties']]:
        """
        Query counts for all dates within the requested range
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class AnalyticsIntervalProperties(dict):
    """
    Details about the interval between time periods in the analytics
    """
    def __init__(__self__, *,
                 max: Optional[int] = None,
                 mean: Optional[float] = None,
                 min: Optional[int] = None):
        """
        Details about the interval between time periods in the analytics
        :param int max: The maximum number of seconds between time periods
        :param float mean: The mean average number of seconds between time periods
        :param int min: The minimum number of seconds between time periods
        """
        if max is not None:
            pulumi.set(__self__, "max", max)
        if mean is not None:
            pulumi.set(__self__, "mean", mean)
        if min is not None:
            pulumi.set(__self__, "min", min)

    @property
    @pulumi.getter
    def max(self) -> Optional[int]:
        """
        The maximum number of seconds between time periods
        """
        return pulumi.get(self, "max")

    @property
    @pulumi.getter
    def mean(self) -> Optional[float]:
        """
        The mean average number of seconds between time periods
        """
        return pulumi.get(self, "mean")

    @property
    @pulumi.getter
    def min(self) -> Optional[int]:
        """
        The minimum number of seconds between time periods
        """
        return pulumi.get(self, "min")


@pulumi.output_type
class AnalyticsLinksProperties(dict):
    def __init__(__self__, *,
                 self: Optional[str] = None):
        """
        :param str self: The URL for these analytics
        """
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def self(self) -> Optional[str]:
        """
        The URL for these analytics
        """
        return pulumi.get(self, "self")


@pulumi.output_type
class AnalyticsStatsProperties(dict):
    """
    Some statistics for these analytics
    """
    def __init__(__self__, *,
                 count: Optional[int] = None,
                 max: Optional[int] = None,
                 mean: Optional[float] = None,
                 min: Optional[int] = None,
                 sum: Optional[int] = None):
        """
        Some statistics for these analytics
        :param int count: The number of time periods
        :param int max: The maximum number of queries for one time period
        :param float mean: The mean average number of queries per time period
        :param int min: The minimum number of queries for one time period
        :param int sum: The sum of queries for this date range
        """
        if count is not None:
            pulumi.set(__self__, "count", count)
        if max is not None:
            pulumi.set(__self__, "max", max)
        if mean is not None:
            pulumi.set(__self__, "mean", mean)
        if min is not None:
            pulumi.set(__self__, "min", min)
        if sum is not None:
            pulumi.set(__self__, "sum", sum)

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        """
        The number of time periods
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter
    def max(self) -> Optional[int]:
        """
        The maximum number of queries for one time period
        """
        return pulumi.get(self, "max")

    @property
    @pulumi.getter
    def mean(self) -> Optional[float]:
        """
        The mean average number of queries per time period
        """
        return pulumi.get(self, "mean")

    @property
    @pulumi.getter
    def min(self) -> Optional[int]:
        """
        The minimum number of queries for one time period
        """
        return pulumi.get(self, "min")

    @property
    @pulumi.getter
    def sum(self) -> Optional[int]:
        """
        The sum of queries for this date range
        """
        return pulumi.get(self, "sum")


@pulumi.output_type
class AnalyticsValuesItemProperties(dict):
    """
    Query count for a particular date
    """
    def __init__(__self__, *,
                 date: Optional[str] = None,
                 value: Optional[int] = None):
        """
        Query count for a particular date
        """
        if date is not None:
            pulumi.set(__self__, "date", date)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def date(self) -> Optional[str]:
        return pulumi.get(self, "date")

    @property
    @pulumi.getter
    def value(self) -> Optional[int]:
        return pulumi.get(self, "value")


