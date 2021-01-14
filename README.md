###Build:

`make build`

Binary will be build inside `$CWD/bin` folder
###Run: 

``hatch --help``

###Run 2 json files comparison:

``hatch json --filePath1 <path> --filePath2 <path>``

Each filepath defaults to `input1.json` and `input2.json` respectively


####Decisions made:
1. Not to unmarshal directly into struct, 
   rather use `map[string]interface` for the sake if speed
2. There are multiple solutions for this task:
   2.1 Iterate through object in the first file and try to found 
       same object in second file. Time complexity will be 
       O(n^2)
       in this case.
   2.2 Organize objects from both files same way , sort them and compare
       first with first, second with second etc. 
   2.3 Hash object into map and count each hash occurrence in both files.
       There is a chance two different items will have sam hashes, but taking into
       account that files may bee only upto few gigabytes, this chance is almost zero.
       In this cae complexity will be linear, exactly what was implemented.
   
   To all the described complexities we should add up a complexity of organizing objects
   themselves(sorting inner arrays etc), but taking into account that this value is much lower than searching itself,
   it can be ignored.
3. Work has been done under assumption that we know structure of each object and we're
   able to write a function for converting from `map[string]interface{}` to object type