# Backend Server Technical Specs:

## Business Goal: 

A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

## Technical requirements

- the backend will be built using golang
- the database will be built using sqlite3
- the API will be built using gin
- the API will always return JSON
- There will no authentication or authorization
- Everything will be treated as a single user

## Database Schema

We have following tables:

- words - store vocabulary words
    - id integer
    - japanese string
    - romaji string
    - english string
    - parts json

- word_groups : join table for words and groups many to many
    - id integer
    - word_id integer
    - group_id integer

- groups: thematic groups of words
    - id integer
    - name string

- study_sessions: records of study sessions grouping word_review_items
    - id integer
    - group_id integer
    - created_at datetime
    - study_activity_id integer

- study_activities: a specific study activity, linking a study session to group
    - id integer
    - study_session_id integer
    - group_id integer
    - created_at datetime

- word_review_items: a record of a word practice, determining if the word was correct or not
    - word_id integer
    - study_session_id integer
    - correct boolean
    - created_at datetime

### API Endpoints


- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick_stats
- GET /api/study_activity/:id
- GET /api/study_activity/:id/study_sessions
- POST /api/study_activities
    - required parameters:
        - group_id
        - study_activity_id

- GET /api/words
- GET /api/words/:id
- GET /api/groups
- GET /api/groups/:id
- GET /api/groups/:id/words
