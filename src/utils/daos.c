/**
 * (C) Copyright 2016-2019 Intel Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
 * The Government's rights to use, modify, reproduce, release, perform, display,
 * or disclose this software are subject to the terms of the Apache License as
 * provided in Contract No. B609815.
 * Any reproduction of computer software, computer software documentation, or
 * portions thereof marked with this legend must also reproduce the markings.
 */
/**
 * daos(8): DAOS Container and Object Management Utility
 */
#include <getopt.h>
#include <stdio.h>
#include <daos.h>
#include <daos/common.h>

const char		*default_group;
typedef int (*command_hdlr_t)(int, char *[]);

enum cont_op {
	CONT_CREATE,
	CONT_DESTROY,
	CONT_LIST,
	CONT_GET_STATUS,
	CONT_GET_STATISTICS,
	CONT_GET_ATTR,
	CONT_SET_ATTR,
	CONT_LIST_ATTRS
};

enum pool_op {
	POOL_GET_STATUS,
	POOL_GET_STATISTICS,
	POOL_GET_PROPERTIES,
	POOL_GET_ATTR,
	POOL_LIST_ATTRS
};

enum obj_op {
	OBJ_LIST,
	OBJ_GET_LAYOUT,
	OBJ_DUMP
};

static enum cont_op
cont_op_parse(const char *str)
{
	if (strcmp(str, "create") == 0)
		return CONT_CREATE;
	else if (strcmp(str, "destroy") == 0)
		return CONT_DESTROY;
	else if (strcmp(str, "list") == 0)
		return CONT_LIST;
	else if (strcmp(str, "get-status") == 0)
		return CONT_GET_STATUS;
	else if (strcmp(str, "get-statistics") == 0)
		return CONT_GET_STATISTICS;
	else if (strcmp(str, "get-attr") == 0)
		return CONT_GET_ATTR;
	else if (strcmp(str, "set-attr") == 0)
		return CONT_SET_ATTR;
	else if (strcmp(str, "list-attrs") == 0)
		return CONT_LIST_ATTRS;
	assert(0);
	return -1;
}

/* Pool operations read-only here. See dmg for full pool management */
static enum pool_op
pool_op_parse(const char *str)
{
	if (strcmp(str, "get-status") == 0)
		return POOL_GET_STATUS;
	else if (strcmp(str, "get-statistics") == 0)
		return POOL_GET_STATISTICS;
	else if (strcmp(str, "get-properties") == 0)
		return POOL_GET_PROPERTIES;
	else if (strcmp(str, "get-attr") == 0)
		return POOL_GET_ATTR;
	else if (strcmp(str, "list-attrs") == 0)
		return POOL_LIST_ATTRS;
	assert(0);
	return -1;
}

#if 0
static enum obj_op
obj_op_parse(const char *str)
{
	if (strcmp(str, "list") == 0)
		return OBJ_LIST;
	else if (strcmp(str, "get-layout") == 0)
		return OBJ_GET_LAYOUT;
	else if (strcmp(str, "dump") == 0)
		return OBJ_DUMP;
	assert(0);
	return -1;
}
#endif

