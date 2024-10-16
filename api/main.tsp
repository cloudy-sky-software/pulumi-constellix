import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./models/models.tsp";

using Http;
using OpenAPI;

/**
 * REST API for interacting with Constellix DNS services.
 *
 * # Overview
 *
 * Welcome to the Constellix REST API. This API allows you to manage your Constellix DNS services. You can add, remove and update domains, records, templates and more.
 *
 * This documentation details all the API endpoints, resource types and contains all the information to help you to use the API.
 *
 * # Using the API
 *
 * ## Making Requests
 *
 * The API is a REST API and supports common HTTP verbs of GET, POST, PUT and DELETE. In general:
 *
 * - **GET** - Retrieve a resource (or resources) from the API
 * - **POST** - Create a new resource, or update an existing resource
 * - **PUT** - Update an existing resource
 * - **DELETE** - Delete a resource
 *
 *
 * Please note that partial data is allowed for PUT requests, only the fields provided in the request will be updated. You may also perform this request using `POST` or `PATCH`.
 *
 *
 * Most request endpoints are authenticated and request the `Authorization: Bearer` header to be provided. More details are in the [authentication section](#section/Authentication).
 *
 * The API uses standard HTTP response codes, 2xx means the request was successful, 3xx is a redirection, 4xx is a client error and 5xx is a server error.
 *
 * ## Formats
 *
 * The API supports `application/json` content-type and accept headers. All responses will be provided as JSON and all requests should specify `application/json` as the `Content-Type` and `Accept` headers.
 *
 * ## Rate Limiting
 *
 * Authenticated API requests are rate limited. The rate limiting details are provided in the response headers.
 *
 * | **Header** | **Example** | **Purpose** |
 * | ---------- | ----------- | ----------- |
 * | X-RateLimit-Limit | 1000 | Your maximum number of API requests |
 * | X-RateLimit-Remaining | 999 | How many requests you have remaining |
 * | X-RateLimit-Reset | 29 | How long until your rate limits are reset |
 *
 * ## Listing resources
 *
 * When you request a list of resources, you will get an object back that contains an array of the resources and information about the total number of resources and pagination.
 *
 * If you want to change the number of results fetched, or specify the page of results, you can use the following query parameters in the URL:
 *
 * | **Parameter** | **Example** | **Purpose** |
 * | -------------- | ----------- | ----------- |
 * | page | 2 | Which page of the results to fetch. By default this is the first page. |
 * | perPage | 50 | How many results to fetch per page. The default value is normally 50. |
 *
 * The results will have a `meta` property, it will contain the following:
 *
 * ```javascript
 * {
 *     "links": {
 *         "self": "https://api.dns.constellix.com/v4/templates?page=1",
 *         "first": "https://api.dns.constellix.com/v4/templates?page=1",
 *         "last": "https://api.dns.constellix.com/v4/templates?page=1",
 *         "previous": null,
 *         "next": null
 *     },
 *     "pagination": {
 *         "total": 3,
 *         "count": 3,
 *         "perPage": 50,
 *         "currentPage": 1
 *     }
 * }
 * ```
 *
 * There are links to other pages in the results as well as total pagination information.
 *
 * ## Creating resources
 *
 * When a new resources is created, it may not be created immediately. In most cases, it will be available immediately, but this is not guaranteed. For this reason, create requests return a `202 Accepted` response code.
 *
 * The response will have a `Location` header set with the location the new resource will exist at. The response body will include this location and also the ID of the resource created.
 *
 * The location may return a `404 Not Found` error until the resource has been created.
 *
 * # Authentication
 *
 * Authentication is performed using `Authorization: Bearer <token>`. The token is a combination of the API Key and request date, encoded using the Secret Key.
 *
 * The individual components of the token are:
 *
 * - **API Key** - The API key from the Constellix Management Portal, eg. `7a3fe9ef-de08-4c87-83a3-5034a8441f07`
 * - **HMAC** - The request timestamp (in milliseconds), HMAC encoded using the Secret Key from the Management Portal. This is then base64 encoded. eg. `UkHJfB5wJpAux8dq518Jfb434ks=`
 * - **Request Timestamp** - The request timestamp (in milliseconds), eg. `1489401874402`
 *
 * The three components are concatenated together, separated by colons. eg. `<API Key>:<HMAC>:<Request Date>`.
 *
 * Example PHP code for achieving this is:
 *
 * ```php
 * define('SECRET_KEY', 'b811608f-7501-4e0c-8a05-4e8d5ab841ec');
 * define ('API_KEY', '15e8b5a1-ef8f-4a16-a7dd-c30898d8affc');
 *
 * $now = new DateTime('now', new \DateTimeZone('UTC'));
 * $timestamp = (string) ($now->getTimestamp() * 1000);
 * $hmac = base64_encode(hash_hmac('sha1', $timestamp, SECRET_KEY, true));
 * $header = 'Authorization: Bearer ' . API_KEY . ':' . $hmac . ':' . $timestamp;
 * ```
 *
 * A pre-request script for Postman which can be used to populate an environment variable is:
 *
 * ```javascript
 * var apiKey = pm.collectionVariables.get("apiKey");
 * var secretKey = pm.collectionVariables.get("secretKey");
 *
 * function epochTime() {
 *   return new Date().getTime() + '';
 * }
 *
 * pm.collectionVariables.unset("constellixAuth");
 * var time = epochTime();
 * var hmac = CryptoJS.HmacSHA1(time, secretKey).toString(CryptoJS.enc.Base64);
 * var token = apiKey + ":" + hmac + ":" + time;
 * pm.collectionVariables.set('constellixAuth', token);
 * ```
 *
 * ## Backwards Compatibility
 *
 * The v4 API also supports authentication using the same headers as the previous Constellix DNS APIs. You can use either separate headers for the API Key, HMAC and Request Timestamp or they can be combined into a single header.
 *
 * It is still recommended that you pass the authentication through as an `Authorization: Bearer <token>` format instead.
 *
 * ### Separate headers
 *
 * ```
 * x-cnsdns-apiKey: 7a3fe9ef-de08-4c87-83a3-5034a8441f07
 * x-cnsdns-hmac: UkHJfB5wJpAux8dq518Jfb434ks=
 * x-cnsdns-requestDate: 1489401874402
 * ```
 *
 * ### Single Header
 *
 * The values should be concatenated with `:`.
 *
 * ```
 * x-cns-security-token: 7a3fe9ef-de08-4c87-83a3-5034a8441f07:UkHJfB5wJpAux8dq518Jfb434ks=:1489401874402
 * ```
 *
 * # Postman Collection
 *
 * If you use Postman, you can [download a Postman Collection](https://api.dns.constellix.com/v4/postman-collection) for this API, with the pre-request script configured. All you need to do is edit variables for the collection and add your secret key and API key.
 *
 * # Changelog
 *
 * ## v4.2.7
 *
 * - Updated documentation to remove incorrect `active` property on failover record value updates.
 *
 * ## v4.2.6
 *
 * - Internal changes.
 *
 * ## v4.2.5
 *
 * - Internal changes.
 *
 * ## v4.2.4
 *
 * - Fix inconsistent API results when adding a domain with tags that doesn't existent. Previously the domain would be created even though an error was returned in the response. Now the domain will not be created.
 *
 * ## v4.2.3
 *
 * - Fix 500 error when attempting to delete a pool that is in use by a template or domain. You will now get a 409 conflict error with an explanation.
 *
 * ## v4.2.2
 *
 * - Fix error that occurs when you update pool values without specifying a `policy` value
 *
 * ## v4.2.1
 *
 * - Fix issue accessing pools with ITO enabled but no ITO configuration
 *
 * ## v4.2.0
 *
 * - New sub-resources on Contact List resources for Slack channels, Microsoft Teams channels and Emails
 * - Fixed documentation on domain snapshots and domain history resources
 *
 * ## v4.1.13
 *
 * - Fix viewing and editing of contact lists
 * - Fix example IP Filter request in documentation
 * - Update documentation for Pools to include `createdAt` and `updatedAt` fields
 * - Fix pool value `speed` property sometimes being returned as `N/A` instead of `null`
 * - Improved error messages for Template Records
 * - Fix error when attempting to update a Template Record but making no changes
 *
 * ## v4.1.12
 *
 * - Correct documentation for geoproximity, domain and domain history resources to match API output
 * - Added version property on API domain resources to match documentation
 *
 * ## v4.1.11
 *
 * - Fix issue where domain notes were not updated
 * - Fix issue when fetching some records
 *
 * ## v4.1.10
 *
 * - Added documentation for domain BIND export
 * - Fixed various errors in documentation
 * - Updated openapi specification
 * - Fix errors when updating or creating records with empty names
 *
 * ## v4.1.9
 *
 * - Fix being unable to update the name on a domain record
 * - Fix for updating and deleting domain and template records
 * - Remove unused field from Vanity Nameserver create and update
 * - Add functionality to enable and disable domains
 *
 * ## v4.1.8
 *
 * - Fix 500 error when updating template records
 * - Fix issue where domain and template ID on record responses was returned as a string instead of an integer
 * - Fix 500 error when supplying invalid security token
 * - Fix 500 error when updating or creating a domain with null contacts or tags
 * - Add documentation for `geoFailover` property on records.
 * - Performance improvements when creating domain records
 *
 * ## v4.1.7
 *
 * - Performance improvements when creating resources
 * - Fix errors when fetching analytics for the current day
 *
 * ## v4.1.6
 *
 * - Updated validation for RP records to allow mailbox to be specified without the domain name
 * - Add support for skipLookup parameter on ANAME records
 *
 * ## v4.1.5
 *
 * - Fixed display of vanity nameservers assigned to a domain
 * - Fixed issue creating/updating IP Filters with an empty continents collection
 *
 * ## v4.1.4
 *
 * - Fixed issue with A, AAAA and ANAME records in standard (round-robin) mode
 * - Fixed documentation for vanity nameservers
 *
 * ## v4.1.3
 *
 * - Fixed issue with CNAME records in certain scenarios
 * - Fixed pagination perPage parameter
 *
 * ## v4.1.2
 *
 * - Fixed issue with Sonar Check IDs in Round-Robin Failover records
 * - Improved documentation around Postman Collection
 *
 * ## v4.1.1
 *
 * - Added domain search endpoint at `/v4/search/domains`
 * - Updated documentation for absolute URLs and hostname used for URLs
 * - Fixed issue with Sonar Check IDs in domain records
 *
 * ## v4.1.0
 *
 * - Added account analytics endpoint at `/v4/analytics`
 * - Added domain analytics endpoint at `/v4/domains/{id}/analytics`
 * - Updated example Postman pre-request script
 *
 */
