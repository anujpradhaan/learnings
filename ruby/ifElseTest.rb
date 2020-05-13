#!/usr/bin/ruby

class Class1
    def run_if_else(x)
        # If evaluates true if the value is not nill or false
        if x
            puts "I am in if with value #{x}"
        else
            puts "I am in else with value #{x}"
        end
    end

    def run_if_els_else(i)
        if i>2
            puts "Inside if"
        elsif
            puts "Inside elsif"
        else
            puts "Inside else"
        end
    end
end

classObject = Class1.new
#classObject.run_if_else(nil)
x=nil
puts "x" if x
puts "x unless" unless x
case 13
when 1..5,12...13,13..15
    puts "First"
when 13..15
    puts "Second"
else 
    puts "else"
end
