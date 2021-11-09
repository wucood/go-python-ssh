#!/usr/bin/python
# -*- coding: UTF-8 -*-
import subprocess

shell = subprocess.Popen(["ping", "www.baidu.com", "-c" "5"])
shell.wait()
