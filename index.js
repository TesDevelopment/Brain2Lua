let code = `+[,.+]`

let base = `
local cells = {0};
`

let instructions = code.split("")
let cur = 1
let set = [1]

instructions.forEach((instruction) => {
    switch (instruction) {
        case "+" :
            base += `cells[${cur}] = cells[${cur}] + 1;\n`
        break
        case "-" :
            base += `cells[${cur}] = cells[${cur}] - 1;\n`
        break
        case ">":
            cur++
            if(!set.includes(cur)){
                base += `cells[${cur}] = 0;\n`
            }
        break
        case "<":
            cur--
            if(cur < 1){
                console.log("ERROR: Grid escape detected at instruction " + instruction)
                //@ts-ignore
                process.exit(1)
            }
        break
        case "[":
            base += `while cells[${cur}] ~= 0 do \n`
        break
        case "]":
            base += `end;\n`
        break
        case ".":
            base += `print(string.char(cells[${cur}]));\n`
        break
        case ",":
            base += `cells[${cur}] = io.read();\nif #cells[${cur}] > 1 then\n error("Expected string of length 1")\nend;\ncells[${cur}] = string.byte(cells[${cur}]);\n`
        break
    }
})

console.log(base)