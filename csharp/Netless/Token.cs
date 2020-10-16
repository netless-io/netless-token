using System;
using System.Collections.Generic;
using System.Security.Cryptography;
using System.Text;
using Newtonsoft.Json;

namespace Netless
{
    public static class TokenRole
    {
        public static readonly string Admin = "0";
        public static readonly string Writter = "1";
        public static readonly string Reader = "2";
    }

    public static class TokenPrefix
    {
        public static readonly string SDK = "NETLESSSDK_";
        public static readonly string ROOM = "NETLESSROOM_";
        public static readonly string TASK = "NETLESSTASK_";
    }

    public class SdkContent
    {
        public string role;

        public SdkContent(string role)
        {
            this.role = role;
        }
    }

    public class RoomContent
    {
        public string role;
        public string uuid;

        public RoomContent(string role, string uuid)
        {
            this.role = role;
            this.uuid = uuid;
        }
    }

    public class TaskContent
    {
        public string role;
        public string uuid;

        public TaskContent(string role, string uuid)
        {
            this.role = role;
            this.uuid = uuid;
        }
    }

    public class NetlessToken
    {
        private static readonly DateTime UnixEpoch = new DateTime(1970, 1, 1, 0, 0, 0, DateTimeKind.Utc);

        private static string CreateToken(string prefix, string accessKey, string secretAccessKey, long lifespan, Dictionary<string, string> content)
        {
            Dictionary<string, string> m = new Dictionary<string, string>
            {
                { "ak", accessKey },
                { "nonce", Guid.NewGuid().ToString() }
            };
            MergeDictionary(content, m);

            if (lifespan > 0)
            {
                m.Add("expireAt", Convert.ToString(Convert.ToInt64((DateTime.Now - UnixEpoch).TotalMilliseconds + lifespan)));
            }

            string infomation = JsonConvert.SerializeObject(m);
            using (HMACSHA256 digest = new HMACSHA256(Encoding.ASCII.GetBytes(secretAccessKey)))
            {
                byte[] hmac = digest.ComputeHash(Encoding.ASCII.GetBytes(infomation));
                m.Add("sig", ByteArrayToString(hmac));
            }
            return prefix + ToBase64(Querify(m));
        }

        private static string ToBase64(string m)
        {
            return Convert.ToBase64String(Encoding.ASCII.GetBytes(m)).TrimEnd('=').Replace('+', '-').Replace('/', '_');
        }

        private static string ByteArrayToString(byte[] a)
        {
            StringBuilder hex = new StringBuilder(a.Length * 2);
            foreach (byte b in a)
            {
                hex.AppendFormat("{0:x2}", b);
            }
            return hex.ToString();
        }

        private static void MergeDictionary(Dictionary<string, string> from, Dictionary<string, string> to)
        {
            foreach (var entry in from)
            {
                to.Add(entry.Key, entry.Value);
            }
        }

        private static string Querify(Dictionary<string, string> a)
        {
            List<string> parts = new List<string>();
            foreach (var e in a)
            {
                parts.Add($"{Uri.EscapeDataString(e.Key)}={Uri.EscapeDataString(e.Value)}");
            }
            return string.Join('&', parts);
        }

        public static string SdkToken(string accessKey, string secretAccessKey, long lifespan, SdkContent content)
        {
            Dictionary<string, string> m = new Dictionary<string, string>
            {
                { "role", content.role }
            };
            return CreateToken(TokenPrefix.SDK, accessKey, secretAccessKey, lifespan, m);
        }

        public static string RoomToken(string accessKey, string secretAccessKey, long lifespan, RoomContent content)
        {
            Dictionary<string, string> m = new Dictionary<string, string>
            {
                { "role", content.role },
                { "uuid", content.uuid }
            };
            return CreateToken(TokenPrefix.ROOM, accessKey, secretAccessKey, lifespan, m);
        }

        public static string TaskToken(string accessKey, string secretAccessKey, long lifespan, TaskContent content)
        {
            Dictionary<string, string> m = new Dictionary<string, string>
            {
                { "role", content.role },
                { "uuid", content.uuid }
            };
            return CreateToken(TokenPrefix.TASK, accessKey, secretAccessKey, lifespan, m);
        }
    }
}