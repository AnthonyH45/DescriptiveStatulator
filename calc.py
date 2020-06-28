def get_data():
    data_str = str(input("Enter the data as comma separated values or whitespace separated values: "))
    if "," in data_str:
        data_str = data_str.split(",")
    elif " " in data_str:
        data_str = data_str.split(" ")
    print(data_str)
    
    data = [int(data_str[i]) for i in range(len(data_str))]
    data.sort()
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
    print("Sum:",sum(data))
    #print_mean(data)
    print("Mean:",sum(data)/len(data))
    # print_median(data)
    print

get_data()