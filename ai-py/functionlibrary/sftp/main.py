# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2024-09-30 11:15:01
@Author : pan
"""
import paramiko
from loguru import logger


def SFTPConn():
    try:
        transport = paramiko.Transport("192.168.100.102", int("22"))
        transport.connect(username="user", password="password")
        sftp = paramiko.SFTPClient.from_transport(transport)
        return transport, sftp
    except Exception as err:
        logger.error(f"sftp conn error: {err}")
        return None, None


def listRmoteFiles():
    _, sftp = SFTPConn()
    files = sftp.listdir("remote_files_path")
    for file in files:
        print(file)


if __name__ == "__main__":
    pass
