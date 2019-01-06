# Suffix Array
This package provides common functionality associated with suffix arrays.

```
//returns [][]byte of all distinct substrings of "azaza"
sa := suffix.NewArray([]byte("azaza"))
sub := sa.DistinctSub()

for k, v := range sub {
		fmt.Printf("%+v\n", string(vv))
}

a
az
aza
azaz
azaza
z
za
zaz
zaza

//returns int count of all distinct substrings 
distinctSubCount := sa.DistinctSubCount()
println(distinctSubCount)
9

//returns int count of all possible substrings 
subcount := sa.SubCount()
println(subcount)
15

sa := NewArray([]byte("abracadabra"))
lrs := sa.LongestRepeatedSubs()
println(lrs[0])
abra
```
