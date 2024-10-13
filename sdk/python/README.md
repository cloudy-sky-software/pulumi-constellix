# Pulumi Provider for Constellix

This Pulumi provider was generated using Constellix [v4 REST API](https://api.dns.constellix.com/v4/docs).

The OpenAPI spec is used as the source to generate Pulumi schema and the Pulumi SDKs:

1. [`pulschema`](https://github.com/cloudy-sky-software/pulschema) to convert an OpenAPI spec to Pulumi schema
2. [`pulumi-provider-framework`](https://github.com/cloudy-sky-software/pulumi-provider-framework) to make HTTP requests against the cloud provider API. It uses the metadata returned by `pulschema` as one of the outputs from
   converting an OpenAPI spec.

## Package SDKs

- Node.js: https://www.npmjs.com/package/@cloudyskysoftware/pulumi-constellix
- Python: https://pypi.org/project/pulumi-constellix/
- .NET: https://www.nuget.org/packages/CloudySkySoftware.Pulumi.Constellix
- Go: `import github.com/cloudyskysoftware/pulumi-constellix/sdk/go/constellix`

## Using The Provider

You'll need an API key and a secret key. Follow Constellix's [docs]([https://render.com/docs/api#getting-started](https://support.constellix.com/support/solutions/articles/47001200127-how-to-generate-an-api-key)) for creating one.
Then set the API key and secret as secret stack configs with `pulumi config set --secret constellix:apiKey` and `pulumi config set --secret constellix:secretKey`.

## Releasing A New Version

:info: Switch to the `main` branch first and get the latest `git pull origin main && git fetch`. Check what the last release tag was.

1. Regular releases should just increment the patch version unless a minor or a major (breaking changes) version bump is warranted.
1. Update the `CHANGELOG.md` with notes about what will be included in this release.
1. Commit the changelog with `git commit -am "vX.Y.Z"` or something similar and push `git push origin main`.
1. Tag the commit with the release version by running

   ```bash
   git tag vX.Y.Z
   git tag sdk/vX.Y.Z
   ```

1. Push the tags.

   ```bash
   git push --tags
   ```
