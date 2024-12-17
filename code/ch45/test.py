my_int = 10
another_int = my_int  # another_int 和 my_int 指向同一个整数对象
another_int += 1
print(my_int)  # 输出: 10
print(another_int)