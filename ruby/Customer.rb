#!/usr/bin/ruby

class Customer
	def initialize(id, name, address)
		@cust_id = id;
		@cust_name = name;
		@cust_addr = address;
	end
	def display_details()
		puts "Customer id is #@cust_id";
	end
end

#create objects
cust1 = Customer.new("1","John","ANB");
cust1.display_details();
