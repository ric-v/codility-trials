//
// Title  : binary_gap_checker
// Author : Richie Varghese
// Date   : 07/12/2020
//
// Binary Gap : A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded
//              by ones at both ends in the binary representation of N.
//              N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.
//              N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps
//
// Functions : binaryGapChecker(int) int ---> for returning the longest gap in a binary formatted integer (longest number of 0's)
//             main() ---> gets user input and calls binary gap checker func for result
//
// Results   : https://app.codility.com/demo/results/trainingCR3RTR-SA4/
//

//
// binaryGapChecker function returns the longest gap between two 1's in a binary converted int input
//
// INPUT  : input int;
// RETURN : longestGap int
//

#include <iostream>
#include <bitset>

using namespace std;

//
// binaryGapChecker function returns the longest number of 0's between two 1's in a binary formatted input
//
// INPUT  : input int;
// RETURN : longestGap int
//
int binaryGapChecker(int input)
{

    int longestGap, localMax;

    // convert the input to binary equivalent
    string binVal = bitset<8>(input).to_string();


    for (int index = 0; index < binVal.length(); index++)
    {



    }

    // return the longest gap
    return longestGap;
}

// main function for binary gap checker
int main()
{

    int input;

    cout << "Insert value for number N : ";
    cin >> input;

    // call binary gap checker for getting the gap value
    int result = binaryGapChecker(input);

    cout << result << endl;

    return 0;
}
