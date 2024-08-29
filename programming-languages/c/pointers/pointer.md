A pointer is a variable that contains the address of a variable. Pointers and arrays are closely related.

 A typical machine ahs an array of consecutively numbered or addresses memory cells that may be manipulated individually on in contiguous groups.
 
* One byte can be a char;
* A pair of one-byte cells can be treated as s short integer;
* 4 adjacent bytes form a long.

A pointer is a group of cells (often two or four) that can hold an address. 

The unary operator & gives the address of an object.

Ex:

`p = &c;`

This expressions assigns the address of c to the variable p, and p points to c.

(The & operator only applies in memory, variables and array elements)

The unary operator * is the indirection or dereferencing operator, when applied to a pointer, it accesses the object the pointer points to.

```
int x = 1, y =2, z[10];
int *ip;  //ip is a pointer to int

ip = &x; //ip now points to x
y = *ip; //y is now 1;
*ip = 0; //x is now 0;
ip = &z[0];  //ip now points to z[0]

```
Since C passes arguments to functions by value, there is no direct way for the called function to alter a variable in the calling function. For instance, a sorting routine might exchange two out-of-order arguments with a function called swap. It is not enough to write

swap(a,b);

where the swap function is defined as 

```
void swap(int x, int y) /* WRONG */
{
    int tempo;
    tempo = x;
    x = y;
    y = tempo;
    
```

Because of call by value, swap can't affect the arguments a and b in the routine that called it. The function above swaps copies of a and b.

The way to obtain the desired effect is for the calling program to pass pointers to the values to
be changed:

````
swap(&a, &b);

````

````
void swap(int *px, int *py)
{
    int temp;
    temp = *px;
    *px = *py;
    *py = temp;
}
````

## Pointers and arrays

In c. there is a strong relationship between pointers and arrays, strong enough that pointers and arrays should be discussed simultaneously. Any operation that can be achieved by array subscripting can also be be done with pointers.
The pointer version will in general be faster but, at least to the uninitiated, somewhat harder to understand.

The declaration 

`int a[10];`

defines an array of size 10, that is, a block od 10 consecutive objects named a[0], a[1],...,a[9].


![pointer.png](../img/pointer.png)

The notation a[i] refers to the i=tih element of the array. If pa is a pointer to an integer, declared as

`int *pa;`
then the assignment

`pa = &a[0];`
sets pa to point to element zero of a; that is, pa contains the address of a[0].

Now the assignment;

`x = *pa;` //x receives the value that pa points to


will copy the content of a[0] into x.

If pa points to a particular element of an array, then by definition `pa+1` points to the next element, `pa+i` points `i` elements after `pa`, and `pa-i`points `i` elements before. Thus, if `pa` points to `a[0]`,

`*(pa+1)` refers to the contents of `a[i]`, `pa+i` is the address of `a[i]`, and `*(pa+i)` is the contents of `a[i]`.
