```
┌─[auugustoliks]@[localhost]:~/golang-zombie-playground/02-functions
└──> $ ./run.sh 

---------------------
Simple Function
10

---------------------
Multiple Returns
carlosneto 10

---------------------
Return Empty
Soma 15 Subtração -5

---------------------
>>> Internal Variable
Before:  Carlos
Before &:  0xc000010210
In Function:  CarlosNETO
In Function &:  0xc000010230
After:  Carlos
After &:  0xc000010210

---------------------
>>> Pointer Function Modify Variables
Before:  Hamilton
Before &: 0xc000010260
In Function:  HamiltonNETO
In Function &:  0xc000010260
After:  HamiltonNETO
After &:  0xc000010260

---------------------
>>> Change Global Variable
Before:  0
Before &: 0x57b9f8
In Function:  456
In Function &:  0x57b9f8
After:  456
After &:  0x57b9f8
```
