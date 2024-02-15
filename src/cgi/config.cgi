#!/usr/bin/python

import json
import os
import stat

from pathlib import Path

config_file = "/var/packages/SimplePermissionManager/etc/config.json"

if __name__ == '__main__':
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user) > 0:
        config = {}

        if os.path.exists(config_file):
            with open(config_file) as file:
                config = json.load(file)

        print(json.dumps(config, indent=4))
    # reject in case of no authentication
    else:
        print ("Security : no user authenticated")
