#
# Title  : binary_gap_checker
# Author : Richie Varghese
# Date   : 06/12/2020
#
# Binary Gap : A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded
#              by ones at both ends in the binary representation of N.
#              N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.
#              N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps
#
# Functions : binaryGapChecker(int) int ---> for returning the longest gap in a binary formatted integer (longest number of 0's)
#             main() ---> gets user input and calls binary gap checker func for result
#
# Results   : https://app.codility.com/demo/results/trainingMSPSSH-C5S/
#


#
# binaryGapChecker func takes an input, convert to its binary form and return the longest gap between two 1's in the gap checker
#
# INPUT  : input int
# OUTPUT : longestGap int
#
def binaryGapChecker(input):

    longestGap, localMax = 0, 0

    # convert the input to it's binary equivalent
    binVal = bin(input)[2:]

    # iterate on each binary element in bin value
    for binElement in binVal:

        # if current binary element is 0, update local maximum by 1
        if int(binElement) == 0:

            localMax += 1

        # if current binary element is 1
        elif int(binElement) == 1:

            # if local max is greater than longest gap, set longest gap as current loca max
            if localMax > longestGap:

                longestGap = localMax

            # reset local max to 0
            localMax = 0

    # return the longest gap
    return longestGap


#
# execute the main function
#
if __name__ == "__main__":

    # get the input value from command line
    input = input("Insert value for number N : ")

    # call binary gap checker function with input to find the longest binary gap
    result = binaryGapChecker(int(input))
    print(f"result = {result}")
