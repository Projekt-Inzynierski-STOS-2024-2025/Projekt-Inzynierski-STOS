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
FILE_UPLOAD_ERROR: ErrorType
FILE_DOWNLOAD_ERROR: ErrorType
WORKER_OPERATION_ERROR: ErrorType
EVALUATOR_ERROR: ErrorType

class File(_message.Message):
    __slots__ = ("name", "data")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    name: str
    data: bytes
    def __init__(self, name: _Optional[str] = ..., data: _Optional[bytes] = ...) -> None: ...

class Files(_message.Message):
    __slots__ = ("content", "uuid")
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    UUID_FIELD_NUMBER: _ClassVar[int]
    content: _containers.RepeatedCompositeFieldContainer[File]
    uuid: str
    def __init__(self, content: _Optional[_Iterable[_Union[File, _Mapping]]] = ..., uuid: _Optional[str] = ...) -> None: ...

class FileInfo(_message.Message):
    __slots__ = ("uuid", "directory_id")
    UUID_FIELD_NUMBER: _ClassVar[int]
    DIRECTORY_ID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    directory_id: int
    def __init__(self, uuid: _Optional[str] = ..., directory_id: _Optional[int] = ...) -> None: ...

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
    __slots__ = ("task_id", "student_id", "directory_id", "files_hash")
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    DIRECTORY_ID_FIELD_NUMBER: _ClassVar[int]
    FILES_HASH_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    student_id: str
    directory_id: int
    files_hash: bytes
    def __init__(self, task_id: _Optional[str] = ..., student_id: _Optional[str] = ..., directory_id: _Optional[int] = ..., files_hash: _Optional[bytes] = ...) -> None: ...
