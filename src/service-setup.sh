
service_postinst ()
{
    # https://help.synology.com/developer-guide/synology_package/script_env_var.html

    for file in config.json permissions.json; do
        if [ ! -e /var/packages/SimplePermissionManager/etc/${file} ]; then
            cp -f ${SYNOPKG_PKGDEST}/etc.defaults/${file} /var/packages/SimplePermissionManager/etc/${file}
        fi
    done

    if [ -e /usr/local/bin/spm-exec ]; then
        st=$(stat -c "%U %G %a" /usr/local/bin/spm-exec)
        if [ "$st" = "root root 6755" ]; then
            /usr/local/bin/spm-exec ${SYNOPKG_PKGDEST}/bin/spm-update
        fi
    fi
}
