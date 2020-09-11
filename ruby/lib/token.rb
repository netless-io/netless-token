require "json"
require "uuidtools"
require "openssl"
require "base64"

module NetlessToken

  module ROLE
    ADMIN = "0"
    WRITER = "1"
    READER = "2"
  end
    
  module PREFIX
    SDK = "NETLESSSDK_"
    ROOM = "NETLESSROOM_"
    TASK = "NETLESSTASK_"
  end

  def self.to_base64(str)
    Base64.urlsafe_encode64(str, padding: false)
  end

  def self.create_token(prefix, access_key, secret_access_key, lifespan, content)
    object = {
      ak: access_key,
      nonce: UUIDTools::UUID.timestamp_create.to_s
    }
    object.merge!(content)
      
    if lifespan > 0
      object.store(:expireAt, ((Time.now.to_f * 1000).to_i + lifespan).to_s)
    end

    infomation = object.sort.to_h.to_json
    digest = OpenSSL::Digest.new('sha256')
    hmac = OpenSSL::HMAC.hexdigest(digest, secret_access_key, infomation)
    object.store(:sig, hmac)
    prefix + to_base64(URI.encode_www_form(object.sort.to_h))
  end

  def self.sdk_token *args
    create_token PREFIX::SDK, *args
  end

  def self.room_token *args
    create_token PREFIX::ROOM, *args
  end

  def self.task_token *args
    create_token PREFIX::TASK, *args
  end
end
