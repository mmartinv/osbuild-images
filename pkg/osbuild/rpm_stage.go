package osbuild

import (
	"slices"

	"github.com/osbuild/images/pkg/rpmmd"
)

type RPMStageOptions struct {
	// Use the given path as RPM database
	DBPath string `json:"dbpath,omitempty"`

	// Array of GPG key contents to import
	GPGKeys []string `json:"gpgkeys,omitempty"`

	// Array of files in the tree containing GPG keys to import
	GPGKeysFromTree []string `json:"gpgkeys.fromtree,omitempty"`

	// Prevent dracut from running
	DisableDracut bool `json:"disable_dracut,omitempty"`

	Exclude *Exclude `json:"exclude,omitempty"`

	// Create the '/run/ostree-booted' marker
	OSTreeBooted *bool `json:"ostree_booted,omitempty"`

	// Set environment variables understood by kernel-install and plugins (kernel-install(8))
	KernelInstallEnv *KernelInstallEnv `json:"kernel_install_env,omitempty"`
}

type Exclude struct {
	// Do not install documentation.
	Docs bool `json:"docs,omitempty"`
}

type KernelInstallEnv struct {
	// Sets $BOOT_ROOT for kernel-install to override
	// $KERNEL_INSTALL_BOOT_ROOT, the installation location for boot entries
	BootRoot string `json:"boot_root,omitempty"`
}

// RPMPackage represents one RPM, as referenced by its content hash
// (checksum). The files source must indicate where to fetch the given
// RPM. If CheckGPG is `true` the RPM must be signed with one of the
// GPGKeys given in the RPMStageOptions.
type RPMPackage struct {
	Checksum string `json:"checksum"`
	CheckGPG bool   `json:"check_gpg,omitempty"`
}

func (RPMStageOptions) isStageOptions() {}

// RPMStageInputs defines a collection of packages to be installed by the RPM
// stage.
type RPMStageInputs struct {
	// Packages to install
	Packages *FilesInput `json:"packages"`
}

func (RPMStageInputs) isStageInputs() {}

// RPM package reference metadata
type RPMStageReferenceMetadata struct {
	// This option defaults to `false`, therefore it does not need to be
	// defined as pointer to bool and can be omitted.
	CheckGPG bool `json:"rpm.check_gpg,omitempty"`
}

func (*RPMStageReferenceMetadata) isFilesInputRefMetadata() {}

// NewRPMStage creates a new RPM stage.
func NewRPMStage(options *RPMStageOptions, inputs *RPMStageInputs) *Stage {
	return &Stage{
		Type:    "org.osbuild.rpm",
		Inputs:  inputs,
		Options: options,
	}
}

// RPMStageMetadata gives the set of packages installed by the RPM stage
type RPMStageMetadata struct {
	Packages []RPMPackageMetadata `json:"packages"`
}

// RPMPackageMetadata contains the metadata extracted from one RPM header
type RPMPackageMetadata struct {
	Name    string  `json:"name"`
	Version string  `json:"version"`
	Release string  `json:"release"`
	Epoch   *string `json:"epoch"`
	Arch    string  `json:"arch"`
	SigMD5  string  `json:"sigmd5"`
	SigPGP  string  `json:"sigpgp"`
	SigGPG  string  `json:"siggpg"`
}

func (pkgmd RPMPackageMetadata) Signature() *string {
	if pkgmd.SigGPG != "" {
		return &pkgmd.SigGPG
	} else if pkgmd.SigPGP != "" {
		return &pkgmd.SigPGP
	}
	return nil
}

func (RPMStageMetadata) isStageMetadata() {}

func NewRpmStageSourceFilesInputs(specs []rpmmd.PackageSpec) *RPMStageInputs {
	input := NewFilesInput(pkgRefs(specs))
	return &RPMStageInputs{Packages: input}
}

func pkgRefs(specs []rpmmd.PackageSpec) FilesInputRef {
	refs := make([]FilesInputSourceArrayRefEntry, len(specs))
	for idx, pkg := range specs {
		var pkgMetadata FilesInputRefMetadata
		if pkg.CheckGPG {
			pkgMetadata = &RPMStageReferenceMetadata{
				CheckGPG: pkg.CheckGPG,
			}
		}
		refs[idx] = NewFilesInputSourceArrayRefEntry(pkg.Checksum, pkgMetadata)
	}
	return NewFilesInputSourceArrayRef(refs)
}

func NewRPMStageOptions(repos []rpmmd.RepoConfig) *RPMStageOptions {
	gpgKeys := make([]string, 0)
	keyMap := make(map[string]bool) // for deduplicating keys
	for _, repo := range repos {
		if len(repo.GPGKeys) == 0 {
			continue
		}
		for _, key := range repo.GPGKeys {
			if !keyMap[key] {
				gpgKeys = append(gpgKeys, key)
				keyMap[key] = true
			}
		}
	}

	slices.Sort(gpgKeys)
	return &RPMStageOptions{
		GPGKeys: gpgKeys,
	}
}
