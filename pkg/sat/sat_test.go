package sat

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/rmohr/bazeldnf/pkg/api"
)

func TestRecursive(t *testing.T) {
	g := NewGomegaWithT(t)
	f, err := os.Open("testdata/perl-pathtools-fc32.xml")
	g.Expect(err).ToNot(HaveOccurred())
	defer f.Close()
	repo := &api.Repository{}
	err = xml.NewDecoder(f).Decode(repo)
	g.Expect(err).ToNot(HaveOccurred())

	for _, pkg := range repo.Packages {
		t.Run(fmt.Sprintf("find solution for %s", pkg.Name), func(t *testing.T) {
			g := NewGomegaWithT(t)
			resolver := NewResolver(false)
			packages := []*api.Package{}
			for i, _ := range repo.Packages {
				packages = append(packages, &repo.Packages[i])
			}
			err = resolver.LoadInvolvedPackages(packages)
			g.Expect(err).ToNot(HaveOccurred())
			err = resolver.ConstructRequirements([]string{pkg.Name})
			g.Expect(err).ToNot(HaveOccurred())
			_, _, err := resolver.Resolve()
			if err != nil {
				t.Fatalf("Failed to solve %s\n", pkg.Name)
			}
		})
	}
}

func Test(t *testing.T) {
	tests := []struct {
		name     string
		requires []string
		installs []string
		repofile string
		nobest   bool
		focus    bool
	}{
		{name: "should resolve bash",
			requires: []string{
				"bash",
				"fedora-release-server",
				"glibc-langpack-en",
			},
			installs: []string{
				"libgcc-0:10.2.1-1.fc32",
				"fedora-gpg-keys-0:32-6",
				"glibc-0:2.31-4.fc32",
				"glibc-langpack-en-0:2.31-4.fc32",
				"fedora-release-common-0:32-3",
				"glibc-common-0:2.31-4.fc32",
				"ncurses-base-0:6.1-15.20191109.fc32",
				"ncurses-libs-0:6.1-15.20191109.fc32",
				"fedora-release-server-0:32-3",
				"tzdata-0:2020a-1.fc32",
				"setup-0:2.13.6-2.fc32",
				"basesystem-0:11-9.fc32",
				"bash-0:5.0.17-1.fc32",
				"filesystem-0:3.14-2.fc32",
				"fedora-repos-0:32-6",
			},
			repofile: "testdata/bash-fc32.xml",
		},
		{
			name: "should resolve fedora-release-container",
			requires: []string{
				"fedora-release-container",
			},
			installs: []string{
				"fedora-gpg-keys-0:32-6",
				"fedora-repos-0:32-6",
				"fedora-release-container-0:32-3",
				"fedora-release-common-0:32-3",
			},
			nobest:   false,
			repofile: "testdata/fedora-release-container.xml",
		},
		{
			name: "should resolve libvirt-daemon-driver-storage-zfs",
			requires: []string{
				"libvirt-daemon-driver-storage-zfs",
				"glibc-langpack-en",
				"fedora-release-container",
				"coreutils-single",
				"libcurl-minimal",
				"libev",
				"libverto-libev",
			},
			installs: []string{
				"xz-libs-0:5.2.5-1.fc32",
				"libcap-ng-0:0.7.11-1.fc32",
				"openldap-0:2.4.47-5.fc32",
				"iproute-tc-0:5.7.0-1.fc32",
				"numactl-libs-0:2.0.12-4.fc32",
				"libssh2-0:1.9.0-5.fc32",
				"systemd-pam-0:245.8-2.fc32",
				"libnetfilter_conntrack-0:1.0.7-4.fc32",
				"glibc-common-0:2.31-4.fc32",
				"shadow-utils-2:4.8.1-2.fc32",
				"elfutils-libelf-0:0.181-1.fc32",
				"linux-atm-libs-0:2.5.1-26.fc32",
				"ncurses-base-0:6.1-15.20191109.fc32",
				"libvirt-daemon-driver-storage-core-0:6.1.0-4.fc32",
				"perl-PathTools-0:3.78-441.fc32",
				"libargon2-0:20171227-4.fc32",
				"util-linux-0:2.35.2-1.fc32",
				"fedora-release-common-0:32-3",
				"libverto-0:0.3.0-9.fc32",
				"libselinux-0:3.0-5.fc32",
				"perl-threads-shared-0:1.60-441.fc32",
				"cracklib-0:2.9.6-22.fc32",
				"libnfsidmap-1:2.5.1-4.rc4.fc32",
				"coreutils-single-0:8.32-4.fc32.1",
				"libtirpc-0:1.2.6-1.rc4.fc32",
				"kmod-libs-0:27-1.fc32",
				"libvirt-daemon-0:6.1.0-4.fc32",
				"systemd-0:245.8-2.fc32",
				"libacl-0:2.2.53-5.fc32",
				"libdb-0:5.3.28-40.fc32",
				"p11-kit-trust-0:0.23.21-2.fc32",
				"dbus-broker-0:24-1.fc32",
				"libnl3-0:3.5.0-2.fc32",
				"polkit-libs-0:0.116-7.fc32",
				"gssproxy-0:0.8.2-8.fc32",
				"python-setuptools-wheel-0:41.6.0-2.fc32",
				"pam-0:1.3.1-26.fc32",
				"zfs-fuse-0:0.7.2.2-14.fc32",
				"cyrus-sasl-lib-0:2.1.27-4.fc32",
				"libgcrypt-0:1.8.5-3.fc32",
				"libstdc++-0:10.2.1-1.fc32",
				"quota-nls-1:4.05-9.fc32",
				"libvirt-daemon-driver-storage-zfs-0:6.1.0-4.fc32",
				"libcollection-0:0.7.0-44.fc32",
				"gawk-0:5.0.1-7.fc32",
				"fuse-libs-0:2.9.9-9.fc32",
				"dbus-1:1.12.20-1.fc32",
				"openssl-libs-1:1.1.1g-1.fc32",
				"nettle-0:3.5.1-5.fc32",
				"libuuid-0:2.35.2-1.fc32",
				"libverto-libev-0:0.3.0-9.fc32",
				"popt-0:1.16-19.fc32",
				"libxcrypt-0:4.4.17-1.fc32",
				"glibc-langpack-en-0:2.31-4.fc32",
				"sed-0:4.5-5.fc32",
				"libref_array-0:0.1.5-44.fc32",
				"libwsman1-0:2.6.8-12.fc32",
				"libxml2-0:2.9.10-7.fc32",
				"ncurses-libs-0:6.1-15.20191109.fc32",
				"libffi-0:3.1-24.fc32",
				"dmidecode-1:3.2-5.fc32",
				"dbus-libs-1:1.12.20-1.fc32",
				"elfutils-libs-0:0.181-1.fc32",
				"lzo-0:2.10-2.fc32",
				"keyutils-libs-0:1.6-4.fc32",
				"fuse-0:2.9.9-9.fc32",
				"libevent-0:2.1.8-8.fc32",
				"python-pip-wheel-0:19.3.1-4.fc32",
				"libpwquality-0:1.4.2-2.fc32",
				"libmount-0:2.35.2-1.fc32",
				"crypto-policies-0:20200619-1.git781bbd4.fc32",
				"libblkid-0:2.35.2-1.fc32",
				"perl-parent-1:0.238-1.fc32",
				"gzip-0:1.10-2.fc32",
				"libpath_utils-0:0.2.1-44.fc32",
				"libsemanage-0:3.0-3.fc32",
				"libutempter-0:1.1.6-18.fc32",
				"perl-Exporter-0:5.74-2.fc32",
				"fedora-repos-0:32-6",
				"pcre2-syntax-0:10.35-6.fc32",
				"qemu-img-2:4.2.1-1.fc32",
				"device-mapper-libs-0:1.02.171-1.fc32",
				"libunistring-0:0.9.10-7.fc32",
				"basesystem-0:11-9.fc32",
				"mozjs60-0:60.9.0-5.fc32",
				"polkit-pkla-compat-0:0.1-16.fc32",
				"qrencode-libs-0:4.0.2-5.fc32",
				"fedora-release-container-0:32-3",
				"python3-0:3.8.5-5.fc32",
				"readline-0:8.0-4.fc32",
				"p11-kit-0:0.23.21-2.fc32",
				"libssh-config-0:0.9.5-1.fc32",
				"perl-libs-4:5.30.3-456.fc32",
				"pcre-0:8.44-1.fc32",
				"libmnl-0:1.0.4-11.fc32",
				"grep-0:3.3-4.fc32",
				"libini_config-0:1.3.1-44.fc32",
				"libev-0:4.31-2.fc32",
				"yajl-0:2.1.0-14.fc32",
				"fuse-common-0:3.9.1-1.fc32",
				"numad-0:0.5-31.20150602git.fc32",
				"gnutls-0:3.6.15-1.fc32",
				"polkit-0:0.116-7.fc32",
				"perl-Scalar-List-Utils-3:1.54-440.fc32",
				"glibc-0:2.31-4.fc32",
				"libaio-0:0.3.111-7.fc32",
				"zlib-0:1.2.11-21.fc32",
				"filesystem-0:3.14-2.fc32",
				"cyrus-sasl-0:2.1.27-4.fc32",
				"bash-0:5.0.17-1.fc32",
				"e2fsprogs-libs-0:1.45.5-3.fc32",
				"parted-0:3.3-3.fc32",
				"perl-File-Path-0:2.17-1.fc32",
				"which-0:2.21-19.fc32",
				"krb5-libs-0:1.18.2-22.fc32",
				"perl-Socket-4:2.030-1.fc32",
				"libtasn1-0:4.16.0-1.fc32",
				"libnfnetlink-0:1.0.1-17.fc32",
				"lz4-libs-0:1.9.1-2.fc32",
				"bzip2-libs-0:1.0.8-2.fc32",
				"cryptsetup-libs-0:2.3.4-1.fc32",
				"json-c-0:0.13.1-13.fc32",
				"perl-Carp-0:1.50-440.fc32",
				"libcurl-minimal-0:7.69.1-6.fc32",
				"libnghttp2-0:1.41.0-1.fc32",
				"perl-interpreter-4:5.30.3-456.fc32",
				"pcre2-0:10.35-6.fc32",
				"libseccomp-0:2.5.0-3.fc32",
				"perl-constant-0:1.33-441.fc32",
				"device-mapper-0:1.02.171-1.fc32",
				"libgpg-error-0:1.36-3.fc32",
				"libnsl2-0:1.2.0-6.20180605git4a062cf.fc32",
				"keyutils-0:1.6-4.fc32",
				"quota-1:4.05-9.fc32",
				"perl-Text-Tabs+Wrap-0:2013.0523-440.fc32",
				"nfs-utils-1:2.5.1-4.rc4.fc32",
				"libcom_err-0:1.45.5-3.fc32",
				"dbus-common-1:1.12.20-1.fc32",
				"python3-libs-0:3.8.5-5.fc32",
				"perl-threads-1:2.22-442.fc32",
				"gdbm-libs-1:1.18.1-3.fc32",
				"expat-0:2.2.8-2.fc32",
				"perl-Unicode-Normalize-0:1.26-440.fc32",
				"libpcap-14:1.9.1-3.fc32",
				"mpfr-0:4.0.2-5.fc32",
				"iptables-libs-0:1.8.4-9.fc32",
				"kmod-0:27-1.fc32",
				"rpcbind-0:1.2.5-5.rc1.fc32.1",
				"libidn2-0:2.3.0-2.fc32",
				"systemd-libs-0:245.8-2.fc32",
				"alternatives-0:1.11-6.fc32",
				"audit-libs-0:3.0-0.19.20191104git1c2f876.fc32",
				"fedora-gpg-keys-0:32-6",
				"libgcc-0:10.2.1-1.fc32",
				"acl-0:2.2.53-5.fc32",
				"cyrus-sasl-gssapi-0:2.1.27-4.fc32",
				"libsigsegv-0:2.11-10.fc32",
				"libbasicobjects-0:0.1.1-44.fc32",
				"libcap-0:2.26-7.fc32",
				"perl-macros-4:5.30.3-456.fc32",
				"iproute-0:5.7.0-1.fc32",
				"libattr-0:2.4.48-8.fc32",
				"glib2-0:2.64.5-1.fc32",
				"gmp-1:6.1.2-13.fc32",
				"nmap-ncat-2:7.80-4.fc32",
				"tzdata-0:2020a-1.fc32",
				"ca-certificates-0:2020.2.41-1.1.fc32",
				"perl-Errno-0:1.30-456.fc32",
				"libsmartcols-0:2.35.2-1.fc32",
				"sqlite-libs-0:3.33.0-1.fc32",
				"libssh-0:0.9.5-1.fc32",
				"perl-IO-0:1.40-456.fc32",
				"libfdisk-0:2.35.2-1.fc32",
				"psmisc-0:23.3-3.fc32",
				"libsepol-0:3.0-4.fc32",
				"elfutils-default-yama-scope-0:0.181-1.fc32",
				"libvirt-libs-0:6.1.0-4.fc32",
				"setup-0:2.13.6-2.fc32",
				"systemd-rpm-macros-0:245.8-2.fc32",
			},
			nobest:   false,
			repofile: "testdata/libvirt-daemon-driver-storage-zfs-fc32.xml",
		},
		{
			name: "should resolve perl-PathTools",
			requires: []string{
				"perl-PathTools",
				"glibc-langpack-en",
				"fedora-release-container",
			},
			installs: []string{
				"perl-Carp-0:1.50-440.fc32",
				"basesystem-0:11-9.fc32",
				"fedora-release-common-0:32-3",
				"perl-interpreter-4:5.30.3-456.fc32",
				"ncurses-base-0:6.1-15.20191109.fc32",
				"glibc-langpack-en-0:2.31-4.fc32",
				"perl-PathTools-0:3.78-441.fc32",
				"libgcc-0:10.2.1-1.fc32",
				"perl-Socket-4:2.030-1.fc32",
				"perl-Scalar-List-Utils-3:1.54-440.fc32",
				"perl-macros-4:5.30.3-456.fc32",
				"perl-libs-4:5.30.3-456.fc32",
				"setup-0:2.13.6-2.fc32",
				"fedora-release-container-0:32-3",
				"perl-Exporter-0:5.74-2.fc32",
				"perl-threads-shared-0:1.60-441.fc32",
				"perl-threads-1:2.22-442.fc32",
				"filesystem-0:3.14-2.fc32",
				"perl-IO-0:1.40-456.fc32",
				"tzdata-0:2020a-1.fc32",
				"perl-Unicode-Normalize-0:1.26-440.fc32",
				"fedora-gpg-keys-0:32-6",
				"bash-0:5.0.17-1.fc32",
				"perl-File-Path-0:2.17-1.fc32",
				"fedora-repos-0:32-6",
				"perl-Text-Tabs+Wrap-0:2013.0523-440.fc32",
				"glibc-common-0:2.31-4.fc32",
				"libxcrypt-0:4.4.17-1.fc32",
				"ncurses-libs-0:6.1-15.20191109.fc32",
				"perl-parent-1:0.238-1.fc32",
				"glibc-0:2.31-4.fc32",
				"perl-constant-0:1.33-441.fc32",
				"perl-Errno-0:1.30-456.fc32",
				"gdbm-libs-1:1.18.1-3.fc32",
			},
			nobest:   false,
			repofile: "testdata/perl-pathtools-fc32.xml",
		},
		{
			name: "should resolve libvirt-devel",
			requires: []string{
				"libvirt-devel",
				"fedora-release-server",
				"glibc-langpack-en",
				"coreutils-single",
				"libcurl-minimal",
			},
			installs: []string{
				"json-c-0:0.13.1-13.fc32",
				"elfutils-libs-0:0.181-1.fc32",
				"libssh-config-0:0.9.5-1.fc32",
				"dbus-1:1.12.20-1.fc32",
				"coreutils-single-0:8.32-4.fc32.1",
				"krb5-libs-0:1.18.2-22.fc32",
				"libnghttp2-0:1.41.0-1.fc32",
				"systemd-0:245.8-2.fc32",
				"lz4-libs-0:1.9.1-2.fc32",
				"glibc-langpack-en-0:2.31-4.fc32",
				"yajl-0:2.1.0-14.fc32",
				"p11-kit-trust-0:0.23.21-2.fc32",
				"libvirt-libs-0:6.1.0-4.fc32",
				"libnetfilter_conntrack-0:1.0.7-4.fc32",
				"gzip-0:1.10-2.fc32",
				"gnutls-0:3.6.15-1.fc32",
				"cyrus-sasl-lib-0:2.1.27-4.fc32",
				"dbus-common-1:1.12.20-1.fc32",
				"ncurses-libs-0:6.1-15.20191109.fc32",
				"readline-0:8.0-4.fc32",
				"zlib-0:1.2.11-21.fc32",
				"libcom_err-0:1.45.5-3.fc32",
				"mpfr-0:4.0.2-5.fc32",
				"keyutils-libs-0:1.6-4.fc32",
				"systemd-rpm-macros-0:245.8-2.fc32",
				"glibc-common-0:2.31-4.fc32",
				"libpwquality-0:1.4.2-2.fc32",
				"libcap-ng-0:0.7.11-1.fc32",
				"libargon2-0:20171227-4.fc32",
				"alternatives-0:1.11-6.fc32",
				"fedora-release-server-0:32-3",
				"basesystem-0:11-9.fc32",
				"cyrus-sasl-0:2.1.27-4.fc32",
				"libnfnetlink-0:1.0.1-17.fc32",
				"cracklib-0:2.9.6-22.fc32",
				"libgcc-0:10.2.1-1.fc32",
				"glibc-0:2.31-4.fc32",
				"kmod-libs-0:27-1.fc32",
				"ncurses-base-0:6.1-15.20191109.fc32",
				"libtasn1-0:4.16.0-1.fc32",
				"cryptsetup-libs-0:2.3.4-1.fc32",
				"fedora-release-common-0:32-3",
				"libseccomp-0:2.5.0-3.fc32",
				"pkgconf-pkg-config-0:1.6.3-3.fc32",
				"openldap-0:2.4.47-5.fc32",
				"device-mapper-libs-0:1.02.171-1.fc32",
				"libssh2-0:1.9.0-5.fc32",
				"openssl-libs-1:1.1.1g-1.fc32",
				"libnl3-0:3.5.0-2.fc32",
				"bzip2-libs-0:1.0.8-2.fc32",
				"systemd-pam-0:245.8-2.fc32",
				"libxml2-0:2.9.10-7.fc32",
				"libidn2-0:2.3.0-2.fc32",
				"nettle-0:3.5.1-5.fc32",
				"device-mapper-0:1.02.171-1.fc32",
				"ca-certificates-0:2020.2.41-1.1.fc32",
				"qrencode-libs-0:4.0.2-5.fc32",
				"acl-0:2.2.53-5.fc32",
				"cyrus-sasl-gssapi-0:2.1.27-4.fc32",
				"libvirt-devel-0:6.1.0-4.fc32",
				"libffi-0:3.1-24.fc32",
				"libsepol-0:3.0-4.fc32",
				"libcap-0:2.26-7.fc32",
				"libsmartcols-0:2.35.2-1.fc32",
				"libmount-0:2.35.2-1.fc32",
				"libtirpc-0:1.2.6-1.rc4.fc32",
				"libverto-0:0.3.0-9.fc32",
				"fedora-repos-0:32-6",
				"gawk-0:5.0.1-7.fc32",
				"elfutils-libelf-0:0.181-1.fc32",
				"libssh-0:0.9.5-1.fc32",
				"libdb-0:5.3.28-40.fc32",
				"systemd-libs-0:245.8-2.fc32",
				"pcre2-syntax-0:10.35-6.fc32",
				"xz-libs-0:5.2.5-1.fc32",
				"pcre-0:8.44-1.fc32",
				"tzdata-0:2020a-1.fc32",
				"libacl-0:2.2.53-5.fc32",
				"libpcap-14:1.9.1-3.fc32",
				"pam-0:1.3.1-26.fc32",
				"pcre2-0:10.35-6.fc32",
				"libgcrypt-0:1.8.5-3.fc32",
				"bash-0:5.0.17-1.fc32",
				"setup-0:2.13.6-2.fc32",
				"libattr-0:2.4.48-8.fc32",
				"iptables-libs-0:1.8.4-9.fc32",
				"filesystem-0:3.14-2.fc32",
				"libunistring-0:0.9.10-7.fc32",
				"libselinux-0:3.0-5.fc32",
				"libmnl-0:1.0.4-11.fc32",
				"shadow-utils-2:4.8.1-2.fc32",
				"glib2-0:2.64.5-1.fc32",
				"libfdisk-0:2.35.2-1.fc32",
				"crypto-policies-0:20200619-1.git781bbd4.fc32",
				"libwsman1-0:2.6.8-12.fc32",
				"libsigsegv-0:2.11-10.fc32",
				"libuuid-0:2.35.2-1.fc32",
				"libcurl-minimal-0:7.69.1-6.fc32",
				"p11-kit-0:0.23.21-2.fc32",
				"util-linux-0:2.35.2-1.fc32",
				"grep-0:3.3-4.fc32",
				"dbus-broker-0:24-1.fc32",
				"fedora-gpg-keys-0:32-6",
				"libgpg-error-0:1.36-3.fc32",
				"dbus-libs-1:1.12.20-1.fc32",
				"numactl-libs-0:2.0.12-4.fc32",
				"elfutils-default-yama-scope-0:0.181-1.fc32",
				"audit-libs-0:3.0-0.19.20191104git1c2f876.fc32",
				"gmp-1:6.1.2-13.fc32",
				"libsemanage-0:3.0-3.fc32",
				"libxcrypt-0:4.4.17-1.fc32",
				"pkgconf-m4-0:1.6.3-3.fc32",
				"pkgconf-0:1.6.3-3.fc32",
				"libpkgconf-0:1.6.3-3.fc32",
				"expat-0:2.2.8-2.fc32",
				"libutempter-0:1.1.6-18.fc32",
				"libnsl2-0:1.2.0-6.20180605git4a062cf.fc32",
				"libblkid-0:2.35.2-1.fc32",
				"sed-0:4.5-5.fc32",
			},
			nobest:   false,
			repofile: "testdata/libvirt-devel-fc32.xml",
		},
		{name: "should resolve pkgconf-pkg-config",
			requires: []string{
				"pkgconf-pkg-config",
				"fedora-release-server",
			},
			installs: []string{
				"glibc-langpack-en-0:2.31-4.fc32",
				"glibc-common-0:2.31-4.fc32",
				"fedora-release-common-0:32-3",
				"fedora-release-server-0:32-3",
				"filesystem-0:3.14-2.fc32",
				"ncurses-libs-0:6.1-15.20191109.fc32",
				"glibc-0:2.31-4.fc32",
				"tzdata-0:2020a-1.fc32",
				"ncurses-base-0:6.1-15.20191109.fc32",
				"libgcc-0:10.2.1-1.fc32",
				"setup-0:2.13.6-2.fc32",
				"pkgconf-0:1.6.3-3.fc32",
				"pkgconf-m4-0:1.6.3-3.fc32",
				"libpkgconf-0:1.6.3-3.fc32",
				"bash-0:5.0.17-1.fc32",
				"fedora-repos-0:32-6",
				"fedora-gpg-keys-0:32-6",
				"pkgconf-pkg-config-0:1.6.3-3.fc32",
				"basesystem-0:11-9.fc32",
			},
			nobest:   false,
			repofile: "testdata/pkgconf-pkg-config-fc32.xml",
		},
		{name: "should resolve libvirt-daemon",
			requires: []string{
				"libvirt-daemon",
				"fedora-release-server",
				"glibc-langpack-en",
				"coreutils-single",
				"libcurl-minimal",
			},
			installs: []string{
				"keyutils-libs-0:1.6-4.fc32",
				"filesystem-0:3.14-2.fc32",
				"iproute-tc-0:5.7.0-1.fc32",
				"zlib-0:1.2.11-21.fc32",
				"bash-0:5.0.17-1.fc32",
				"libuuid-0:2.35.2-1.fc32",
				"qrencode-libs-0:4.0.2-5.fc32",
				"libcom_err-0:1.45.5-3.fc32",
				"lz4-libs-0:1.9.1-2.fc32",
				"alternatives-0:1.11-6.fc32",
				"openssl-libs-1:1.1.1g-1.fc32",
				"expat-0:2.2.8-2.fc32",
				"libssh-config-0:0.9.5-1.fc32",
				"libgpg-error-0:1.36-3.fc32",
				"libattr-0:2.4.48-8.fc32",
				"ncurses-base-0:6.1-15.20191109.fc32",
				"libselinux-0:3.0-5.fc32",
				"libunistring-0:0.9.10-7.fc32",
				"libnl3-0:3.5.0-2.fc32",
				"libpcap-14:1.9.1-3.fc32",
				"gnutls-0:3.6.15-1.fc32",
				"libxml2-0:2.9.10-7.fc32",
				"setup-0:2.13.6-2.fc32",
				"libffi-0:3.1-24.fc32",
				"libmount-0:2.35.2-1.fc32",
				"libvirt-libs-0:6.1.0-4.fc32",
				"dbus-libs-1:1.12.20-1.fc32",
				"ca-certificates-0:2020.2.41-1.1.fc32",
				"psmisc-0:23.3-3.fc32",
				"xz-libs-0:5.2.5-1.fc32",
				"iproute-0:5.7.0-1.fc32",
				"openldap-0:2.4.47-5.fc32",
				"libidn2-0:2.3.0-2.fc32",
				"acl-0:2.2.53-5.fc32",
				"libacl-0:2.2.53-5.fc32",
				"libgcc-0:10.2.1-1.fc32",
				"libmnl-0:1.0.4-11.fc32",
				"libtirpc-0:1.2.6-1.rc4.fc32",
				"libargon2-0:20171227-4.fc32",
				"libcap-ng-0:0.7.11-1.fc32",
				"device-mapper-0:1.02.171-1.fc32",
				"mpfr-0:4.0.2-5.fc32",
				"libcurl-minimal-0:7.69.1-6.fc32",
				"linux-atm-libs-0:2.5.1-26.fc32",
				"systemd-pam-0:245.8-2.fc32",
				"sed-0:4.5-5.fc32",
				"systemd-libs-0:245.8-2.fc32",
				"gzip-0:1.10-2.fc32",
				"krb5-libs-0:1.18.2-22.fc32",
				"crypto-policies-0:20200619-1.git781bbd4.fc32",
				"dbus-common-1:1.12.20-1.fc32",
				"nettle-0:3.5.1-5.fc32",
				"numactl-libs-0:2.0.12-4.fc32",
				"libtasn1-0:4.16.0-1.fc32",
				"libsigsegv-0:2.11-10.fc32",
				"libsepol-0:3.0-4.fc32",
				"polkit-0:0.116-7.fc32",
				"json-c-0:0.13.1-13.fc32",
				"glibc-common-0:2.31-4.fc32",
				"dbus-broker-0:24-1.fc32",
				"audit-libs-0:3.0-0.19.20191104git1c2f876.fc32",
				"nmap-ncat-2:7.80-4.fc32",
				"tzdata-0:2020a-1.fc32",
				"device-mapper-libs-0:1.02.171-1.fc32",
				"shadow-utils-2:4.8.1-2.fc32",
				"fedora-repos-0:32-6",
				"libfdisk-0:2.35.2-1.fc32",
				"pcre2-0:10.35-6.fc32",
				"fedora-release-common-0:32-3",
				"libsmartcols-0:2.35.2-1.fc32",
				"dmidecode-1:3.2-5.fc32",
				"pcre2-syntax-0:10.35-6.fc32",
				"bzip2-libs-0:1.0.8-2.fc32",
				"libstdc++-0:10.2.1-1.fc32",
				"libnghttp2-0:1.41.0-1.fc32",
				"libxcrypt-0:4.4.17-1.fc32",
				"libdb-0:5.3.28-40.fc32",
				"glib2-0:2.64.5-1.fc32",
				"elfutils-default-yama-scope-0:0.181-1.fc32",
				"systemd-rpm-macros-0:245.8-2.fc32",
				"libnsl2-0:1.2.0-6.20180605git4a062cf.fc32",
				"glibc-0:2.31-4.fc32",
				"libsemanage-0:3.0-3.fc32",
				"libgcrypt-0:1.8.5-3.fc32",
				"systemd-0:245.8-2.fc32",
				"libvirt-daemon-0:6.1.0-4.fc32",
				"polkit-libs-0:0.116-7.fc32",
				"p11-kit-trust-0:0.23.21-2.fc32",
				"fedora-gpg-keys-0:32-6",
				"iptables-libs-0:1.8.4-9.fc32",
				"kmod-libs-0:27-1.fc32",
				"libutempter-0:1.1.6-18.fc32",
				"libcap-0:2.26-7.fc32",
				"ncurses-libs-0:6.1-15.20191109.fc32",
				"gmp-1:6.1.2-13.fc32",
				"cyrus-sasl-0:2.1.27-4.fc32",
				"libverto-0:0.3.0-9.fc32",
				"readline-0:8.0-4.fc32",
				"polkit-pkla-compat-0:0.1-16.fc32",
				"libnetfilter_conntrack-0:1.0.7-4.fc32",
				"coreutils-single-0:8.32-4.fc32.1",
				"libssh-0:0.9.5-1.fc32",
				"kmod-0:27-1.fc32",
				"util-linux-0:2.35.2-1.fc32",
				"libseccomp-0:2.5.0-3.fc32",
				"grep-0:3.3-4.fc32",
				"glibc-langpack-en-0:2.31-4.fc32",
				"p11-kit-0:0.23.21-2.fc32",
				"libblkid-0:2.35.2-1.fc32",
				"libwsman1-0:2.6.8-12.fc32",
				"cryptsetup-libs-0:2.3.4-1.fc32",
				"mozjs60-0:60.9.0-5.fc32",
				"elfutils-libelf-0:0.181-1.fc32",
				"libpwquality-0:1.4.2-2.fc32",
				"fedora-release-server-0:32-3",
				"cyrus-sasl-gssapi-0:2.1.27-4.fc32",
				"gawk-0:5.0.1-7.fc32",
				"basesystem-0:11-9.fc32",
				"numad-0:0.5-31.20150602git.fc32",
				"libssh2-0:1.9.0-5.fc32",
				"elfutils-libs-0:0.181-1.fc32",
				"dbus-1:1.12.20-1.fc32",
				"yajl-0:2.1.0-14.fc32",
				"pam-0:1.3.1-26.fc32",
				"cyrus-sasl-lib-0:2.1.27-4.fc32",
				"libnfnetlink-0:1.0.1-17.fc32",
				"pcre-0:8.44-1.fc32",
				"cracklib-0:2.9.6-22.fc32",
			},
			repofile: "testdata/libvirt-daemon-fc32.xml",
		},
	}

	focus := false
	for _, tt := range tests {
		if tt.focus == true {
			focus = true
			break
		}
	}
	for _, tt := range tests {
		if focus != tt.focus {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			f, err := os.Open(tt.repofile)
			g.Expect(err).ToNot(HaveOccurred())
			defer f.Close()
			repo := &api.Repository{}
			err = xml.NewDecoder(f).Decode(repo)
			g.Expect(err).ToNot(HaveOccurred())

			resolver := NewResolver(tt.nobest)
			packages := []*api.Package{}
			for i, _ := range repo.Packages {
				packages = append(packages, &repo.Packages[i])
			}
			err = resolver.LoadInvolvedPackages(packages)
			g.Expect(err).ToNot(HaveOccurred())
			err = resolver.ConstructRequirements(tt.requires)
			g.Expect(err).ToNot(HaveOccurred())
			install, _, err := resolver.Resolve()
			g.Expect(pkgToString(install)).To(ConsistOf(tt.installs))
			g.Expect(err).ToNot(HaveOccurred())
		})
	}
}

