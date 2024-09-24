class LineSegment {
    constructor(start, end, origin) {
        this.start = start;
        this.end = end;
        this.origin = origin.copy();

        this.initial = {
            start: start.copy(),
            end: end.copy(),
        };
        this.isDone = false;
        this.angle = 0;
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

    duplicate(origin) {
        return new LineSegment(this.start.copy(), this.end.copy(), origin);
    }

    update() {
        this.angle += rotationSpeed;
        if (this.angle >= 0.5 * PI) {
            this.angle = 0.5 * PI;
            this.isDone = true;
        }
        let vs = p5.Vector.sub(this.initial.start, this.origin);
        let ve = p5.Vector.sub(this.initial.end, this.origin);
        vs.rotate(-this.angle);
        ve.rotate(-this.angle);
        this.start = p5.Vector.add(this.origin, vs);
        this.end = p5.Vector.add(this.origin, ve);
    }

}
