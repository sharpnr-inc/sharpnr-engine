import { cn } from "@/lib/aceternity-ui";
import React, { useEffect, useRef } from "react";

export const BackgroundBeams = ({
  className,
}: {
  className?: string;
}) => {
  const containerRef = useRef<HTMLDivElement>(null);
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const animationFrameRef = useRef<number>(0);

  useEffect(() => {
    const canvas = canvasRef.current;
    const ctx = canvas?.getContext("2d");
    if (!canvas || !ctx) return;

    const resizeCanvas = () => {
      canvas.width = containerRef.current?.clientWidth || window.innerWidth;
      canvas.height = containerRef.current?.clientHeight || window.innerHeight;
    };

    window.addEventListener("resize", resizeCanvas);
    resizeCanvas();

    const beams: Beam[] = [];
    const numBeams = 10;

    class Beam {
      x: number;
      y: number;
      angle: number;
      speed: number;
      length: number;
      width: number;
      color: string;

      constructor() {
        this.x = Math.random() * canvas.width;
        this.y = Math.random() * canvas.height;
        this.angle = Math.random() * 2 * Math.PI;
        this.speed = Math.random() * 0.2 + 0.1;
        this.length = Math.random() * 200 + 100;
        this.width = Math.random() * 2 + 1;
        this.color = `rgba(128, 128, 128, ${Math.random() * 0.3 + 0.1})`;
      }

      update() {
        this.x += Math.cos(this.angle) * this.speed;
        this.y += Math.sin(this.angle) * this.speed;
        if (this.x < 0 || this.x > canvas.width || this.y < 0 || this.y > canvas.height) {
          this.x = Math.random() * canvas.width;
          this.y = Math.random() * canvas.height;
        }
      }

      draw() {
        ctx.beginPath();
        ctx.moveTo(this.x, this.y);
        ctx.lineTo(
          this.x + Math.cos(this.angle) * this.length,
          this.y + Math.sin(this.angle) * this.length
        );
        ctx.strokeStyle = this.color;
        ctx.lineWidth = this.width;
        ctx.stroke();
      }
    }

    for (let i = 0; i < numBeams; i++) {
      beams.push(new Beam());
    }

    const animate = () => {
      ctx.clearRect(0, 0, canvas.width, canvas.height);
      beams.forEach((beam) => {
        beam.update();
        beam.draw();
      });
      animationFrameRef.current = requestAnimationFrame(animate);
    };

    animate();

    return () => {
      window.removeEventListener("resize", resizeCanvas);
      cancelAnimationFrame(animationFrameRef.current);
    };
  }, []);

  return (
    <div ref={containerRef} className={cn("absolute inset-0 w-full h-full -z-10", className)}>
      <canvas ref={canvasRef} className="w-full h-full" />
    </div>
  );
};