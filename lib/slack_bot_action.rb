# frozen_string_literal: true

require 'json'
require 'bundler/setup'

require_relative 'slack_bot_action/github_client'
require_relative 'slack_bot_action/slack_client'

class SlackActionBot
  attr_accessor :github_client, :slack_client, :inputs, :mapping

  def initialize
    parse_inputs

    @github_client = GithubClient.new(inputs['github-token'])
    @slack_client = SlackClient.new(inputs['slack-token'])

    load_mapping
    parse_destination
    parse_message

    self
  end

  def load_mapping
    puts inputs.inspect

    raw_mapping = github_client.contents(
      inputs['username-mapping-repository'],
      inputs['username-mapping-path']
    )

    @mapping = YAML.safe_load(raw_mapping)
  end

  def parse_inputs
    @inputs = {}

    ENV.each do |key, value|
      next unless key.start_with?('INPUT_')

      inputs[key] = value

      key = key.split('INPUT_').last.downcase

      inputs[key] = value.is_a?(String) ? value.strip : value
    end

    inputs
  end

  def slack_id_for_github_user(github_username)
    mapping.dig(github_username)
  end

  # If the destination is set to committer determine which user to send to.
  # "Committer" is a bit of a misnomer, as it could be the user who triggered
  # a rebuild of a workflow (triggering_actor), the user who started the workflow (actor), or the
  # user who pushed the commit (head commit author), in that order. If the destination is not
  # set to committer then it is assumed to be a channel.
  def parse_destination
    # Assume # is a channel and short circuit
    return if inputs['destination'].first == '#'

    # Don't allow hardcoding slack user ids
    if inputs['destination'].match?(/^U\w{8,10}$/)
      raise "Sending to users directly is not supported, please use destination: 'committer'"
    end

    if inputs['destination'] == 'committer'
      # set when someone rebuilds a workflow
      if ENV['GITHUB_TRIGGERING_ACTOR']
        inputs['destination'] = slack_id_for_github_user(context.triggering_actor)
      # set to the user who started the workflow manually/dispatched (maybe commit author on push)
      elsif ENV['GITHUB_ACTOR']
        inputs['destination'] = slack_id_for_github_user(context.actor)
      elsif ENV['GITHUB_EVENT_NAME'] == 'push'
        inputs['destination'] = slack_id_for_github_user(context.event.head_commit.author.username)
      else
        inputs['destination'] = inputs['fallback_destination']
      end
    end

    if ['U', '#'].include?(inputs['destination'].first)
      raise "Destination and fallback-destination did not resolve to a valid slack user or channel"
    end
  end

  def parse_message
    # TODO: port parser.py
  end

  def send_message
    slack_client.post(inputs['message'], inputs['destination'])
  end

  def self.send_message
    new.send_message
  end
end
