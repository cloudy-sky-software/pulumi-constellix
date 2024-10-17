import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

using Http;
using OpenAPI;

/**
 * Metadata for list responses
 */
model ListMetadata {
  /**
   * Relevant links for this list
   */
  links?: {
    @format("url")
    first?: string;

    @format("url")
    last?: string;

    @format("url")
    next?: string;

    @format("url")
    previous?: string;

    @format("url")
    self?: string;
  };

  /**
   * Pagination details
   */
  pagination?: {
    count?: integer;
    currentPage?: integer;
    perPage?: integer;
    total?: integer;
    totalPages?: integer;
  };
}

model link {
  /**
   * Links for the domain record
   */
  links?: {
    @format("url")
    self?: string;
  };
}

model a {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-a`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    failover?: {
      mode?: "normal" | "off" | "one-way";
      values?: unknown[];
    };
    pools?: `simple-pool`[];
    roundRobinFailover?: unknown[];
    standard?: {
      enabled?: boolean;

      @format("ipv4")
      value?: string;
    }[];
  };

  /**
   * How the record should work
   */
  mode?: "standard" | "failover" | "pools" | "roundRobinFailover";
}

model aaaa {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-aaaa`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    failover?: {
      mode?: "normal" | "off" | "one-way";
      values?: unknown[];
    };
    pools?: `simple-pool`[];
    roundRobinFailover?: unknown[];
    standard?: {
      enabled?: boolean;

      @format("ipv6")
      value?: string;
    }[];
  };

  /**
   * How the record should work
   */
  mode?: "standard" | "failover" | "pools" | "roundrobin-failover";
}

/**
 * Analytics for your account
 */
model analytics {
  /**
   * The end date for the analytics
   */
  end?: utcDateTime;

  /**
   * Details about the interval between time periods in the analytics
   */
  interval?: {
    max?: integer;
    mean?: numeric;
    min?: integer;
  };

  links?: {
    self?: string;
  };

  /**
   * The start date for the analytics
   */
  start?: utcDateTime;

  /**
   * Some statistics for these analytics
   */
  stats?: {
    count?: integer;
    max?: integer;
    mean?: numeric;
    min?: integer;
    sum?: integer;
  };

  /**
   * Query counts for all dates within the requested range
   */
  values?: {
    date?: utcDateTime;
    value?: integer;
  }[];
}

model aname {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-aname`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    failover?: {
      mode?: "normal" | "off" | "one-way";
      values?: unknown[];
    };
    pools?: `simple-pool`[];
    standard?: {
      enabled?: boolean;
      value?: string;
    }[];
  };

  /**
   * How the record should work
   */
  mode?: "standard" | "failover" | "pools";
}

model announcement {
  /**
   * A numeric ID for the Announcement
   */
  id?: integer;

  /**
   * A link for the announcement
   */
  link?: string;

  /**
   * Links for announcements
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The announcement
   */
  title?: string;

  /**
   * The type of Announcement
   */
  type: "DashboardLinks";
}

model caa {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-caa`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: items[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";
}

model cert {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-cert`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "CAA";
}

model cname {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-cname`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    failover?: {
      mode?: "normal" | "off" | "one-way";
      values?: unknown[];
    };
    pools?: `simple-pool`[];
    standard?: {
      enabled?: boolean;
      value?: string;
    }[];
  };

  /**
   * How the record should work
   */
  mode?: "standard" | "failover" | "pools";
}

/**
 * Lists of email addresses used for notifications and messages about domains and records.
 */
model contactlist {
  /**
   * The number of emails in this contact list
   */
  emailCount?: integer;

  /**
   * The emails in this list
   */
  emails?: {
    @format("email")
    address?: string;

    verified?: boolean;
  }[];

  /**
   * Unique ID for the contact list
   */
  id?: integer;

  /**
   * Links for the object
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * A name for this contact list
   */
  name?: string;
}

/**
 * Lists of email addresses used for notifications and messages about domains and records.
 */
model `contactlist-createupdate` {
  /**
   * The emails in this list
   */
  emails?: string[];

  /**
   * A name for this contact list
   */
  name?: string;
}

model `contactlist-email-createupdate` {
  /**
   * The email address
   */
  @format("email") address?: string;
}

model `contactlist-slack-createupdate` {
  /**
   * Slack channel to send notifications to
   */
  channel?: string;

  /**
   * Incoming Webhook URL
   */
  @format("url") webhook?: string;
}

model `contactlist-teams-createupdate` {
  /**
   * Teams channel to send notifications to
   */
  channel?: string;

  /**
   * Incoming Webhook URL
   */
  @format("url") webhook?: string;
}

model contactlist_email {
  @format("email") address?: string;
  contactlist?: {
    id?: integer;
    links?: {
      @format("url")
      self?: string;
    };
  };
  id?: integer;

  /**
   * Links for the email address
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * Has the email been verified or not
   */
  verified?: boolean;
}

model contactlist_slack {
  /**
   * The channel to send the message to
   */
  channel?: string;

  contactlist?: {
    id?: integer;
    links?: {
      @format("url")
      self?: string;
    };
  };
  id?: integer;

  /**
   * Links for the Slack webhook
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The inbound webhook URL for Slack
   */
  @format("url") webhook?: string;
}

model contactlist_teams {
  /**
   * The channel to send the message to
   */
  channel?: string;

  contactlist?: {
    id?: integer;
    links?: {
      @format("url")
      self?: string;
    };
  };
  id?: integer;

  /**
   * Links for the MS Teams Webhook
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The inbound webhook URL for MS Teams
   */
  @format("url") webhook?: string;
}

/**
 * Lists of email addresses used for notifications and messages about domains and records.
 */
model contactlistwithoutemails {
  /**
   * The number of emails in this contact list
   */
  emailCount?: integer;

  /**
   * Unique ID for the contact list
   */
  id?: integer;

  /**
   * Links for the object
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * A name for this contact list
   */
  name?: string;
}

