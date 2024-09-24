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
    let canvas = {"w": 800, "h": 600};
    createCanvas(canvas.w, canvas.h);
    let start = createVector(0.5 * canvas.w, 0.51 * canvas.h);
    let end = createVector(0.5 * canvas.w, 0.49 * canvas.h);
    let segment = new LineSegment(start, end);
    endpoint = end;
    segments.push(segment);
}


function draw() {
    background(220);
    for (const s of segments) {
        s.draw();
        //s.drawDebugInfo();
    }
    // drawDebugInfo();
}

function drawDebugInfo() {
    // Draw the global endpoint:
    strokeWeight(0);
    textSize(10);
    text("E", endpoint.x + 2, endpoint.y + 8);

    // Draw the statistics - number of line segments:
    text(`segments: ${segments.length}`, 10, 10);
}
