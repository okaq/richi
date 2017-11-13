# -*- coding: utf-8 -*-
"""
Created on Sun Nov  5 14:05:19 2017

@author: ahmad
"""

# test op tools
# n queens problem
from __future__ import print_function
import sys
from ortools.constraint_solver import pywrapcp

def main():
    solver = pywrapcp.Solver("simple_example")
    
    num_vals = 3
    x = solver.IntVar(0, num_vals - 1, "x")
    y = solver.IntVar(0, num_vals - 1, "y")
    z = solver.IntVar(0, num_vals - 1, "z")
    
    solver.Add(x != y)
    solver.Add(y != z)
    solver.Add(x != z)
    
    db = solver.Phase([x,y,z], solver.CHOOSE_FIRST_UNBOUND, solver.ASSIGN_MIN_VALUE)
    
    solver.Solve(db)
    count = 0

    while solver.NextSolution():
        count += 1
        print("Solution", count, '\n')
        print("x = ", x.Value())
        print("y = ", y.Value())
        print("z = ", z.Value())
        print()
    print("Number of solutions: ", count)
   
   
if __name__ == "__main__":
    main()
    
"""

    by adding all possible constraints
    we build a powerful engine
    for enumerating complex combinatorial spaces

"""
