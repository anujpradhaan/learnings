#https://www.practicepython.org/exercise/2014/02/05/02-odd-or-even.html
# Problem 1
# from datetime import date
# name = input("What's your name : ")
# age = int(input("What's your age : "))
# currentdate = date.today()
# nextDate = currentdate.replace(year=currentdate.year + (100-age))
# print("Hello %s, you will tun 100 on %s" %(name,nextDate))

#Problem 2
# number = int(input("Input a number"))
# if(number%2==0):
#     print("Even")
# else:
#     print("Odd")

#Problem 3

# a = [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89]
# newlist  = []
# for i in a:
#     if(i<5):
#         newlist.append(i)

# print(newlist)

# print([i for i in a if i < 5])

#Problem 4
# number = int(input("Get Number : "))
# allnumbers = range(2,number+1)
# print([x for x in allnumbers if number%x ==0])

#Problem 5
# a = [1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89]
# b = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]
# print({i for i in a if i in b})

#Problem 6 : Palindrome 

# x = input("Enter Palindrome")
# y = []
# length = len(x) - 1
# while length >= 0: 
#     y.append(x[length])
#     length -= 1

# if x == y : 
#     print("Palindrome")
# else:
#     print("Not Palindrome")

#Problem 7 : Even element from a list
a = [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
print([x for x in a if x%2==0])

