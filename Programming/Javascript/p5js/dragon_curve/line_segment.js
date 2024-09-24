class LineSegment {
    constructor(start, end) {
        this.start = start;
        this.end = end;
    }

    draw() {
        stroke(0);
        strokeWeight(2);
        line(this.start.x, this.start.y, this.end.x, this.end.y);
    }

    drawDebugInfo() {
        strokeWeight(0);
        textSize(8);
        text("s", this.start.x + 2, this.start.y + 8);
        text("e", this.end.x + 2, this.end.y + 8);
    }

    rotate(origin) {
        let vs = p5.Vector.sub(this.start, origin);
        let ve = p5.Vector.sub(this.end, origin);
        vs.rotate(-0.5 * PI);
        ve.rotate(-0.5 * PI);
        let rotatedStart = p5.Vector.add(origin, vs);
        let rotatedEnd = p5.Vector.add(origin, ve);
        let newSegment = new LineSegment(rotatedStart, rotatedEnd);
        return newSegment;
    }
}
