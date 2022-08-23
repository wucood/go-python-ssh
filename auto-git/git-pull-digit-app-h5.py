#!/usr/bin/python
# -*- coding: UTF-8 -*-
import sys
import paramiko


def main(username, hostname, port):
    # 连接远程服务器
    ssh = paramiko.SSHClient()
    # 允许连接不在know_hosts文件中的主机
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy)
    ssh.connect(hostname=hostname, port=port, username=username, password="")

    # 服务器上的执行命令
    cmd = "bash /server/scripts/git-pull-digit-app-h5.sh"
    sshCMD(ssh, cmd)
    ssh.close()
    print("任务已结束")


# 执行远程命令，实时打印命令结果
def sshCMD(ssh, cmd):
    # 执行远程服务器命令
    stdin, stdout, stderr = ssh.exec_command(cmd)
    # 实时获取远程命令结果
    while True:
        line = stdout.readline()
        print(line, end="")
        if not line:
            break


if __name__ == '__main__':
    userName = sys.argv[1]
    serverIP = sys.argv[2]
    sshPort = sys.argv[3]
    main(userName, serverIP, sshPort)
    print("已经运行！")
