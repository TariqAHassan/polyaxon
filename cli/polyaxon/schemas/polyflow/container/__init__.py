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

from hestia.list_utils import to_list
from hestia.string_utils import strip_spaces
from marshmallow import ValidationError, fields, validates_schema

from polyaxon.schemas.base import BaseConfig, BaseSchema
from polyaxon.schemas.fields import ObjectOrListObject
from polyaxon.schemas.fields.docker_image import validate_image


def get_container_command_args(config):
    def sanitize_str(value):
        if not value:
            return
        value = strip_spaces(value=value, join=False)
        value = [c.strip().strip("\\") for c in value if (c and c != "\\")]
        value = [c for c in value if (c and c != "\\")]
        return " ".join(value)

    def sanitize(value):
        return (
            [sanitize_str(v) for v in value]
            if isinstance(value, list)
            else to_list(sanitize_str(value), check_none=True)
        )

    return to_list(config.command, check_none=True), sanitize(config.args)


class ContainerSchema(BaseSchema):
    image = fields.Str(required=True)
    image_pull_policy = fields.Str(allow_none=True)
    command = ObjectOrListObject(fields.Str, allow_none=True)
    args = ObjectOrListObject(fields.Str, allow_none=True)

    @staticmethod
    def schema_config():
        return ContainerConfig

    @validates_schema
    def validate_container(self, values):
        validate_image(values.get("image"))


class ContainerConfig(BaseConfig):
    SCHEMA = ContainerSchema
    IDENTIFIER = "container"
    REDUCED_ATTRIBUTES = ["image_pull_policy", "command", "args"]

    def __init__(self, image=None, image_pull_policy=None, command=None, args=None):
        validate_image(image)
        self.image = image
        self.image_pull_policy = image_pull_policy
        self.command = command
        self.args = args

    def get_container_command_args(self):
        return get_container_command_args(self)
