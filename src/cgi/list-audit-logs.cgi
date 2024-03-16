#!/usr/bin/python

import io
import os
import json
import sqlite3

from urllib.parse import parse_qs

audit_db_file = '/var/packages/SimplePermissionManager/var/audit.db'

def print_audit_logs():
    offset = 0
    limit = 100

    if "QUERY_STRING" in os.environ:
        query = os.environ['QUERY_STRING']
        pairs = parse_qs(query)
        if 'offset' in pairs:
            offset = int(pairs['offset'][0])
        if 'limit' in pairs:
            limit = int(pairs['limit'][0])

    conn = sqlite3.connect(audit_db_file)
    conn.row_factory = sqlite3.Row
    db = conn.cursor()

    rows = db.execute('SELECT * FROM audits ORDER BY id DESC LIMIT ? OFFSET ?', (limit, offset)).fetchall()
    total = db.execute("SELECT COUNT(*) FROM audits").fetchone()[0]

    conn.commit()
    conn.close()

    output = {
        'result': [dict(ix) for ix in rows],
        'success': True,
        'total': total
    }
    print(json.dumps(output, indent=4))

if __name__ == '__main__':
    print("Content-type: application/json\n")
    f = os.popen('/usr/syno/synoman/webman/modules/authenticate.cgi','r')
    user = f.read()

    # check user is authenticated
    if len(user)>0:
        print_audit_logs()
    # reject in case of no authentication
    else:
        print ("Security : no user authenticated")
