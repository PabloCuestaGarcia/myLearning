"""
Python data classes.
--------------------

"""

import random
import string

from dataclasses import dataclass, field


def generate_id() -> str:
    return "".join(random.choices(string.ascii_uppercase, k=12))

class Person:
    def __init__(self, name: str, address: str):
        self.name = name
        self.address = address
    def __str__(self) -> str:
        return f"{self.name}, {self.address}"

# data class example
@dataclass
class DataPerson:
    name: str
    address: str
    active: bool = True
    email_addresses: list[str] = field(default_factory=list)
    id: str = field(init=False, default_factory=generate_id)
    search_string: str = field(init=False, repr=False)

    def __post_init__(self) -> None:
        self.search_string = f"{self.name}, {self.address}"

# Frozen class is like a sealed record or sealed class
@dataclass(frozen=True)
class FrozenPerson:
    name: str
    address: str
    active: bool = True        

@dataclass(kw_only=True)
class DataPerson:
    name: str
    address: str
    active: bool = True

def main() -> None:
    person = Person(name="Pablo", address="Costa del Sol")
    print(person)

    another_person0 = DataPerson("Pablo", "Costa del Sol")
    another_person1 = DataPerson("Cristina", "Costa del Sol")
    print(another_person0)
    print(another_person1)

    frozen_person = FrozenPerson(name="Pablo", address="Costa del Sol")





if __name__ == "__main__":
    main()
