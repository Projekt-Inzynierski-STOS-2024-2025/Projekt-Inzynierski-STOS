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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0emessages.proto\x12\rstos.messages\"\"\n\x04\x46ile\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"U\n\x05\x45rror\x12\x0c\n\x04time\x18\x01 \x01(\t\x12\'\n\x05\x65rror\x18\x02 \x01(\x0e\x32\x18.stos.messages.ErrorType\x12\x15\n\rerror_message\x18\x03 \x01(\t\"j\n\x0cTaskDispatch\x12\x0f\n\x07task_id\x18\x01 \x01(\t\x12\x12\n\nstudent_id\x18\x02 \x01(\t\x12\x12\n\nfiles_hash\x18\x03 \x01(\x0c\x12!\n\x04\x64\x61ta\x18\x04 \x03(\x0b\x32\x13.stos.messages.File\"O\n\x08LogEvent\x12\x0c\n\x04time\x18\x01 \x01(\t\x12$\n\x04type\x18\x02 \x01(\x0e\x32\x16.stos.messages.LogType\x12\x0f\n\x07\x63ontent\x18\x03 \x01(\t\";\n\x05\x46iles\x12$\n\x07\x63ontent\x18\x01 \x03(\x0b\x32\x13.stos.messages.File\x12\x0c\n\x04uuid\x18\x02 \x01(\t*l\n\tErrorType\x12\x15\n\x11\x46ILE_UPLOAD_ERROR\x10\x00\x12\x17\n\x13\x46ILE_DOWNLOAD_ERROR\x10\x01\x12\x1a\n\x16WORKER_OPERATION_ERROR\x10\x02\x12\x13\n\x0f\x45VALUATOR_ERROR\x10\x03*+\n\x07LogType\x12\x08\n\x04INFO\x10\x00\x12\t\n\x05\x45RROR\x10\x01\x12\x0b\n\x07WARNING\x10\x02\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'messages_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_ERRORTYPE']._serialized_start=406
  _globals['_ERRORTYPE']._serialized_end=514
  _globals['_LOGTYPE']._serialized_start=516
  _globals['_LOGTYPE']._serialized_end=559
  _globals['_FILE']._serialized_start=33
  _globals['_FILE']._serialized_end=67
  _globals['_ERROR']._serialized_start=69
  _globals['_ERROR']._serialized_end=154
  _globals['_TASKDISPATCH']._serialized_start=156
  _globals['_TASKDISPATCH']._serialized_end=262
  _globals['_LOGEVENT']._serialized_start=264
  _globals['_LOGEVENT']._serialized_end=343
  _globals['_FILES']._serialized_start=345
  _globals['_FILES']._serialized_end=404
# @@protoc_insertion_point(module_scope)
