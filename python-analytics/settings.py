from subjects import *


SCHOOLS = {
    'default': {
        'grades': {
            '5/5': 9,
            '5': 7,
            '4': 3,
            '3': 1,
            '2': 0,
            'gm': 0,
        },
        'areas': {
            'tebigy': {
                ASTRONOMY: 20,
                CHEMISTRY: 25,
                BIOLOGY: 45,
                PHYSICS: 50,
            },
            'takyk': {
                MATH: 60,
                INFORMATICS: 40,
                GEOMETRY: 50,
            },
            'ynsanperwer': {
                ENGLISH: 25,
                HISTORY: 30,
                LITERATURE: 35,
            },
            # 'programmirleme': {
            #     ENGLISH: 40,
            #     MATH: 25,
            #     INFORMATICS: 50,
            #     GEOMETRY: 25,
            # },
        }
    }
}


def get_school():
    return SCHOOLS['default']
