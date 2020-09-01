import { createHmac } from "crypto";
import { v1 as uuidv1 } from "uuid";

export enum TokenRole {
    // 数字越小，权限越大
    Admin = "0",
    Writer = "1",
    Reader = "2",
}

export enum TokenPrefix {
    SDK = "NETLESSSDK_",
    ROOM = "NETLESSROOM_",
    TASK = "NETLESSTASK_",
}

// buffer 转 base64，且格式化字符
const bufferToBase64 = (buffer: Buffer): string => {
    return buffer.toString("base64").replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "");
};

// 排序，以确保最终生成的 string 与顺序无关
// keys 的顺序不应该影响 hash 的值
const formatJSON = <T extends StrAndIntByObj>(object: T): StrByObj => {
    const keys = Object.keys(object).sort();
    const target: StrByObj = {};

    for (const key of keys) {
        target[key] = String(object[key]);
    }
    return target;
};

// 序列化对象
const stringify = (object: StrByObj): string => {
    return Object.keys(object)
        .map(key => {
            const value = object[key];

            if (value === undefined) {
                return "";
            }

            if (value === null) {
                return "null";
            }

            return `${encodeURIComponent(key)}=${encodeURIComponent(value)}`;
        })
        .join("&");
};

// 根据相关 prefix 生成相应的token
const createToken = <T extends {}>(
    prefix: TokenPrefix,
): ((accessKey: string, secretAccessKey: string, lifespan: number, content: T) => string) => {
    return (accessKey: string, secretAccessKey: string, lifespan: number, content: T) => {
        const object: StrAndIntByObj = {
            ...content,
            ak: accessKey,
            nonce: uuidv1(),
        };

        if (lifespan > 0) {
            object.expireAt = `${Date.now() + lifespan}`;
        }

        const information = JSON.stringify(formatJSON(object));
        const hmac = createHmac("sha256", secretAccessKey);
        object.sig = hmac.update(information).digest("hex");

        const query = stringify(formatJSON(object));
        const buffer = Buffer.from(query, "utf8");

        return prefix + bufferToBase64(buffer);
    };
};

// 生成 sdk token
export const SdkToken = createToken<SdkTokenTags>(TokenPrefix.SDK);

// 生成 room token
export const RoomToken = createToken<RoomTokenTags>(TokenPrefix.ROOM);

// 生成 task token
export const TaskToken = createToken<TaskTokenTags>(TokenPrefix.TASK);

export type SdkTokenTags = {
    readonly role?: TokenRole;
};

export type RoomTokenTags = {
    readonly uuid?: string;
    readonly role?: TokenRole;
};

export type TaskTokenTags = {
    readonly uuid?: string;
    readonly role?: TokenRole;
};

type StrAndIntByObj = Record<string, string | number>;
type StrByObj = Record<string, string>;
