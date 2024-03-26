package com.stos.worker;

import com.google.protobuf.ByteString;
import com.google.protobuf.InvalidProtocolBufferException;
import com.stos.worker.proto.Messages;

import java.util.ArrayList;
import java.util.Random;

public class Helpers {

    public byte[] generateTaskDispatch(){
        Random random = new Random();
        Messages.File file = Messages.File.newBuilder()
                .setName("testFile")
                .setData(ByteString.copyFrom(new byte[] { 0x01, 0x02, 0x03 }))
                .build();
        Messages.TaskDispatch taskDispatch = Messages.TaskDispatch.newBuilder()
                .setTaskId(Integer.toString(random.nextInt(1000)))
                .setStudentId(Integer.toString(random.nextInt(1000)))
                .setFilesHash(ByteString.copyFrom(new byte[] { 0x01, 0x02, 0x03 }))
                .addData(file)
                .build();
        return taskDispatch.toByteArray();
    }

    public Messages.TaskDispatch deserializeTaskDispatch(byte[] serializedMessage) throws InvalidProtocolBufferException {
        return Messages.TaskDispatch.parseFrom(serializedMessage);
    }
}
