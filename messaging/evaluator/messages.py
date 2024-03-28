import messages_pb2


def file_message(data: dict):
    message = messages_pb2.File()
    message.name = data['name']
    message.data = str(data['data']).encode()

    return message


def files_message(data: dict):
    message = messages_pb2.Files()
    message.content = data['files']
    message.uuid = data['uuid']

    return message


def files_info_message(data: dict):
    message = messages_pb2.FileInfo()
    message.uuid = data['uuid']
    message.directory_id = data['directory_id']

    return message


def error_message(data: dict):
    message = messages_pb2.Error()
    message.time = data['time']
    message.error = data['error']
    message.error_message = data['error_message']

    return message


def task_dispatch_message(data: dict):
    message = messages_pb2.TaskDispatch()
    message.task_id = data['task_id']
    message.student_id = data['student_id']
    message.directory_id = data['directory_id']
    message.files_hash = str(data['files_hash']).encode()

    return message
