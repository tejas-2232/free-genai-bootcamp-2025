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

## API Endpoints


### GET /api/dashboard/last_study_session
Returns information about the most recent study session.

#### JSON Response
```json
{
  "id": 123,
  "group_id": 456,
  "created_at": "2025-02-08T17:20:23-05:00",
  "study_activity_id": 789,
  "group_id": 456,
  "group_name": "Basic Greetings"
}
```

### GET /api/dashboard/study_progress
Returns study progress statistics.
Please note that the frontend will determine progress bar basedon total words studied and total available words.

#### JSON Response

```json
{
  "total_words_studied": 3,
  "total_available_words": 124,
}
```

### GET /api/dashboard/quick_stats

Returns quick statistics about the user's study progress.

#### JSON Response

```json
{
  "success_rate": 30,
  "total_study_sessions": 4,
  "total_active_groups": 3,
  "study_streak_days": 14
}
```

### GET /api/study_activity/:id

#### JSON Response
```json
{
  "id": 1,
  "name": "Vocabulary Quiz",
  "thumbnail_url": "https://example.com/thumbnail.jpg",
  "description": "Practice your vocabulary with flashcards"
}
```

### GET /api/study_activities/:id/study_sessions

- pagination with 100 items per page

```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-08T17:20:23-05:00",
      "end_time": "2025-02-08T17:30:23-05:00",
      "review_items_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 20
  }
}
```

### POST /api/study_activities

#### Request Params
- group_id integer
- study_activity_id integer

#### JSON Response
{
  "id": 124,
  "group_id": 123
}

### GET /api/words
- pagination with 100 items per page

#### JSON Response
```json
{
  "items": [
    {
      "japanese": "こんにちは",
      "romaji": "konnichiwa",
      "english": "hello",
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 500,
    "items_per_page": 100
  }
}
```

### GET /api/words/:id
#### JSON Response
```json
{
  "japanese": "こんにちは",
  "romaji": "konnichiwa",
  "english": "hello",
  "stats": {
    "correct_count": 5,
    "wrong_count": 2
  },
  "groups": [
    {
      "id": 1,
      "name": "Basic Greetings"
    }
  ]
}
```

### GET /api/groups
    - pagination with 100 items per page

### GET /api/groups/:id

- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions
- GET /api/study_sessions
    - pagination with 100 items per page

-GET /api/study_sessions/:id

- POST /api/reset_history
- POST /api/full_reset
- POST /api/study_sessions/:id/words/:word_id/review
    - required parameters:
        - correct 