model domain {
  /**
   * An array of contact lists that will be notified when records in this domain change.
   */
  contacts?: `simple-contactlist`[];

  createdAt?: utcDateTime;

  /**
   * Is the domain enabled
   */
  enabled?: boolean;

  /**
   * Is GeoIP functionality enabled for the domain
   */
  geoip?: boolean;

  /**
   * Is Global Traffic Director enabled for the domain
   */
  gtd?: boolean;

  /**
   * A unique numeric ID for this domain
   */
  id?: integer;

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    records?: string;

    @format("url")
    self?: string;
  };

  /**
   * The name of the domain
   */
  name?: string;

  /**
   * The nameservers for the domain
   */
  nameservers?: string[];

  /**
   * A note for the domain
   */
  note?: string;

  soa?: soa;
  status?: "ACTIVE" | "SUSPENDED" | "TERMINATED";

  /**
   * An array of tags for this domain.
   */
  tags?: tag[];

  template?: `simple-template`;
  updatedAt?: utcDateTime;
  vanityNameserver?: `simple-vanitynameserver`;

  /**
   * The version of the domain resource
   */
  version?: integer;
}

model `domain-create` {
  /**
   * Contactlists to be notified if the domain is updated
   */
  contacts?: integer[];

  /**
   * Is the domain enabled
   */
  enabled?: boolean;

  /**
   * Is GeoIP functionality enabled for the domain
   */
  geoip?: boolean;

  /**
   * Is Global Traffic Director enabled for the domain
   */
  gtd?: boolean;

  /**
   * The name of the domain
   */
  name?: string;

  /**
   * A note for the domain
   */
  note?: string;

  /**
   * The SOA details for the domain
   */
  soa?: {
    email?: string;
    expire?: integer;
    negativeCache?: integer;
    primaryNameserver?: string;
    refresh?: integer;
    retry?: integer;
    ttl?: integer;
  };

  /**
   * The numeric IDs of tags you want to apply to this domain
   */
  tags?: integer[];

  /**
   * The template to use for creating this domain. It will be linked to this template so any changes made to the template will apply to this domain.
   */
  template?: integer;

  /**
   * The vanity nameserver to use for this domain.
   */
  vanityNameserver?: integer;
}

model `domain-update` {
  /**
   * Contactlists to be notified if the domain is updated
   */
  contacts?: integer[];

  /**
   * Is the domain enabled
   */
  enabled?: boolean;

  /**
   * Is GeoIP functionality enabled for the domain
   */
  geoip?: boolean;

  /**
   * Is Global Traffic Director enabled for the domain
   */
  gtd?: boolean;

  /**
   * A note for the domain
   */
  note?: string;

  /**
   * The SOA details for the domain
   */
  soa?: {
    email?: string;
    expire?: integer;
    negativeCache?: integer;
    primaryNameserver?: string;
    refresh?: integer;
    retry?: integer;
    ttl?: integer;
  };

  /**
   * The numeric IDs of tags you want to apply to this domain
   */
  tags?: integer[];

  /**
   * The template to use for creating this domain. It will be linked to this template so any changes made to the template will apply to this domain.
   */
  template?: integer;

  /**
   * The vanity nameserver to use for this domain.
   */
  vanityNameserver?: integer;
}

/**
 * Analytics for a specific domain
 */
model domainanalytics {
  domain?: `simple-domain`;

  /**
   * The end date for the analytics
   */
  end?: utcDateTime;

  /**
   * Details about the interval between time periods in the analytics
   */
  interval?: {
    max?: integer;
    mean?: numeric;
    min?: integer;
  };

  links?: {
    domain?: string;
    self?: string;
  };

  /**
   * An array of analytics data for different types of query
   */
  queries?: {
    stats?: {
      count?: integer;
      max?: integer;
      mean?: numeric;
      min?: integer;
      sum?: integer;
    };
    type: "geo_proximity" | "standard" | "geo_filter";
    values?: {
      date?: utcDateTime;
      value?: integer;
    }[];
  }[];

  /**
   * The start date for the analytics
   */
  start?: utcDateTime;

  /**
   * Some statistics for these analytics
   */
  stats?: {
    count?: integer;
    max?: integer;
    mean?: numeric;
    min?: integer;
    sum?: integer;
  };
}

model domainhistory {
  domain?: `simple-domain`;

  /**
   * The name of the domain
   */
  name?: string;

  updatedAt?: utcDateTime;

  /**
   * The version of the domain resource
   */
  version?: integer;
}

/**
 * Represents a search result when searching for domains
 */
model domainsearchresult {
  /**
   * The ID for the domain name
   */
  id?: integer;

  /**
   * Links for domain search results
   */
  links?: {
    @format("url")
    domain?: string;
  };

  /**
   * The name of the domain
   */
  name?: string;
}

model domainsnapshot {
  domain?: `simple-domain`;

  /**
   * The name of the domain
   */
  name?: string;

  updatedAt?: utcDateTime;

  /**
   * The version of the domain resource
   */
  version?: integer;
}

model failoverrecordvalue {
  /**
   * Whether this entry is considered active or not
   */
  active?: boolean;

  /**
   * Whether the failover entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * Does sonar consider this entry to be failed
   */
  failed?: boolean;

  /**
   * The sort order of the entry. Lower order entries are preferred over higher order entries
   */
  order?: integer;

  /**
   * The ID in Sonar to use for checking if the record should be used
   */
  sonarCheckId?: integer;

  /**
   * The status of the value in Sonar
   */
  status?: string;
}

/**
 * Geo Proximity Location
 */
model geoproximity {
  /**
   * ID of the city
   */
  city?: integer;

  /**
   * 2 digit ISO country code
   */
  country?: string;

  /**
   * The unique ID for the Geo Proximity location
   */
  id?: integer;

  /**
   * Latitude of the location
   */
  latitude?: numeric;

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * Longitude of the location
   */
  longitude?: numeric;

  /**
   * The name of the Geo Proximity location
   */
  name?: string;

  /**
   * Region, state or province code
   */
  region?: string;
}

/**
 * Geo Proximity Location
 */
