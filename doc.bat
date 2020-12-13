
set MODULE=github.com/yodstar/dueros-dpl/dpl

godoc2md %MODULE%/command Command > doc/command.md
godoc2md %MODULE%/component Component > doc/component.md
godoc2md %MODULE% Document Template RenderConfig > doc/document.md
godoc2md %MODULE% ExecuteCommands > doc/execute-commands.md
godoc2md %MODULE% RenderDocument > doc/render-document.md
godoc2md %MODULE%/style Styles Stylesheet > doc/style.md
