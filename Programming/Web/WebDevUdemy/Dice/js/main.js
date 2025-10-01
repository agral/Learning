let i1 = document.getElementById("i1")
let i2 = document.getElementById("i2")
let r = document.getElementById("result")

let roll1 = 1 + Math.floor(Math.random() * 6);
let roll2 = 1 + Math.floor(Math.random() * 6);

i1.src = "img/dice" + roll1 + ".png"
i2.src = "img/dice" + roll2 + ".png"

if (roll1 > roll2) {
    r.innerText = "Player 1 wins!"
} else if (roll1 < roll2) {
    r.innerText = "Player 2 wins!"
} else {
    r.innerText = "That's a draw!"
}
