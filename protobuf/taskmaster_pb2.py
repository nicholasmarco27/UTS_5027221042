# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: taskmaster.proto
# Protobuf Python Version: 5.27.0-rc1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    0,
    '-rc1',
    'taskmaster.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import Empty_pb2 as google_dot_protobuf_dot_Empty__pb2
from google.protobuf import Wrappers_pb2 as google_dot_protobuf_dot_Wrappers__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10taskmaster.proto\x12\x08protoapi\x1a\x1bgoogle/protobuf/Empty.proto\x1a\x1egoogle/protobuf/Wrappers.proto\"W\n\x04Task\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\r\n\x05price\x18\x04 \x01(\x05\x12\x10\n\x08quantity\x18\x05 \x01(\x05\"(\n\x08TaskList\x12\x1c\n\x04list\x18\x01 \x03(\x0b\x32\x0e.protoapi.Task2\xee\x01\n\x07TaskApi\x12.\n\nCreateTask\x12\x0e.protoapi.Task\x1a\x0e.protoapi.Task\"\x00\x12\x39\n\tListTasks\x12\x16.google.protobuf.Empty\x1a\x12.protoapi.TaskList\"\x00\x12.\n\nUpdateTask\x12\x0e.protoapi.Task\x1a\x0e.protoapi.Task\"\x00\x12H\n\nDeleteTask\x12\x1c.google.protobuf.StringValue\x1a\x1a.google.protobuf.BoolValue\"\x00\x42]Z[github.com/nicholasmarco27/UTS_5027221042_Nicholas-Marco-Weinandra/master/common/taskmasterb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'taskmaster_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z[github.com/nicholasmarco27/UTS_5027221042_Nicholas-Marco-Weinandra/master/common/taskmaster'
  _globals['_TASK']._serialized_start=91
  _globals['_TASK']._serialized_end=178
  _globals['_TASKLIST']._serialized_start=180
  _globals['_TASKLIST']._serialized_end=220
  _globals['_TASKAPI']._serialized_start=223
  _globals['_TASKAPI']._serialized_end=461
# @@protoc_insertion_point(module_scope)
