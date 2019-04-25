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

#include "dfuse_common.h"
#include "dfuse.h"

#define REQ_NAME request
#define POOL_NAME fgh_da
#define TYPE_NAME common_req
#include "dfuse_ops.h"

static bool
dfuse_getattr_result_fn(struct dfuse_request *request)
{
	struct dfuse_attr_out *out = request->out;

	DFUSE_REQUEST_RESOLVE(request, out);

	if (request->rc == 0)
		DFUSE_REPLY_ATTR(request->req, &out->stat);
	else
		DFUSE_REPLY_ERR(request, request->rc);

	dfuse_da_release(request->fsh->POOL_NAME, CONTAINER(request));
	return false;
}

static const struct dfuse_request_api getattr_api = {
	.on_result	= dfuse_getattr_result_fn,
};

void
dfuse_cb_getattr(fuse_req_t req, struct dfuse_inode_entry *inode)
{
	struct stat	stat = {};
	int		rc;

	rc = dfs_ostat(inode->ie_dfs->dffs_dfs, inode->obj, &stat);
	if (rc != -DER_SUCCESS) {
		D_GOTO(err, 0);
	}

	DFUSE_REPLY_ATTR(req, &stat);

	return;
err:
	DFUSE_REPLY_ERR_RAW(inode, req, rc);
}

void
Xdfuse_cb_getattr(fuse_req_t req, fuse_ino_t ino, struct fuse_file_info *fi)
{
	struct dfuse_projection_info	*fs_handle = fuse_req_userdata(req);
	struct dfuse_file_handle	*handle = NULL;
	struct common_req		*desc = NULL;
	int rc;

	if (fi)
		handle = (void *)fi->fh;

	DFUSE_TRA_INFO(fs_handle, "inode %lu handle %p", ino, handle);

	DFUSE_REQ_INIT_REQ(desc, fs_handle, getattr_api, req, rc);
	if (rc)
		D_GOTO(err, rc);

	if (handle) {
		desc->request.ir_ht = RHS_FILE;
		desc->request.ir_file = handle;
		rc = dfuse_fs_send(&desc->request);
		if (rc != 0)
			D_GOTO(err, rc);
	} else {
		if (rc != -DER_SUCCESS) {
			D_GOTO(err, rc = ENOTSUP);
		}
	}

	return;
err:
	DFUSE_REPLY_ERR_RAW(fs_handle, req, rc);
	if (desc) {
		DFUSE_TRA_DOWN(&desc->request);
		dfuse_da_release(fs_handle->POOL_NAME, desc);
	}
}
