# Translator

Translator is a CLI application to translate words or sentences. I'm using this little programing to put into practice all the new stuff
that I'm learning in Go.
Translator uses [libreTranslate API](https://libretranslate.com/docs/), during my tests a got some instability issues, so is possible that happens occasionally.

### Usage:

```
translator [command]

Available commands:
languages	 get available available languages to translate
translate	 has three arguments source language, target language, and return (enter) command.
                 after that, you can type a word or sentence to translate
help		 show this

Take a look at translator command, e.g.

translator en pt [return]
Translate
hello
Olá.
```