model `geoproximity-createupdate` {
  /**
   * The numeric ID for a city
   */
  city?: integer;

  /**
   * 2 digit ISO country code
   */
  country?: string;

  /**
   * Latitude of the location
   */
  latitude?: numeric;

  /**
   * Longitude of the location
   */
  longitude?: numeric;

  /**
   * The name of the Geo Proximity location
   */
  name?: string;

  /**
   * Region, state or province code
   */
  region?: string;
}

model hinfo {
  ...record;
  ...domainrecord;
  ...templaterecord;
  ...`value-hinfo`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "HINFO";
}

model http {
  ...record;
  ...domainrecord;
  ...templaterecord;
  ...`value-http`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: value;
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "HTTP";
}

model ipfilter {
  /**
   * AS Numbers that this rule applies to. Values must be between 0 and 4,294,967,295
   */
  asn?: integer[];

  /**
   * A list of continents in this filter
   */
  continents?: ("AF" | "AN" | "AS" | "EU" | "NA" | "OC" | "SA")[];

  /**
   * 2 digit ISO code for countries in this filter - https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
   */
  countries?: string[];

  /**
   * The unique ID for this IP filter
   */
  id?: integer;

  /**
   * IPv4 addresses and CIDRs that this filter applies to
   */
  ipv4?: string[];

  /**
   * IPv6 addresses and CIDRs that this filter applies to
   */
  ipv6?: string[];

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The name for this IP filter
   */
  name?: string;

  /**
   * Regions for this IP filter.
   */
  regions?: region[];

  /**
   * Values should be multiples of 100. Upper limit is the quota assigned to the account.
   */
  rulesLimit?: integer;
}

model `ipfilter-createupdate` {
  /**
   * AS Numbers that this rule applies to. Values must be between 0 and 4,294,967,295
   */
  asn?: integer[];

  /**
   * A list of continents in this filter
   */
  continents?: ("world" | "AF" | "AN" | "AS" | "EU" | "NA" | "OC" | "SA")[];

  /**
   * 2 digit ISO code for countries in this filter - https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
   */
  countries?: string[];

  /**
   * IPv4 addresses and CIDRs that this filter applies to
   */
  ipv4?: string[];

  /**
   * IPv6 addresses and CIDRs that this filter applies to
   */
  ipv6?: string[];

  /**
   * The name for this IP filter
   */
  name?: string;

  /**
   * Combinations of Regions, Countries, Continents and optionally ASNs.
   */
  regions?: {
    asn?: integer;
    continent?: "world" | "AF" | "AN" | "AS" | "EU" | "NA" | "OC" | "SA";
    country?: string;
    region?: string;
  }[];

  /**
   * Values should be multiples of 100. Upper limit is the quota assigned to the account.
   */
  rulesLimit?: integer;
}

model items {
  /**
   * The value for this tag
   */
  data?: string;

  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * A flag byte for any flags on the record.
   */
  flags?: integer;

  /**
   * Foo
   */
  tag?: "issue" | "issuewild" | "iodef";
}

model mx {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-mx`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-mx_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "MX";
}

model naptr {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-naptr`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-naptr_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "NAPTR";
}

model ns {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-ns`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-ns_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "NS";
}

model ping {
  ip?: string;
  timestamp?: utcDateTime;
  version?: string;
}

model pool {
  /**
   * Contactlists that will be notified when the pool is updated
   */
  contacts?: `simple-contactlist`[];

  createdAt?: utcDateTime;

  /**
   * Domains that are using this pool
   */
  domains?: `simple-domain`[];

  /**
   * Whether the pool is enabled or not
   */
  enabled?: boolean;

  /**
   * Whether the pool is considered failed or not
   */
  failed?: boolean;

  id?: integer;
  ito?: poolito;

  /**
   * Links for the pool
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The number of records that must be available for this pool to be used for failover. Between 0 and 64.
   */
  minimumFailover?: integer;

  name?: string;

  /**
   * The minimum number of entries from the pool to return when queried. Between 0 and 64.
   */
  `return`?: integer;

  /**
   * Templates that are using this pool
   */
  templates?: `simple-template`[];

  /**
   * The type of pool
   */
  type: "A" | "AAAA" | "CNAME";

  updatedAt?: utcDateTime;

  /**
   * The values for this pool
   */
  values?: {
    activated?: boolean;
    enabled?: boolean;
    failed?: boolean;
    handicap?: numeric;
    policy?: "follow_sonar" | "always_off" | "always_on" | "off_on_failure";
    sonarCheckId?: integer;
    speed?: numeric;

    @oneOf
    value?: string | string | string;

    weight?: integer;
  }[];
}

model `pool-createupdate` {
  /**
   * Contact lists to be emailed when this pool changes
   */
  contacts?: integer[];

  /**
   * Whether the pool is enabled or not
   */
  enabled?: boolean;

  ito?: poolito;

  /**
   * The number of records that must be available for this pool to be used for failover. Between 0 and 64.
   */
  minimumFailover?: integer;

  /**
   * A name for the pool
   */
  name?: string;

  /**
   * The minimum number of entries from the pool to return when queried. Between 0 and 64.
   */
  `return`?: integer;

  /**
   * The type of pool, either A, AAAA or CNAME
   */
  type: "A" | "AAAA" | "CNAME";

  /**
   * The values for this pool
   */
  values?: {
    enabled?: boolean;
    handicap?: integer;
    policy?: "follow_sonar" | "always_off" | "always_on" | "off_on_failure";
    sonarCheckId?: integer;

    @oneOf
    value?: string | string | string;

    weight?: integer;
  }[];
}

model poolindex {
  /**
   * Contactlists that will be notified when the pool is updated
   */
  contacts?: `simple-contactlist`[];

  createdAt?: utcDateTime;

  /**
   * Domains that are using this pool
   */
  domains?: `simple-domain`[];

  /**
   * Whether the pool is enabled or not
   */
  enabled?: boolean;

  /**
   * Whether the pool is considered failed or not
   */
  failed?: boolean;

  id?: integer;
  ito?: poolito;

  /**
   * Links for the pool
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The number of records that must be available for this pool to be used for failover. Between 0 and 64.
   */
  minimumFailover?: integer;

  name?: string;

  /**
   * The minimum number of entries from the pool to return when queried. Between 0 and 64.
   */
  `return`?: integer;

  /**
   * Templates that are using this pool
   */
  templates?: `simple-template`[];

  /**
   * The type of pool
   */
  type: "A" | "AAAA" | "CNAME";

  updatedAt?: utcDateTime;

  /**
   * The values for this pool
   */
  values?: {
    @oneOf
    value?: string | string | string;

    weight?: integer;
  }[];
}

model poolito {
  /**
   * The Ito configuration
   */
  config?: {
    deviationAllowance?: 10 | 20 | 30 | 40 | 50 | 60 | 70 | 80 | 90;
    handicapFactor?: "none" | "percent" | "speed";
    maximumNumberOfResults?: integer;
    monitoringRegion?:
      | "world"
      | "asiapac"
      | "europe"
      | "nacentral"
      | "naeast"
      | "nawest"
      | "oceania"
      | "southamerica";
    period?: 30 | 60 | 120 | 180 | 240 | 300;
  };

  /**
   * Is Ito enabled for this pool?
   */
  enabled?: boolean;
}

model `properties-value-items` {
  /**
   * Text representing the CPU
   */
  cpu?: string;

  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * Text representing the OS
   */
  os?: string;
}

model ptr {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-ptr`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-ptr_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "PTR";
}

