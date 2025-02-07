#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8
from __future__ import absolute_import, division, print_function

import json
import os

from polyaxon_sdk.rest import ApiException
from urllib3.exceptions import HTTPError

from polyaxon.client import PolyaxonClient
from polyaxon.containers.contexts import CONTEXT_MOUNT_AUTH
from polyaxon.exceptions import PolyaxonClientException
from polyaxon.logger import logger
from polyaxon.schemas.api.authentication import AccessTokenConfig


def create_polyaxon_tmp():
    base_path = os.path.join("/tmp", ".polyaxon")
    if not os.path.exists(base_path):
        try:
            os.makedirs(base_path)
        except OSError:
            # Except permission denied and potential race conditions
            # in multi-threaded environments.
            logger.warning("Could not create config directory `%s`", base_path)
    return base_path


def create_context_auth(access_token, context_auth_path=None):
    context_auth_path = context_auth_path or CONTEXT_MOUNT_AUTH
    with open(context_auth_path, "w") as config_file:
        config_file.write(access_token.to_dict(dump=True))


def impersonate(owner, project, run_uuid):
    try:
        response = PolyaxonClient().runs_v1.impersonate_token(owner, project, run_uuid)
        polyaxon_client = PolyaxonClient(token=response.token)
        user = polyaxon_client.users_v1.get_user()
        access_token = AccessTokenConfig(username=user.username, token=response.token)
        create_context_auth(access_token)
    except (ApiException, HTTPError) as e:
        raise PolyaxonClientException(
            "This worker is not allowed to run this job %s." % e
        )
