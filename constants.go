package phonenumbers

const MAX_INPUT_STRING_LENGTH = 250
const MIN_LENGTH_FOR_NSN = 2

const CAPTURING_EXTN_DIGITS = "(" + DIGITS + "{1,7})"
const EXTN_PATTERNS_FOR_PARSING = RFC3966_EXTN_PREFIX + CAPTURING_EXTN_DIGITS + "|" + "[ \u00A0\\t,]*" +
	"(?:e?xt(?:ensi(?:o\u0301?|\u00F3))?n?|\uFF45?\uFF58\uFF54\uFF4E?|" +
	"[,x\uFF58#\uFF03~\uFF5E]|int|anexo|\uFF49\uFF4E\uFF54)" +
	"[:\\.\uFF0E]?[ \u00A0\\t,-]*" + CAPTURING_EXTN_DIGITS + "#?|" +
	"[- ]+(" + DIGITS + "{1,5})#"
const PLUS_CHARS = "+\uFF0B"
const PLUS_SIGN = "+"
const DIGITS = "\\p{Nd}"
const STAR_SIGN = "*"
const VALID_ALPHA = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const VALID_PHONE_NUMBER = DIGITS + "{" + string(MIN_LENGTH_FOR_NSN) + "}" + "|" +
	"([" + PLUS_CHARS + "]*)+(?:[" + VALID_PUNCTUATION + STAR_SIGN + "]*" + DIGITS + "){3,}[" +
	VALID_PUNCTUATION + STAR_SIGN + VALID_ALPHA + DIGITS + "]*"
const VALID_PUNCTUATION = "-x\u2010-\u2015\u2212\u30FC\uFF0D-\uFF0F " +
	"\u00A0\u00AD\u200B\u2060\u3000()\uFF08\uFF09\uFF3B\uFF3D.\\[\\]/~\u2053\u223C\uFF5E"

const RFC3966_EXTN_PREFIX = ";ext="
const RFC3966_ISDN_SUBADDRESS = ";isub="
const RFC3966_PHONE_CONTEXT = ";phone-context="
const RFC3966_PREFIX = "tel:"
