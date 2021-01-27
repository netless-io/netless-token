// you should also add this line:
// <script src="https://cdn.jsdelivr.net/npm/uuid/dist/umd/uuidv1.min.js"></script>
// which imports `uuidv1()` function

/**
 * Usage:
 * ```js
 * await NetlessToken.sdkToken("ak", "sk", 60 * 1000,
 *     { role: NetlessToken.TokenRole.Admin });
 * await NetlessToken.roomToken("ak", "sk", 60 * 1000,
 *     { role: NetlessToken.TokenRole.Writer, uuid: <room uuid> });
 * await NetlessToken.taskToken("ak", "sk", 60 * 1000,
 *     { role: NetlessToken.TokenRole.Reader, uuid: <task uuid> });
 * ```
 */
var NetlessToken = (function (exports, uuidv1) {
    const TokenRole = {
        Admin: "0",
        Writer: "1",
        Reader: "2",
    };
    exports.TokenRole = TokenRole;

    const TokenPrefix = {
        SDK: "NETLESSSDK_",
        ROOM: "NETLESSROOM_",
        TASK: "NETLESSTASK_",
    };
    exports.TokenPrefix = TokenPrefix;

    function formatJSON(obj) {
        const keys = Object.keys(obj).sort();
        const target = {};

        for (const key of keys) {
            target[key] = String(obj[key]);
        }

        return target;
    }

    function stringify(obj) {
        return Object.keys(obj)
            .map((key) => {
                const value = obj[key];
                if (value === undefined) return "";
                if (value === null) return "null";
                return `${encodeURIComponent(key)}=${encodeURIComponent(
                    value
                )}`;
            })
            .join("&");
    }

    const encoder = new TextEncoder("utf-8");

    async function createHmac(hash, secretKey) {
        if (hash === "sha256") hash = "SHA-256";
        const key = await crypto.subtle.importKey(
            "raw",
            encoder.encode(secretKey),
            {
                name: "HMAC",
                hash: hash,
            },
            true,
            ["sign", "verify"]
        );
        return {
            data: "",
            update(str) {
                this.data += str;
                return this;
            },
            async digest() {
                const data = encoder.encode(this.data);
                const sig = await crypto.subtle.sign("HMAC", key, data);
                const b = Array.from(new Uint8Array(sig));
                return b.map((x) => x.toString(16).padStart(2, "0")).join("");
            },
        };
    }

    function createToken(prefix) {
        return async (accessKey, secretAccessKey, lifespan, content) => {
            const object = {
                ...content,
                ak: accessKey,
                nonce: uuidv1(),
            };

            if (lifespan > 0) {
                object.expireAt = `${Date.now() + lifespan}`;
            }

            const information = JSON.stringify(formatJSON(object));
            const hmac = await createHmac("sha256", secretAccessKey);
            object.sig = await hmac.update(information).digest("hex");
            const query = stringify(formatJSON(object));

            return (
                prefix +
                btoa(query)
                    .replace(/\+/g, "-")
                    .replace(/\//g, "_")
                    .replace(/=+$/, "")
            );
        };
    }

    exports.sdkToken = createToken(TokenPrefix.SDK);
    exports.roomToken = createToken(TokenPrefix.ROOM);
    exports.taskToken = createToken(TokenPrefix.TASK);

    return exports;
})({}, uuidv1);
