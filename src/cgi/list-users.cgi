#!/usr/bin/python

import json
import os
import pwd

from urllib.parse import parse_qs

def list_users():
    pwall = pwd.getpwall()
    users = []
    for user in pwall:
        users.append({'name': user.pw_name, 'uid': user.pw_uid, 'gid': user.pw_gid})

    users = sorted(users, key=lambda x: x['uid'])
    for i in range(len(users)):
        users[i]['id'] = i+1

    return users

def print_users():
    offset = 0
    limit = -1

    if "QUERY_STRING" in os.environ:
        query = os.environ['QUERY_STRING']
        pairs = parse_qs(query)
        if 'offset' in pairs:
            offset = int(pairs['offset'][0])
        if 'limit' in pairs:
            limit = int(pairs['limit'][0])

    all_users = list_users()
    users = []

    if limit == -1:
        users = all_users
    else:
        users = all_users[offset:offset+limit]

    output = {
        'result': users,
        'success': True,
        'total': len(all_users)
    }
    print(json.dumps(output, indent=4))

if __name__ == '__main__':
    print_users()
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user)>0:
        print_users()
    # reject in case of no authentication
    else:
        print ("Security : no user authenticated")
