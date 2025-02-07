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

from polyaxon import kinds
from polyaxon.exceptions import PolyaxonSchemaError
from polyaxon.schemas.polyflow.op import OpConfig
from polyaxon.specs.base import BaseSpecification


class OpSpecificationMixin(object):
    @property
    def dependencies(self):
        return self.config.dependencies

    @property
    def trigger(self):
        return self.config.trigger

    @property
    def conditions(self):
        return self.config.conditions

    @property
    def skip_on_upstream_skip(self):
        return self.config.skip_on_upstream_skip


class OpSpecification(BaseSpecification, OpSpecificationMixin):
    """The polyaxonfile specification for ops."""

    _SPEC_KIND = kinds.OP

    CONFIG = OpConfig

    def apply_context(self):
        raise PolyaxonSchemaError("This method is not allowed on this specification.")

    def apply_container_contexts(self, contexts=None):
        raise PolyaxonSchemaError("This method is not allowed on this specification.")

    def generate_run_data(self, override=None, override_post=True):
        values = [
            {"kind": kinds.COMPONENT, "version": self.config.version},
            self.config.component.to_light_dict(),
        ]
        op_override = {}
        for field in [
            self.NAME,
            self.DESCRIPTION,
            self.TAGS,
            self.ENVIRONMENT,
            self.TERMINATION,
            self.INIT,
            self.MOUNTS,
            self.SERVICE,
            self.PROFILE,
            self.NOCACHE,
            self.WORKFLOW,
        ]:
            override_field = getattr(self.config, field)
            if hasattr(override_field, "to_dict"):
                op_override[field] = override_field.to_dict()
            elif override_field:
                op_override[field] = override_field
        if override_post:
            if op_override:
                values.append(op_override)
            if override:
                values.append(override)
        else:
            if override:
                values.append(override)
            if op_override:
                values.append(op_override)
        return values
