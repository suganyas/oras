/*
Copyright The ORAS Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package foobar

import (
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/test/e2e/internal/utils/match"
)

var (
	Tag              = "foobar"
	Digest           = "sha256:fd6ed2f36b5465244d5dc86cb4e7df0ab8a9d24adc57825099f522fe009a22bb"
	ManifestStateKey = match.StateKey{Digest: "fd6ed2f36b54", Name: "application/vnd.oci.image.manifest.v1+json"}

	FileLayerNames = []string{
		"foobar/foo1",
		"foobar/foo2",
		"foobar/bar",
	}
	FileConfigName     = "foobar/config.json"
	FileConfigStateKey = match.StateKey{
		Digest: "46b68ac1696c", Name: "application/vnd.unknown.config.v1+json",
	}
	FileStateKeys = []match.StateKey{
		{Digest: "2c26b46b68ff", Name: FileLayerNames[0]},
		{Digest: "2c26b46b68ff", Name: FileLayerNames[1]},
		{Digest: "fcde2b2edba5", Name: FileLayerNames[2]},
	}

	ConfigDesc         = "{\"mediaType\":\"application/vnd.unknown.config.v1+json\",\"digest\":\"sha256:44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a\",\"size\":2}"
	AttachFileName     = "foobar/to-be-attached"
	AttachFileMedia    = "test/oras.e2e"
	AttachFileStateKey = match.StateKey{
		Digest: "d3b29f7d12d9", Name: AttachFileName,
	}

	ImageLayerNames = []string{
		"foo1",
		"foo2",
		"bar",
	}
	ImageLayerStateKeys = []match.StateKey{
		{Digest: "2c26b46b68ff", Name: ImageLayerNames[0]},
		{Digest: "2c26b46b68ff", Name: ImageLayerNames[1]},
		{Digest: "fcde2b2edba5", Name: ImageLayerNames[2]},
	}
	ImageConfigName = "config.json"

	FooBlobDigest     = "sha256:2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
	FooBlobContent    = "foo"
	FooBlobDescriptor = `{"mediaType":"application/octet-stream","digest":"sha256:2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae","size":3}`
)

func ImageConfigStateKey(configName string) match.StateKey {
	return match.StateKey{Digest: "44136fa355b3", Name: configName}
}

// referrers
var (
	SBOMImageReferrer = ocispec.Descriptor{
		MediaType: "application/vnd.oci.image.manifest.v1+json",
		Digest:    digest.Digest("sha256:32b78bd00723cd7d5251d4586f84d252530b7b5fe1c4104532767e6da4e04e47"),
		Size:      660,
		Annotations: map[string]string{
			"org.opencontainers.image.created": "2023-01-18T08:37:42Z",
		},
		ArtifactType: "test.sbom.file",
	}
	SignatureImageReferrer = ocispec.Descriptor{
		MediaType: "application/vnd.oci.image.manifest.v1+json",
		Digest:    digest.Digest("sha256:0e007dcb9ded7f49c4dc8e3eed4a446712eb6fdf08a665a4f2352d6d2f8bdf17"),
		Size:      670,
	}
	SBOMArtifactReferrer = ocispec.Descriptor{
		MediaType: "application/vnd.oci.artifact.manifest.v1+json",
		Digest:    digest.Digest("sha256:8d7a27ff2662dae183f762d281f46d626ba7b6e56a72cc9959cdbcd91aad7fbc"),
		Size:      547,
		Annotations: map[string]string{
			"org.opencontainers.artifact.created": "2023-01-16T05:49:46Z",
		},
		ArtifactType: "test.sbom.file",
	}
	SignatureArtifactReferrer = ocispec.Descriptor{
		MediaType: "application/vnd.oci.artifact.manifest.v1+json",
		Digest:    digest.Digest("sha256:2dbea575a3490375f5052fbeb380a2f498866d99eb809b4168e49e224a274a39"),
		Size:      560,
	}
	ArtifactReferrerStateKeys = []match.StateKey{
		{Digest: "8d7a27ff2662", Name: "application/vnd.oci.artifact.manifest.v1+json"},
		{Digest: "2dbea575a349", Name: "application/vnd.oci.artifact.manifest.v1+json"},
	}
	ImageReferrersStateKeys = []match.StateKey{
		{Digest: "0e007dcb9ded", Name: "application/vnd.oci.image.manifest.v1+json"},
		{Digest: "32b78bd00723", Name: "application/vnd.oci.image.manifest.v1+json"},
	}
	ImageReferrerConfigStateKeys = []match.StateKey{
		{Digest: "44136fa355b3", Name: "test.signature.file"},
		{Digest: "44136fa355b3", Name: "test.sbom.file"},
	}
	FallbackImageReferrersStateKeys = []match.StateKey{
		{Digest: "316405db72cc", Name: "application/vnd.oci.image.manifest.v1+json"},
		{Digest: "8b3f7e000c4a", Name: "application/vnd.oci.image.manifest.v1+json"},
	}
)

// fallback referrers
var (
	FallbackSignatureImageReferrer = ocispec.Descriptor{
		MediaType: "application/vnd.oci.image.manifest.v1+json",
		Digest:    digest.Digest("sha256:8b3f7e000c4a6d32cd6bfcabfe874ed470d470501a09adc65afaf1c342f988ff"),
		Size:      670,
	}

	FallbackSBOMImageReferrer = ocispec.Descriptor{
		MediaType: "application/vnd.oci.image.manifest.v1+json",
		Digest:    digest.Digest("sha256:316405db72cc8f0212c19db23b498f9af8a456c9cd288f9e33acd1ba9e7cd534"),
		Size:      660,
		Annotations: map[string]string{
			"org.opencontainers.image.created": "2023-01-29T02:32:18Z",
		},
		ArtifactType: "test.sbom.file",
	}
)
