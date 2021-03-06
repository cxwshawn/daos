#!/usr/bin/python
"""
  (C) Copyright 2018-2019 Intel Corporation.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

  GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
  The Government's rights to use, modify, reproduce, release, perform, display,
  or disclose this software are subject to the terms of the Apache License as
  provided in Contract No. B609815.
  Any reproduction of computer software, computer software documentation, or
  portions thereof marked with this legend must also reproduce the markings.
"""

import os
import sys
import json
import subprocess

from avocado       import Test
from avocado.utils import process

sys.path.append('./util')
sys.path.append('../util')
sys.path.append('../../../utils/py')
sys.path.append('./../../utils/py')
# pylint: disable=wrong-import-position,import-error
import AgentUtils
import server_utils
import write_host_file
# pylint: enable=wrong-import-position,import-error

CLIENT_LOG = "client_daos.log"

class DaosCoreTest(Test):
    """
    Runs the daos_test subtests with multiple servers.

    """
    # pylint: disable=too-many-instance-attributes
    subtest_name = None

    def setUp(self):
        self.agent_sessions = None
        self.hostlist = None
        self.subtest_name = self.params.get("test_name",
                                            '/run/daos_tests/Tests/*')
        self.subtest_name = self.subtest_name.replace(" ", "_")
        # get paths from the build_vars generated by build
        with open('../../../.build_vars.json') as build_vars:
            build_paths = json.load(build_vars)
        self.basepath = os.path.normpath(build_paths['PREFIX'] + "/../")

        self.server_group = self.params.get("name", '/server_config/',
                                            'daos_server')
        self.daos_test = self.basepath + '/install/bin/daos_test'
        self.orterun = self.basepath + '/install/bin/orterun'
        self.hostlist = self.params.get("test_machines", '/run/hosts/*')

        self.hostfile = write_host_file.write_host_file(self.hostlist,
                                                        self.workdir)
        logfile_env = os.environ['D_LOG_FILE']
        self.log_dir, self.server_log = os.path.split(logfile_env)
        self.client_log = os.path.join(self.log_dir,
                                       self.subtest_name + "_" + CLIENT_LOG)
        #To generate the seperate client log file
        self.orterun_env = '-x D_LOG_FILE={}'.format(self.client_log)

        self.agent_sessions = AgentUtils.run_agent(self.basepath, self.hostlist)
        server_utils.run_server(self.hostfile, self.server_group, self.basepath)

    def tearDown(self):
        if self.agent_sessions:
            AgentUtils.stop_agent(self.hostlist, self.agent_sessions)
        server_utils.stop_server(hosts=self.hostlist)

        # collect up a debug log so that we have a separate one for each
        # subtest
        if self.subtest_name:
            try:
                new_logfile = os.path.join(self.log_dir,
                                           self.subtest_name + "_" + \
                                           self.server_log)
                # rename on each of the servers
                for host in self.hostlist:
                    subprocess.check_call(['ssh', host,
                                           '[ -f \"{0}\" ] && '
                                           ' mv \"{0}\" '
                                           ' \"{1}\"'.format("/tmp/server.log",
                                                             new_logfile)])
            except KeyError:
                pass

    def test_subtest(self):
        """
        Test ID: DAOS-1568

        Test Description: Run daos_test with a subtest argument

        Use Cases: core tests for daos_test

        :avocado: tags=daos_test,multiserver,vm,regression
        """

        subtest = self.params.get("daos_test", '/run/daos_tests/Tests/*')
        num_clients = self.params.get("num_clients",
                                      '/run/daos_tests/num_clients/*')
        num_replicas = self.params.get("num_replicas",
                                       '/run/daos_tests/num_replicas/*')
        args = self.params.get("args", '/run/daos_tests/Tests/*')
        if not args:
            args = ""

        cmd = "{} -n {} {} {} -s {} -{} {}".format(self.orterun, num_clients,
                                                   self.orterun_env,
                                                   self.daos_test,
                                                   num_replicas, subtest, args)

        env = {}
        env['CMOCKA_XML_FILE'] = "%g_results.xml"
        env['CMOCKA_MESSAGE_OUTPUT'] = "xml"

        try:
            process.run(cmd, env=env)
        except process.CmdError as result:
            if result.result.exit_status is not 0:
                # fake a JUnit failure output
                with open(self.subtest_name +
                          "_results.xml", "w") as results_xml:
                    results_xml.write('''<?xml version="1.0" encoding="UTF-8"?>
<testsuite name="{0}" errors="1" failures="0" skipped="0" tests="1" time="0.0">
  <testcase name="ALL" time="0.0" >
    <error message="Test failed to start up"/>
    <system-out>
<![CDATA[{1}]]>
    </system-out>
    <system-err>
<![CDATA[{2}]]>
    </system-err>
  </testcase>
</testsuite>'''.format(self.subtest_name, result.result.stdout,
                       result.result.stderr))
                self.fail("{0} failed with return code={1}.\n"
                          .format(cmd, result.result.exit_status))