model domainrecord {
  domain?: `simple-domain`;
}

/**
 * An object representing a DNS record
 */
model record {
  contacts?: `simple-contactlist`;

  /**
   * Whether the record is enabled or not. A disabled record will return an NXDOMAIN response.
   */
  enabled?: boolean;

  /**
   * Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
   */
  geoFailover?: boolean;

  geoproximity?: `simple-geoproximity`;

  /**
   * A unique ID for this domain record
   */
  id?: integer;

  ipfilter?: `simple-ipfilter`;

  /**
   * If the requesting IP matches the IP filter, don't return a response
   */
  ipfilterDrop?: boolean;

  /**
   * The name of the record
   */
  name?: string;

  /**
   * A note about the record. Max 512 characters.
   */
  notes?: string;

  /**
   * The region for this record
   */
  region?:
    | "default"
    | "europe"
    | "us-east"
    | "us-west"
    | "asia-pacific"
    | "oceania"
    | "south-america";

  /**
   * Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
   */
  skipLookup?: boolean | null;

  /**
   * The time to live in seconds for this record. must be between 0 and 2147483647
   */
  ttl?: integer;
}

model `record-create-details` {
  /**
   * Contact lists to be notified if a failover happens in a failover mode.
   */
  contacts?: integer[];

  /**
   * Whether the record is enabled
   */
  enabled?: boolean;

  /**
   * Disable the record if all hosts fail. If all hosts fail, another matching IP Filter, nearest Proximity or World (Default) record will be used instead.
   */
  geoFailover?: boolean;

  /**
   * The integer ID of a GeoProximity to use for this record. Cannot be used with IP Filter.
   */
  geoproximity?: integer;

  /**
   * The integer ID of an IP Filter to use for this record. Cannot be used with GeoPeoximity.
   */
  ipfilter?: integer;

  /**
   * If the requesting IP matches the IP filter, don't return a response
   */
  ipfilterDrop?: boolean;

  /**
   * The name for the record
   */
  name?: string;

  /**
   * A description of the record. It must be 512 characters or less.
   */
  notes?: string;

  /**
   * Optional region for this record. Will default to 'default'.
   */
  region?:
    | "default"
    | "europe"
    | "us-east"
    | "us-west"
    | "asia-pacific"
    | "oceania"
    | "south-america";

  /**
   * Only used on POST or PATCH requests for ANAME records, used to specify whether the hostname should be looked up immediately. Will be null otherwise.
   */
  skipLookup?: boolean | null;

  /**
   * How long DNS servers should cache the record for
   */
  ttl?: integer;
}

model `record-create-update-a` {
  ...`record-create-details`;
  ...`value-a-request`;

  /**
   * The current mode for this record
   */
  mode?: "standard" | "failover" | "roundrobin-failover" | "pools";

  /**
   * The type of record
   */
  type: "A";
}

model `record-create-update-aaaa` {
  ...`record-create-details`;
  ...`value-aaaa-request`;

  /**
   * The current mode for this record
   */
  mode?: "standard" | "failover" | "roundrobin-failover" | "pools";

  /**
   * The type of record
   */
  type: "AAAA";
}

model `record-create-update-cname` {
  ...`record-create-details`;
  ...`value-cname-request`;

  /**
   * The current mode for this record
   */
  mode?: "standard" | "failover" | "pools";

  /**
   * The type of record
   */
  type: "CNAME";
}

model `record-create-update-aname` {
  ...`record-create-details`;
  ...`value-aname-request`;

  /**
   * The current mode for this record
   */
  mode?: "standard" | "failover" | "pools";

  /**
   * The type of record
   */
  type: "ANAME";
}

model `record-create-update-caa` {
  ...`record-create-details`;
  ...`value-caa`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "CAA";
}

model `record-create-update-cert` {
  ...`record-create-details`;
  ...`value-cert`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "CERT";
}

model `record-create-update-hinfo` {
  ...`record-create-details`;
  ...`value-hinfo`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "HINFO";
}

model `record-create-update-mx` {
  ...`record-create-details`;
  ...`value-mx`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "MX";
}

model `record-create-update-http` {
  ...`record-create-details`;
  ...`value-http`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "HTTP";
}

model `record-create-update-naptr` {
  ...`record-create-details`;
  ...`value-naptr`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "NAPTR";
}

model `record-create-update-ns` {
  ...`record-create-details`;
  ...`value-ns`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "NS";
}

model `record-create-update-ptr` {
  ...`record-create-details`;
  ...`value-ptr`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "PTR";
}

