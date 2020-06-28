import math

def print_stats():
    data_str = str(input("Enter the sample data as comma separated values or whitespace separated values: "))
    if "," in data_str:
        data_str = data_str.split(",")
    elif " " in data_str:
        data_str = data_str.split(" ")

    data = [int(data_str[i]) for i in range(len(data_str))]
    data.sort()
    print(data)

    print("Descriptive Statistics:")

    # print_min(data)
    print("Min: ", data[0])

    # print_max(data)
    print("Max:", data[len(data)-1])

    # print_range(data)
    print("Range:", data[len(data)-1] - data[0])

    # print_size(data)
    print("Size:",len(data))

    #print_sum(data)
    s = sum(data)
    print("Sum:",s)

    #print_mean(data)
    m = sum(data)/len(data)
    print("Mean:",m)

    # print_median(data)
    if len(data) % 2 == 1: # size is odd
        print("Median:",data[int(len(data)/2)])
    else:
        print("Median:",(data[int(len(data))-1] + data[int(len(data)/2)]) /2)

    # print_mode(data)
    print("Mode:",max(set(data), key=data.count))

    # print_sd(data)
    # sum the diff x_i & xbar
    x_diff = 0
    for i in data:
        x_diff += ((i - m) ** 2)
    variance = (x_diff / (len(data) - 1))
    print("Standard Deviation:", math.sqrt(variance))
    print("Variance",variance)

    # print_quarters(data)
    # print Q1, Q2, Q3, IQR
    if len(data) < 4:
        print("No quartile information as there are less than 4 data points")
        exit(1)
    print("Q1, Q2, Q3, IQR, Outliers, Extreme Outliers coming soon!")

print_stats()
