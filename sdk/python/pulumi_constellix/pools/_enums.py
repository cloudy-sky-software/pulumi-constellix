# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PoolType',
    'PoolValuesItemPropertiesPolicy',
    'PoolindexType',
    'PoolitoConfigPropertiesDeviationAllowance',
    'PoolitoConfigPropertiesHandicapFactor',
    'PoolitoConfigPropertiesMonitoringRegion',
    'PoolitoConfigPropertiesPeriod',
    'SimpleDomainStatus',
    'Type',
    'ValuesItemPropertiesPolicy',
]


class PoolType(str, Enum):
    """
    The type of pool
    """
    A = "A"
    AAAA = "AAAA"
    CNAME = "CNAME"


class PoolValuesItemPropertiesPolicy(str, Enum):
    """
    The failover/check policy for this value
    """
    FOLLOW_SONAR = "follow_sonar"
    ALWAYS_OFF = "always_off"
    ALWAYS_ON = "always_on"
    OFF_ON_FAILURE = "off_on_failure"


class PoolindexType(str, Enum):
    """
    The type of pool
    """
    A = "A"
    AAAA = "AAAA"
    CNAME = "CNAME"


class PoolitoConfigPropertiesDeviationAllowance(int, Enum):
    """
    Percentage of how much is the response time allowed to deviate?
    """
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_10_ = 10
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_20_ = 20
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_30_ = 30
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_40_ = 40
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_50_ = 50
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_60_ = 60
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_70_ = 70
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_80_ = 80
    POOLITO_CONFIG_PROPERTIES_DEVIATION_ALLOWANCE_D_FLOAT64_90_ = 90


class PoolitoConfigPropertiesHandicapFactor(str, Enum):
    NONE = "none"
    PERCENT = "percent"
    SPEED = "speed"


class PoolitoConfigPropertiesMonitoringRegion(str, Enum):
    """
    Where monitoring should be performed from
    """
    WORLD = "world"
    ASIAPAC = "asiapac"
    EUROPE = "europe"
    NACENTRAL = "nacentral"
    NAEAST = "naeast"
    NAWEST = "nawest"
    OCEANIA = "oceania"
    SOUTHAMERICA = "southamerica"


class PoolitoConfigPropertiesPeriod(int, Enum):
    """
    The number of seconds between each check
    """
    POOLITO_CONFIG_PROPERTIES_PERIOD_D_FLOAT64_30_ = 30
    POOLITO_CONFIG_PROPERTIES_PERIOD_D_FLOAT64_60_ = 60
    POOLITO_CONFIG_PROPERTIES_PERIOD_D_FLOAT64_120_ = 120
    POOLITO_CONFIG_PROPERTIES_PERIOD_D_FLOAT64_180_ = 180
    POOLITO_CONFIG_PROPERTIES_PERIOD_D_FLOAT64_240_ = 240
    POOLITO_CONFIG_PROPERTIES_PERIOD_D_FLOAT64_300_ = 300


class SimpleDomainStatus(str, Enum):
    ACTIVE = "ACTIVE"
    SUSPENDED = "SUSPENDED"
    TERMINATED = "TERMINATED"


class Type(str, Enum):
    """
    The type of pool, either A, AAAA or CNAME
    """
    A = "A"
    AAAA = "AAAA"
    CNAME = "CNAME"


class ValuesItemPropertiesPolicy(str, Enum):
    """
    The failover/check policy for this value
    """
    FOLLOW_SONAR = "follow_sonar"
    ALWAYS_OFF = "always_off"
    ALWAYS_ON = "always_on"
    OFF_ON_FAILURE = "off_on_failure"
