from inputs import get_input
from settings import get_school
from subjects import *
from bgcolors import BG_COLORS


def main():
    school = get_school()
    input = get_input()

    subjects_capacity = {
        subject: (
            input['subjects_in_a_month'][subject] * 
            school['grades']['5']
        )
        for subject in SUBJECTS
    }
    subjects_progress_in_percentage = {
        subject: round(
            input['grades_sum_per_subjects'][subject] / 
            subjects_capacity[subject] * 100,
            2
        )
        for subject in SUBJECTS
    }
    
    areas_progress = {}
    for area, subjects_required_progress in school['areas'].items():
        area_progress = []
        for subject, required_progress in subjects_required_progress.items():
            if subjects_progress_in_percentage[subject] < required_progress:
                success_percentage = (
                    subjects_progress_in_percentage[subject] / 
                    required_progress
                )
                area_progress.append(success_percentage)
            else:
                area_progress.append(1)

        areas_progress[area] = round(
            sum(area_progress) / 
            len(area_progress) * 100, 
            2
        )

    areas_progress_in_percentage = {
        subject: round(progress / sum(areas_progress.values()) * 100, 2)
        for subject, progress in areas_progress.items()
        if progress > 0
    }

    print(f'{BG_COLORS["info"]}[Subjects Success Percentage]{BG_COLORS["end"]}')
    for subject, percentage in subjects_progress_in_percentage.items():
        print(f'{subject.capitalize()} {percentage}%')

    print()
    print(f'{BG_COLORS["info"]}[Area Progress]{BG_COLORS["end"]}')
    for area, progress in areas_progress.items():
        print(f'{area.capitalize()} {progress}%')

    print()
    print(f'{BG_COLORS["info"]}[Analytics]{BG_COLORS["end"]}')
    for area, percentage in areas_progress_in_percentage.items():
        print(f'{area.capitalize()} {percentage}%')


if __name__ == '__main__':
    main()