model `record-create-update-rp` {
  ...`record-create-details`;
  ...`value-rp`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "RP";
}

model `record-create-update-spf` {
  ...`record-create-details`;
  ...`value-spf`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "SPF";
}

model `record-create-update-srv` {
  ...`record-create-details`;
  ...`value-srv`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "SRV";
}

model `record-create-update-txt` {
  ...`record-create-details`;
  ...`value-txt`;

  /**
   * The current mode for this record
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "TXT";
}

@discriminator("type")
@oneOf
union `record-create` {
  `record-create-update-a`,
  `record-create-update-aaaa`,
  `record-create-update-cname`,
  `record-create-update-aname`,
  `record-create-update-caa`,
  `record-create-update-cert`,
  `record-create-update-hinfo`,
  `record-create-update-mx`,
  `record-create-update-http`,
  `record-create-update-naptr`,
  `record-create-update-ptr`,
  `record-create-update-rp`,
  `record-create-update-spf`,
  `record-create-update-srv`,
  `record-create-update-txt`,
}

@discriminator("type")
@oneOf
union `record-update` {
  `record-create-update-a`,
  `record-create-update-aaaa`,
  `record-create-update-cname`,
  `record-create-update-aname`,
  `record-create-update-caa`,
  `record-create-update-cert`,
  `record-create-update-hinfo`,
  `record-create-update-mx`,
  `record-create-update-http`,
  `record-create-update-naptr`,
  `record-create-update-ptr`,
  `record-create-update-rp`,
  `record-create-update-spf`,
  `record-create-update-srv`,
  `record-create-update-txt`,
}

/**
 * IP Filter Region
 */
model region {
  /**
   * AS Number for this region
   */
  asn?: integer;

  /**
   * The continent for this region
   */
  continent?: string;

  /**
   * The country for the region
   */
  country?: string;

  /**
   * The region code
   */
  region?: string;
}

model rp {
  ...record;
  ...domainrecord;
  ...templaterecord;
  ...`value-rp`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-rp_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "RP";
}

/**
 * A simple version of a contact list when inclued with other resources
 */
model `simple-contactlist` {
  /**
   * Unique ID for the contact list
   */
  id?: integer;

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    self?: string;
  };
}

model `simple-domain` {
  createdAt?: utcDateTime;

  /**
   * Is the domain enabled
   */
  enabled?: boolean;

  /**
   * Is GeoIP functionality enabled for the domain
   */
  geoip?: boolean;

  /**
   * Is Global Traffic Director enabled for the domain
   */
  gtd?: boolean;

  /**
   * A unique numeric ID for this domain
   */
  id?: integer;

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    records?: string;

    @format("url")
    self?: string;
  };

  /**
   * The name of the domain
   */
  name?: string;

  /**
   * A note for the domain
   */
  note?: string;

  status?: "ACTIVE" | "SUSPENDED" | "TERMINATED";

  /**
   * An array of tags for this domain.
   */
  tags?: tag[];

  updatedAt?: utcDateTime;

  /**
   * The version of the domain resource
   */
  version?: integer;
}

/**
 * Geo Proximity Location
 */
model `simple-geoproximity` {
  /**
   * The unique ID for the Geo Proximity location
   */
  id?: integer;

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The name of the Geo Proximity location
   */
  name?: string;
}

model `simple-ipfilter` {
  /**
   * The unique ID for this IP filter
   */
  id?: integer;

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * The name for this IP filter
   */
  name?: string;
}

model `simple-pool` {
  id?: integer;

  /**
   * Links for the pool
   */
  links?: {
    @format("url")
    self?: string;
  };
}

model `simple-template` {
  /**
   * The unique ID for this template
   */
  id?: integer;

  /**
   * Links for the template
   */
  links?: {
    @format("url")
    records?: string;

    @format("url")
    self?: string;
  };

  /**
   * The name for the template
   */
  name?: string;

  /**
   * The version of the template resource
   */
  version?: integer;
}

/**
 * A simplied version of a rebranded nameserver using your own domain name.
 */
model `simple-vanitynameserver` {
  /**
   * A unique ID for the vanity nameserver
   */
  id?: integer;

  /**
   * Links relevant to this object
   */
  links?: {
    @format("url")
    self?: string;
  };
}

/**
 * SOA (Start of Authority) Record
 */
model soa {
  @doc("Email of the administrator for the domain. @ should be replaced with .")
  email?: string;

  /**
   * Number of seconds after which secondary nameservers should stop responding to queries, if the master does not respond
   */
  expire?: integer;

  /**
   * How long NXDOMAIN responses should be cached for
   */
  negativeCache?: integer;

  /**
   * Primary master nameserver for the domain
   */
  primaryNameserver?: string;

  /**
   * The interval for secondary nameservers should query for the SOA record
   */
  refresh?: integer;

  /**
   * The number of seconds after which secondary servers should retry to request the serial number if the master does not respond
   */
  retry?: integer;

  /**
   * A serial number for this version of the records for this domain
   */
  serial?: integer;

  /**
   * The Time To Live (TTL) in seconds for the SOA record
   */
  ttl?: integer;
}

model spf {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-spf`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-spf_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "SPF";
}

model srv {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-srv`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-srv_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "SRV";
}

/**
 * A tag is used to group resources together
 */
model tag {
  /**
   * Unique ID for the tag
   */
  id?: integer;

  /**
   * Links for tags
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * A name for this tag
   */
  name?: string;
}

model `tag-createupdate` {
  /**
   * A name for this tag
   */
  name?: string;
}

/**
 * A domain template
 */
model template {
  createdAt?: utcDateTime;

  /**
   * Is GeoIP functionality enabled for the template
   */
  geoip?: boolean;

  /**
   * Is Global Traffic Director enabled for the template
   */
  gtd?: boolean;

  /**
   * The unique ID for this template
   */
  id?: integer;

  /**
   * Links for the template
   */
  links?: {
    @format("url")
    records?: string;

    @format("url")
    self?: string;
  };

  /**
   * The name for the template
   */
  name?: string;

  updatedAt?: utcDateTime;

  /**
   * The version of the template resource
   */
  version?: integer;
}

