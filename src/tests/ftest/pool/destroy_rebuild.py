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
import time
import traceback
import sys
import json
from avocado import Test

sys.path.append('./util')
sys.path.append('../util')
sys.path.append('./../../utils/py')
sys.path.append('../../../utils/py')

import AgentUtils
import server_utils
import write_host_file
from daos_api import DaosContext, DaosPool, DaosServer, DaosApiError

class DestroyRebuild(Test):

    """
    Test Class Description:
    This test verifies destruction of a pool that is rebuilding.

    :avocado: tags=pool,pooldestroy,rebuild,desreb
    """

    build_paths = []
    server_group = ""
    context = None
    pool = None
    hostfile = ""

    def setUp(self):
        """ setup for the test """
        self.agent_sessions = None
        # get paths from the build_vars generated by build
        with open('../../../.build_vars.json') as build_file:
            build_paths = json.load(build_file)
        self.context = DaosContext(build_paths['PREFIX'] + '/lib/')
        self.basepath = os.path.normpath(build_paths['PREFIX'] + "/../")

        # generate a hostfile
        self.hostlist = self.params.get("test_machines", '/run/hosts/')
        self.hostfile = write_host_file.write_host_file(self.hostlist,
                                                        self.workdir)

        # fire up the DAOS servers
        self.server_group = self.params.get("name", '/run/server_config/',
                                            'daos_server')
        self.agent_sessions = AgentUtils.run_agent(self.basepath, self.hostlist)
        server_utils.run_server(self.hostfile, self.server_group,
                                build_paths['PREFIX'] + '/../')

        # create a pool to test with
        createmode = self.params.get("mode", '/run/pool/createmode/')
        createuid = self.params.get("uid", '/run/pool/createuid/')
        creategid = self.params.get("gid", '/run/pool/creategid/')
        createsetid = self.params.get("setname", '/run/pool/createset/')
        createsize = self.params.get("size", '/run/pool/createsize/')
        self.pool = DaosPool(self.context)
        self.pool.create(createmode, createuid, creategid, createsize,
                         createsetid)
        self.pool.get_uuid_str()

        time.sleep(2)

    def tearDown(self):
        """ cleanup after the test """

        try:
            os.remove(self.hostfile)
            if self.pool:
                self.pool.destroy(1)
        finally:
            if self.agent_sessions:
                AgentUtils.stop_agent(self.hostlist, self.agent_sessions)
            server_utils.stop_server(hosts=self.hostlist)


    def test_destroy_while_rebuilding(self):
        """
        :avocado: tags=pool,pooldestroy,rebuild,desreb
        """
        try:
            print("\nsetup complete, starting test\n")

            # create a server object that references on of our pool target hosts
            # and then kill it
            svr_to_kill = int(self.params.get("rank_to_kill",
                                              '/run/testparams/ranks/'))
            server = DaosServer(self.context, bytes(self.server_group),
                                svr_to_kill)

            print("created server ")

            # BUG if you don't connect the rebuild doesn't start correctly
            self.pool.connect(1 << 1)
            status = self.pool.pool_query()
            if not status.pi_ntargets == len(self.hostlist):
                self.fail("target count wrong.\n")
            if not status.pi_ndisabled == 0:
                self.fail("disabled target count wrong.\n")

            print("connect ")

            time.sleep(1)
            server.kill(1)

            print("killed server ")

            # exclude the target from the dead server
            self.pool.exclude([svr_to_kill])

            print("exclude target ")

            #self.pool.disconnect()
            #print "disconnect "

            # the rebuild won't take long since there is no data so do
            # the destroy quickly
            self.pool.destroy(1)
            print("destroy ")

        except DaosApiError as excep:
            print(excep)
            print(traceback.format_exc())
            self.fail("Expecting to pass but test has failed.\n")
