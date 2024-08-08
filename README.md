# Bouncing Dot Animation

![](screenshot.PNG)

## Overview

This project implements a simple animation of a bouncing dot with a colorful trail using the Ebitengine game library in Go. It demonstrates basic concepts of physics simulation and color theory in computer graphics.

## Features

1. A dot that moves and bounces off screen boundaries
2. A colorful, fading trail that follows the dot
3. Smooth color cycling for the trail

## Implementation Details

### 1. Dot Movement

The dot's motion is implemented using basic linear motion equations:

```go
g.dot.x += g.dot.dx
g.dot.y += g.dot.dy
```

This approach was chosen for its simplicity and effectiveness in simulating constant velocity motion.

### 2. Boundary Collision

Collision with screen boundaries is handled by reversing the dot's velocity:

```go
if g.dot.x <= dotRadius || g.dot.x >= screenWidth-dotRadius {
    g.dot.dx = -g.dot.dx
}
```

This method provides a straightforward approximation of elastic collision, suitable for this animation's purposes.

### 3. Trail Generation

The trail is generated by storing recent positions of the dot:

```go
g.trails = append(g.trails, [2]float64{g.dot.x, g.dot.y})
if len(g.trails) > trailLength {
    g.trails = g.trails[1:]
}
```

I opted for a fixed-length queue to maintain memory efficiency while creating a fading effect.

### 4. Color Cycling

Smooth color transitions are achieved using trigonometric functions:

```go
r := uint8((math.Sin(g.dot.angle) + 1) * 127)
b := uint8((math.Cos(g.dot.angle) + 1) * 127)
```

This technique was selected for its ability to produce smooth, continuous color variations with minimal computational overhead.

## Mathematical Concepts

1. **Linear Motion**: Used for basic dot movement.
2. **Elastic Collision**: Simplified for boundary interactions.
3. **Trigonometric Functions**: Employed for smooth color cycling.
4. **2D Coordinate System**: Utilized for positioning elements on the screen.

## Algorithms

1. **Bouncing**: Implemented via simple boundary checking and velocity reversal.
2. **Trail Generation**: Utilizes a queue-like data structure with a fixed maximum length.
3. **Color Cycling**: Achieved through continuous angle updates and trigonometric color mapping.

## Setup and Execution

1. Ensure Go is installed on your system.
2. Install the Ebitengine library:
   ```
   go get github.com/hajimehoshi/ebiten/v2
   ```
3. Save the provided code in a file named `main.go`.
4. Execute the program:
   ```
   go run main.go
   ```

## Educational Value

This project serves as an introductory example to:
- Basic motion simulation in computer graphics
- Simple physics modeling
- Color representation and manipulation in code
- Practical applications of trigonometric functions in graphics

## Customization

Users are encouraged to experiment with the constants and parameters in the code to observe different effects. Potential modifications include:
- Altering dot speed and size
- Modifying trail length and fade rate
- Adjusting color cycling speed and range

## Conclusion

This animation demonstrates how fundamental mathematical concepts and straightforward algorithms can create engaging visual effects in computer graphics. It provides a foundation for understanding more complex simulations and graphics programming techniques.

## Future Enhancements

Potential areas for expansion include:
- Implementing more realistic physics (e.g., gravity, friction)
- Adding user interaction (e.g., mouse control of the dot)
- Introducing multiple dots with collision detection

## License

MIT