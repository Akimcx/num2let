const go = new Go();
let wasm = undefined
const number = document.getElementById("number")
const result = document.getElementById("result")

function cstrlen(mem, ptr) {
    let len = 0;
    let ptr_cp = ptr
    while (mem[ptr_cp] != 0) {
	len += 1
	ptr_cp += 1
    }
    return len;
}

function trim(str) {
    const arr = Array.from(str)
    return arr.filter(x => x !== " ").join("")
}
function cstr_by_ptr(mem_buffer, ptr) {
    const mem = new Uint8Array(mem_buffer);
    const len = cstrlen(mem, ptr);
    const bytes = new Uint8Array(mem_buffer, ptr, len);
    return new TextDecoder().decode(bytes);
}

WebAssembly.instantiateStreaming(fetch('num2let.wasm'),go.importObject)
    .then(res => {
	wasm = res.instance
	// let ptr = res.instance.exports.add3(0)
	// const buffer = wasm.exports.memory.buffer
	// const mem = new Uint8Array(buffer)
	// const bytes = new Uint8Array(buffer, ptr, 1);
	// const td = new TextDecoder()
	// console.log(buffer,ptr)
	// console.log(`mem[ptr]: ${mem[ptr+1]}`)
	// console.log(bytes)
	// console.log(td.decode(bytes))

	number.addEventListener("input", () => {
	    const val = Number.parseInt(trim(number.value))
	    if(isNaN(val)) {
		result.value = ``
		return
	    }
	    const ptr = wasm.exports.convert(val)
	    const buffer = wasm.exports.memory.buffer
	    const str =  cstr_by_ptr(buffer,ptr)
	    result.value = str
	    // const buffer = res.instance.exports.memory.buffer
	    // console.log(cstr_by_ptr(buffer,result.value))
	})
    })
    .catch(err => {
	console.log(err)
})
