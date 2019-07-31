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
from v1 import job_pb2 as v1_dot_job__pb2


class JobServiceStub(object):
  """Service to manage jobs
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListJobs = channel.unary_unary(
        '/v1.JobService/ListJobs',
        request_serializer=v1_dot_base__pb2.ProjectBodyRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.ListJobsResponse.FromString,
        )
    self.ListBookmarkedJobs = channel.unary_unary(
        '/v1.JobService/ListBookmarkedJobs',
        request_serializer=v1_dot_base__pb2.OwnerBodyRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.ListJobsResponse.FromString,
        )
    self.ListArchivedJobs = channel.unary_unary(
        '/v1.JobService/ListArchivedJobs',
        request_serializer=v1_dot_base__pb2.OwnerBodyRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.ListJobsResponse.FromString,
        )
    self.CreateJob = channel.unary_unary(
        '/v1.JobService/CreateJob',
        request_serializer=v1_dot_job__pb2.JobBodyRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.Job.FromString,
        )
    self.GetJob = channel.unary_unary(
        '/v1.JobService/GetJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.Job.FromString,
        )
    self.UpdateJob = channel.unary_unary(
        '/v1.JobService/UpdateJob',
        request_serializer=v1_dot_job__pb2.JobBodyRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.Job.FromString,
        )
    self.DeleteJob = channel.unary_unary(
        '/v1.JobService/DeleteJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.DeleteJobs = channel.unary_unary(
        '/v1.JobService/DeleteJobs',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.StopJob = channel.unary_unary(
        '/v1.JobService/StopJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.StopJobs = channel.unary_unary(
        '/v1.JobService/StopJobs',
        request_serializer=v1_dot_base__pb2.ProjectBodyRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.RestartJob = channel.unary_unary(
        '/v1.JobService/RestartJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.Job.FromString,
        )
    self.ResumeJob = channel.unary_unary(
        '/v1.JobService/ResumeJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.Job.FromString,
        )
    self.ArchiveJob = channel.unary_unary(
        '/v1.JobService/ArchiveJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.RestoreJob = channel.unary_unary(
        '/v1.JobService/RestoreJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.BookmarkJob = channel.unary_unary(
        '/v1.JobService/BookmarkJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.UnBookmarkJob = channel.unary_unary(
        '/v1.JobService/UnBookmarkJob',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetJobStatuses = channel.unary_unary(
        '/v1.JobService/GetJobStatuses',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_base__pb2.StatusResponse.FromString,
        )
    self.ListJobStatuses = channel.unary_unary(
        '/v1.JobService/ListJobStatuses',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.ListJobStatusesResponse.FromString,
        )
    self.CreateJobStatus = channel.unary_unary(
        '/v1.JobService/CreateJobStatus',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=v1_dot_job__pb2.JobStatus.FromString,
        )
    self.GetJobCodeRef = channel.unary_unary(
        '/v1.JobService/GetJobCodeRef',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GreateJobCodeRef = channel.unary_unary(
        '/v1.JobService/GreateJobCodeRef',
        request_serializer=v1_dot_base__pb2.OwnedEntityIdRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )


class JobServiceServicer(object):
  """Service to manage jobs
  """

  def ListJobs(self, request, context):
    """List jobs
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListBookmarkedJobs(self, request, context):
    """List bookmarked jobs
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListArchivedJobs(self, request, context):
    """List archived jobs
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateJob(self, request, context):
    """Create new job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetJob(self, request, context):
    """Get job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateJob(self, request, context):
    """Update job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteJob(self, request, context):
    """Delete job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteJobs(self, request, context):
    """Delete jobs
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StopJob(self, request, context):
    """Stop job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StopJobs(self, request, context):
    """Stop jobs
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RestartJob(self, request, context):
    """Restart job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ResumeJob(self, request, context):
    """Resume job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ArchiveJob(self, request, context):
    """Archive job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RestoreJob(self, request, context):
    """Restore job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def BookmarkJob(self, request, context):
    """Bookmark job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UnBookmarkJob(self, request, context):
    """UnBookmark job
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetJobStatuses(self, request, context):
    """Get job status
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListJobStatuses(self, request, context):
    """List job statuses
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateJobStatus(self, request, context):
    """Create new job status
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetJobCodeRef(self, request, context):
    """Get job code ref
    TODO: should be an code ref
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GreateJobCodeRef(self, request, context):
    """Get job code ref
    TODO: should be an code ref
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_JobServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListJobs': grpc.unary_unary_rpc_method_handler(
          servicer.ListJobs,
          request_deserializer=v1_dot_base__pb2.ProjectBodyRequest.FromString,
          response_serializer=v1_dot_job__pb2.ListJobsResponse.SerializeToString,
      ),
      'ListBookmarkedJobs': grpc.unary_unary_rpc_method_handler(
          servicer.ListBookmarkedJobs,
          request_deserializer=v1_dot_base__pb2.OwnerBodyRequest.FromString,
          response_serializer=v1_dot_job__pb2.ListJobsResponse.SerializeToString,
      ),
      'ListArchivedJobs': grpc.unary_unary_rpc_method_handler(
          servicer.ListArchivedJobs,
          request_deserializer=v1_dot_base__pb2.OwnerBodyRequest.FromString,
          response_serializer=v1_dot_job__pb2.ListJobsResponse.SerializeToString,
      ),
      'CreateJob': grpc.unary_unary_rpc_method_handler(
          servicer.CreateJob,
          request_deserializer=v1_dot_job__pb2.JobBodyRequest.FromString,
          response_serializer=v1_dot_job__pb2.Job.SerializeToString,
      ),
      'GetJob': grpc.unary_unary_rpc_method_handler(
          servicer.GetJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_job__pb2.Job.SerializeToString,
      ),
      'UpdateJob': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateJob,
          request_deserializer=v1_dot_job__pb2.JobBodyRequest.FromString,
          response_serializer=v1_dot_job__pb2.Job.SerializeToString,
      ),
      'DeleteJob': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'DeleteJobs': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteJobs,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'StopJob': grpc.unary_unary_rpc_method_handler(
          servicer.StopJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'StopJobs': grpc.unary_unary_rpc_method_handler(
          servicer.StopJobs,
          request_deserializer=v1_dot_base__pb2.ProjectBodyRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'RestartJob': grpc.unary_unary_rpc_method_handler(
          servicer.RestartJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_job__pb2.Job.SerializeToString,
      ),
      'ResumeJob': grpc.unary_unary_rpc_method_handler(
          servicer.ResumeJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_job__pb2.Job.SerializeToString,
      ),
      'ArchiveJob': grpc.unary_unary_rpc_method_handler(
          servicer.ArchiveJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'RestoreJob': grpc.unary_unary_rpc_method_handler(
          servicer.RestoreJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'BookmarkJob': grpc.unary_unary_rpc_method_handler(
          servicer.BookmarkJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'UnBookmarkJob': grpc.unary_unary_rpc_method_handler(
          servicer.UnBookmarkJob,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetJobStatuses': grpc.unary_unary_rpc_method_handler(
          servicer.GetJobStatuses,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_base__pb2.StatusResponse.SerializeToString,
      ),
      'ListJobStatuses': grpc.unary_unary_rpc_method_handler(
          servicer.ListJobStatuses,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_job__pb2.ListJobStatusesResponse.SerializeToString,
      ),
      'CreateJobStatus': grpc.unary_unary_rpc_method_handler(
          servicer.CreateJobStatus,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=v1_dot_job__pb2.JobStatus.SerializeToString,
      ),
      'GetJobCodeRef': grpc.unary_unary_rpc_method_handler(
          servicer.GetJobCodeRef,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GreateJobCodeRef': grpc.unary_unary_rpc_method_handler(
          servicer.GreateJobCodeRef,
          request_deserializer=v1_dot_base__pb2.OwnedEntityIdRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'v1.JobService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
