# gungi

Hunter x Hunter - Gungi |  獵人 - 軍儀

## API

### User

#### `GET /api/users/:id`

Get user information.

Parameters:

- `id`: User ID.

Response:

```json
{
  "id": 1,
  "name": "LoveSnowEx",
}
```

#### `POST /api/users`

Create a new user.

Request:

```json
{
  "name": "LoveSnowEx",
}
```

Response:

```json
{
    "id": 1,
    "name": "LoveSnowEx",
}
```

### Game

#### `GET /api/games/:id`

Get game information.

Parameters:

- `id`: Game ID.

Response:

```json
{
    "id": 1,
    "players": [
        {
            "id": 1,
            "name": "LoveSnowEx",
        },
        {
            "id": 2,
            "name": "GungiMaster",
        },
    ],
    "board_pieces": [
        {
            "id": 4,
            "type": "Pawn",
            "color": "Black",
            "position": [2, 5, 0],
        },
        {
            "id": 11,
            "type": "Archer",
            "color": "White",
            "position": [3, 4, 1],
        },
        ...
    ],
    "reserve_pieces": [
        {
            "id": 23,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "dicard_pieces": [
        {
            "id": 7,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "turn": 1,
    "phase": 1,
    "winner": null,
}
```

#### `POST /api/games`

Create a new game.

Request:

None.

Response:

```json
{
    "id": 1,
    "players": [],
    "board_pieces": [],
    "reserve_pieces": [],
    "dicard_pieces": [],
    "turn": 1,
    "phase": 1,
    "winner": null,
}
```

#### `POST /api/games/:id/join`

Join a game.

Parameters:

- `id`: Game ID.

Request:

```json
{
    "user_id": 1,
}
```

Response:

```json
{
    "id": 1,
    "players": [
        {
            "id": 1,
            "name": "LoveSnowEx",
        },
    ],
    "board_pieces": [],
    "reserve_pieces": [],
    "dicard_pieces": [],
    "turn": 1,
    "phase": 0,
    "winner": null,
}
```

#### `POST /api/games/:id/leave`

Leave a game.

Parameters:

- `id`: Game ID.

Request:

```json
{
    "user_id": 1,
}
```

Response:

```json
{
    "id": 1,
    "players": [],
    "board_pieces": [],
    "reserve_pieces": [],
    "dicard_pieces": [],
    "turn": 1,
    "phase": 0,
    "winner": null,
}
```

#### `POST /api/games/:id/start`

Start a game.

Parameters:

- `id`: Game ID.

Response:

```json
{
    "id": 1,
    "players": [
        {
            "id": 1,
            "name": "LoveSnowEx",
            "color": "Black",
        },
        {
            "id": 2,
            "name": "GungiMaster",
            "color": "White",
        },
    ],
    "board_pieces": [
        {
            "id": 4,
            "type": "Pawn",
            "color": "Black",
            "position": [2, 5, 0],
        },
        {
            "id": 11,
            "type": "Archer",
            "color": "White",
            "position": [3, 4, 1],
        },
        ...
    ],
    "reserve_pieces": [
        {
            "id": 23,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "dicard_pieces": [
        {
            "id": 7,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "turn": 1,
    "phase": 1,
    "winner": null,
}
```

#### `POST /api/games/:id/actions/place`

Place a piece from reserve to board.

Parameters:

- `id`: Game ID.

Request:

```json
{
    "user_id": 1,
    "piece_id": 4,
    "position": [2, 5, 0],
}
```

Response:

```json
{
    "id": 1,
    "players": [
        {
            "id": 1,
            "name": "LoveSnowEx",
            "color": "Black",
        },
        {
            "id": 2,
            "name": "GungiMaster",
            "color": "White",
        },
    ],
    "board_pieces": [
        {
            "id": 4,
            "type": "Pawn",
            "color": "Black",
            "position": [2, 5, 0],
        },
        {
            "id": 11,
            "type": "Archer",
            "color": "White",
            "position": [3, 4, 1],
        },
        ...
    ],
    "reserve_pieces": [
        {
            "id": 23,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "dicard_pieces": [
        {
            "id": 7,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "turn": 2,
    "phase": 1,
    "winner": null,
}
```

#### `POST /api/games/:id/actions/move`

Move a piece.

Parameters:

- `id`: Game ID.

Request:

```json
{
    "user_id": 1,
    "piece_id": 4,
    "position": [2, 6, 0],
}
```

Response:

```json
{
    "id": 1,
    "players": [
        {
            "id": 1,
            "name": "LoveSnowEx",
            "color": "Black",
        },
        {
            "id": 2,
            "name": "GungiMaster",
            "color": "White",
        },
    ],
    "board_pieces": [
        {
            "id": 4,
            "type": "Pawn",
            "color": "Black",
            "position": [2, 6, 0],
        },
        {
            "id": 11,
            "type": "Archer",
            "color": "White",
            "position": [3, 4, 1],
        },
        ...
    ],
    "reserve_pieces": [
        {
            "id": 23,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "dicard_pieces": [
        {
            "id": 7,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "turn": 2,
    "phase": 1,
    "winner": null,
}
```

#### `POST /api/games/:id/actions/capture`

Capture a piece.

Parameters:

- `id`: Game ID.

Request:

```json
{
    "user_id": 1,
    "piece_id": 11,
    "position": [3, 6, 0],
}
```

Response:

```json
{
    "id": 1,
    "players": [
        {
            "id": 1,
            "name": "LoveSnowEx",
            "color": "Black",
        },
        {
            "id": 2,
            "name": "GungiMaster",
            "color": "White",
        },
    ],
    "board_pieces": [
        {
            "id": 4,
            "type": "Pawn",
            "color": "Black",
            "position": [2, 6, 0],
        },
        ...
    ],
    "reserve_pieces": [
        {
            "id": 23,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "dicard_pieces": [
        {
            "id": 7,
            "type": "Pawn",
            "color": "Black",
        },
        ...
    ],
    "turn": 2,
    "phase": 1,
    "winner": null,
}
```
