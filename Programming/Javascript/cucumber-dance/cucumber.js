const cucumber = document.querySelector("#cucumber");
const cucumberBox = cucumber.getBoundingClientRect()
const maxWidth = document.body.clientWidth - cucumberBox.width;

let counter = 0;
let rate = 3;

function loop() {
    counter += rate;
    cucumber.style.left = counter + "px";
    cucumber.style.top = (Math.cos(counter * 0.05) * 10 + Math.sin(counter * 0.13) * 15) + "px";
    if ((counter < 0) || (counter > maxWidth)) {
        rate *= -1;
    }

    requestAnimationFrame(loop);
}

function start() {
    console.log("start()")
    loop();
}

start();
