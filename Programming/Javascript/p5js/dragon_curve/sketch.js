let segments = [];
let endpoint;

function mousePressed() {
    let endpoint = segments[segments.length - 1].start;
    let s = segments[0];
    let newS = s.rotate(s.end);
    segments.push(newS);
    // todo: update the endpoint...?
}

function setup() {
    createCanvas(800, 600);
    let start = createVector(200, 250);
    let end = createVector(200, 150);
    let segment = new LineSegment(start, end);
    endpoint = end;
    segments.push(segment);
}


function draw() {
    background(220);
    for (const s of segments) {
        s.draw();
        s.drawDebugInfo();
    }
    // Draw the global endpoint:
    strokeWeight(0);
    textSize(10);
    text("E", endpoint.x + 2, endpoint.y + 8);
}
