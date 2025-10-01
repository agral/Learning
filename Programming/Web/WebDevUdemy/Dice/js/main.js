let roll1 = 1 + Math.floor(Math.random() * 6);
let roll2 = 1 + Math.floor(Math.random() * 6);
console.log(roll1)
console.log(roll2)

let i1 = document.getElementById("i1")
let i2 = document.getElementById("i2")

i1.src = "img/dice" + roll1 + ".png"
i2.src = "img/dice" + roll2 + ".png"
