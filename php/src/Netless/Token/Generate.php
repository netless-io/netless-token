<?php

namespace Netless\Token;

date_default_timezone_set("UTC");

use Ramsey\Uuid\Uuid;

class Generate
{
    // 数字越小，权限越大
    public const AdminRole = "0";
    public const WriterRole = "1";
    public const ReaderRole = "2";

    /**
     * 生成 sdk token
     * @param string $accessKey netless ak
     * @param string $secretAccessKey netless sk
     * @param int $lifespan 过期时长，为 0 则永不过期
     * @param array $content 额外补充信息
     * @return string
     */
    public function sdkToken(string $accessKey, string $secretAccessKey, int $lifespan, array $content): string
    {
        return $this->createToken("NETLESSSDK_")($accessKey, $secretAccessKey, $lifespan, $content);
    }

    /**
     * 生成 room token
     * @param string $accessKey netless ak
     * @param string $secretAccessKey netless sk
     * @param int $lifespan 过期时长，为 0 则永不过期
     * @param array $content 额外补充信息
     * @return string
     */
    public function roomToken(string $accessKey, string $secretAccessKey, int $lifespan, array $content): string
    {
        return $this->createToken("NETLESSROOM_")($accessKey, $secretAccessKey, $lifespan, $content);
    }

    /**
     * 生成 task token
     * @param string $accessKey netless ak
     * @param string $secretAccessKey netless sk
     * @param int $lifespan 过期时长，为 0 则永不过期
     * @param array $content 额外补充信息
     * @return string
     */
    public function taskToken(string $accessKey, string $secretAccessKey, int $lifespan, array $content): string
    {
        return $this->createToken("NETLESSTASK_")($accessKey, $secretAccessKey, $lifespan, $content);
    }

    /**
     * bufferToBase64 buffer 转 base64
     * 并格式化字符
     * @param string $str 需要转义的字符串
     * @return string
     */
    private function bufferToBase64(string $str): string
    {
        $result = base64_encode($str);
        $result = preg_replace("/\+/", "-", $result, -1);
        $result = preg_replace("/\//", "_", $result, -1);
        return preg_replace("/=+$/", "", $result, -1);
    }

    /**
     * encodeURIComponent 基于 url 编码对字符串进行编码
     * 最终实现和 JavaScript 中的 encodeURIComponent 一致
     * @see https://stackoverflow.com/a/1734255/6596777
     * @param string $str 需要转换字符
     * @return string
     */
    private function encodeURIComponent(string $str): string
    {
        $revert = array('%21'=>'!', '%2A'=>'*', '%27'=>"'", '%28'=>'(', '%29'=>')');
        return strtr(rawurlencode($str), $revert);
    }

    /**
     * stringify 序列化 array
     * @param array<string> $obj 需要转义的数组
     * @return string
     */
    private function stringify(array $obj): string
    {
        $result = array();
        foreach ($obj as $k => $v) {
            if ($v !== "") {
                array_push($result, $this->encodeURIComponent($k) . "=" . $this->encodeURIComponent($v));
            }
        }
        return join("&", $result);
    }

    /**
     * 根据 prefix 生成相应的 generate
     * @param string $prefix 必须为: NETLESSSDK_ / NETLESSROOM_ / NETLESSTASK_
     * @return \Closure
     */
    private function createToken(string $prefix): \Closure
    {
        return function (string $accessKey, string $secretAccessKey, int $lifespan, array $content) use ($prefix): string {
            $map = array_replace($content, array(
                "ak" => $accessKey,
                "nonce" => Uuid::uuid4()->toString()
            ));

            if ($lifespan > 0) {
                $map += array(
                    "expireAt" => strval((int)(microtime(true) * 1000) + $lifespan),
                );
            }
            ksort($map);

            $map += array(
                "sig" => hash_hmac("sha256", json_encode($map), $secretAccessKey)
            );
            ksort($map);

            return $prefix . $this->bufferToBase64($this->stringify($map));
        };
    }
}
