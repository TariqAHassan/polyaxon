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

import rhea

from polyaxon import kinds
from polyaxon.specs.base import BaseSpecification
from polyaxon.specs.op import OpSpecification
from polyaxon.specs.component import ComponentSpecification

SPECIFICATION_BY_KIND = {
    kinds.OP: OpSpecification,
    kinds.COMPONENT: ComponentSpecification,
}


def get_specification(data):
    data = rhea.read(data)
    kind = BaseSpecification.get_kind(data=data)
    return SPECIFICATION_BY_KIND[kind](data)
