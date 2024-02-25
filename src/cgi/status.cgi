#!/usr/bin/python

import json
import os
import stat
import io
import configparser
import subprocess

from pathlib import Path

exec_path = "/usr/local/bin/spm-exec"

if __name__ == '__main__':
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user) > 0:
        acvite = True

        if not os.path.exists(exec_path):
            acvite = False
        else:
            file_owner = Path(exec_path).owner()
            file_stat = os.stat(exec_path)

            # 1. check owner
            if file_owner != 'root':
                acvite = False

            # 2. check mode
            file_mode = stat.S_IMODE(file_stat.st_mode)
            if not (stat.S_ISUID & file_mode):
                acvite = False

            if not (stat.S_IRWXU & file_mode):
                acvite = False

            # 3. check verify output
            if acvite is True:
                desire = "Simple Permission Manager is awesome"
                real = subprocess.run(["/usr/local/bin/spm-exec", "echo", desire], stdout=subprocess.PIPE, check=False).stdout.decode()
                if desire not in real:
                    acvite = False

        version_path = '/etc/VERSION'
        version_str = '[root]\n' + open(version_path, 'r').read()
        fp = io.StringIO(version_str)
        config = configparser.RawConfigParser()
        config.read_file(fp)

        major = config['root']['major'].strip('"')
        minor = config['root']['minor'].strip('"')
        base = config['root']['base'].strip('"')

        print(json.dumps(
            {
                'active': acvite,
                'major': major,
                'minor': minor,
                'base': base,
            }, indent=4))
    # reject in case of no authentication
    else:
        print ("Security : no user authenticated")
