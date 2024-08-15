from settings import (
    get_school,
)
from subjects import *


school = get_school()
grades = school['grades']
areas = school['areas']


INPUT = {
    'subjects_in_a_month': {
        ASTRONOMY: 4,
        CHEMISTRY: 4,
        BIOLOGY: 4,
        PHYSICS: 12,
        MATH: 12,
        INFORMATICS: 8,
        GEOMETRY: 4,
        ENGLISH: 12,
        HISTORY: 4,
        LITERATURE: 8,
    },
    'grades_sum_per_subjects': {
        ASTRONOMY: sum([
            grades['5/5'] * 1,
            grades['4'] * 1,
        ]),
        CHEMISTRY: sum([
            grades['5'] * 2,
        ]),
        BIOLOGY: sum([
            grades['4'] * 2,
        ]),
        PHYSICS: sum([
            grades['5'] * 6,
            grades['4'] * 1,
        ]),

        MATH: sum([
            grades['5'] * 9,
        ]),
        INFORMATICS: sum([
            grades['5'] * 6,
        ]),
        GEOMETRY: sum([
            grades['5'] * 1,
            grades['3'] * 1,
        ]),
        
        ENGLISH: sum([
            grades['5'] * 7,
        ]),
        HISTORY: sum([
            grades['5'] * 1,
        ]),
        LITERATURE: sum([
            grades['2'] * 1,
        ]),
    },
    
}

INPUT_2 = {}
INPUT_3 = {}
INPUT_4 = {}


def get_input():
    return INPUT
    # return INPUT_2
    # return INPUT_3
    # return INPUT_4
