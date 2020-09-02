<?php

use PHPUnit\Framework\TestCase;
use Netless\Token\Generate;

class TokenTest extends TestCase
{
    public function testGenerateSDKToken()
    {
        $netlessToken = new Generate;
        $sdkToken = $netlessToken->sdkToken("a", "b", 1000 * 60 * 10, array(
            "role" => Generate::AdminRole
        ));

        $this->assertStringContainsString("NETLESSSDK_", $sdkToken);
    }

    public function testGenerateRoomToken()
    {
        $netlessToken = new Generate;
        $sdkToken = $netlessToken->roomToken("a", "b", 1000 * 60 * 10, array(
            "role" => Generate::ReaderRole,
            "uuid" => "this is uuid",
        ));

        $this->assertStringContainsString("NETLESSROOM_", $sdkToken);
    }

    public function testGenerateTaskToken()
    {
        $netlessToken = new Generate;
        $sdkToken = $netlessToken->taskToken("a", "b", 1000 * 60 * 10, array(
            "role" => Generate::WriterRole,
            "uuid" => "this is uuid",
        ));

        $this->assertStringContainsString("NETLESSTASK_", $sdkToken);
    }
}
