let segments = [];
let endpoint;

function mousePressed() {
    let newSegments = [];
    for (const s of segments) {
        let newSegment = s.rotate(endpoint);
        newSegments.push(newSegment);
    }
    segments = segments.concat(newSegments);
    endpoint = newSegments[0].start;
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

    // Draw the statistics - number of line segments:
    text(`segments: ${segments.length}`, 10, 10);
}
