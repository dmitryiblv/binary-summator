# binary-summator

Sample of Binary Summator, Little Indian

One day, I thought about how I could clearly explain to a person not from the IT sphere why the main numeral system for computers is binary.

This is easy to understand if you work closely with computers yourself. Otherwise, you need to show the person a real example that you can kind of touch.

Adding two numbers illustrates this feature very well. If you want, you can assemble a binary summator (adder) with analog components or even with mechanical ones.

<pre>

~Helper table: dec to bin
    0:  0000
    1:  0001
    2:  0010
    3:  0011
    4:  0100
    5:  0101
    6:  0110
    7:  0111
    8:  1000
    9:  1001
    10: 1010
    ...
    15: 01111
    16: 10000

~Helper bitwise operations: manual sum
    sum 0 + 1: 0000 + 0001 -> 0001
        0 xor 1 -> 1

    sum 1 + 1: 0001 + 0001 -> 0010
        1 xor 1 -> 0, 1 mov         0
        0 xor 0 -> 0, xor 1 -> 1    1
        
    sum 1 + 3: 0001 + 0011 -> 0100
        1 xor 1 -> 0, 1 mov         0
        0 xor 1 -> 1, xor 1 -> 0    0
        -> 1                        1

    sum 2 + 2: 0010 + 0010 -> 0100
        0 xor 0 -> 0            0
        1 xor 1 -> 0, 1 mov     0
        -> 1                    1

    sum 2 + 3: 0010 + 0011 -> 0101
        0 xor 1 -> 1                1
        1 xor 1 -> 0, 1 mov         0
        0 xor 0 -> 0, xor 1 -> 1    1   

    sum 2 + 5: 0010 + 0101 -> 0111
        0 xor 1 -> 1
        1 xor 0 -> 1
        0 xor 1 -> 1

    sum 3 + 4: 0011 + 0100 -> 0111
        1 xor 0 -> 1
        1 xor 0 -> 1
        0 xor 1 -> 1

    sum 1 + 15: 0001 + 1111 -> 10000
        1 xor 1 -> 0, 1 mov                 0
        0 xor 1 -> 1, xor 1 -> 0, 1 mov     0
        0 xor 1 -> 1, xor 1 -> 0, 1 mov     0
        0 xor 1 -> 1, xor 1 -> 0, 1 mov     0    
        -> 1                                1

    sum 17 + 19: 10001 + 10011 -> 100100
        1 xor 1 -> 0, 1 mov                 0
        0 xor 1 -> 1, xor 1 -> 0, 1 mov     0
        0 xor 0 -> 0, xor 1 -> 1            1
        0 xor 0 -> 0                        0
        1 xor 1 -> 0, 1 mov                 0
        -> 1                                1

Output, sumMaxBits = 3 :
    2025/02/14 17:02:07     Num A max: 4, num B max: 3
    2025/02/14 17:02:07     000 (0) + 000 (0) -> 000 (0)
    2025/02/14 17:02:07     000 (0) + 001 (1) -> 001 (1)
    2025/02/14 17:02:07     000 (0) + 010 (2) -> 010 (2)
    2025/02/14 17:02:07     000 (0) + 011 (3) -> 011 (3)
    2025/02/14 17:02:07     001 (1) + 000 (0) -> 001 (1)
    2025/02/14 17:02:07     001 (1) + 001 (1) -> 010 (2)
    2025/02/14 17:02:07     001 (1) + 010 (2) -> 011 (3)
    2025/02/14 17:02:07     001 (1) + 011 (3) -> 100 (4)
    2025/02/14 17:02:07     010 (2) + 000 (0) -> 010 (2)
    2025/02/14 17:02:07     010 (2) + 001 (1) -> 011 (3)
    2025/02/14 17:02:07     010 (2) + 010 (2) -> 100 (4)
    2025/02/14 17:02:07     010 (2) + 011 (3) -> 101 (5)
    2025/02/14 17:02:07     011 (3) + 000 (0) -> 011 (3)
    2025/02/14 17:02:07     011 (3) + 001 (1) -> 100 (4)
    2025/02/14 17:02:07     011 (3) + 010 (2) -> 101 (5)
    2025/02/14 17:02:07     011 (3) + 011 (3) -> 110 (6)
    2025/02/14 17:02:07     100 (4) + 000 (0) -> 100 (4)
    2025/02/14 17:02:07     100 (4) + 001 (1) -> 101 (5)
    2025/02/14 17:02:07     100 (4) + 010 (2) -> 110 (6)
    2025/02/14 17:02:07     100 (4) + 011 (3) -> 111 (7)
    2025/02/14 17:02:07     OK
</pre>
