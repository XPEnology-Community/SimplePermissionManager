#!/usr/bin/python

import json
import os
import pwd
import sys

permissions_file = '/var/packages/SimplePermissionManager/etc/permissions.json'

def get_package(package):
    pwall = pwd.getpwall()

    # for packages which not in /etc/passwd, default user is root
    name = 'root'
    uid = 0
    gid = 0

    for user in pwall:
        if user.pw_dir == '/var/packages/%s/home' % package:
            name = user.pw_name
            uid = user.pw_uid
            gid = user.pw_gid

    return {'name': name, 'uid': uid, 'gid': gid, 'package': package}

def update_package(package):
    permissions = {'users':{},'packages':{}}

    if os.path.exists(permissions_file):
        with open(permissions_file) as file:
            permissions = json.load(file)

    permissions['packages'][package['package']] = package

    file = open(permissions_file, 'w')
    file.write(json.dumps(permissions, indent=4))
    file.close()

if __name__ == '__main__':
    package = get_package("SimplePermissionManager")
    package['enabled'] = True
    update_package(package)