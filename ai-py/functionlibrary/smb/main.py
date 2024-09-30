# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2024-09-30 11:19:58
@Author : pan
"""
from smb.SMBConnection import SMBConnection  # pip install pysmb


def list_smb_files(server_name, shared_name, username, password):
    conn = SMBConnection(username, password, "myclient", server_name, use_mtlm_v1=True)
    conn.connect(server_name, 139)

    files = conn.listPath(shared_name, "path/")
    for file in files:
        print(file)

    conn.close()


if __name__ == "__main__":
    pass
