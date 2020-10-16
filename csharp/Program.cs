using System;
using Netless;

class Program
{
    static void Main(string[] args)
    {
        string token = NetlessToken.SdkToken("ak", "sk", 1000 * 60 * 10, new SdkContent(TokenRole.Admin));

        Console.WriteLine(token);
    }
}
