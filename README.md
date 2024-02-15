# Simple Permission Manager

## Generate Middle Public Key Signature

1. Export middle public key

```shell
gpg --output public.pgp --export 'Hello World <a@b.com>'
```

2. Sign by root key

```shell
gpg --detach-sign public.pgp
```

3. Save middle public key signature

Save `public.pgp` and `public.pgp.pgp`

## Generate Binary Signature

1. Sign by middle key

```shell
gpg --detach-sign hello-world.sh
```

2. Save binary signature

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

## Run Signatured Command

```
spm-exec /path/to/phello-world.sh
```