model `template-createupdate` {
  /**
   * Is GeoIP functionality enabled for the template
   */
  geoip?: boolean;

  /**
   * Is Global Traffic Director enabled for the template
   */
  gtd?: boolean;

  /**
   * The name of the template
   */
  name?: string;
}

model templaterecord {
  template?: `simple-template`;
}

model txt {
  ...record;
  ...link;
  ...domainrecord;
  ...templaterecord;
  ...`value-txt`;

  /**
   * The previous values of the record in the different modes
   */
  lastValues?: {
    standard?: `value-txt_properties-value-items`[];
  };

  /**
   * How the record should work
   */
  mode?: "standard";

  /**
   * The type of record
   */
  type: "TXT";
}

model value {
  /**
   * The description to include in meta tags
   */
  description?: string;

  /**
   * If this is a frame redirect, don't include the URL entered in the browser
   */
  hard?: boolean;

  /**
   * Meta keywords for the page containing the frame
   */
  keywords?: string;

  /**
   * The type of redirect, either 301, 302 or frame.
   */
  redirecttype: string;

  /**
   * The title of the page containing the redirect frame
   */
  title?: string;

  /**
   * The URL to redirect to
   */
  @format("url") url?: string;
}

/**
 * Values for an A record
 */
model `value-a` {
  /**
   * The type of record
   */
  type: "A";

  @oneOf value?:
    | {
        enabled?: boolean;

        @format("ipv4")
        value?: string;
      }[]
    | {
        mode?: "normal" | "off" | "one-way";
        values?: unknown[];
      }
    | unknown[]
    | `simple-pool`[];
}

/**
 * Values for an A record
 */
model `value-a-request` {
  @oneOf value?:
    | {
        enabled?: boolean;

        @format("ipv4")
        value?: string;
      }[]
    | {
        mode?: "normal" | "off" | "one-way";
        values?: {
          enabled?: boolean;
          order?: integer;
          sonarCheckId?: integer;

          @format("ipv4")
          value?: string;
        }[];
      }
    | {
        enabled?: boolean;
        order?: integer;
        sonarCheckId?: integer;

        @format("ipv4")
        value?: string;
      }[]
    | {
        id?: integer;
      }[];
}

/**
 * Values for an AAAA record
 */
model `value-aaaa` {
  /**
   * The type of record
   */
  type: "AAAA";

  @oneOf value?:
    | {
        enabled?: boolean;

        @format("ipv6")
        value?: string;
      }[]
    | {
        enabled?: boolean;
        mode?: "normal" | "off" | "one-way";
        values?: unknown[];
      }
    | unknown[]
    | `simple-pool`[];
}

/**
 * Values for an AAAA record
 */
model `value-aaaa-request` {
  @oneOf value?:
    | {
        enabled?: boolean;

        @format("ipv6")
        value?: string;
      }[]
    | {
        enabled?: boolean;
        mode?: "normal" | "off" | "one-way";
        values?: {
          active?: boolean;
          enabled?: boolean;
          order?: integer;
          sonarCheckId?: integer;

          @format("ipv6")
          value?: string;
        }[];
      }
    | {
        active?: boolean;
        enabled?: boolean;
        order?: integer;
        sonarCheckId?: integer;

        @format("ipv6")
        value?: string;
      }[]
    | {
        id?: integer;
      }[];
}

/**
 * ANAME Record
 */
model `value-aname` {
  /**
   * The type of record
   */
  type: "ANAME";

  @oneOf value?: {
    enabled?: boolean;
    value?: string;
  }[] | {
    mode?: "normal" | "off" | "one-way";
    values?: unknown[];
  } | `simple-pool`[];
}

/**
 * ANAME Record
 */
model `value-aname-request` {
  @oneOf value?: {
    enabled?: boolean;

    @format("ipv6")
    value?: string;
  }[] | {
    enabled?: boolean;
    mode?: "normal" | "off" | "one-way";
    values?: {
      active?: boolean;
      enabled?: boolean;
      order?: integer;
      sonarCheckId?: integer;

      @format("ipv6")
      value?: string;
    }[];
  } | {
    id?: integer;
  }[];
}

/**
 * CAA Record
 */
model `value-caa` {
  /**
   * Standard record mode
   */
  value?: {
    data?: string;
    enabled?: boolean;
    flags?: integer;
    tag?: "issue" | "issuewild" | "iodef";
  }[];
}

/**
 * CERT Record
 */
model `value-cert` {
  /**
   * Standard record mode
   */
  value?: {
    algorithm?: integer;

    @format("base64 encoded")
    certificate?: string;

    certificatetype: integer;
    enabled?: boolean;
    keyTag?: integer;
  }[];
}

/**
 * CNAME Record
 */
model `value-cname` {
  /**
   * The type of record
   */
  type: "CNAME";

  @oneOf value?: {
    enabled?: boolean;
    value?: string;
  }[] | {
    mode?: "normal" | "off" | "one-way";
    values?: unknown[];
  } | `simple-pool`[];
}

/**
 * CNAME Record
 */
model `value-cname-request` {
  @oneOf value?: {
    enabled?: boolean;

    @format("ipv6")
    value?: string;
  }[] | {
    enabled?: boolean;
    mode?: "normal" | "off" | "one-way";
    values?: {
      active?: boolean;
      enabled?: boolean;
      order?: integer;
      sonarCheckId?: integer;

      @format("ipv6")
      value?: string;
    }[];
  } | {
    id?: integer;
  }[];
}

/**
 * HINFO Record
 */
model `value-hinfo` {
  /**
   * Standard record mode
   */
  value?: {
    cpu?: string;
    enabled?: boolean;
    os?: string;
  }[];
}

/**
 * HTTP Record
 */
model `value-http` {
  value?: {
    description?: string;
    hard?: boolean;
    keywords?: string;
    redirecttype: string;
    title?: string;

    @format("url")
    url?: string;
  };
}

