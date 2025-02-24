# frozen_string_literal: true

require 'base64'
require 'octokit'

class GithubClient
  attr_accessor :client

  def initialize(token)
    ::Octokit.configure do |c|
      c.access_token = token
    end

    @client = ::Octokit::Client.new

    puts client.inspect

    self
  end

  def contents(repo, path)
    _contents = client.contents(repo, path: path)

    Base64.decode64(_contents.content)
  end
end
