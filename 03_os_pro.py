import os


directory_path = "/home/rohit"


contents = os.listdir(directory_path)


for item in contents:
    print(item)
