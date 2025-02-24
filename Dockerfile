FROM ruby:3.4

WORKDIR /app

COPY Gemfile Gemfile.lock /app/

RUN bundle install

COPY action.rb /app/
COPY lib /app/lib

CMD ["ruby", "/app/action.rb"]
