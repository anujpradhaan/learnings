#!/usr/bin/ruby
#Creating a Global Variable
$global_variable=1

class Class1
    #Creating a class level variable
    @@total_number_of_objects = 0
    CONSTANT_VALUE = 12
   
    #Initialize the value into the object on call of new for object creation
    def initialize
        @@total_number_of_objects += 1
        puts @@total_number_of_objects
        puts CONSTANT_VALUE
    end

    def print_global_values
        puts "Global value is #$global_variable"
    end

    #Using a local variable
    def print_global_value(_value)
        puts "Global value is #{$global_variable} and local value is #{_value}"
    end

    def print_file_name
        puts "File name is #{__FILE__}"
    end

    def print_line_number
        puts "Line number here is #{__LINE__}"
    end

    def try_arrays
        arry = ["fred",1,1.5,"last"]
        arry.each do |i|
            puts i
        end
    end

    def test_ranges
        (1..5).each do |value|
            puts value
        end

        (1...5).each do |value|
            puts value
        end
    end
end

class Class2
    def print_global_values
        puts "Global Value is #$global_variable"
    end
end

class1Object = Class1.new
#class1Object.print_global_value(2)
#class2Object = Class2.new
#class1Object2 = Class1.new
#class2Object.print_global_values
#class1Object.print_line_number
#class1Object.print_file_name
#class1Object.try_arrays
class1Object.test_ranges