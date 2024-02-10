#!/usr/bin/python

import json
import os
import stat

from pathlib import Path

if __name__ == '__main__':
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user) > 0:
        file_path = "/var/packages/SimplePermissionManager/target/bin/spm-exec"
        file_owner = Path(file_path).owner()
        file_stat = os.stat(file_path)

        acvite = True
        if file_owner != 'root':
            acvite = False

        file_mode = stat.S_IMODE(file_stat.st_mode)
        if not (stat.S_ISUID & file_mode):
            acvite = False

        if not (stat.S_IRWXU & file_mode):
            acvite = False

        print(json.dumps({ 'active': acvite }, indent=4))
    # reject in case of no authentication
    else:
        print ("Security : no user authenticated")
