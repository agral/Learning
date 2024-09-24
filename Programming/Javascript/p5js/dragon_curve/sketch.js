let segments = [];

function mousePressed() {
    let s = segments[0];
    let newS = s.rotate(s.end);
    segments.push(newS);
}

function setup() {
    createCanvas(800, 600);
    let start = createVector(200, 250);
    let end = createVector(200, 150);
    let segment = new LineSegment(start, end);
    segments.push(segment);
}


function draw() {
    background(220);
    for (const s of segments) {
        s.draw();
        s.drawDebugInfo();
    }
}
