from dataclasses import dataclass
from enum import Enum

# Basic syntax


def http_error(status: int) -> str:
    match status:
        case 400:
            return "Bad request"
        case 401 | 403:
            return "Not allowed"
        case 404:
            return "Not found"
        case _:
            return "Unknown status code"



@dataclass
class Point:
    x: int
    y: int

def location(point: Point) -> None:
    match point:
        case Point(x=0, y=0):
            print("Origin is the point's location.")
        case Point(x=0, y=y):
            print(f"Y={y} and the point is on the y-axis.")
        case Point(x=x, y=0):
            print(f"X={x} and the point is on the x-axis.")


points = [Point(0,0), Point(1,1), Point(1,12)]

def matching_points(point_list: list[Point]) -> None:
    match point_list:
        case []:
            print("No points in the list.")
        case [Point(0,0)]:
            print("The origin is the only point in the list")
        case [Point(x, y)] if x == y:
            print("A point x == y is in the list.")
        case [*points] if Point(1,12) in points:
            print(f"Point(1,12) is in the list -> {points}")
        case _:
            print("Any point matched.")

class Color(Enum):
    RED = 0
    GREEN = 1
    BLUE = 2

def matching_colors(color: Color) -> None:
    match color:
        case Color.RED:
            print("The color is red.")
        case Color.GREEN:
            print("The color is green")
        case Color.BLUE:
            print("The color is blue.")


if __name__ == "__main__":

    print(http_error(400))
    location(Point(x=0, y=0))
    location(Point(0, 1))
    matching_points([])
    matching_points([Point(30,30)])
    matching_points(points)
    matching_colors(Color.BLUE)