@service({
  title: "Constellix DNS API",
})
@info({
  version: "4.2.7",
  contact: {
    email: "support@constellix.com",
    name: "Constellix",
  },
})
namespace ConstellixDNSAPI;

/**
 * Retrieve query analytics for your account
 *
 */
@tag("Analytics")
@route("/analytics")
@get
op listAnalytics(): listAnalytics200ApplicationJsonResponse;

/**
 * A list of all announcements.
 *
 */
@tag("Announcements")
@route("/announcements")
@get
op listAnnouncements(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listAnnouncements200ApplicationJsonResponse;

/**
 * Fetches a single announcement.
 *
 */
@tag("Announcements")
@route("/announcements/{id}")
@get
op getAnnouncement(): getAnnouncement200ApplicationJsonResponse;

/**
 * Fetches a paginated list of Contact Lists for the account.
 *
 */
@tag("Contact Lists")
@route("/contactlists")
@get
op listContactLists(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listContactLists200ApplicationJsonResponse;

/**
 * Creates a new Contact List
 *
 */
@tag("Contact Lists")
@route("/contactlists")
@post
op createContactList(
  @bodyRoot body: `contactlist-createupdate`,
): createContactList202ApplicationJsonResponse;

/**
 * Remove the Contact List.
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}")
@delete
op deleteContactList(): deleteContactList204Response;

/**
 * Fetches a single Contact List for the account.
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}")
@get
op getContactList(): getContactList200ApplicationJsonResponse;

/**
 * Updates the Contact List object. Fields left out of the request will not be updated.
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}")
@put
op updateContactList(
  @bodyRoot body: `contactlist-createupdate`,
): updateContactList200ApplicationJsonResponse;

/**
 * A list of all emails in the contact list. The list is paginated.
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/emails")
@get
op listContactListEmails(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listContactListEmails200ApplicationJsonResponse;

/**
 * Adds an email to the contact list
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/emails")
@post
op addContactListEmail(
  @bodyRoot body: `contactlist-email-createupdate`,
): addContactListEmail202ApplicationJsonResponse;

/**
 * Remove the Email Address from the contact list
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/emails/{id}")
@delete
op deleteContactListEmail(): deleteContactListEmail204Response;

/**
 * Fetches a single Email Address in a contact list
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/emails/{id}")
@get
op getContactListEmail(): getContactListEmail200ApplicationJsonResponse;

/**
 * A list of all Slack webhooks in the contact list. The list is paginated.
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/slack")
@get
op listContactListSlackWebhooks(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listContactListSlackWebhooks200ApplicationJsonResponse;

/**
 * Adds a Slack webhook to the contact list
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/slack")
@post
op createContactListSlackWebhook(
  @bodyRoot body: `contactlist-slack-createupdate`,
): createContactListSlackWebhook202ApplicationJsonResponse;

/**
 * Remove the Slack webhook from the contact list
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/slack/{id}")
@delete
op deleteContactListSlackWebhook(): deleteContactListSlackWebhook204Response;

/**
 * Fetches a single Slack Webhook in a contact list
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/slack/{id}")
@get
op getContactListSlackWebhook(): getContactListSlackWebhook200ApplicationJsonResponse;

/**
 * A list of all MS Teams webhooks in the contact list. The list is paginated.
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/teams")
@get
op listContactListTeams(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listContactListTeams200ApplicationJsonResponse;

/**
 * Adds a Microsoft Teams webhook to the contact list
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/teams")
@post
op createContactListTeam(
  @bodyRoot body: `contactlist-teams-createupdate`,
): createContactListTeam202ApplicationJsonResponse;

/**
 * Remove the Microsoft Teams webhook from the contact list
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/teams/{id}")
@delete
op deleteContactListTeam(): deleteContactListTeam204Response;

/**
 * Fetches a single Microsoft Teams Webhook in a contact list
 *
 */
@tag("Contact Lists")
@route("/contactlists/{contactlist_id}/teams/{id}")
@get
op getContactListTeam(): getContactListTeam200ApplicationJsonResponse;

/**
 * A list of all domains in the account. The list is paginated and can also be searched and filtered.
 *
 */
@tag("Domains")
@route("/domains")
@get
op listDomains(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listDomains200ApplicationJsonResponse;

/**
 * Creates a new domain in your account.
 */
@tag("Domains")
@route("/domains")
@post
op createDomain(@bodyRoot body: `domain-create`): createDomain202ApplicationJsonResponse;

/**
 * Remove the domain and all records for it
 */
@tag("Domains")
@route("/domains/{domain_id}")
@delete
op deleteDomain(): deleteDomain204Response;

/**
 * Fetches a single domain in your account. This will contain all the details of the domain
 * and links to the records.
 *
 */
@tag("Domains")
@route("/domains/{domain_id}")
@get
op getDomain(): getDomain200ApplicationJsonResponse;

/**
 * Updates the domain object. Fields left out of the request will not be updated.
 *
 */
@tag("Domains")
@route("/domains/{domain_id}")
@put
op updateDomain(@bodyRoot body: `domain-update`): updateDomain200ApplicationJsonResponse;

/**
 * Retrieve query analytics for the domain
 *
 */
@tag("Domains")
@route("/domains/{domain_id}/analytics")
@get
op getDomainAnalytics(): getDomainAnalytics200ApplicationJsonResponse;

/**
 * Exports the entire domain in the BIND format.
 *
 */
@tag("Domains")
@route("/domains/{domain_id}/bind")
@get
op getDomainBindExport(): getDomainBindExport200TextPlainResponse;

/**
 * Fetches a history of the domain to allow you to see when it has changed.
 *
 */
@tag("Domain History")
@route("/domains/{domain_id}/history")
@get
op listDomainHistory(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listDomainHistory200ApplicationJsonResponse;

/**
 * Fetches a single history version for the domain.
 *
 */
@tag("Domain History")
@route("/domains/{domain_id}/history/{version}")
@get
op getDomainHistoryVersion(): getDomainHistoryVersion200ApplicationJsonResponse;

/**
 * Update the domain to the specified version in the history of the domain
 *
 */
@tag("Domain History")
@route("/domains/{domain_id}/history/{version}/apply")
@post
op applyDomainHistory(): applyDomainHistory204Response;

/**
 * Snapshot this history version. Snapshots are persisted and kept until they are deleted by the user.
 *
 */
@tag("Domain History")
@route("/domains/{domain_id}/history/{version}/snapshot")
@post
op snapshotDomainHistory(): snapshotDomainHistory202ApplicationJsonResponse;

/**
 * Fetches the current nameservers the domain is using. This may be different from the ones assigned to it by Constellix.
 *
 */
@tag("Domains")
@route("/domains/{domain_id}/nameservers")
@get
op getDomainNameservers(): getDomainNameservers200ApplicationJsonResponse;

/**
 * Fetches a paginated list of all records for the domain. This list can be ordered and filtered.
 *
 */
@tag("Domain Records")
@route("/domains/{domain_id}/records")
@get
op listDomainRecords(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listDomainRecords200ApplicationJsonResponse;

/**
 * Creates a new record for the domain
 */
@tag("Domain Records")
@route("/domains/{domain_id}/records")
@post
op createDomainRecord(
  @bodyRoot body: `record-create`,
): createDomainRecord202ApplicationJsonResponse;

/**
 * Remove the record
 */
@tag("Domain Records")
@route("/domains/{domain_id}/records/{id}")
@delete
op deleteDomainRecord(@bodyRoot body: `record-update`): deleteDomainRecord204Response;

/**
 * Fetches a single record for the domain.
 *
 */
@tag("Domain Records")
@route("/domains/{domain_id}/records/{id}")
@get
op getDomainRecord(): getDomainRecord200ApplicationJsonResponse;

/**
 * Updates the record object. Fields left out of the request will not be updated.
 *
 * If the mode is updated and value is not provided with the request, the current value will be validated alongside the new mode. If it is invalid, a 422 response will be issued.
 *
 */
@tag("Domain Records")
@route("/domains/{domain_id}/records/{id}")
@put
op updateDomainRecord(
  @bodyRoot body: `record-update`,
): updateDomainRecord200ApplicationJsonResponse;

/**
 * Fetches saved snapshots of the history the domain to allow you to see when it has changed.
 *
 */
@tag("Domain Snapshots")
@route("/domains/{domain_id}/snapshots")
@get
op listDomainSnapshots(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listDomainSnapshots200ApplicationJsonResponse;

/**
 * Remove the snapshot.
 */
@tag("Domain Snapshots")
@route("/domains/{domain_id}/snapshots/{version}")
@delete
op deleteDomainSnapshot(): deleteDomainSnapshot204Response;

/**
 * Fetches a single snapshot for the domain.
 *
 */
@tag("Domain Snapshots")
@route("/domains/{domain_id}/snapshots/{version}")
@get
op getDomainSnapshot(): getDomainSnapshot200ApplicationJsonResponse;

/**
 * Update the domain to the specified snapshot in the history of the domain
 *
 */
@tag("Domain Snapshots")
@route("/domains/{domain_id}/snapshots/{version}/apply")
@post
op applyDomainSnapshot(): applyDomainSnapshot204Response;

/**
 * Fetches a paginated list of Geo Proximity locations for the account.
 *
 */
@tag("Geo Proximity Locations")
@route("/geoproximities")
@get
op listGeoProximityLocations(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listGeoProximityLocations200ApplicationJsonResponse;

/**
 * Creates a new Geo Proximity Location.
 *
 * Only name, longitude and latitude are required and used, the other
 * properties are just for information.
 *
 */
@tag("Geo Proximity Locations")
@route("/geoproximities")
@post
op createGeoProximityLocation(
  @bodyRoot body: `geoproximity-createupdate`,
): createGeoProximityLocation202ApplicationJsonResponse;

/**
 * Remove the Geo Proximity Location.
 */
@tag("Geo Proximity Locations")
@route("/geoproximities/{id}")
@delete
op deleteGeoProximityLocation(): deleteGeoProximityLocation204Response;

/**
 * Fetches a single Geo Proximity location for the account.
 *
 */
@tag("Geo Proximity Locations")
@route("/geoproximities/{id}")
@get
op getGeoProximityLocation(): getGeoProximityLocation200ApplicationJsonResponse;

/**
 * Updates the Geo Proximity Location object. Fields left out of the request will not be updated.
 * Only name, longitude and latitude are required and used, the other properties are just for information.
 *
 */
@tag("Geo Proximity Locations")
@route("/geoproximities/{id}")
@put
op updateGeoProximityLocation(
  @bodyRoot body: `geoproximity-createupdate`,
): updateGeoProximityLocation200ApplicationJsonResponse;

/**
 * Fetches a paginated list of IP filters for the account.
 *
 */
@tag("IP Filters")
@route("/ipfilters")
@get
op listIPFilters(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listIPFilters200ApplicationJsonResponse;

/**
 * Creates a new IP filter.
 *
 */
@tag("IP Filters")
@route("/ipfilters")
@post
op createIPFilter(
  @bodyRoot body: `ipfilter-createupdate`,
): createIPFilter202ApplicationJsonResponse;

/**
 * Remove the IP Filter.
 */
@tag("IP Filters")
@route("/ipfilters/{id}")
@delete
op deleteIPFilter(): deleteIPFilter204Response;

/**
 * Fetches a single IP Filter for the account.
 *
 */
@tag("IP Filters")
@route("/ipfilters/{id}")
@get
op getIPFilter(): getIPFilter200ApplicationJsonResponse;

/**
 * Updates the IP Filter object. Fields left out of the request will not be updated.
 *
 */
@tag("IP Filters")
@route("/ipfilters/{id}")
@put
op updateIPFilter(
  @bodyRoot body: `ipfilter-createupdate`,
): updateIPFilter200ApplicationJsonResponse;

/**
 * Returns basic information about the API and the current user. This will include your IP address, the API timestamp and the version.
 *
 */
@tag("General")
@route("/ping")
@get
op getPing(): getPing200ApplicationJsonResponse;

/**
 * Returns a paginated list of pools. This includes all types of pool.
 *
 */
@tag("Pools")
@route("/pools")
@get
op listPools(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listPools200ApplicationJsonResponse;

/**
 * Creates a new A, AAAA or CNAME pool.
 *
 */
@tag("Pools")
@route("/pools")
@post
op createPool(@bodyRoot body: `pool-createupdate`): createPool202ApplicationJsonResponse;

/**
 * Remove the pool.
 */
@tag("Pools")
@route("/pools/{type}/{id}")
@delete
op deletePool(): deletePool204Response;

/**
 * Fetches a single Pool for the account.
 *
 */
@tag("Pools")
@route("/pools/{type}/{id}")
@get
op getPool(): getPool200ApplicationJsonResponse;

/**
 * Updates the pool object. Fields left out of the request will not be updated.
 *
 */
@tag("Pools")
@route("/pools/{type}/{id}")
@put
op updatePool(@bodyRoot body: `pool-createupdate`): updatePool200ApplicationJsonResponse;

/**
 * Search for domains within your account
 *
 */
@tag("Search")
@route("/search/domains")
@get
op searchDomains(): searchDomains200ApplicationJsonResponse;

/**
 * Fetches a paginated list ofTags for the account.
 *
 */
@tag("Tags")
@route("/tags")
@get
op listTags(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listTags200ApplicationJsonResponse;

/**
 * Creates a new Tag
 *
 */
@tag("Tags")
@route("/tags")
@post
op createTag(@bodyRoot body: `tag-createupdate`): createTag202ApplicationJsonResponse;

/**
 * Remove the Tag.
 */
@tag("Tags")
@route("/tags/{id}")
@delete
op deleteTag(): deleteTag204Response;

/**
 * Fetches a single Tag for the account.
 *
 */
@tag("Tags")
@route("/tags/{id}")
@get
op getTag(): getTag200ApplicationJsonResponse;

/**
 * Updates the Tag object. Fields left out of the request will not be updated.
 *
 */
@tag("Tags")
@route("/tags/{id}")
@put
op updateTag(@bodyRoot body: `tag-createupdate`): updateTag200ApplicationJsonResponse;

/**
 * Returns a paginated list of domain templates.
 *
 */
@tag("Templates")
@route("/templates")
@get
op listTemplates(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listTemplates200ApplicationJsonResponse;

/**
 * Adds a new template.
 *
 */
@tag("Templates")
@route("/templates")
@post
op createTemplate(
  @bodyRoot body: `template-createupdate`,
): createTemplate202ApplicationJsonResponse;

/**
 * Remove the template and all records for it
 */
@tag("Templates")
@route("/templates/{template_id}")
@delete
op deleteTemplate(): deleteTemplate204Response;

/**
 * Retrieves a single template.
 *
 */
@tag("Templates")
@route("/templates/{template_id}")
@get
op getTemplate(): getTemplate200ApplicationJsonResponse;

/**
 * Updates the properties on the template. Properties that aren't included in the request will not be updated.
 *
 */
@tag("Templates")
@route("/templates/{template_id}")
@put
op updateTemplate(
  @bodyRoot body: `template-createupdate`,
): updateTemplate200ApplicationJsonResponse;

/**
 * Fetches a paginated list of all records for the template. This list can be ordered and filtered.
 *
 */
@tag("Template Records")
@route("/templates/{template_id}/records")
@get
op listTemplateRecords(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listTemplateRecords200ApplicationJsonResponse;

/**
 * Creates a new record for the template
 */
@tag("Template Records")
@route("/templates/{template_id}/records")
@post
op createTemplateRecord(
  @bodyRoot body: `record-create`,
): createTemplateRecord202ApplicationJsonResponse;

/**
 * Remove the record.
 */
@tag("Template Records")
@route("/templates/{template_id}/records/{id}")
@delete
op deleteTemplateRecord(@bodyRoot body: `record-update`): deleteTemplateRecord204Response;

/**
 * Fetches a single record for the template.
 *
 */
@tag("Template Records")
@route("/templates/{template_id}/records/{id}")
@get
op getTemplateRecord(): getTemplateRecord200ApplicationJsonResponse;

/**
 * Updates the record object. Fields left out of the request will not be updated.
 *
 * If the mode is updated and value is not provided with the request, the current value will be validated alongside the new mode. If it is invalid, a 422 response will be issued.
 *
 */
@tag("Template Records")
@route("/templates/{template_id}/records/{id}")
@put
op updateTemplateRecord(
  @bodyRoot body: `record-update`,
): updateTemplateRecord200ApplicationJsonResponse;

/**
 * Fetches a paginated list of Vanity Nameservers for the account.
 *
 */
@tag("Vanity Nameservers")
@route("/vanitynameservers")
@get
op listVanityNameservers(
  ...Parameters.PaginationPage,
  ...Parameters.PaginationPerPage,
): listVanityNameservers200ApplicationJsonResponse;

/**
 * Creates a new Vanity Nameserver
 *
 */
@tag("Vanity Nameservers")
@route("/vanitynameservers")
@post
op createVanityNameserver(
  @bodyRoot body: `vanitynameserver-create`,
): createVanityNameserver202ApplicationJsonResponse;

/**
 * Remove the Vanity Nameserver.
 */
@tag("Vanity Nameservers")
@route("/vanitynameservers/{id}")
@delete
op deleteVanityNameserver(): deleteVanityNameserver204Response;

/**
 * Fetches a single Vanity Nameserver for the account.
 *
 */
@tag("Vanity Nameservers")
@route("/vanitynameservers/{id}")
@get
op getVanityNameserver(): getVanityNameserver200ApplicationJsonResponse;

/**
 * Updates the Vanity Nameserver object. Fields left out of the request will not be updated.
 *
 */
@tag("Vanity Nameservers")
@route("/vanitynameservers/{id}")
@put
op updateVanityNameserver(
  @bodyRoot body: `vanitynameserver-update`,
): updateVanityNameserver200ApplicationJsonResponse;

namespace Parameters {
  model PaginationPage {
    /**
     * The page of results to fetch
     */
    @query page?: integer = 1;
  }
  model PaginationPerPage {
    /**
     * The number of results per page
     */
    @query perPage?: integer = 50;
  }
}
