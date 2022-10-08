# csgo-death-trigger

Triggers a POST HTTP requests when you die in CS:GO

## Install

- Download binary from Releases
- Add to CS:GO launch options: `-condebug -netconport 2121` ([how?](https://support.steampowered.com/kb_article.php?ref=1040-JWMT-2947))

## Usage

Open the game, then run:

```bash
csgo-death-trigger.exe "s1mple" "http://127.0.0.1:3000" "{ \"a_value\": 123 }"
```

## Actions

- `make run`
- `make build`
- `make compress`
