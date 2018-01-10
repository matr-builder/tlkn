## tlkn

matr builder basic helper util funcs

## Debug mode

To enable debug mode set debug to true

```go
  tlkn.Debug = true
```

## Bash(ctx context.Context, cmd string) (out []byte, err error) 
**Note: leading whitespace will be trimmed.
On multi-line strings each line's leading whitespace will be trimmed**

```go
  // one liner
  err := tlkn.Bash(context.Background(), "ls -la")

  // multi line
  // note: each line will be trimmed of leading whitespace
  err := tlkn.Bask(`ls -la
    echo "hello"
  `)
```

## BashCmd(ctx context.Context, cmd string) func() error

Simple wrapper around Bash func for use with Parallel

```go
  cmd := tlkn.BashCmd(context.Background(), "ls -la")

  if err := cmd(); err != nil {
    log.Fatal(err)
  }
```

## Parallel(fns ...func error) error

Run multiple funcs in parallel. Exits if any funcs return an error

```go
  ctx := context.Background()

  err := tlkn.Parallel(
    tlkn.BashCmd(ctx, "echo \"0\"")
    tlkn.BashCmd(ctx, "echo \"1\"")
    tlkn.BashCmd(ctx, "echo \"2\"")
    tlkn.BashCmd(ctx, "echo \"3\"")
    tlkn.BashCmd(ctx, "echo \"4\"")
  )

  if err != nil {
    log.Fatal(err)
  }
```

## Prompt(prompt string, args ...interface) (s string)

Prompt prompts user for input with default value.

```go
  v := tlkn.Prompt("Username:")
  // returns string value or "" if no value given
  fmt.Println(v) // prints string value 
```

## PromptRequired(prompt string, args ...interface) (s string)

Prompt prompts user for input with default value and requires an input.

```go
  v := tlkn.PromptRequired("state:")
  // returns string value if none present prompt is re-displayed
  fmt.Println(v)
```

## PromptConfirm(prompt string, args ...interface) bool 

PromptConfirm continues prompting until the input is boolean-ish.

```go
  v := tlkn.PromptConfirm("do you hate writing docs:")
  // accepted cli inputs:
  //   "Yes", "yes", "y", "Y"
  //   "No", "no", "n", "N"

  // prints bool response
  fmt.Println(v)
```

## PromptChoose(prompt string, list []string) int

Choose prompts for a single selection from `list`, returning in the index.

```go
  choices := []string{"dev", "qa", "staging", "prod"}

  // return selected choice index
  v := tlkn.PromptChoose("Where would you like to deploy?", choices)

  // prints choice string
  fmt.Println("Deploying to:", choices[v])
```

## Credits

Special thanks to the many packages that were used directly or as inspiration.

prompts : [segmentio/go-prompt](https://github.com/segmentio/go-prompt) 

<a href="https://github.com/bleveinc">
  <p align="center">
    <img src="https://raw.githubusercontent.com/matr-builder/matr-builder.github.io/master/assets/matr_footer.png">
  </p>
</a>
