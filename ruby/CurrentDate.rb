#!/usr/bin/ruby
# frozen_string_literal: true
require 'date'
puts "Ruby Version: #{RUBY_VERSION}"
puts "Ruby Patch Level: #{RUBY_PATCHLEVEL}"
current_date = DateTime.now
puts current_date.strftime "%d/%m/%Y %H:%M"
