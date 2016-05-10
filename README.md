# hound

Go CLI utility for the Data Dog API.

So far the only subcommand is `create-event`, but other API calls will be
added in the future.

## Commands

| command | description |
|---------|-------------|
|`create-event` | Creates an event |

## Examples

```bash
# Create an event with only a Title
hound create-event "First event"
```

```bash
# Create an event with a normal priority and an alert type "Error"
hound create-event \
    --priority "normal" \
    --alert-type "error" \
    --title "A Normal Error"
```

```bash
# Create an event with some text
hound create-event \
    --text "It looks like there was an error" \
    --title "This has a description"
```
