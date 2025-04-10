marks = {"Harry": 100, "Rohit": 66, "Rohan": 23}

print(marks, type(marks))

print(marks["Harry"])

print(marks.items())
print(marks.keys())
print(marks.values())

marks.update({"Harry": 99})

print(marks)

print(marks.get("Rohit"))
