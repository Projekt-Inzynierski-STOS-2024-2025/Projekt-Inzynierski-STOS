from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ErrorType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    FILE_UPLOAD_ERROR: _ClassVar[ErrorType]
    FILE_DOWNLOAD_ERROR: _ClassVar[ErrorType]
    WORKER_OPERATION_ERROR: _ClassVar[ErrorType]
    EVALUATOR_ERROR: _ClassVar[ErrorType]

class LogType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INFO: _ClassVar[LogType]
    ERROR: _ClassVar[LogType]
    WARNING: _ClassVar[LogType]
FILE_UPLOAD_ERROR: ErrorType
FILE_DOWNLOAD_ERROR: ErrorType
WORKER_OPERATION_ERROR: ErrorType
EVALUATOR_ERROR: ErrorType
INFO: LogType
ERROR: LogType
WARNING: LogType

class File(_message.Message):
    __slots__ = ("name", "data")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    name: str
    data: bytes
    def __init__(self, name: _Optional[str] = ..., data: _Optional[bytes] = ...) -> None: ...

class Error(_message.Message):
    __slots__ = ("time", "error", "error_message")
    TIME_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    ERROR_MESSAGE_FIELD_NUMBER: _ClassVar[int]
    time: str
    error: ErrorType
    error_message: str
    def __init__(self, time: _Optional[str] = ..., error: _Optional[_Union[ErrorType, str]] = ..., error_message: _Optional[str] = ...) -> None: ...

class TaskDispatch(_message.Message):
    __slots__ = ("task_id", "student_id", "files_hash", "data")
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    FILES_HASH_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    student_id: str
    files_hash: bytes
    data: _containers.RepeatedCompositeFieldContainer[File]
    def __init__(self, task_id: _Optional[str] = ..., student_id: _Optional[str] = ..., files_hash: _Optional[bytes] = ..., data: _Optional[_Iterable[_Union[File, _Mapping]]] = ...) -> None: ...

class LogEvent(_message.Message):
    __slots__ = ("time", "type", "content")
    TIME_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    time: str
    type: LogType
    content: str
    def __init__(self, time: _Optional[str] = ..., type: _Optional[_Union[LogType, str]] = ..., content: _Optional[str] = ...) -> None: ...

class Files(_message.Message):
    __slots__ = ("content", "uuid")
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    content: _containers.RepeatedCompositeFieldContainer[File]
    uuid: str
    def __init__(self, content: _Optional[_Iterable[_Union[File, _Mapping]]] = ..., uuid: _Optional[str] = ...) -> None: ...
