import messages_pb2


def file_message(name: str, data: bytes):
    message = messages_pb2.File()
    message.name = name
    message.data = data

    return message


def files_message(files, uuid: str):
    message = messages_pb2.Files()
    message.content = files
    message.uuid = uuid

    return message


def files_info_message(uuid: str, directory_id: int):
    message = messages_pb2.FileInfo()
    message.uuid = uuid
    message.directory_id = directory_id

    return message


def error_message(time: str, error: str, error_message: str):
    message = messages_pb2.Error()
    message.time = time
    message.error = error
    message.error_message = error_message

    return message


def task_dispatch_message(task_id: str, student_id: str, directory_id: int, files_hash: bytes):
    message = messages_pb2.TaskDispatch()
    message.task_id = task_id
    message.student_id = student_id
    message.directory_id = directory_id
    message.files_hash = files_hash

    return message
