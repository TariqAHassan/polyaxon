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
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from v1 import base_pb2 as v1_dot_base__pb2
from v1 import experiment_pb2 as v1_dot_experiment__pb2


class ExperimentServiceStub(object):
  """Service to manage experiments
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListExperiments = channel.unary_unary(
        '/v1.ExperimentService/ListExperiments',
        request_serializer=v1_dot_base__pb2.ProjectBodyRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.ListExperimentsResponse.FromString,
        )
    self.ListBookmarkedExperiments = channel.unary_unary(
        '/v1.ExperimentService/ListBookmarkedExperiments',
        request_serializer=v1_dot_base__pb2.OwnerBodyRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.ListExperimentsResponse.FromString,
        )
    self.ListArchivedExperiments = channel.unary_unary(
        '/v1.ExperimentService/ListArchivedExperiments',
        request_serializer=v1_dot_base__pb2.OwnerBodyRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.ListExperimentsResponse.FromString,
        )
    self.CreateExperiment = channel.unary_unary(
        '/v1.ExperimentService/CreateExperiment',
        request_serializer=v1_dot_experiment__pb2.ExperimentBodyRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.Experiment.FromString,
        )
    self.GetExperiment = channel.unary_unary(
        '/v1.ExperimentService/GetExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.Experiment.FromString,
        )
    self.UpdateExperiment = channel.unary_unary(
        '/v1.ExperimentService/UpdateExperiment',
        request_serializer=v1_dot_experiment__pb2.ExperimentBodyRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.Experiment.FromString,
        )
    self.DeleteExperiment = channel.unary_unary(
        '/v1.ExperimentService/DeleteExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.DeleteExperiments = channel.unary_unary(
        '/v1.ExperimentService/DeleteExperiments',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.StopExperiment = channel.unary_unary(
        '/v1.ExperimentService/StopExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.StopExperiments = channel.unary_unary(
        '/v1.ExperimentService/StopExperiments',
        request_serializer=v1_dot_base__pb2.ProjectBodyRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.RestartExperiment = channel.unary_unary(
        '/v1.ExperimentService/RestartExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.Experiment.FromString,
        )
    self.ResumeExperiment = channel.unary_unary(
        '/v1.ExperimentService/ResumeExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.Experiment.FromString,
        )
    self.ArchiveExperiment = channel.unary_unary(
        '/v1.ExperimentService/ArchiveExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.RestoreExperiment = channel.unary_unary(
        '/v1.ExperimentService/RestoreExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.BookmarkExperiment = channel.unary_unary(
        '/v1.ExperimentService/BookmarkExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.UnBookmarkExperiment = channel.unary_unary(
        '/v1.ExperimentService/UnBookmarkExperiment',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.StartExperimentTensorboard = channel.unary_unary(
        '/v1.ExperimentService/StartExperimentTensorboard',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.StopExperimentTensorboard = channel.unary_unary(
        '/v1.ExperimentService/StopExperimentTensorboard',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetExperimentStatuses = channel.unary_unary(
        '/v1.ExperimentService/GetExperimentStatuses',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_base__pb2.StatusResponse.FromString,
        )
    self.ListExperimentStatuses = channel.unary_unary(
        '/v1.ExperimentService/ListExperimentStatuses',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.ListExperimentStatusesResponse.FromString,
        )
    self.CreateExperimentStatus = channel.unary_unary(
        '/v1.ExperimentService/CreateExperimentStatus',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_experiment__pb2.ExperimentStatus.FromString,
        )
    self.GetExperimentCodeRef = channel.unary_unary(
        '/v1.ExperimentService/GetExperimentCodeRef',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GreateExperimentCodeRef = channel.unary_unary(
        '/v1.ExperimentService/GreateExperimentCodeRef',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )


class ExperimentServiceServicer(object):
  """Service to manage experiments
  """

  def ListExperiments(self, request, context):
    """List experiments
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListBookmarkedExperiments(self, request, context):
    """List bookmarked experiments
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListArchivedExperiments(self, request, context):
    """List archived experiments
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateExperiment(self, request, context):
    """Create new experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetExperiment(self, request, context):
    """Get experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateExperiment(self, request, context):
    """Update experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteExperiment(self, request, context):
    """Delete experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteExperiments(self, request, context):
    """Delete experiments
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StopExperiment(self, request, context):
    """Stop experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StopExperiments(self, request, context):
    """Stop experiments
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RestartExperiment(self, request, context):
    """Restart experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ResumeExperiment(self, request, context):
    """Resume experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ArchiveExperiment(self, request, context):
    """Archive experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RestoreExperiment(self, request, context):
    """Restore experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def BookmarkExperiment(self, request, context):
    """Bookmark experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UnBookmarkExperiment(self, request, context):
    """UnBookmark experiment
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StartExperimentTensorboard(self, request, context):
    """Start experiment tensorboard
    TODO: should be a tensorboard object
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StopExperimentTensorboard(self, request, context):
    """Stop experiment tensorboard
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetExperimentStatuses(self, request, context):
    """Get experiment status
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListExperimentStatuses(self, request, context):
    """List experiment statuses
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateExperimentStatus(self, request, context):
    """Create new experiment status
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetExperimentCodeRef(self, request, context):
    """Get experiment code ref
    TODO: should be an code ref
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GreateExperimentCodeRef(self, request, context):
    """Get experiment code ref
    TODO: should be an code ref
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ExperimentServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListExperiments': grpc.unary_unary_rpc_method_handler(
          servicer.ListExperiments,
          request_deserializer=v1_dot_base__pb2.ProjectBodyRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.ListExperimentsResponse.SerializeToString,
      ),
      'ListBookmarkedExperiments': grpc.unary_unary_rpc_method_handler(
          servicer.ListBookmarkedExperiments,
          request_deserializer=v1_dot_base__pb2.OwnerBodyRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.ListExperimentsResponse.SerializeToString,
      ),
      'ListArchivedExperiments': grpc.unary_unary_rpc_method_handler(
          servicer.ListArchivedExperiments,
          request_deserializer=v1_dot_base__pb2.OwnerBodyRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.ListExperimentsResponse.SerializeToString,
      ),
      'CreateExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.CreateExperiment,
          request_deserializer=v1_dot_experiment__pb2.ExperimentBodyRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.Experiment.SerializeToString,
      ),
      'GetExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.GetExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.Experiment.SerializeToString,
      ),
      'UpdateExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateExperiment,
          request_deserializer=v1_dot_experiment__pb2.ExperimentBodyRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.Experiment.SerializeToString,
      ),
      'DeleteExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'DeleteExperiments': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteExperiments,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'StopExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.StopExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'StopExperiments': grpc.unary_unary_rpc_method_handler(
          servicer.StopExperiments,
          request_deserializer=v1_dot_base__pb2.ProjectBodyRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'RestartExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.RestartExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.Experiment.SerializeToString,
      ),
      'ResumeExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.ResumeExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.Experiment.SerializeToString,
      ),
      'ArchiveExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.ArchiveExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'RestoreExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.RestoreExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'BookmarkExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.BookmarkExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'UnBookmarkExperiment': grpc.unary_unary_rpc_method_handler(
          servicer.UnBookmarkExperiment,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'StartExperimentTensorboard': grpc.unary_unary_rpc_method_handler(
          servicer.StartExperimentTensorboard,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'StopExperimentTensorboard': grpc.unary_unary_rpc_method_handler(
          servicer.StopExperimentTensorboard,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetExperimentStatuses': grpc.unary_unary_rpc_method_handler(
          servicer.GetExperimentStatuses,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_base__pb2.StatusResponse.SerializeToString,
      ),
      'ListExperimentStatuses': grpc.unary_unary_rpc_method_handler(
          servicer.ListExperimentStatuses,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.ListExperimentStatusesResponse.SerializeToString,
      ),
      'CreateExperimentStatus': grpc.unary_unary_rpc_method_handler(
          servicer.CreateExperimentStatus,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_experiment__pb2.ExperimentStatus.SerializeToString,
      ),
      'GetExperimentCodeRef': grpc.unary_unary_rpc_method_handler(
          servicer.GetExperimentCodeRef,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GreateExperimentCodeRef': grpc.unary_unary_rpc_method_handler(
          servicer.GreateExperimentCodeRef,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'v1.ExperimentService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