func TestNewResolver(t *testing.T) {
	tests := []struct {
		name     string
		packages []*api.Package
		requires []string
		install  []string
		exclude  []string
		solvable bool
		focus    bool
	}{
		{name: "with indirect dependency", packages: []*api.Package{
			newPkg("testa", "1", []string{"testa", "a", "b"}, []string{"d", "g"}, []string{}),
			newPkg("testb", "1", []string{"testb", "c"}, []string{}, []string{}),
			newPkg("testc", "1", []string{"testc", "d"}, []string{}, []string{}),
			newPkg("testd", "1", []string{"testd", "e", "f", "g"}, []string{"h"}, []string{}),
			newPkg("teste", "1", []string{"teste", "h"}, []string{}, []string{}),
		}, requires: []string{
			"testa",
		},
			install:  []string{"testa-0:1", "testc-0:1", "testd-0:1", "teste-0:1"},
			exclude:  []string{"testb-0:1"},
			solvable: true,
		},
		{name: "with circular dependency", packages: []*api.Package{
			newPkg("testa", "1", []string{"testa", "a", "b"}, []string{"d", "g"}, []string{}),
			newPkg("testb", "1", []string{"testb", "c"}, []string{}, []string{}),
			newPkg("testc", "1", []string{"testc", "d"}, []string{}, []string{}),
			newPkg("testd", "1", []string{"testd", "e", "f", "g"}, []string{"h"}, []string{}),
			newPkg("teste", "1", []string{"teste", "h"}, []string{"a"}, []string{}),
		}, requires: []string{
			"testa",
		},
			install:  []string{"testa-0:1", "testc-0:1", "testd-0:1", "teste-0:1"},
			exclude:  []string{"testb-0:1"},
			solvable: true,
		},
		{name: "with an unresolvable dependency", packages: []*api.Package{
			newPkg("testa", "1", []string{"testa", "a", "b"}, []string{"d"}, []string{}),
		}, requires: []string{
			"testa",
		},
			solvable: false,
		},
		{name: "with two sources to choose from, should use the newer one", packages: []*api.Package{
			newPkg("testa", "1", []string{"testa", "a", "b"}, []string{"d"}, []string{"testa", "a"}),
			newPkg("testb", "1", []string{"testb", "d"}, []string{}, []string{}),
			newPkg("testb", "2", []string{"testb", "d"}, []string{}, []string{}),
		}, requires: []string{
			"testa",
		},
			install:  []string{"testa-0:1", "testb-0:2"},
			exclude:  []string{},
			solvable: true,
		},
		{name: "with one source only referencing itself", packages: []*api.Package{
			newPkg("testa", "1", []string{"testa"}, []string{}, []string{}),
		}, requires: []string{
			"testa",
		},
			install:  []string{"testa-0:1"},
			exclude:  []string{},
			solvable: true,
		},
		// TODO: Add test cases.
	}
	focus := false
	for _, tt := range tests {
		if tt.focus == true {
			focus = true
			break
		}
	}
	for _, tt := range tests {
		if focus != tt.focus {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			resolver := NewResolver(false)
			err := resolver.LoadInvolvedPackages(tt.packages)
			if err != nil {
				t.Fail()
			}
			err = resolver.ConstructRequirements(tt.requires)
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
			install, exclude, err := resolver.Resolve()
			g := NewGomegaWithT(t)
			if tt.solvable {
				g.Expect(err).ToNot(HaveOccurred())
			} else {
				g.Expect(err).To(HaveOccurred())
				return
			}
			g.Expect(pkgToString(install)).To(ConsistOf(tt.install))
			g.Expect(pkgToString(exclude)).To(ConsistOf(tt.exclude))
		})
	}
}

func newPkg(name string, version string, provides []string, requires []string, conflicts []string) *api.Package {
	pkg := &api.Package{}
	pkg.Name = name
	pkg.Version = api.Version{Ver: version}
	for _, req := range requires {
		pkg.Format.Requires.Entries = append(pkg.Format.Requires.Entries, api.Entry{Name: req})
	}
	pkg.Format.Provides.Entries = append(pkg.Format.Provides.Entries, api.Entry{
		Name:  name,
		Flags: "EQ",
		Ver:   version,
	})
	for _, req := range provides {
		pkg.Format.Provides.Entries = append(pkg.Format.Provides.Entries, api.Entry{Name: req})
	}
	for _, req := range conflicts {
		pkg.Format.Conflicts.Entries = append(pkg.Format.Conflicts.Entries, api.Entry{Name: req})
	}

	return pkg
}

func strToPkg(wanted []string, given []*api.Package) (resolved []*api.Package) {
	m := map[string]*api.Package{}
	for _, p := range given {
		m[p.String()] = p
	}
	for _, w := range wanted {
		resolved = append(resolved, m[w])
	}
	return resolved
}

func pkgToString(given []*api.Package) (resolved []string) {
	for _, p := range given {
		resolved = append(resolved, p.String())
	}
	return
}