model `value-items` {
  /**
   * An integer representing the algorithm
   */
  algorithm?: integer;

  /**
   * A base 64 encoded string containing the certificate information
   */
  @format("base64 encoded") certificate?: string;

  /**
   * An integer representing the type of certificate
   */
  certificatetype: integer;

  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * An integer representing the key tag
   */
  keyTag?: integer;
}

/**
 * MX Record
 */
model `value-mx` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    priority?: integer;
    server?: string;
  }[];
}

model `value-mx_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * A number representing the priority of the mail server
   */
  priority?: integer;

  /**
   * The hostname of the mail server
   */
  server?: string;
}

/**
 * NAPTR Record
 */
model `value-naptr` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    flags?: string;
    order?: integer;
    preference?: integer;
    regularExpression?: string;
    replacement?: string;
    service?: string;
  }[];
}

model `value-naptr_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * Any flags for this record
   */
  flags?: string;

  /**
   * The order of the record
   */
  order?: integer;

  /**
   * The preference for the record
   */
  preference?: integer;

  /**
   * A regular expression to use
   */
  regularExpression?: string;

  /**
   * The replacement for the regular expression
   */
  replacement?: string;

  /**
   * The service the record is used for
   */
  service?: string;
}

/**
 * NS Record
 */
model `value-ns` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    host?: string;
  }[];
}

model `value-ns_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * The hostname for the nameserver
   */
  host?: string;
}

/**
 * PTR Record
 */
model `value-ptr` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    system?: string;
  }[];
}

model `value-ptr_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * The hostname for the IP address
   */
  system?: string;
}

/**
 * RP Record
 */
model `value-rp` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    mailbox?: string;
    txt?: string;
  }[];
}

model `value-rp_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  @doc("The email of the person responsible. Replace @ with .")
  mailbox?: string;

  /**
   * The name of a TXT record containing more information
   */
  txt?: string;
}

/**
 * SPF Record
 */
model `value-spf` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    value?: string;
  }[];
}

model `value-spf_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * The SPF record value
   */
  value?: string;
}

/**
 * SRV Record
 */
model `value-srv` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    host?: string;
    port?: integer;
    priority?: integer;
    weight?: integer;
  }[];
}

