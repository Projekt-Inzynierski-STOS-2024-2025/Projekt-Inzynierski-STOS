# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: messages.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0emessages.proto\x12\rstos.messages\"\"\n\x04\x46ile\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\";\n\x05\x46iles\x12$\n\x07\x63ontent\x18\x01 \x03(\x0b\x32\x13.stos.messages.File\x12\x0c\n\x04uuid\x18\x02 \x01(\t\".\n\x08\x46ileInfo\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x14\n\x0c\x64irectory_id\x18\x02 \x01(\x03\"U\n\x05\x45rror\x12\x0c\n\x04time\x18\x01 \x01(\t\x12\'\n\x05\x65rror\x18\x02 \x01(\x0e\x32\x18.stos.messages.ErrorType\x12\x15\n\rerror_message\x18\x03 \x01(\t\"]\n\x0cTaskDispatch\x12\x0f\n\x07task_id\x18\x01 \x01(\t\x12\x12\n\nstudent_id\x18\x02 \x01(\t\x12\x14\n\x0c\x64irectory_id\x18\x03 \x01(\x03\x12\x12\n\nfiles_hash\x18\x04 \x01(\x0c*l\n\tErrorType\x12\x15\n\x11\x46ILE_UPLOAD_ERROR\x10\x00\x12\x17\n\x13\x46ILE_DOWNLOAD_ERROR\x10\x01\x12\x1a\n\x16WORKER_OPERATION_ERROR\x10\x02\x12\x13\n\x0f\x45VALUATOR_ERROR\x10\x03\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'messages_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_ERRORTYPE']._serialized_start=360
  _globals['_ERRORTYPE']._serialized_end=468
  _globals['_FILE']._serialized_start=33
  _globals['_FILE']._serialized_end=67
  _globals['_FILES']._serialized_start=69
  _globals['_FILES']._serialized_end=128
  _globals['_FILEINFO']._serialized_start=130
  _globals['_FILEINFO']._serialized_end=176
  _globals['_ERROR']._serialized_start=178
  _globals['_ERROR']._serialized_end=263
  _globals['_TASKDISPATCH']._serialized_start=265
  _globals['_TASKDISPATCH']._serialized_end=358
# @@protoc_insertion_point(module_scope)
