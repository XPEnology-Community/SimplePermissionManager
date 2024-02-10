#!/usr/bin/python

import json
import os

from pathlib import Path

def find_owner(file_path):
    file = Path(file_path)
    return file.owner()

if __name__ == '__main__':
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user)>0:
        file_path = "/var/packages/SimplePermissionManager/target/bin/spm-exec"
        owner = find_owner(file_path)
        print(json.dumps({ 'active': owner == 'root'}, indent=4))
    # reject in case of no authentication
    else:
        print ("Security : no user authenticated")
