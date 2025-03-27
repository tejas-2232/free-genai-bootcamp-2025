# Frontend Technical Specs

## pages

### Dashboard `/dashboard`

#### Purpose

The purpose of this page is to show a summary of learning and act as a the default page when user visits the web app

#### Components

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

#### Needed API Endpoints

- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick_stats

<hr>

### Study Activity `/study_activity`

#### Purpose

The purpose of this page is to show a collection of study activities with a thumbnail and it's name, to either launch or view the study activity

#### Components

- Study activity card
    - shows thumbnail of the study activity
    - the name of the study activity
    - a launch button to take us to launch page
    - the view page to view more information about past study sessions for this study activity

#### Needed API Endpoints

- GET /api/study_activities
    - shows the study activity progress
    - shows the study activity summary
    - has a link to the group ( to view group)

#### Needed API Endpoints

- GET /api/study_activities

<hr>

### Study Activity Show `/study_activity/:id`

#### Purpose

The purpose of this page is to show a specific study activity with it's details and past study sessions

#### Components

- Name of study activity
- Thumbnail of study activity
- Description of study activity
- Launch Button
- Study activities Paginated List
    - id
    - activity_name
    - group name
    - start time
    - end time (inferred by the last word_review_item submitted)
    - number of review items

#### Needed API Endpoints

- GET /api/study_activity/:id
- GET /api/study_activity/:id/study_sessions

<hr>

### Study Activity Launch `/study_activity/:id/launch`

#### Purpose

The purpose of this page is to Launch a study activity

#### Components

- Name of study activity
- Launch form
    - Select field for group
    - Launch now button

## Bahavior

After the form is submitted, a new tab opens with the study acitivty baes on its URL provided in the database

Also after the form is submitted the page will redirect to the study session show page

#### Needed API Endpoints

- POST /api/study_activities

<hr>

### Words `/words`

#### Purpose

The purpose of this page is to show a collection of all words in our database

#### Components

- paginated word list
    - Columns
        - Japanese
        - romaji
        - english
        - Study Statistics
            - Correct Count
            - Wrong Count
    - Pagination with 100 items per page
    - Clicking the Japanese word will take us to the word show page

#### Needed API Endpoints

- GET /api/words/

<hr>

### Word Show `/words/:id`

#### Purpose
The purpose of this page is to show information about a specific word.

#### Components
- Japanese
- Romaji
- English
- Study Statistics
    - Correct Count
    - Wrong Count
- Word Groups 
    - show an a series of pills eg. tags
    - when group name is clicked it will take us to the group show page

#### Needed API Endpoints
- GET /api/words/:id

<hr>

### Word Groups Index `/groups`

#### Purpose
The purpose of this page is to show a list of groups in our database.

#### Components
- Paginated Group List
    - Columns
        - Group Name
        - Word Count
    - Clicking the group name will take us to the group show page

#### Needed API Endpoints
- GET /api/groups

<hr>

### Word Group Show `/groups/:id`

#### Purpose
The purpose of this page is to show a specific group with it's details and words

#### Components
- Name of group
- Description of group
- Words in group
    - Columns
        - Japanese
        - romaji
        - english
        - Study Statistics
            - Correct Count
            - Wrong Count
        - Pagination with 100 items per page
    - Clicking the Japanese word will take us to the word show page

#### Needed API Endpoints
- GET /api/groups/:id (the name and groups stats)
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions

<hr>

### Study Sessions Index `/study_sessions`

#### Purpose
The purpose of this page is to show a list of study sessions in our database.

#### Components
- Paginated study session list
    - Columns
        - Id
        - Activity Name
        - Group Name
        - Start Time
        - End Time
        - Number of review items
    - Clicking the session Id will take us to the study session show page

#### Needed API Endpoints
- GET /api/study_sessions

<hr>

### Study Session Show `/study_sessions/:id`

#### Purpose
The purpose of this page is to show a specific study session information

#### Components

- Study Session Details
    - Activity Name
    - Group Name
    - Start Time
    - End Time
    - Number of review items
- Words Review Items (Paginated List of Words)
    - should use the same component as the words index page

#### Needed API Endpoints
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words

