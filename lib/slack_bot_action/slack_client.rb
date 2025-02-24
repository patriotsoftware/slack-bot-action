# frozen_string_literal: true

require 'slack-ruby-client'

class SlackClient
  attr_accessor :client

  def initialize(token)
    ::Slack.configure do |config|
      config.token = token
      raise 'Missing ENV[SLACK_API_TOKEN]!' unless config.token
    end

    @client = ::Slack::Web::Client.new
    client.auth_test

    self
  end

  def post(message, channel)
    client.chat_postMessage(channel: channel, text: message)
  end
end
