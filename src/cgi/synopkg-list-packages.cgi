#!/usr/bin/python

import io
import json
import os
import configparser

from urllib.parse import parse_qs

permissions_file = '/var/packages/SimplePermissionManager/etc/permissions.json'

def chunks(lst, n):
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]

def list_packages():
    packages = []
    permissions = {}

    if os.path.exists(permissions_file):
        with open(permissions_file) as file:
            permissions = json.load(file)

    directories = [d for d in os.listdir('/var/packages') if os.path.isdir('/var/packages/'+d)]
    for dir in directories:
        package = ''
        version = ''
        description = ''

        ini_path = '/var/packages/%s/INFO' % dir
        ini_str = '[root]\n' + open(ini_path, 'r').read()
        ini_fp = io.StringIO(ini_str)
        config = configparser.RawConfigParser()
        config.read_file(ini_fp)

        package = config['root']['package'].strip('"')
        version = config['root']['version'].strip('"')
        if 'description' in config['root']:
            description = config['root']['description'].strip('"')

        enabled = False
        if package in permissions and 'enabled' in permissions[package]:
            enabled = permissions[package]['enabled']

        packages.append({
            'package': package,
            'version': version,
            'enabled': enabled,
            'description': description
        })
    packages = sorted(packages, key=lambda x: x['package'])
    for i in range(len(packages)):
        packages[i]['id'] = i+1

    return packages

def print_packages():
    offset = 0
    limit = -1

    if "QUERY_STRING" in os.environ:
        query = os.environ['QUERY_STRING']
        pairs = parse_qs(query)
        if 'offset' in pairs:
            offset = int(pairs['offset'][0])
        if 'limit' in pairs:
            limit = int(pairs['limit'][0])

    all_packages = list_packages()
    packages = []

    if limit == -1:
        packages = all_packages
    else:
        packages = all_packages[offset:offset+limit]

    output = {
        'result': packages,
        'success': True,
        'total': len(all_packages)
    }
    print(json.dumps(output, indent=4))

if __name__ == '__main__':
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user)>0:
        print_packages()
    # reject in case of no authentication
    else:
        print ("Security : no user authenticated")
