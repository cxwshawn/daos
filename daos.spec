# Needed because of the GO binaries
%undefine _missing_build_ids_terminate_build

%define daoshome %{_exec_prefix}/lib/%{name}

Name:          daos
Version:       0.4.0
Release:       1%{?dist}
Summary:       DAOS Storage Engine

License:       Apache
URL:           https//github.com/daos-stack/daos
Source0:       %{name}-%{version}.tar.gz
Source1:       scons_local-%{version}.tar.gz

BuildRequires: scons
BuildRequires: cart-devel
BuildRequires: argobots-devel
BuildRequires: libpmem-devel, libpmemobj-devel
BuildRequires: fuse-devel >= 3.4.2
BuildRequires: protobuf-c-devel
BuildRequires: spdk-devel, spdk-tools
BuildRequires: fio-devel, fio
BuildRequires: isa-l-devel
BuildRequires: raft-devel
BuildRequires: mercury-devel
BuildRequires: openpa-devel
BuildRequires: libfabric-devel
BuildRequires: openssl-devel
BuildRequires: ompi-devel
BuildRequires: pmix-devel
BuildRequires: hwloc-devel
BuildRequires: libevent-devel
BuildRequires: libyaml-devel
BuildRequires: libcmocka-devel
BuildRequires: readline-devel
BuildRequires: numactl-devel
BuildRequires: golang-bin
BuildRequires: libipmctl-devel
BuildRequires: CUnit-devel
Requires: cart
Requires: argobots
Requires: libpmem, libpmemobj
Requires: fuse >= 3.4.2
Requires: protobuf-c
Requires: spdk
Requires: fio
Requires: isa-l

%description
The Distributed Asynchronous Object Storage (DAOS) is an open-source
software-defined object store designed from the ground up for
massively distributed Non Volatile Memory (NVM). DAOS takes advantage
of next generation NVM technology like Storage Class Memory (SCM) and
NVM express (NVMe) while presenting a key-value storage interface and
providing features such as transactional non-blocking I/O, advanced
data protection with self healing on top of commodity hardware, end-
to-end data integrity, fine grained data control and elastic storage
to optimize performance and cost.

%package server
Summary: The DAOS server
Requires: %{name} = %{version}-%{release}

%description server
This is the package needed to run a DAOS server

%package client
Summary: The DAOS client
Requires: %{name} = %{version}-%{release}

%description client
This is the package needed to run a DAOS client

%package tests
Summary: The DAOS test suite
Requires: %{name}-client = %{version}-%{release}

%description tests
This is the package needed to run the DAOS test suite

%package devel
Summary: The DAOS development libraries and headers

%description devel
This is the package needed to build software with the DAOS library.

%prep
%setup -q
%setup -q -a 1


%build
# remove rpathing from the build
rpath_files="utils/daos_build.py"
rpath_files+=" $(find . -name SConscript)"
sed -i -e '/AppendUnique(RPATH=.*)/d' $rpath_files

scons %{?no_smp_mflags}    \
      --config=force       \
      USE_INSTALLED=all    \
      PREFIX=%{?buildroot}

%install
scons %{?no_smp_mflags}              \
      --config=force                 \
      install                        \
      USE_INSTALLED=all              \
      PREFIX=%{?buildroot}%{_prefix}
BUILDROOT="%{?buildroot}"
PREFIX="%{?_prefix}"
sed -i -e s/${BUILDROOT//\//\\/}[^\"]\*/${PREFIX//\//\\/}/g %{?buildroot}%{_prefix}/TESTING/.build_vars.*
mv %{?buildroot}%{_prefix}/lib{,64}
#mv %{?buildroot}/{usr/,}etc
mkdir -p %{?buildroot}/%{daoshome}
mv %{?buildroot}%{_prefix}/{TESTING,lib/%{name}/}
cp -al ftest.sh src/tests/ftest %{?buildroot}%{daoshome}/TESTING
find %{?buildroot}%{daoshome}/TESTING/ftest -name \*.py[co] -print0 | xargs -r0 rm -f
#ln %{?buildroot}%{daoshome}/{TESTING/.build_vars,.build_vars-Linux}.sh

%files
%defattr(-, root, root, -)
%{_bindir}/daos_gen_io_conf
%{_bindir}/daos_run_io_conf
%{_bindir}/daos_shell
%{_bindir}/daosctl
%{_bindir}/dcont
%{_bindir}/duns
%{_bindir}/hello_drpc
%{_bindir}/io_conf
%{_bindir}/obj_ctl
%{_bindir}/pl_map
%{_bindir}/rdbt
%{_bindir}/vos_size
%{_bindir}/vos_size.py
%{_libdir}/libvos.so
%{_prefix}%{_sysconfdir}/daos.yml
%{_prefix}%{_sysconfdir}/vos_dfs_sample.yaml
%{_prefix}%{_sysconfdir}/vos_size_input.yaml
%{_libdir}/libdaos_common.so
%{_libdir}/daos_srv/libplacement.so
%doc

%files server
%{_prefix}%{_sysconfdir}/daos_server.yml
%{_bindir}/daos_server
%{_bindir}/daos_io_server
%{_libdir}/daos_srv/libbio.so
%{_libdir}/daos_srv/libcont.so
%{_libdir}/daos_srv/libdtx.so
%{_libdir}/daos_srv/libmgmt.so
%{_libdir}/daos_srv/libobj.so
%{_libdir}/daos_srv/libpool.so
%{_libdir}/daos_srv/librdb.so
%{_libdir}/daos_srv/librdbt.so
%{_libdir}/daos_srv/librebuild.so
%{_libdir}/daos_srv/librsvc.so
%{_libdir}/daos_srv/libsecurity.so
%{_libdir}/daos_srv/libvos_srv.so
%{_datadir}/%{name}

%files client
%{_bindir}/daos_agent
%{_bindir}/dfuse
%{_bindir}/dmg
%{_bindir}/dfuse_hl
%{_libdir}/*.so.*
%{_libdir}/libdfs.so
%{_libdir}/libduns.so
%{_libdir}/libdfuse.so
%{_libdir}/libioil.so
%{_datadir}/%{name}/ioil-ld-opts

%files tests
%{daoshome}/TESTING
%{_bindir}/*_test*
%{_bindir}/io_conf/daos_io_conf_1
%{_bindir}/io_conf/daos_io_conf_2
%{_bindir}/smd_ut
%{_bindir}/vea_ut
%{_bindir}/daosbench
%{_bindir}/daos_perf
%{_bindir}/evt_ctl
%{_libdir}/libdaos_tests.so

%files devel
%{_includedir}/*
%{_libdir}/libdaos.so
%{_libdir}/*.a

%changelog
* Wed Apr  3 2019 Brian J. Murrell <brian.murrell@intel.com>
- initial package
