clearArray = (arr) ->
    arr.splice 0, arr.length
    arr
console.log clearArray [0,1]

run = (fun, a...) -> fun.apply this, a

