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

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/artifacts_store.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='v1/artifacts_store.proto',
  package='v1',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x18v1/artifacts_store.proto\x12\x02v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf4\x02\n\x0e\x41rtifactsStore\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x0e\n\x06readme\x18\x04 \x01(\t\x12\x0c\n\x04tags\x18\x05 \x03(\t\x12.\n\ncreated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06\x66rozen\x18\x08 \x01(\x08\x12\x10\n\x08\x64isabled\x18\t \x01(\x08\x12\x0f\n\x07\x64\x65leted\x18\n \x01(\x08\x12\x12\n\nk8s_secret\x18\x0b \x01(\t\x12\x0c\n\x04type\x18\x0c \x01(\t\x12\x12\n\nmount_path\x18\r \x01(\t\x12\x11\n\thost_path\x18\x0e \x01(\t\x12\x14\n\x0cvolume_claim\x18\x0f \x01(\t\x12\x0e\n\x06\x62ucket\x18\x10 \x01(\t\x12\x11\n\tread_only\x18\x11 \x01(\x08\"V\n\x19\x41rtifactsStoreBodyRequest\x12\r\n\x05owner\x18\x01 \x01(\t\x12*\n\x0e\x61rtifact_store\x18\x02 \x01(\x0b\x32\x12.v1.ArtifactsStore\"q\n\x1bListArtifactsStoresResponse\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\x12#\n\x07results\x18\x02 \x03(\x0b\x32\x12.v1.ArtifactsStore\x12\x10\n\x08previous\x18\x03 \x01(\t\x12\x0c\n\x04next\x18\x04 \x01(\tb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_ARTIFACTSSTORE = _descriptor.Descriptor(
  name='ArtifactsStore',
  full_name='v1.ArtifactsStore',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='v1.ArtifactsStore.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='v1.ArtifactsStore.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='v1.ArtifactsStore.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='readme', full_name='v1.ArtifactsStore.readme', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='v1.ArtifactsStore.tags', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='v1.ArtifactsStore.created_at', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='v1.ArtifactsStore.updated_at', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='frozen', full_name='v1.ArtifactsStore.frozen', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='disabled', full_name='v1.ArtifactsStore.disabled', index=8,
      number=9, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='deleted', full_name='v1.ArtifactsStore.deleted', index=9,
      number=10, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='k8s_secret', full_name='v1.ArtifactsStore.k8s_secret', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='v1.ArtifactsStore.type', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mount_path', full_name='v1.ArtifactsStore.mount_path', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='host_path', full_name='v1.ArtifactsStore.host_path', index=13,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='volume_claim', full_name='v1.ArtifactsStore.volume_claim', index=14,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bucket', full_name='v1.ArtifactsStore.bucket', index=15,
      number=16, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='read_only', full_name='v1.ArtifactsStore.read_only', index=16,
      number=17, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=66,
  serialized_end=438,
)


_ARTIFACTSSTOREBODYREQUEST = _descriptor.Descriptor(
  name='ArtifactsStoreBodyRequest',
  full_name='v1.ArtifactsStoreBodyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner', full_name='v1.ArtifactsStoreBodyRequest.owner', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='artifact_store', full_name='v1.ArtifactsStoreBodyRequest.artifact_store', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=440,
  serialized_end=526,
)


_LISTARTIFACTSSTORESRESPONSE = _descriptor.Descriptor(
  name='ListArtifactsStoresResponse',
  full_name='v1.ListArtifactsStoresResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='count', full_name='v1.ListArtifactsStoresResponse.count', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='results', full_name='v1.ListArtifactsStoresResponse.results', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='previous', full_name='v1.ListArtifactsStoresResponse.previous', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next', full_name='v1.ListArtifactsStoresResponse.next', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=528,
  serialized_end=641,
)

_ARTIFACTSSTORE.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_ARTIFACTSSTORE.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_ARTIFACTSSTOREBODYREQUEST.fields_by_name['artifact_store'].message_type = _ARTIFACTSSTORE
_LISTARTIFACTSSTORESRESPONSE.fields_by_name['results'].message_type = _ARTIFACTSSTORE
DESCRIPTOR.message_types_by_name['ArtifactsStore'] = _ARTIFACTSSTORE
DESCRIPTOR.message_types_by_name['ArtifactsStoreBodyRequest'] = _ARTIFACTSSTOREBODYREQUEST
DESCRIPTOR.message_types_by_name['ListArtifactsStoresResponse'] = _LISTARTIFACTSSTORESRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ArtifactsStore = _reflection.GeneratedProtocolMessageType('ArtifactsStore', (_message.Message,), {
  'DESCRIPTOR' : _ARTIFACTSSTORE,
  '__module__' : 'v1.artifacts_store_pb2'
  # @@protoc_insertion_point(class_scope:v1.ArtifactsStore)
  })
_sym_db.RegisterMessage(ArtifactsStore)

ArtifactsStoreBodyRequest = _reflection.GeneratedProtocolMessageType('ArtifactsStoreBodyRequest', (_message.Message,), {
  'DESCRIPTOR' : _ARTIFACTSSTOREBODYREQUEST,
  '__module__' : 'v1.artifacts_store_pb2'
  # @@protoc_insertion_point(class_scope:v1.ArtifactsStoreBodyRequest)
  })
_sym_db.RegisterMessage(ArtifactsStoreBodyRequest)

ListArtifactsStoresResponse = _reflection.GeneratedProtocolMessageType('ListArtifactsStoresResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTARTIFACTSSTORESRESPONSE,
  '__module__' : 'v1.artifacts_store_pb2'
  # @@protoc_insertion_point(class_scope:v1.ListArtifactsStoresResponse)
  })
_sym_db.RegisterMessage(ListArtifactsStoresResponse)


# @@protoc_insertion_point(module_scope)
