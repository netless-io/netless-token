import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.*;

public class Token {
    public enum TokenPrefix {
        SDK("NETLESSSDK_"),
        ROOM("NETLESSROOM_"),
        TASK("NETLESSTASK_");

        private String value;
        TokenPrefix(String name) {
            this.value = name;
        }
        public String getValue() {
            return value;
        }
    }

    /**
     * 数字越小，权限越大
     */
    public enum TokenRole {
        Admin("0"),
        Writer("1"),
        Reader("2");

        private String value;
        TokenRole(String name) {
            this.value = name;
        }
        public String getValue() {
            return value;
        }
    }

    public static String sdkToken(String accessKey, String secretAccessKey, long lifespan, Map<String, String> content) throws Exception {
        return createToken(TokenPrefix.SDK.getValue(), accessKey, secretAccessKey, lifespan, content);
    }

    public static String roomToken(String accessKey, String secretAccessKey, long lifespan, Map<String, String> content) throws Exception {
        return createToken(TokenPrefix.ROOM.getValue(), accessKey, secretAccessKey, lifespan, content);
    }

    public static String taskToken(String accessKey, String secretAccessKey, long lifespan, Map<String, String> content) throws Exception {
        return createToken(TokenPrefix.TASK.getValue(), accessKey, secretAccessKey, lifespan, content);
    }

    private static String createToken(String prefix, String accessKey, String secretAccessKey, long lifespan, Map<String, String> content) throws Exception {
        LinkedHashMap<String, String> map = new LinkedHashMap<>();
        map.putAll(content);
        map.put("ak", accessKey);
        map.put("nonce", UUID.randomUUID().toString());

        if (lifespan > 0) {
            map.put("expireAt", System.currentTimeMillis() + lifespan + "");
        }

        String information = toJson(sortMap(map));
        map.put("sig", createHmac(secretAccessKey, information));

        String query = sortAndStringifyMap(map);


        return prefix + stringToBase64(query);
    }

    private static LinkedHashMap<String, String> sortMap(Map<String, String> object) {
        List<String> keys = new ArrayList<>(object.keySet());
        keys.sort(null);

        LinkedHashMap<String, String> linkedHashMap = new LinkedHashMap<>();
        for (int i = 0; i < keys.size(); i++) {
            linkedHashMap.put(keys.get(i), object.get(keys.get(i)));
        }
        return linkedHashMap;
    }

    /**
     * 因程序本身的 map 只有一层，而非嵌套 map，所以此方法没有实现嵌套 map 转 string
     * 可自行替换为其他 json stringify 实现
     */
    private static String toJson(LinkedHashMap<String, String> map) {
        Iterator<Map.Entry<String, String>> iterator= map.entrySet().iterator();

        List<String> result = new ArrayList<>();
        while(iterator.hasNext()) {
            Map.Entry<String, String> entry = iterator.next();
            String value;
            if (entry.getValue() == null) {
                value = "null";
            } else {
                value = entry.getValue();
            }
            result.add("\"" + entry.getKey() + "\"" + ":" + "\"" + value + "\"");
        }
        return "{" + String.join(",", result) + "}";
    }

    private static String createHmac(String key, String data) throws Exception {
        Mac sha256_HMAC = Mac.getInstance("HmacSHA256");
        SecretKeySpec secret_key = new SecretKeySpec(key.getBytes("UTF-8"), "HmacSHA256");
        sha256_HMAC.init(secret_key);

        return byteArrayToHexString(sha256_HMAC.doFinal(data.getBytes("UTF-8")));
    }

    private static String sortAndStringifyMap(Map<String, String> object) {
        List<String> keys = new ArrayList<>(object.keySet());
        keys.sort(null);

        List<String> kvStrings = new ArrayList<>();
        for (int i = 0; i < keys.size(); i++) {
            if (object.get(keys.get(i)) == null) {
                continue;
            } else {
                kvStrings.add(encodeURIComponent(keys.get(i)) + "=" + encodeURIComponent(object.get(keys.get(i))));
            }
        }
        return String.join("&", kvStrings);
    }

    private static String stringToBase64(String str) throws UnsupportedEncodingException {
        return Base64.getEncoder().encodeToString(str.getBytes("utf-8")).replace("+", "-").replace("/", "_").replaceAll("=+$", "");
    }

    private static String byteArrayToHexString(byte[] b) {
        StringBuilder hs = new StringBuilder();
        String stmp;
        for (int n = 0; b!=null && n < b.length; n++) {
            stmp = Integer.toHexString(b[n] & 0XFF);
            if (stmp.length() == 1)
                hs.append('0');
            hs.append(stmp);
        }
        return hs.toString().toLowerCase();
    }

    /**
     * encodeURIComponent 基于 url 编码对字符串进行编码
     * 最终实现和 JavaScript 中的 encodeURIComponent 一致
     *
     * https://stackoverflow.com/questions/607176/java-equivalent-to-javascripts-encodeuricomponent-that-produces-identical-outpu
     */
    private static String encodeURIComponent(String s) {
        String result = null;

        try {
            result = URLEncoder.encode(s, "UTF-8")
                    .replaceAll("\\+", "%20")
                    .replaceAll("%21", "!")
                    .replaceAll("%27", "'")
                    .replaceAll("%28", "(")
                    .replaceAll("%29", ")")
                    .replaceAll("%7E", "~");
        }
        // This exception should never occur.
        catch (UnsupportedEncodingException e) {
            result = s;
        }

        return result;
    }
}
