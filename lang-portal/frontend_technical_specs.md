# Frontend Technical Specs

## pages

### Dashboard /dashboard




This page contains the following components:
- Last study session
    - shows last activity used
    - shows when last activity used
    - summarizes wrong and correct words from last study activity
    - has a link to the group ( to view group)

- study progress
    - total words studied eg. 3/124
    - across all study sessions shows total words studied out of all possible words in our database
    - display mastery progress and percentage

- quick stats
    - success rate eg 30%
    - total study sessions eg. 4
    - total active groups eg. 3
    - study streak eg. 14 days
    
- start studying button
    - goes to study activity page


we'll need following API endpoints to power this page
- GET /dashboard/last_study_session
- GET /dashboard/study_progress
- GET /dashboard/quick_stats


