let segments = [];
let endSegment;
let endpoint;

function mousePressed() {
    let newSegments = [];
    for (let s of segments) {
        let origin = (segments.length == 1 ? endSegment.end : endSegment.start)
        let newSegment = s.duplicate(origin);
        newSegments.push(newSegment);
    }
    segments = segments.concat(newSegments);
    endSegment = newSegments[0];
}

function setup() {
    let canvas = {"w": 800, "h": 600};
    createCanvas(canvas.w, canvas.h);
    let start = createVector(0.5 * canvas.w, 0.55 * canvas.h);
    let end = createVector(0.5 * canvas.w, 0.45 * canvas.h);
    endSegment = new LineSegment(start, end, end);
    endSegment.isDone = true;
    segments.push(endSegment);
}


function draw() {
    background(220);
    for (let s of segments) {
        if (!s.isDone) {
            s.update();
        }
        s.draw();
        //s.drawDebugInfo();
    }
    drawDebugInfo();
}

function drawDebugInfo() {
    // Draw the global endpoint:
    strokeWeight(0);
    textSize(10);
    text("E", endSegment.end.x + 2, endSegment.end.y + 8);

    // Draw the statistics - number of line segments:
    text(`segments: ${segments.length}`, 10, 10);
}
