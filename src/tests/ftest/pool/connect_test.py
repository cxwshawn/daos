#!/usr/bin/python
'''
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
'''
from __future__ import print_function

import os
import traceback
import sys
import json

from avocado import Test
from avocado.utils import process

sys.path.append('./util')
import AgentUtils
import server_utils
import check_for_pool
import write_host_file

class ConnectTest(Test):
    """
    Tests DAOS pool creation, calling it repeatedly one after another

    :avocado: tags=pool,poolconnect
    """

    # super wasteful since its doing this for every variation
    def setUp(self):

        # get paths from the build_vars generated by build
        with open('../../../.build_vars.json') as build_file:
            build_paths = json.load(build_file)
        self.basepath = os.path.normpath(build_paths['PREFIX'] + "/../")

        self.hostlist = self.params.get("test_machines", '/run/hosts/')
        self.hostfile = write_host_file.write_host_file(self.hostlist,
                                                        self.workdir)

        server_group = self.params.get("server_group", '/server/',
                                       'daos_server')

        self.agent_sessions = AgentUtils.run_agent(self.basepath, self.hostlist)
        server_utils.run_server(self.hostfile, server_group, self.basepath)

    def tearDown(self):
        if self.agent_sessions:
            AgentUtils.stop_agent(self.hostlist, self.agent_sessions)
        server_utils.stop_server(hosts=self.hostlist)

    def test_connect(self):
        """
        Test connecting to a pool.

        :avocado: tags=pool,poolconnect,quick
        """

        # Accumulate a list of pass/fail indicators representing what is
        # expected for each parameter then "and" them to determine the
        # expected result of the test
        expected_for_param = []

        setidlist = self.params.get("setname", '/run/tests/setnames/*')
        setid = setidlist[0]
        expected_for_param.append(setidlist[1])

        # if any parameter results in failure then the test should FAIL
        expected_result = 'PASS'
        for result in expected_for_param:
            if result == 'FAIL':
                expected_result = 'FAIL'
                break
        try:
            uid = os.geteuid()
            gid = os.getegid()

            # TODO make these params in the yaml
            daosctl = self.basepath + '/install/bin/daosctl'

            host1 = self.hostlist[0]
            host2 = self.hostlist[1]

            create_cmd = (
                "{0} create-pool "
                "-m {1} "
                "-u {2} "
                "-g {3} "
                "-s {4} "
                "-c 1".format(daosctl, "0731", uid, gid, setid))
            uuid_str = """{0}""".format(process.system_output(create_cmd))
            print("uuid is {0}\n".format(uuid_str))

            exists = check_for_pool.check_for_pool(host1, uuid_str)
            if exists != 0:
                self.fail("Pool {0} not found on host {1}.\n".
                          format(uuid_str, host1))
            exists = check_for_pool.check_for_pool(host2, uuid_str)
            if exists != 0:
                self.fail("Pool {0} not found on host {1}.\n".
                          format(uuid_str, host2))

            connect_cmd = ('{0} connect-pool -i {1} '
                           '-s {2} -r -l 0,1'.format(daosctl,
                                                     uuid_str, setid))
            process.system(connect_cmd)


            if expected_result == 'FAIL':
                self.fail("Expected to fail but passed.\n")

        except Exception as excep:
            print(excep)
            print(traceback.format_exc())
            if expected_result == 'PASS':
                self.fail("Expecting to pass but test has failed.\n")
