# Simple Permission Manager

## Generate Middle Public Key Signature

1. Export middle public key

```shell
gpg --output public.pgp --export 'Hello World <a@b.com>'
```

2. Sign by root key(Contact Jim)

```shell
gpg --output public.pgp.sig --detach-sign public.pgp
```

3. Save middle public key signature

Save `public.pgp` and `public.pgp.sig`

## Generate Binary Signature

1. Sign by middle key

```shell
gpg --output hello-world.sh.gpg.sig --detach-sign hello-world.sh
```

2. Save binary signature

File name: `hello-world.sh.sig`

```json
{
    "version": 1,
    "signature": "<base64 format read binary signature>",
    "publicKeys": [
        {
            "publicKey": "<base64 format middle public key>",
            "signature": "<base64 format middle public signature>"
        }
    ]
}
```

### Sign Script

```shell
file=hello-world.sh

pub_key=$(base64 -w 0 public.pgp)
pub_sig=$(base64 -w 0 public.pgp.sig)

sha256=$(sha256sum $file | awk '{print $1}')
gpg --output "$file".gpg.sig --detach-sign "$file"
sig=$(base64 -w 0 "$file".gpg.sig)
rm -f "$file".gpg.sig

cat << EOF > "$file".sig
{
    "version": 1,
    "sha256": "${sha256}",
    "signature": "${sig}",
    "publicKeys": [
        {
            "publicKey": "${pub_key}",
            "signature": "${pub_sig}"
        }
    ]
}
EOF
```

## Run Signatured Command

```
spm-exec /path/to/hello-world.sh
```