/* For operations that take <pool_uuid, pool_group, pool_svc_ranks>. */
static int
pool_op_hdlr(int argc, char *argv[])
{
	struct option		options[] = {
		{"group",	required_argument,	NULL,	'G'},
		{"pool",	required_argument,	NULL,	'p'},
		{"mdsrv",	required_argument,	NULL,	'v'},
		{NULL,		0,			NULL,	0}
	};
	const char	       *group = default_group;
	uuid_t			pool_uuid;
	daos_handle_t		pool;
	const char	       *mdsrv_str = NULL;
	d_rank_list_t	       *mdsrv;
	enum pool_op		op = pool_op_parse(argv[2]);
	int			rc;
	int			rc2;

	uuid_clear(pool_uuid);

	while ((rc = getopt_long(argc, argv, "", options, NULL)) != -1) {
		switch (rc) {
		case 'G':
			group = optarg;
			break;
		case 'p':
			if (uuid_parse(optarg, pool_uuid) != 0) {
				fprintf(stderr,
					"failed to parse pool UUID: %s\n",
					optarg);
				return 2;
			}
			break;
		case 'v':
			mdsrv_str = optarg;
			break;
		default:
			printf("unknown option : %d\n", rc);
			return 2;
		}
	}

	/* Check the pool UUID. */
	if (uuid_is_null(pool_uuid)) {
		fprintf(stderr, "pool UUID required\n");
		return 2;
	}
	/* Check the pool service ranks. */
	if (mdsrv_str == NULL) {
		fprintf(stderr, "--mdsrv must be specified\n");
		return 2;
	}
	mdsrv = daos_rank_list_parse(mdsrv_str, ":");
	if (mdsrv == NULL) {
		fprintf(stderr, "failed to parse metadata service ranks\n");
		return 2;
	}

	if (mdsrv->rl_nr == 0) {
		fprintf(stderr, "--mdsrv mustn't be empty\n");
		rc = 2;
		goto bad_opt_free_mdsrv;
	}

	if (op == POOL_GET_STATUS) {
		daos_pool_info_t		 pinfo;
		struct daos_pool_space		*ps = &pinfo.pi_space;
		struct daos_rebuild_status	*rstat = &pinfo.pi_rebuild_st;
		int				 i;

		rc = daos_pool_connect(pool_uuid, group, mdsrv, DAOS_PC_RO, &pool,
				       NULL /* info */, NULL /* ev */);
		if (rc != 0) {
			fprintf(stderr, "failed to connect to pool: %d\n", rc);
			goto bad_pool_connect;
		}

		pinfo.pi_bits = DPI_ALL;
		rc = daos_pool_query(pool, NULL, &pinfo, NULL, NULL);
		if (rc != 0) {
			fprintf(stderr, "pool get-status failed: %d\n", rc);
			goto bad_pool_query;
		}
		D_PRINT("Pool "DF_UUIDF", ntarget=%u, disabled=%u\n",
			DP_UUID(pinfo.pi_uuid), pinfo.pi_ntargets,
			pinfo.pi_ndisabled);

		D_PRINT("Pool space info:\n");
		D_PRINT("- Target(VOS) count:%d\n", ps->ps_ntargets);
		for (i = DAOS_MEDIA_SCM; i < DAOS_MEDIA_MAX; i++) {
			D_PRINT("- %s:\n",
				i == DAOS_MEDIA_SCM ? "SCM" : "NVMe");
			D_PRINT("  Total size: "DF_U64"\n",
				ps->ps_space.s_total[i]);
			D_PRINT("  Free: "DF_U64", min:"DF_U64", max:"DF_U64", "
				"mean:"DF_U64"\n", ps->ps_space.s_free[i],
				ps->ps_free_min[i], ps->ps_free_max[i],
				ps->ps_free_mean[i]);
		}

		if (rstat->rs_errno == 0) {
			char	*sstr;

			if (rstat->rs_version == 0)
				sstr = "idle";
			else if (rstat->rs_done)
				sstr = "done";
			else
				sstr = "busy";

			D_PRINT("Rebuild %s, "DF_U64" objs, "DF_U64" recs\n",
				sstr, rstat->rs_obj_nr, rstat->rs_rec_nr);
		} else {
			D_PRINT("Rebuild failed, rc=%d, status=%d\n",
				rc, rstat->rs_errno);
		}
	}

	if (op == POOL_GET_STATISTICS) {
		fprintf(stdout, "Pool get-statistics operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}

	if (op == POOL_GET_PROPERTIES) {
		fprintf(stdout, "Pool get-properties operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}

	if (op == POOL_GET_ATTR) {
		fprintf(stdout, "Pool get-attr operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}
	if (op == POOL_LIST_ATTRS) {
		fprintf(stdout, "Pool list-attrs operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}

bad_op_no_impl:
bad_pool_query:
	/* Pool disconnect  in normal and error flows: preserve rc */
	if (op == POOL_GET_STATUS) {
		rc2 = daos_pool_disconnect(pool, NULL);
		if (rc2 != 0) {
			fprintf(stderr, "Pool disconnect failed : %d\n", rc2);
		}
		if (rc == 0)
			rc = rc2;
	}
bad_pool_connect:
bad_opt_free_mdsrv:
	daos_rank_list_free(mdsrv);
	return rc;
}
static int
cont_op_hdlr(int argc, char *argv[])
{
	struct option		options[] = {
		{"group",	required_argument,	NULL, 'G'},
		{"pool",	required_argument,	NULL, 'p'},
		{"mdsrv",	required_argument,	NULL, 'v'},
		{"cont",	required_argument,	NULL, 'c'},
		{NULL,		0,			NULL,  0}
	};
	const char		*group = default_group;
	uuid_t			pool_uuid;
	uuid_t			cont_uuid;
	daos_handle_t		pool;
	daos_handle_t		coh;
	const char		*mdsrv_str = NULL;
	d_rank_list_t		*mdsrv;
	daos_cont_info_t	cont_info;
	enum cont_op		op = cont_op_parse(argv[2]);
	int			rc;
	int			rc2;

	uuid_clear(pool_uuid);
	uuid_clear(cont_uuid);

	while ((rc = getopt_long(argc, argv, "", options, NULL)) != -1) {
		switch (rc) {
		case 'G':
			group = optarg;
			break;
		case 'p':
			if (uuid_parse(optarg, pool_uuid) != 0) {
				fprintf(stderr,
					"failed to parse pool UUID: %s\n",
					optarg);
				return 2;
			}
			break;
		case 'v':
			mdsrv_str = optarg;
			break;
		case 'c':
			if (uuid_parse(optarg, cont_uuid) != 0) {
				fprintf(stderr,
					"failed to parse cont UUID: %s\n",
					optarg);
				return 2;
			}
			break;
		default:
			printf("unknown option : %d\n", rc);
			return 2;
		}
	}

	/* Check the pool UUID. */
	if (uuid_is_null(pool_uuid)) {
		fprintf(stderr, "pool UUID required\n");
		return 2;
	}
	/* Check the pool service ranks. */
	if (mdsrv_str == NULL) {
		fprintf(stderr, "--mdsrv must be specified\n");
		return 2;
	}

	mdsrv = daos_rank_list_parse(mdsrv_str, ":");
	if (mdsrv == NULL) {
		fprintf(stderr, "failed to parse metadata service ranks\n");
		return 2;
	}

	if (mdsrv->rl_nr == 0) {
		fprintf(stderr, "--mdsrv mustn't be empty\n");
		rc = 2;
		goto bad_opt_free_mdsrv;
	}

	if (uuid_is_null(cont_uuid)) {
		fprintf(stderr, "valid cont uuid required\n");
		rc = 2;
		goto bad_opt_free_mdsrv;
	}

	/*
	 * all cont operations require a pool handle, lets make a =
	 * pool connection
	 */
	rc = daos_pool_connect(pool_uuid, group, mdsrv, DAOS_PC_RW, &pool,
			       NULL /* info */, NULL /* ev */);
	if (rc != 0) {
		fprintf(stderr, "failed to connect to pool: %d\n", rc);
		goto bad_pool_connect;
	}

	if (op == CONT_CREATE) {
		rc = daos_cont_create(pool, cont_uuid, NULL, NULL);
		if (rc != 0) {
			fprintf(stderr, "failed to create container: %d\n", rc);
			goto bad_cont_create;
		}
		fprintf(stdout, "Successfully created container "DF_UUIDF"\n",
			DP_UUID(cont_uuid));
	}

	if (op != CONT_DESTROY) {
		rc = daos_cont_open(pool, cont_uuid, DAOS_COO_RW, &coh,
				    &cont_info, NULL);
		if (rc != 0) {
			fprintf(stderr,"cont open failed: %d\n", rc);
			goto bad_cont_open;
		}
	}


	if (op != CONT_DESTROY) {
		rc = daos_cont_close(coh, NULL);
		if (rc != 0) {
			fprintf(stderr, "failed to close container: %d\n", rc);
			goto bad_cont_close;
		}
	}

	if (op == CONT_DESTROY) {
		rc = daos_cont_destroy(pool, cont_uuid, 1, NULL);
		if (rc != 0) {
			fprintf(stderr, "failed to destroy container: %d\n",
				rc);
			goto bad_cont_destroy;
		}
		fprintf(stdout, "Successfully destroyed container "DF_UUIDF"\n",
			DP_UUID(cont_uuid));
	}

	if (op == CONT_GET_STATUS) {
		fprintf(stdout, "Container get-status operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}

	if (op == CONT_GET_STATISTICS) {
		fprintf(stdout, "Container get-statistics operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}

	if (op == CONT_GET_ATTR) {
			fprintf(stdout, "Container get-attr operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}

	if (op == CONT_SET_ATTR) {
			fprintf(stdout, "Container set-attr operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}

	if (op == CONT_LIST_ATTRS) {
		fprintf(stdout, "Container list-attrs operation not yet implemented\n");
		rc = 1;
		goto bad_op_no_impl;
	}


bad_op_no_impl:
bad_cont_destroy:
bad_cont_close:
bad_cont_open:
bad_cont_create:
	/* Pool disconnect  in normal and error flows: preserve rc */
	rc2 = daos_pool_disconnect(pool, NULL);
	if (rc2 != 0) {
		fprintf(stderr, "Pool disconnect failed : %d\n", rc2);
	}
	if (rc == 0)
		rc = rc2;
bad_pool_connect:
bad_opt_free_mdsrv:
	daos_rank_list_free(mdsrv);
	return rc;
}

static int
help_hdlr(int argc, char *argv[])
{
	printf("\
usage: daos RESOURCE COMMAND [OPTIONS]\n\
resources:\n\
	pool            pool\n\
	cont            container\n\
	obj             object\n\
	help            print this message and exit\n");

	printf("\n\
pool commands:\n\
	list-containers list all containers in pool\n\
	get-status      query a pool\n\
	get-statistics  get pool statistics\n\
	list-attrs      list pool user-defined attributes\n\
	get-attr        get pool user-defined attribute\n\
	set-attr        set pool user-defined attribute\n");

	printf("\
pool command options:\n\
	--pool=UUID     pool UUID \n\
	--group=STR     pool server process group (\"%s\")\n\
	--mdsrv=RANKS   pool metadata service replicas like 1:2:3\n",
	default_group);

	printf("\
pool get-attr options:\n\
	--attr=ATTRNAME\n");

	printf("\
pool set-attr options:\n\
	--attr=ATTRNAME:VALUE\n");

	printf("\n\
cont commands:\n\
	create          create a container\n\
	destroy         destroy a container\n\
	list-objs       list all objects in container\n\
	get-status      query a container\n\
	get-statistics  get container statistics\n\
	list-attrs      list container user-defined attributes\n\
	get-attr        get container user-defined attribute\n\
	set-attr        set container user-defined attribute\n");

	printf("\
cont options:\n\
	--pool=UUID     pool UUID \n\
	--group=STR     pool server process group (\"%s\")\n\
	--mdsrv=RANKS   pool metadata service replicas like 1:2:3\n\
	--cont=UUID     cont UUID\n", default_group);

	printf("\
cont get-attr options:\n\
	--attr=ATTRNAME\n");

	printf("\
cont set-attr options:\n\
	--attr=ATTRNAME:VALUE\n");

	printf("\n\
obj commands:\n\
	get-layout      get object layout\n\
	dump            dump object\n");

	printf("\
obj options:\n\
	--pool=UUID     pool UUID \n\
	--group=STR     pool server process group (\"%s\")\n\
	--mdsrv=RANKS   pool metadata service replicas like 1:2:3\n\
	--cont=UUID     cont UUID\n\
	--oid=OID       object ID\n", default_group);

	return 0;
}

int
main(int argc, char *argv[])
{
	int			rc = 0;
	command_hdlr_t		hdlr = NULL;

	/* argv[1] is RESOURCE or "help"; argv[2] if provided is a resource-specific command */
	if (argc <= 2 || strcmp(argv[1], "help") == 0)
		hdlr = help_hdlr;
	else if (strcmp(argv[1], "cont") == 0)
		hdlr = cont_op_hdlr;
	else if (strcmp(argv[1], "pool") == 0)
		hdlr = pool_op_hdlr;

	if (hdlr == NULL || hdlr == help_hdlr) {
		help_hdlr(argc, argv);
		return hdlr == NULL ? 2 : 0;
	}

	rc = daos_init();
	if (rc != 0) {
		fprintf(stderr, "failed to initialize daos: %d\n", rc);
		return 1;
	}

	rc = hdlr(argc, argv);
	daos_fini();

	if (rc < 0) {
		return 1;
	} else if (rc > 0) {
		printf("rc: %d\n", rc);
		help_hdlr(argc, argv);
		return 2;
	}

	return 0;
}
