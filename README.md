https://github.com/golang/go/issues/16426

This is a stupid simple library that's hoping to accomplish the tag-based validations suggested in the issue raised above. This first iteration does require two Unmarshal's because the current `json.UnmarshalJSON` interface acts on the whole struct and there are no places in the `json` package for a field-by-field interface call.

I'm using `github.com/jeffail/gabs` package for the KISS mentality.

# Abstract

This document proposes to support simple validation tags in the encoding/json module.

# Background

It's very common for the 2 ends, which communicate in JSON, to validate the received data before processing it. Current approaches include making the validation by introducing a schema validation package or making an in-line validation. Either solution has its shortcomings. A schema validation package seems too heavy for a JSON message containing only simple validations and is not efficient because it causes an extra decoding cycle of the message. An in-line validation will need extra efforts to maintain when the message structure changes.

# Proposal

Simple validations can be specified in the tags of corresponding Go's struct. When json.Unmarshal is called to decode the data, these tags will be considered and the data passed in will be checked against them. If one of the validations fails, json.Unmarshal will return with an error describing this situation.

For instance, given the Go struct:

    type Message struct {
        Type    int      `json:"type,required" json.enum:"1,2,3"`
        Content string   `json:"content" json.len:"[5,20]"`
    }

only a JSON object with a type field, which is of integer type and has a value between 1 and 3, and an optional content field, which is of string type and has a value containing 5 to 20 characters, can pass the validation.

Suggested tags are:

<table>

<thead>

<tr>

<th>Tag</th>

<th>Value</th>

<th>Effect Scope</th>

<th>Usage</th>

<th>Effect</th>

</tr>

</thead>

<tbody>

<tr>

<td>noExtraFields</td>

<td>-</td>

<td>struct</td>

<td>json="[...],noExtraFields"</td>

<td>the encoded data cannot contain extra fields other than listed in the struct</td>

</tr>

<tr>

<td>required</td>

<td>-</td>

<td>field</td>

<td>json="[...],required"</td>

<td>fields with this tag must appear in the encoded data</td>

</tr>

<tr>

<td>json.default</td>

<td>field default value</td>

<td>field</td>

<td>json.default="..."</td>

<td>provide a default value for that field if missing in the encoded data</td>

</tr>

<tr>

<td>json.enum</td>

<td>value list seperated by ','</td>

<td>field</td>

<td>json.enum="value1,..."</td>

<td>restrict the field value to be one of them</td>

</tr>

<tr>

<td>json.pattern</td>

<td>a regular expression</td>

<td>field</td>

<td>json.pattern="..."</td>

<td>restrict the field value must match the given regular expression</td>

</tr>

<tr>

<td>json.multipleOf</td>

<td>a number</td>

<td>field</td>

<td>json.multipleOf="..."</td>

<td>only valid to integer fields; restrict the field value must be a multiple of the given number</td>

</tr>

<tr>

<td>json.range</td>

<td>a range</td>

<td>field</td>

<td>json.range="[min,max]"</td>

<td>only valid to numeric fields; [ or ] means inclusive, ( or ) means exclusive; restrict the field value must be within the given range</td>

</tr>

<tr>

<td>json.len</td>

<td>a range</td>

<td>field</td>

<td>json.range="[min,max]"</td>

<td>only valid to string or array fields; [ or ] means inclusive, ( or ) means exclusive; restrict the field length must be within the given range</td>

</tr>

<tr>

<td>json.format</td>

<td>date-time, email, hostname, ipv4, ipv6 or uri</td>

<td>field</td>

<td>json.format="..."</td>

<td>restrict the field value must be in the given format</td>

</tr>

</tbody>

</table>

# Impact

The proposal won't break current behaviours. It only has little performance penalty and the implementation will be a little bit complex. It will also bring in package dependency to the regexp package if the pattern or format tag is supported.

# References

[JSON Schema Validation](http://json-schema.org/latest/json-schema-validation.html)