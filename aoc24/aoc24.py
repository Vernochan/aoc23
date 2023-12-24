import z3

infile = open("test.txt")

positions = []
velocities = []

for line in infile:
    line = line.strip()
    p, v = line.split(" @ ")
    # using float, becuase thats much faster in z3
    positions.append(tuple(map(float, p.split(", "))))
    velocities.append(tuple(map(float, v.split(", "))))


# https://ericpony.github.io/z3py-tutorial/guide-examples.htm

# Add variables to solve for
# using Real, becuase it's _MUCH_ faster than using Int (oders of magnitude)
stone_x = z3.Real('x')
stone_y = z3.Real('y')
stone_z = z3.Real('z')
stone_vx = z3.Real('vx')
stone_vy = z3.Real('vy')
stone_vz = z3.Real('vz')

solver = z3.Solver()

for i in range(len(positions)):
    current_hail_x, current_hail_y, current_hail_z = positions[i]
    current_hail_vx, current_hail_vy, current_hail_vz = velocities[i]

    # Add variable for time, since times for a hit can (will/should?) be different,
    # it needs to have a unique name for every hail
    current_hail_hit_time = z3.Real(f"t_{i}")

    # add constaints: For x, y and z: 
    # Left side:  Current Hail Position + velocity * time
    # Right side: Stone Position
    solver.add(current_hail_x + current_hail_vx * current_hail_hit_time == stone_x + stone_vx * current_hail_hit_time)
    solver.add(current_hail_y + current_hail_vy * current_hail_hit_time == stone_y + stone_vy * current_hail_hit_time)
    solver.add(current_hail_z + current_hail_vz * current_hail_hit_time == stone_z + stone_vz * current_hail_hit_time)

solver.check()
model = solver.model()
# only prints the values, does not calculate them.. could be converted somehow, but just pipe the output to bc
print(model[stone_x] + model[stone_y] + model[stone_z])