model `value-srv_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * The hostname for the service
   */
  host?: string;

  /**
   * The port the service runs on
   */
  port?: integer;

  /**
   * A priority for this record
   */
  priority?: integer;

  /**
   * A weight for this record
   */
  weight?: integer;
}

/**
 * TXT Record
 */
model `value-txt` {
  /**
   * Standard record mode
   */
  value?: {
    enabled?: boolean;
    value?: string;
  }[];
}

model `value-txt_properties-value-items` {
  /**
   * Whether the entry is enabled or not. Disabled entries will not be included in a response
   */
  enabled?: boolean;

  /**
   * The text record value
   */
  value?: string;
}

/**
 * A rebranded nameserver using your own domain name.
 */
model vanitynameserver {
  /**
   * Is this the default nameserver for domains in the account
   */
  default?: boolean;

  /**
   * A unique ID for the vanity nameserver
   */
  id?: integer;

  /**
   * Links for domain objects
   */
  links?: {
    @format("url")
    self?: string;
  };

  /**
   * A unique name for this vanity nameserver
   */
  name?: string;

  nameserverGroup?: {
    id?: integer;
    name?: string;
  };

  /**
   * The nameserver hostnames
   */
  nameservers?: string[];

  /**
   * Is this nameserver public?
   */
  public?: boolean;
}

/**
 * A rebranded nameserver using your own domain name.
 */
model `vanitynameserver-create` {
  /**
   * Is this the default nameserver for domains in the account
   */
  default?: boolean;

  /**
   * A unique name for this vanity nameserver
   */
  name?: string;

  nameserverGroup?: {
    id?: integer;
  };

  /**
   * The nameserver hostnames
   */
  nameservers?: string[];
}

/**
 * A rebranded nameserver using your own domain name.
 */
model `vanitynameserver-update` {
  /**
   * Is this the default nameserver for domains in the account
   */
  default?: boolean;

  /**
   * A unique name for this vanity nameserver
   */
  name?: string;

  /**
   * The nameserver hostnames
   */
  nameservers?: string[];
}

/**
 * Retrieve query analytics for your account
 */
model listAnalytics200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * List of announcements
 */
model listAnnouncements200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * Fetch a single announcement
 */
model getAnnouncement200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The list of Contact Lists
 */
model listContactLists200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The new Contact List resource
 */
model createContactList202ApplicationJsonResponse {
  /**
   * The URL for the newly created object
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The Contact List was deleted successfully
 */
model deleteContactList204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a Contact List
 */
model getContactList200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The updated Contact List resource
 */
model updateContactList200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * List of emails in the contactlist
 */
model listContactListEmails200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The email address has been added
 */
model addContactListEmail202ApplicationJsonResponse {
  /**
   * The URL for the newly created email
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The email address was deleted successfully
 */
model deleteContactListEmail204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch an Email Address
 */
model getContactListEmail200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * List of Slack webhooks in the contactlist
 */
model listContactListSlackWebhooks200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The Slack webhook has been added
 */
model createContactListSlackWebhook202ApplicationJsonResponse {
  /**
   * The URL for the newly created Slack webhook
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The webhook was deleted successfully
 */
model deleteContactListSlackWebhook204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a Slack Webhook
 */
model getContactListSlackWebhook200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * List of MS Teams webhooks in the contactlist
 */
model listContactListTeams200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The Microsoft Teams webhook has been added
 */
model createContactListTeam202ApplicationJsonResponse {
  /**
   * The URL for the newly created MS Teams webhook
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The webhook was deleted successfully
 */
model deleteContactListTeam204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch an MS Teams Webhook
 */
model getContactListTeam200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * List of domains in your account
 */
model listDomains200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The domain has been accepted for creation
 */
model createDomain202ApplicationJsonResponse {
  /**
   * The URL for the newly created domain
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The domain was deleted successfully
 */
model deleteDomain204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a single domain from your account
 */
model getDomain200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The updated domain resource
 */
model updateDomain200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * Retrieve query analytics for the domain
 */
model getDomainAnalytics200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * Export of the domain in BIND format
 */
model getDomainBindExport200TextPlainResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: string;
  @header contentType: "text/plain";
}

/**
 * Fetch the list of history for the domain
 */
model listDomainHistory200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * Fetch a history version for the domain
 */
model getDomainHistoryVersion200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The historical version has been applied
 */
model applyDomainHistory204Response {
  @statusCode statusCode: 204;
}

/**
 * The snapshot has accepted for creation.
 */
model snapshotDomainHistory202ApplicationJsonResponse {
  /**
   * The URL for the newly created snapshot
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * Fetch the nameservers for the domain
 */
model getDomainNameservers200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * A list of records for the domain
 */
model listDomainRecords200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The record has been accepted for creation
 */
model createDomainRecord202ApplicationJsonResponse {
  /**
   * The URL for the newly created record
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The record was deleted successfully
 */
model deleteDomainRecord204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a record for the domain
 */
model getDomainRecord200ApplicationJsonResponse {
  @statusCode statusCode: 200;

  @bodyRoot
  @discriminator("type")
  @oneOf
  body:
    | a
    | aaaa
    | cname
    | aname
    | caa
    | cert
    | hinfo
    | mx
    | naptr
    | ns
    | rp
    | ptr
    | txt
    | spf
    | srv
    | http;
}

/**
 * The updated record resource
 */
model updateDomainRecord200ApplicationJsonResponse {
  @statusCode statusCode: 200;

  @bodyRoot
  @discriminator("type")
  @oneOf
  body:
    | a
    | aaaa
    | cname
    | aname
    | caa
    | cert
    | hinfo
    | mx
    | naptr
    | ns
    | rp
    | ptr
    | txt
    | spf
    | srv
    | http;
}

/**
 * Fetch the list of snapshots for the domain
 */
model listDomainSnapshots200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The snapshot was deleted successfully
 */
model deleteDomainSnapshot204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a snapshot for the domain
 */
model getDomainSnapshot200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The snapshot has been applied
 */
model applyDomainSnapshot204Response {
  @statusCode statusCode: 204;
}

/**
 * The list of Geo Proximity locations
 */
model listGeoProximityLocations200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The new Geo Proximity location resource
 */
model createGeoProximityLocation202ApplicationJsonResponse {
  /**
   * The URL for the newly created object
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The Geo Proximity location was deleted successfully
 */
model deleteGeoProximityLocation204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a Geo Proximity location
 */
model getGeoProximityLocation200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The updated Geo Proximity location resource
 */
model updateGeoProximityLocation200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The list of IP filters
 */
model listIPFilters200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The new IPFilter resource
 */
model createIPFilter202ApplicationJsonResponse {
  /**
   * The URL for the newly created object
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The IP Filter was deleted successfully
 */
model deleteIPFilter204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch an IP Filter
 */
model getIPFilter200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The updated IP filter resource
 */
model updateIPFilter200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * API Ping response
 */
model getPing200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: ping;
}

/**
 * List of pools
 */
model listPools200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The new Pool resource
 */
model createPool202ApplicationJsonResponse {
  /**
   * The URL for the newly created object
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The pool was deleted successfully
 */
model deletePool204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a pool
 */
model getPool200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The updated pool resource
 */
model updatePool200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * Search for domains within your account
 */
model searchDomains200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The list of Tags
 */
model listTags200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The new Tag resource
 */
model createTag202ApplicationJsonResponse {
  /**
   * The URL for the newly created object
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The Tag was deleted successfully
 */
model deleteTag204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a Tag
 */
model getTag200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The updated Tag resource
 */
model updateTag200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * List of templates
 */
model listTemplates200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The template has been created successfully
 */
model createTemplate202ApplicationJsonResponse {
  /**
   * The URL for the newly created template
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The template was deleted successfully
 */
model deleteTemplate204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a single template from your account
 */
model getTemplate200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The template has been updated successfully
 */
model updateTemplate200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * A list of records for the template
 */
model listTemplateRecords200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The record has been accepted for creation
 */
model createTemplateRecord202ApplicationJsonResponse {
  /**
   * The URL for the newly created record
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The record was deleted successfully
 */
model deleteTemplateRecord204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a record for the template
 */
model getTemplateRecord200ApplicationJsonResponse {
  @statusCode statusCode: 200;

  @bodyRoot
  @discriminator("type")
  @oneOf
  body:
    | a
    | aaaa
    | cname
    | aname
    | caa
    | cert
    | hinfo
    | mx
    | naptr
    | ns
    | rp
    | ptr
    | txt
    | spf
    | srv
    | http;
}

/**
 * The updated record resource
 */
model updateTemplateRecord200ApplicationJsonResponse {
  @statusCode statusCode: 200;

  @bodyRoot
  @discriminator("type")
  @oneOf
  body:
    | a
    | aaaa
    | cname
    | aname
    | caa
    | cert
    | hinfo
    | mx
    | naptr
    | ns
    | rp
    | ptr
    | txt
    | spf
    | srv
    | http;
}

/**
 * The list of Vanity Nameservers
 */
model listVanityNameservers200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The new Vanity Nameserver resource
 */
model createVanityNameserver202ApplicationJsonResponse {
  /**
   * The URL for the newly created object
   */
  @header @format("url") Location?: string;

  @statusCode statusCode: 202;
  @bodyRoot body: unknown;
}

/**
 * The Vanity Nameserver was deleted successfully
 */
model deleteVanityNameserver204Response {
  @statusCode statusCode: 204;
}

/**
 * Fetch a Vanity Nameserver
 */
model getVanityNameserver200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}

/**
 * The updated Vanity Nameserver resource
 */
model updateVanityNameserver200ApplicationJsonResponse {
  @statusCode statusCode: 200;
  @bodyRoot body: unknown;
}
