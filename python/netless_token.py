# NOTE: one can still import "private" functions in this module
#       by `from netless_token import _create_token` or access `netless_token._create_token`,
#       the `__all__` only takes effect when `from netless_token import *`
import urllib.parse
import uuid
import time
import json
import hmac
import hashlib
import base64

__all__ = ["sdk_token", "room_token", "task_token", "ADMIN", "WRITER", "READER"]

ADMIN = "0"
WRITER = "1"
READER = "2"


def _create_token(prefix: str, access_key: str, secret_access_key: str, lifespan: int, content: dict) -> str:
    data = {"ak": access_key, "nonce": str(uuid.uuid1())}
    data.update(content)

    if lifespan > 0:
        data["expireAt"] = str(int(round(time.time() * 1000)) + lifespan)

    info = json.dumps(data, sort_keys=True, separators=(",", ":"))
    digest = hmac.new(bytes(secret_access_key, encoding="utf-8"), digestmod="SHA256")
    digest.update(bytes(info, encoding="utf-8"))
    data["sig"] = digest.hexdigest()
    encoded = base64.urlsafe_b64encode(bytes(urllib.parse.urlencode(data), encoding="utf-8")).decode("utf-8").rstrip("=")
    return prefix + encoded


def sdk_token(access_key: str, secret_access_key: str, lifespan: int, role: str):
    return _create_token("NETLESSSDK_", access_key, secret_access_key, lifespan, content={'role': role})


def room_token(access_key: str, secret_access_key: str, lifespan: int, role: str, uuid: str):
    return _create_token("NETLESSROOM_", access_key, secret_access_key, lifespan, content={'role': role, 'uuid': uuid})


def task_token(access_key: str, secret_access_key: str, lifespan: int, role: str, uuid: str):
    return _create_token("NETLESSTASK_", access_key, secret_access_key, lifespan, content={'role': role, 'uuid': uuid})
