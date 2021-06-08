# Напишите программу, которая выводит числа от 1 до 100.
# Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz»,
# а для кратных как трём, так и пяти — «FizzBuzz».

for i in range(101):
    if i % 3 == 0 and i % 5 == 0:
        print(i, '- FizzBuzz')
    elif i % 3 == 0:
        print(i, '- Fizz')
    elif i % 5 == 0:
        print(i, '- Buzz')
    else:
        print(i)
