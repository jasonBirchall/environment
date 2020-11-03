#!/usr/bin/env ruby

HOST = "1.1.1.1"

def ping_host()
  cmd = `ping -c 5 #{HOST} | tail -1 | awk '{print $4}' | cut -d '/' -f 2`
  if cmd.empty?
    puts "No signal"
  else
    puts "ping: #{cmd}"
  end
end

ping_host
