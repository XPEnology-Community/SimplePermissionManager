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
    permissions = {}

    if os.path.exists(permissions_file):
        with open(permissions_file) as file:
            permissions = json.load(file)

    permissions[package['package']] = package

    file = open(permissions_file, 'w')
    file.write(json.dumps(permissions, indent=4))
    file.close()

def process():
    content_len = int(os.environ.get('CONTENT_LENGTH', 0))
    # method = os.environ.get('REQUEST_METHOD', '')
    # query_string = os.environ.get('QUERY_STRING', '')
    # x_header = os.environ.get('HTTP_X_MARVIN_STATUS', '')
    body = sys.stdin.read(content_len)
    request = json.loads(body)

    package = get_package(request['package'])
    package['enabled'] = request['enabled']

    update_package(package)

if __name__ == '__main__':
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user)>0:
        process()
    # reject in case of no authentication
    else:
        print("Security : no user authenticated")