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

from unittest import TestCase

from marshmallow import ValidationError

from polyaxon.deploy.schemas.intervals import IntervalsConfig


class TestIntervalsConfig(TestCase):
    def test_intervals_config(self):
        bad_config_dicts = [
            {"experimentsScheduler": "dsf"},
            {"experimentsSync": "dsf"},
            {"clustersUpdateSystemInfo": "dsf"},
            {"clustersUpdateSystemNodes": "dsf"},
            {"operationsDefaultRetryDelay": "dsf"},
            {"operationsMaxRetryDelay": ["dsf"]},
        ]

        for config_dict in bad_config_dicts:
            with self.assertRaises(ValidationError):
                IntervalsConfig.from_dict(config_dict)

        config_dict = {
            "experimentsScheduler": 12,
            "experimentsSync": 12,
            "clustersUpdateSystemInfo": 12,
            "clustersUpdateSystemNodes": 12,
            "pipelinesScheduler": 12,
            "operationsDefaultRetryDelay": 12,
            "operationsMaxRetryDelay": 12,
        }
        config = IntervalsConfig.from_dict(config_dict)
        assert config.to_light_dict() == config_dict
