┌─────────────────────┐       ┌─────────────────────┐       ┌─────────────────────┐
│       users         │       │    user_yoga_flows  │       │     yoga_flows      │
├─────────────────────┤       ├─────────────────────┤       ├─────────────────────┤
│ id (PK)             │──┐    │ id (PK)             │    ┌──│ id (PK)             │
│ username            │  │    │ user_id (FK)        │────┘  │ type                │
│ email               │  └───>│ yoga_flow_id (FK)   │───────│ timelength          │
│ passwordhash        │       │ created_at          │       │ numberofposes       │
│ firstname           │       └─────────────────────┘       │ poselist (JSONB)    │
│ lastname            │                                     │ averagestrength     │
│ bio                 │                                     │ averageflexibility  │
│ avatarurl           │                                     │ averagedifficulty   │
│ created_at          │                                     └─────────────────────┘
│ updated_at          │
│ role                │       ┌─────────────────────┐       ┌─────────────────────┐
│ user_type           │       │   user_favorites    │       │     yoga_poses      │
│ tier                │       ├─────────────────────┤       ├─────────────────────┤
│ isactive            │──┐    │ id (PK)             │       │ id (PK)             │
└─────────────────────┘  │    │ user_id (FK)        │───┐   │ name                │
                         └───>│ favorite_poses      │   │   │ sanskrit            │
                              │ favorite_flows      │   │   │ category            │
┌─────────────────────┐       │ created_at          │   │   │ strength            │
│    user_profile     │       │ updated_at          │   │   │ flexibility         │
├─────────────────────┤       └─────────────────────┘   │   │ difficulty          │
│ id (PK)             │                                 │   │ level               │
│ user_id (FK)        │───┐   ┌─────────────────────┐   │   └─────────────────────┘
│ fitness_level       │   │   │  workout_session    │   │
│ flexibility_level   │   │   ├─────────────────────┤   │
│ strength_level      │   │   │ id (PK)             │   │
│ injuries (JSONB)    │   │   │ user_id (FK)        │───┤
│ goals (JSONB)       │   │   │ yoga_flow_id (FK)   │   │
│ created_at          │   │   │ started_at          │   │
│ updated_at          │   │   │ completed_at        │   │
└─────────────────────┘   │   │ duration            │   │
                          │   │ rating              │   │
┌─────────────────────┐   │   │ feedback            │   │
│   user_progress     │   │   │ created_at          │   │
├─────────────────────┤   │   └─────────────────────┘   │
│ id (PK)             │   │                             │
│ user_id (FK)        │───┤   ┌─────────────────────┐   │
│ date                │   │   │  pose_performance   │   │
│ strength_improvement│   │   ├─────────────────────┤   │
│ flexibility_improve │   │   │ id (PK)             │   │
│ sessions_completed  │   │   │ user_id (FK)        │───┤
│ total_time          │   │   │ pose_id (FK)        │───┘
└─────────────────────┘   │   │ attempts            │
                          │   │ success_rate        │
                          │   │ difficulty_rating   │
                          └──>│ last_attempted      │
                              └─────────────────────┘