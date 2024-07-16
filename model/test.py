import sys

if len(sys.argv) > 1:
    print("Arguments received:")
    for arg in sys.argv[1:]:
        print(arg)
else:
    print("No arguments received.")