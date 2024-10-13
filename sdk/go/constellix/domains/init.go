// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domains

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/cloudy-sky-software/pulumi-constellix/sdk/go/constellix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "constellix:domains:ApplyDomainHistory":
		r = &ApplyDomainHistory{}
	case "constellix:domains:ApplyDomainSnapshot":
		r = &ApplyDomainSnapshot{}
	case "constellix:domains:Domain":
		r = &Domain{}
	case "constellix:domains:DomainRecord":
		r = &DomainRecord{}
	case "constellix:domains:SnapshotDomainHistory":
		r = &SnapshotDomainHistory{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"constellix",
		"domains",
		&module{version},
	)
}
