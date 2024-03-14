import paramiko


def exec_shell(ssh, command):
    res_info = ""
    try:
        stdin, stdout, stderr = ssh.exec_command(command)
        for line in stdout.readlines():
            res_info += line
        print(f"command:{command} res_info:\n {res_info}")
    except Exception as e:
        print(f"command:{command} 执行失败！{e}")


def tinyproxy_deploy(ip, password):
    print(f"ip:{ip} action!")
    # 创建ssh对象
    ssh = paramiko.SSHClient()
    # 允许链接不在know_hosts文件中的主机
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    # 连接服务器
    ssh.connect(hostname=ip, username="root", password=password)

    # 执行脚本1
    command_01 = "mkdir tinyproxy"
    exec_shell(ssh, command_01)

    # sftp连接
    try:
        transport = paramiko.Transport((ip, 22))
        transport.connect(username="root", password=password)
        sftp = paramiko.SFTPClient.from_transport(transport)
    except Exception as e:
        print(f"连接失败，密码不正确？\n{e}")
        return False
    file_name = "tinyproxy.sh"
    # 上传脚本
    sftp.put(f"/Users/mac/Desktop/tmp/{file_name}", f"/root/tinyproxy/{file_name}")
    # 关闭上传连接
    transport.close()

    # 执行脚本2
    command_02 = f"bash /root/tinyproxy/{file_name}"
    exec_shell(ssh, command_02)


if __name__ == "__main__":
    server_list = [
        {
            "ip": "localhost",
            "password": "password",
        },
        {
            "ip": "localhost",
            "password": "password",
        },
    ]
    for server in server_list:
        tinyproxy_deploy(ip=server["ip"], password=server["password"])
