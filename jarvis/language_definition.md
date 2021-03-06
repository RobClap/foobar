# Language definition
Operators are:
* `:` The column operator declares the name of the command
* `!` The exclamation mark means the order of the words in the following block matter (default is order don't matter)
* `?` The question mark means the following block may not be present
 * Two such blocks cannot appear in sequence
* `*` The asterisk mark means there might be words, but they are not relevant
* `()` Round parentheses can be used to specify the validity of an operator
* `[,]` Square brackets specify a comma separated list of synonyms that can be used
 * Example: `[add,sum,put together]`
* `{:}` Braces can be used to specify a parameter and its type. Supported types are `string` (can be omitted) `integer` and `date`
 * Example: `{amount:integer}` or `{person_name}` or `{when:date}`.
 * Two such blocks cannot appear in sequence
 * This can't be the first block of a command
 * A command cannot be constituted only of such blocks
* `<:|>` Chevrons can be used to specify alternative words, much like a multiple value parameter
 * Example: `<variable_name:value|othervalue>`
* Punctuation is not allowed
* The language is case insensitive

## Examples

### Volume changing
```
volumeHandler: <what:increase|[decrease,lower]> * volume ?( * {percentage:integer} ?percent)
```
The above example understands sentences like:
* Increase volume
* Decrease volume
* Increase the volume of ten percent
* Lower the volume of one hundred
* Volume Increase twenty

To handle such command a method with the following signature must be specified.
As an example some lines of the method have been implemented.
```python
def volumeHandler(what,percentage="10",type_error):
	if what=="increase":
		operator = "+"
	else
		operator = "-"
	volumeChange(operator + percentage + "%")
```
Where `what` is the word after the chevron and before the colon in the example and `percentage` is the word between braces.

Arguments are passed in order of appearance, even for non mandatory argument.

Some attempt to convert numbers in their numeric form will be performed. In any case the value passed will always be a string, so some internal checks should be implemented to make sure the string can be parsed as the expected type.

type_error is a list of integers that contains the index of the parameters where a type error occurred. In this example it can be `[1]` or `[]`

### Shopping List
```
!(<action:add|[remove,delete]> {what} [to,from] {when:date} * shopping list))
```
The above example understands sentences like:
* Add potatoes to tomorrow's shopping list
* Remove garlic from Wednesday's shopping list

And does not match the following ones:
* potatoes remove from Wednesday's shopping list

To handle such command a method with the following signature must be specified:
```python
def volumeHandler(action,what,when,type_error):
```
In this case the `when` parameter will be a string that Jarvis will attempt to transform so that the following line can parse it into a date:
```python
d = datetime.datetime.strptime( when , "%Y-%m-%dT%H:%M:%SZ" )
```
type_error is a list of integers that contains the index of the parameters where a type error occurred. In this example it can be `[1]`, `[1,2]`, `[]` and so on.
