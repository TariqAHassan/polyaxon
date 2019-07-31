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

import sys


def main(argv):
    pass


if __name__ == '__main__':
    main(sys.argv)
# coding: utf-8

"""
    Experiment service

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: 1.0
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from swagger_client.models.v1_experiment import V1Experiment  # noqa: F401,E501


class V1ExperimentBodyRequest(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'owner': 'str',
        'project': 'str',
        'experiment': 'V1Experiment'
    }

    attribute_map = {
        'owner': 'owner',
        'project': 'project',
        'experiment': 'experiment'
    }

    def __init__(self, owner=None, project=None, experiment=None):  # noqa: E501
        """V1ExperimentBodyRequest - a model defined in Swagger"""  # noqa: E501

        self._owner = None
        self._project = None
        self._experiment = None
        self.discriminator = None

        if owner is not None:
            self.owner = owner
        if project is not None:
            self.project = project
        if experiment is not None:
            self.experiment = experiment

    @property
    def owner(self):
        """Gets the owner of this V1ExperimentBodyRequest.  # noqa: E501


        :return: The owner of this V1ExperimentBodyRequest.  # noqa: E501
        :rtype: str
        """
        return self._owner

    @owner.setter
    def owner(self, owner):
        """Sets the owner of this V1ExperimentBodyRequest.


        :param owner: The owner of this V1ExperimentBodyRequest.  # noqa: E501
        :type: str
        """

        self._owner = owner

    @property
    def project(self):
        """Gets the project of this V1ExperimentBodyRequest.  # noqa: E501


        :return: The project of this V1ExperimentBodyRequest.  # noqa: E501
        :rtype: str
        """
        return self._project

    @project.setter
    def project(self, project):
        """Sets the project of this V1ExperimentBodyRequest.


        :param project: The project of this V1ExperimentBodyRequest.  # noqa: E501
        :type: str
        """

        self._project = project

    @property
    def experiment(self):
        """Gets the experiment of this V1ExperimentBodyRequest.  # noqa: E501


        :return: The experiment of this V1ExperimentBodyRequest.  # noqa: E501
        :rtype: V1Experiment
        """
        return self._experiment

    @experiment.setter
    def experiment(self, experiment):
        """Sets the experiment of this V1ExperimentBodyRequest.


        :param experiment: The experiment of this V1ExperimentBodyRequest.  # noqa: E501
        :type: V1Experiment
        """

        self._experiment = experiment

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(V1ExperimentBodyRequest, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1ExperimentBodyRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
