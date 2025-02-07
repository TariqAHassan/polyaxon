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

import polyaxon_sdk

from marshmallow import EXCLUDE, fields

from polyaxon.env_vars.keys import (
    POLYAXON_KEYS_API_HOST,
    POLYAXON_KEYS_API_VERSION,
    POLYAXON_KEYS_ASSERT_HOSTNAME,
    POLYAXON_KEYS_AUTHENTICATION_TYPE,
    POLYAXON_KEYS_CERT_FILE,
    POLYAXON_KEYS_CONNECTION_POOL_MAXSIZE,
    POLYAXON_KEYS_DEBUG,
    POLYAXON_KEYS_HEADER,
    POLYAXON_KEYS_HEADER_SERVICE,
    POLYAXON_KEYS_IN_CLUSTER,
    POLYAXON_KEYS_INTERVAL,
    POLYAXON_KEYS_IS_LOCAL,
    POLYAXON_KEYS_IS_MANAGED,
    POLYAXON_KEYS_IS_OFFLINE,
    POLYAXON_KEYS_IS_SERVICE,
    POLYAXON_KEYS_K8S_NAMESPACE,
    POLYAXON_KEYS_K8S_POD_ID,
    POLYAXON_KEYS_KEY_FILE,
    POLYAXON_KEYS_LOG_LEVEL,
    POLYAXON_KEYS_NO_OP,
    POLYAXON_KEYS_SSL_CA_CERT,
    POLYAXON_KEYS_TIMEOUT,
    POLYAXON_KEYS_VERIFY_SSL,
    POLYAXON_KEYS_WATCH_INTERVAL,
)
from polyaxon.exceptions import PolyaxonClientException
from polyaxon.managers.auth import AuthConfigManager
from polyaxon.schemas.base import BaseConfig, BaseSchema
from polyaxon.services.auth import AuthenticationTypes


class ClientSchema(BaseSchema):
    host = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_API_HOST)
    version = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_API_VERSION)
    debug = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_DEBUG)
    log_level = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_LOG_LEVEL)
    authentication_type = fields.Str(
        allow_none=True, data_key=POLYAXON_KEYS_AUTHENTICATION_TYPE
    )
    is_managed = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_IS_MANAGED)
    is_service = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_IS_SERVICE)
    in_cluster = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_IN_CLUSTER)
    is_local = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_IS_LOCAL)
    is_offline = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_IS_OFFLINE)
    no_op = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_NO_OP)
    timeout = fields.Float(allow_none=True, data_key=POLYAXON_KEYS_TIMEOUT)
    watch_interval = fields.Int(allow_none=True, data_key=POLYAXON_KEYS_WATCH_INTERVAL)
    interval = fields.Float(allow_none=True, data_key=POLYAXON_KEYS_INTERVAL)
    verify_ssl = fields.Bool(allow_none=True, data_key=POLYAXON_KEYS_VERIFY_SSL)
    ssl_ca_cert = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_SSL_CA_CERT)
    cert_file = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_CERT_FILE)
    key_file = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_KEY_FILE)
    assert_hostname = fields.Bool(
        allow_none=True, data_key=POLYAXON_KEYS_ASSERT_HOSTNAME
    )
    connection_pool_maxsize = fields.Int(
        allow_none=True, data_key=POLYAXON_KEYS_CONNECTION_POOL_MAXSIZE
    )

    header = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_HEADER)
    header_service = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_HEADER_SERVICE)

    pod_id = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_K8S_POD_ID)
    namespace = fields.Str(allow_none=True, data_key=POLYAXON_KEYS_K8S_NAMESPACE)

    @staticmethod
    def schema_config():
        return ClientConfig


class ClientConfig(BaseConfig):
    SCHEMA = ClientSchema
    IDENTIFIER = "global"

    PAGE_SIZE = 20
    BASE_URL = "{}/api/{}"

    UNKNOWN_BEHAVIOUR = EXCLUDE

    def __init__(
        self,
        host=None,
        token=None,
        debug=None,
        log_level=None,
        version=None,
        authentication_type=None,
        is_managed=None,
        is_service=None,
        is_local=None,
        is_offline=None,
        in_cluster=None,
        no_op=None,
        timeout=None,
        watch_interval=None,
        interval=None,
        verify_ssl=None,
        ssl_ca_cert=None,
        cert_file=None,
        key_file=None,
        assert_hostname=None,
        connection_pool_maxsize=None,
        header=None,
        header_service=None,
        pod_id=None,
        namespace=None,
        **kwargs
    ):

        self.host = host or "https://cloud.polyaxon.com"
        self.token = token
        self.debug = self._get_bool(debug, False)
        self.log_level = log_level
        self.version = version or "v1"
        self.is_managed = self._get_bool(is_managed, False)
        self.is_service = self._get_bool(is_service, False)
        self.is_local = self._get_bool(is_local, False)
        self.is_offline = self._get_bool(is_offline, False)
        self.in_cluster = self._get_bool(in_cluster, False)
        self.no_op = self._get_bool(no_op, False)
        self.verify_ssl = self._get_bool(verify_ssl, None)
        self.ssl_ca_cert = ssl_ca_cert
        self.cert_file = cert_file
        self.key_file = key_file
        self.assert_hostname = self._get_bool(assert_hostname, None)
        self.connection_pool_maxsize = connection_pool_maxsize
        self.header = header
        self.header_service = header_service
        self.timeout = timeout or 20
        self.interval = interval or 5
        self.watch_interval = watch_interval or 5
        self.pod_id = pod_id
        self.namespace = namespace
        self.authentication_type = authentication_type or AuthenticationTypes.TOKEN

        self.client_header = {}

        if all([self.header, self.header_service]):
            self.client_header["header_name"] = self.header
            self.client_header["header_value"] = self.header_service

        self.base_url = self.BASE_URL.format(self.host, self.version)

    @staticmethod
    def get_config_from_manager():
        from polyaxon.managers.client import ClientConfigManager

        config = ClientConfigManager.get_config_or_default()
        config.token = AuthConfigManager.get_value("token")
        return config

    @property
    def sdk_config(self):
        if not self.host and not self.in_cluster:
            raise PolyaxonClientException(
                "Api config requires at least a host if not running in-cluster."
            )

        config = polyaxon_sdk.Configuration()
        config.debug = self.debug
        config.host = self.host
        config.verify_ssl = self.verify_ssl
        config.ssl_ca_cert = self.ssl_ca_cert
        config.cert_file = self.cert_file
        config.key_file = self.key_file
        config.assert_hostname = self.assert_hostname
        if self.connection_pool_maxsize:
            config.connection_pool_maxsize = self.connection_pool_maxsize
        if self.token:
            config.api_key["Authorization"] = self.token
            config.api_key_prefix["Authorization"] = self.authentication_type
        return config

    @staticmethod
    def _get_bool(value, default_value):
        if isinstance(value, bool):
            return value

        return default_value
