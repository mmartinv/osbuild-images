#!/usr/bin/env bash
#
# Build an image and store the hash of the manifest in S3 if successful.

set -euxo pipefail

distro="${1}"
arch=$(uname -m)
imgtype="${2}"
config="${3}"

echo "Building image ${distro}/${imgtype} using config ${config}"
cat "${config}"  # print the config for logging
sudo go run ./cmd/build -output ./build -distro "${distro}" -image "${imgtype}" -config "${config}"

echo "Build finished!!"

echo "Registering successful build in S3"

u() {
    echo "${1}" | tr - _
}


config_name=$(jq -r .name "${config}")
builddir="./build/$(u "${distro}")-$(u "${arch}")-$(u "${imgtype}")-$(u "${config_name}")"
manifest="${builddir}/manifest.json"

# calculate manifest ID based on hash of concatenated stage IDs
manifest_id=$(osbuild --inspect "${manifest}" | jq -r '.pipelines[-1].stages[-1].id')

osbuild_ver=$(osbuild --version)

# TODO: Include osbuild commit hash
filename="${manifest_id}.json"
cat << EOF > "${filename}"
{
  "distro": "${distro}",
  "arch": "${arch}",
  "image-type": "${imgtype}",
  "config": "${config_name}",
  "manifest-checksum": "${manifest_id}",
  "obuild-version": "${osbuild_ver}",
  "commit": "${CI_COMMIT_SHA}"
}
EOF

s3url="s3://image-builder-ci-artifacts/images/builds/${distro}/${arch}/${filename}"

echo "Uploading ${filename} to ${s3url}"
AWS_SECRET_ACCESS_KEY="$V2_AWS_SECRET_ACCESS_KEY" \
AWS_ACCESS_KEY_ID="$V2_AWS_ACCESS_KEY_ID" \
s3cmd --acl-private put "${filename}" "${s3url}"
