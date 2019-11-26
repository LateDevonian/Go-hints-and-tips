if else statements:
b := true
if b {
    //do a thing if it's evaluated as true
    } else {
        //do the other thing
    }

Operands:
== equal to 
!= not equal to
< less than
<= less than or equal to
> gt
>= gt or et

+ sum /add 
- difference/ minus
* product/times
/ quotient/division
% reaminder 

switch statement and case
example

s := "c"
switch s {
    case "a";
    //do a thing
    case "b";
    //do another thing
    default:
    //do the default
}

loop with a for
i := 0
for i < 10 {
    i ++
    fmt.Println("i is", i)
}

for statement with init and post
can iniitalise i, evaluate it, and incremented
excecution returns a boolean that checks to see if it's less than ten, we do it again
for i := 0, i < 10; i ++ {
    //do a thing
}

loop over a data structure: a range clause
 numbers := []int(1,2,3)
 //for statement assigns an iteraion variable to i, holds the index
 //updated at end of loop
 //for assigns interation variable to n. holds value from teh slice
 //for the iteration in the loop. updated at end of loop
    for i, n := range numbers {
        fmt.Println("the indeex of the loop is", i)
        fmt.Println("the falcue from the array is", n)
    }

Defer statement, allows a fucntion to e excuted after the surrounding function returns. used as cleanup or ensuring network calls chave completed.