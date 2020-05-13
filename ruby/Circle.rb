
def get_parameter(radius)
  2*Math::PI*radius
end

def get_area(radius)
  Math::PI*(radius**2)
end

printf("%5.2f\n", get_parameter(2))
puts get_area(1)